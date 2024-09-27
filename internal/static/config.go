package static

import (
	"os"
	"strconv"
	"strings"

	"github.com/agent-base/agentbase-sandbox/internal/types"
	"github.com/agent-base/agentbase-sandbox/internal/utils/log"
	"gopkg.in/yaml.v3"
)

var agentbaseSandboxGlobalConfigurations types.AgentBaseSandboxGlobalConfigurations

func InitConfig(path string) error {
	agentbaseSandboxGlobalConfigurations = types.AgentBaseSandboxGlobalConfigurations{}

	// read config file
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer configFile.Close()

	// parse config file
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&agentbaseSandboxGlobalConfigurations)
	if err != nil {
		return err
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err == nil {
		agentbaseSandboxGlobalConfigurations.App.Debug = debug
	}

	max_workers := os.Getenv("MAX_WORKERS")
	if max_workers != "" {
		agentbaseSandboxGlobalConfigurations.MaxWorkers, _ = strconv.Atoi(max_workers)
	}

	max_requests := os.Getenv("MAX_REQUESTS")
	if max_requests != "" {
		agentbaseSandboxGlobalConfigurations.MaxRequests, _ = strconv.Atoi(max_requests)
	}

	port := os.Getenv("SANDBOX_PORT")
	if port != "" {
		agentbaseSandboxGlobalConfigurations.App.Port, _ = strconv.Atoi(port)
	}

	timeout := os.Getenv("WORKER_TIMEOUT")
	if timeout != "" {
		agentbaseSandboxGlobalConfigurations.WorkerTimeout, _ = strconv.Atoi(timeout)
	}

	api_key := os.Getenv("API_KEY")
	if api_key != "" {
		agentbaseSandboxGlobalConfigurations.App.Key = api_key
	}

	python_path := os.Getenv("PYTHON_PATH")
	if python_path != "" {
		agentbaseSandboxGlobalConfigurations.PythonPath = python_path
	}

	if agentbaseSandboxGlobalConfigurations.PythonPath == "" {
		agentbaseSandboxGlobalConfigurations.PythonPath = "/usr/local/bin/python3"
	}

	python_lib_path := os.Getenv("PYTHON_LIB_PATH")
	if python_lib_path != "" {
		agentbaseSandboxGlobalConfigurations.PythonLibPaths = strings.Split(python_lib_path, ",")
	}

	if len(agentbaseSandboxGlobalConfigurations.PythonLibPaths) == 0 {
		agentbaseSandboxGlobalConfigurations.PythonLibPaths = DEFAULT_PYTHON_LIB_REQUIREMENTS
	}

	python_pip_mirror_url := os.Getenv("PIP_MIRROR_URL")
	if python_pip_mirror_url != "" {
		agentbaseSandboxGlobalConfigurations.PythonPipMirrorURL = python_pip_mirror_url
	}
	nodejs_path := os.Getenv("NODEJS_PATH")
	if nodejs_path != "" {
		agentbaseSandboxGlobalConfigurations.NodejsPath = nodejs_path
	}

	if agentbaseSandboxGlobalConfigurations.NodejsPath == "" {
		agentbaseSandboxGlobalConfigurations.NodejsPath = "/usr/local/bin/node"
	}

	enable_network := os.Getenv("ENABLE_NETWORK")
	if enable_network != "" {
		agentbaseSandboxGlobalConfigurations.EnableNetwork, _ = strconv.ParseBool(enable_network)
	}

	allowed_syscalls := os.Getenv("ALLOWED_SYSCALLS")
	if allowed_syscalls != "" {
		strs := strings.Split(allowed_syscalls, ",")
		ary := make([]int, len(strs))
		for i := range ary {
			ary[i], err = strconv.Atoi(strs[i])
			if err != nil {
				return err
			}
		}
		agentbaseSandboxGlobalConfigurations.AllowedSyscalls = ary
	}

	if agentbaseSandboxGlobalConfigurations.EnableNetwork {
		log.Info("network has been enabled")
		socks5_proxy := os.Getenv("SOCKS5_PROXY")
		if socks5_proxy != "" {
			agentbaseSandboxGlobalConfigurations.Proxy.Socks5 = socks5_proxy
		}

		if agentbaseSandboxGlobalConfigurations.Proxy.Socks5 != "" {
			log.Info("using socks5 proxy: %s", agentbaseSandboxGlobalConfigurations.Proxy.Socks5)
		}

		https_proxy := os.Getenv("HTTPS_PROXY")
		if https_proxy != "" {
			agentbaseSandboxGlobalConfigurations.Proxy.Https = https_proxy
		}

		if agentbaseSandboxGlobalConfigurations.Proxy.Https != "" {
			log.Info("using https proxy: %s", agentbaseSandboxGlobalConfigurations.Proxy.Https)
		}

		http_proxy := os.Getenv("HTTP_PROXY")
		if http_proxy != "" {
			agentbaseSandboxGlobalConfigurations.Proxy.Http = http_proxy
		}

		if agentbaseSandboxGlobalConfigurations.Proxy.Http != "" {
			log.Info("using http proxy: %s", agentbaseSandboxGlobalConfigurations.Proxy.Http)
		}
	}
	return nil
}

// avoid global modification, use value copy instead
func GetAgentBaseSandboxGlobalConfigurations() types.AgentBaseSandboxGlobalConfigurations {
	return agentbaseSandboxGlobalConfigurations
}

type RunnerDependencies struct {
	PythonRequirements string
}

var runnerDependencies RunnerDependencies

func GetRunnerDependencies() RunnerDependencies {
	return runnerDependencies
}

func SetupRunnerDependencies() error {
	file, err := os.ReadFile("dependencies/python-requirements.txt")
	if err != nil {
		if err == os.ErrNotExist {
			return nil
		}
		return err
	}

	runnerDependencies.PythonRequirements = string(file)

	return nil
}
