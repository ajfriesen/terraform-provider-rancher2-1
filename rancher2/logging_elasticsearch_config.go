package rancher2

import (
	"github.com/hashicorp/terraform/helper/schema"
	managementClient "github.com/rancher/types/client/management/v3"
)

const (
	loggingElasticsearchKind = "elasticsearch"
)

//Schemas

func elasticsearchConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"endpoint": {
			Type:     schema.TypeString,
			Required: true,
		},
		"auth_password": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"auth_username": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"certificate": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"client_cert": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"client_key": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"client_key_pass": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"date_format": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "YYYY-MM-DD",
		},
		"index_prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "local",
		},
		"ssl_verify": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"ssl_version": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}

	return s
}

// Flatteners

func flattenElasticsearchConfig(in *managementClient.ElasticsearchConfig) ([]interface{}, error) {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}, nil
	}

	obj["endpoint"] = in.Endpoint
	obj["date_format"] = in.DateFormat
	obj["index_prefix"] = in.IndexPrefix

	if len(in.AuthPassword) > 0 {
		obj["auth_password"] = in.AuthPassword
	}

	if len(in.AuthUserName) > 0 {
		obj["auth_username"] = in.AuthUserName
	}

	if len(in.Certificate) > 0 {
		obj["certificate"] = in.Certificate
	}

	if len(in.ClientCert) > 0 {
		obj["client_cert"] = in.ClientCert
	}

	if len(in.ClientKey) > 0 {
		obj["client_key"] = in.ClientKey
	}

	if len(in.ClientKeyPass) > 0 {
		obj["client_key_pass"] = in.ClientKeyPass
	}

	obj["ssl_verify"] = in.SSLVerify

	if len(in.SSLVersion) > 0 {
		obj["ssl_version"] = in.SSLVersion
	}

	return []interface{}{obj}, nil
}

// Expanders

func expandElasticsearchConfig(p []interface{}) (*managementClient.ElasticsearchConfig, error) {
	obj := &managementClient.ElasticsearchConfig{}

	if len(p) == 0 || p[0] == nil {
		return obj, nil
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["endpoint"].(string); ok && len(v) > 0 {
		obj.Endpoint = v
	}

	if v, ok := in["date_format"].(string); ok && len(v) > 0 {
		obj.DateFormat = v
	}

	if v, ok := in["index_prefix"].(string); ok && len(v) > 0 {
		obj.IndexPrefix = v
	}

	if v, ok := in["auth_password"].(string); ok && len(v) > 0 {
		obj.AuthPassword = v
	}

	if v, ok := in["auth_username"].(string); ok && len(v) > 0 {
		obj.AuthUserName = v
	}

	if v, ok := in["certificate"].(string); ok && len(v) > 0 {
		obj.Certificate = v
	}

	if v, ok := in["client_cert"].(string); ok && len(v) > 0 {
		obj.ClientCert = v
	}

	if v, ok := in["client_key"].(string); ok && len(v) > 0 {
		obj.ClientKey = v
	}

	if v, ok := in["client_key_pass"].(string); ok && len(v) > 0 {
		obj.ClientKeyPass = v
	}

	if v, ok := in["ssl_verify"].(bool); ok {
		obj.SSLVerify = v
	}

	if v, ok := in["ssl_version"].(string); ok && len(v) > 0 {
		obj.SSLVersion = v
	}

	return obj, nil
}
