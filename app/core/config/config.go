package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"strconv"
)

/*
*
也可以了解viper包
*/
type Model struct {
	Name  string            `yaml:name`
	HTTP  http              `yaml:"http" env:"http"`
	Log   log               `yaml:"log"`
	DB    map[string]dbItem `yaml:"db"`
	Redis map[string]Redis  `yaml:"redis"`
}

func (m *Model) LoadConfig(prjHome string) error {
	confDef := prjHome + "/config/config-default.yml"
	confOverride := prjHome + "/config/config.yml"

	err := m.parseConfig(confDef)
	if err != nil {
		return err
	}

	err = m.parseConfig(confOverride)
	if err != nil {
		return err
	}

	err = m.MergeEnv()
	if err != nil {
		return err
	}
	return nil

}

func (m *Model) parseConfig(conf string) error {

	data, err := os.ReadFile(conf)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, m)

	if err != nil {
		return err
	}

	return nil
}

func (m *Model) MergeEnv() error {
	assign(reflect.ValueOf(m))
	return nil
}

func assign(val reflect.Value) {
	v := reflect.Indirect(val)

	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Tag.Get("env")

		processOne(v.Field(i), key)
	}
}

func processOne(field reflect.Value, key string) {

	envVal, envOk := os.LookupEnv(key)
	switch field.Type().Kind() {
	case reflect.String:
		if !envOk {
			return
		}

		field.SetString(envVal)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		if !envOk {
			return
		}

		intVal, e := strconv.ParseInt(envVal, 0, field.Type().Bits())
		if e != nil {
			return
		}
		field.SetInt(intVal)

	case reflect.Float64, reflect.Float32:

		if !envOk {
			return
		}

		val, e := strconv.ParseInt(envVal, 0, field.Type().Bits())
		if e != nil {
			return
		}
		field.SetInt(val)
	case reflect.Bool:

		if !envOk {
			return
		}

		val, e := strconv.ParseBool(envVal)
		if e != nil {
			return
		}
		field.SetBool(val)

	case reflect.Struct:

		assign(field)

	}

}
