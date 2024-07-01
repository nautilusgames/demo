package server

import (
	"bytes"
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/nautilusgames/demo/internal/logger"
	"github.com/vrischmann/envconfig"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"

	conf "github.com/nautilusgames/demo/pkg/config"
)

// Flags ...
type Flags struct {
	ConfigPath string
	Template   bool
	Validate   bool
}

// ParseFlags command line arguments.
func ParseFlags() *Flags {
	f := &Flags{}
	flag.StringVar(&f.ConfigPath, "c", "config.yaml", "path to YAML configuration")
	flag.BoolVar(&f.Template, "template", false, "executes go templates on the configuration file")
	flag.BoolVar(&f.Validate, "validate", false, "validates the configuration file and exits")
	flag.Parse()
	return f
}

// ParseFile ...
func ParseFile(path string, pb proto.Message, template bool) error {
	// Get absolute path representation for better error message in case file not found.
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Read file.
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Execute templates if Enabled.
	if template {
		contents, err = executeTemplate(contents)
		if err != nil {
			return err
		}
	}

	err = parseYAML(contents, pb)
	if err != nil {
		return err
	}
	//Load config from env variable with CONFIG_ prefix
	err = loadFromEnv(pb)
	return err
}

func loadFromEnv(pb proto.Message) error {
	return envconfig.InitWithOptions(pb, envconfig.Options{
		AllOptional:     true,
		AllowUnexported: true,
		LeaveNil:        true,
		Prefix:          "CONFIG"})
}

func parseYAML(contents []byte, pb proto.Message) error {
	// Decode YAML.
	var rawConfig map[string]interface{}
	if err := yaml.Unmarshal(contents, &rawConfig); err != nil {
		return err
	}

	// Encode YAML to JSON.
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(rawConfig); err != nil {
		return err
	}

	// Unmarshal JSON to proto object.
	if err := protojson.Unmarshal(jsonBuffer.Bytes(), pb); err != nil {
		return err
	}

	// All good!
	return nil
}

func executeTemplate(contents []byte) ([]byte, error) {
	tmpl := template.New("config").Funcs(map[string]interface{}{
		"getenv": os.Getenv,
		"getboolenv": func(key string) bool {
			b, _ := strconv.ParseBool(os.Getenv(key))
			return b
		},
	})

	tmpl, err := tmpl.Parse(string(contents))
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	if err := tmpl.Execute(&b, nil); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func loadConfig(f *Flags) *conf.Config {
	// Use a temporary logger to parse the configuration and output.
	tmpLogger := logger.New().With(zap.String("filename", f.ConfigPath))

	var cfg conf.Config
	if err := ParseFile(f.ConfigPath, &cfg, f.Template); err != nil {
		tmpLogger.Fatal("parsing configuration failed", zap.Error(err))
	}

	if err := cfg.Validate(); err != nil {
		tmpLogger.Fatal("validating configuration failed", zap.Error(err))
	}

	if f.Validate {
		tmpLogger.Info("configuration validation was successful")
		os.Exit(0)
	}

	return &cfg
}
