package config

import (
	"fmt"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) Marshal() ([]byte, error) {
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

type Variables struct {
	ChatChannel                           string    `mapstructure:"CHAT_CHANNEL"`
	ChatInput                             string    `mapstructure:"CHAT_INPUT"`
	ChatUserId                            int16     `mapstructure:"CHAT_USER_ID"`
	CI                                    string    `mapstructure:"CI"`
	ApiUrl                                string    `mapstructure:"CI_API_V4_URL"`
	BuildsDir                             string    `mapstructure:"CI_BUILDS_DIR"`
	CommitAuthor                          string    `mapstructure:"CI_COMMIT_AUTHOR"`
	CommitBeforeSha                       string    `mapstructure:"CI_COMMIT_BEFORE_SHA"`
	CommitBranch                          string    `mapstructure:"CI_COMMIT_BRANCH"`
	CommitDescription                     string    `mapstructure:"CI_COMMIT_DESCRIPTION"`
	CommitMessage                         string    `mapstructure:"CI_COMMIT_MESSAGE"`
	CommitRefName                         string    `mapstructure:"CI_COMMIT_REF_NAME"`
	CommitRefProtected                    bool      `mapstructure:"CI_COMMIT_REF_PROTECTED"`
	CommitRefSlug                         string    `mapstructure:"CI_COMMIT_REF_SLUG"`
	CommiSha                              string    `mapstructure:"CI_COMMIT_SHA"`
	CommiShortSha                         string    `mapstructure:"CI_COMMIT_SHORT_SHA"`
	CommiTag                              string    `mapstructure:"CI_COMMIT_TAG"`
	CommiTagMessage                       string    `mapstructure:"CI_COMMIT_TAG_MESSAGE"`
	CommiTimespamp                        Timestamp `mapstructure:"CI_COMMIT_TIMESTAMP"`
	CommiTitle                            string    `mapstructure:"CI_COMMIT_TITLE"`
	CommiConcurrentId                     int16     `mapstructure:"CI_CONCURRENT_ID"`
	CommiConcurrentProjectId              int16     `mapstructure:"CI_CONCURRENT_PROJECT_ID"`
	ConfigPath                            string    `mapstructure:"CI_CONFIG_PATH"`
	DebugTrace                            bool      `mapstructure:"CI_DEBUG_TRACE"`
	DebugService                          bool      `mapstructure:"CI_DEBUG_SERVICE"`
	DefaultBranch                         string    `mapstructure:"CI_DEFAULT_BRANCH"`
	DependencyProxyGroupImagePrefix       string    `mapstructure:"CI_DEPENDENCY_PROXY_GROUP_IMAGE_PREFIX"`
	DependencyProxyDirectGroupImagePrefix string    `mapstructure:"CI_DEPENDENCY_PROXY_DIRECT_GROUP_IMAGE_PREFIX"`
	DependencyProxyPassword               string    `mapstructure:"CI_DEPENDENCY_PROXY_PASSWORD"`
	DependencyProxyServer                 string    `mapstructure:"CI_DEPENDENCY_PROXY_SERVER"`
	DependencyProxyUser                   string    `mapstructure:"CI_DEPENDENCY_PROXY_USER"`
	DeployFreeze                          bool      `mapstructure:"CI_DEPLOY_FREEZE"`
	DeployPassword                        string    `mapstructure:"CI_DEPLOY_PASSWORD"`
	DeployUser                            string    `mapstructure:"CI_DEPLOY_USER"`
	DisposableEnvironment                 bool      `mapstructure:"CI_DISPOSABLE_ENVIRONMENT"`
	EnvironmentName                       string    `mapstructure:"CI_ENVIRONMENT_NAME"`
	EnvironmentSlug                       string    `mapstructure:"CI_ENVIRONMENT_SLUG"`
	EnvironmentUrl                        string    `mapstructure:"CI_ENVIRONMENT_URL"`
	EnvironmentAction                     string    `mapstructure:"CI_ENVIRONMENT_ACTION"`
	EnvironmentTier                       string    `mapstructure:"CI_ENVIRONMENT_TIER"`
	ReleaseDescription                    string    `mapstructure:"CI_RELEASE_DESCRIPTION"`
	FipsMode                              bool      `mapstructure:"CI_GITLAB_FIPS_MODE"`
	HasOpenRequirements                   bool      `mapstructure:"CI_HAS_OPEN_REQUIREMENTS"`
	JobId                                 int32     `mapstructure:"CI_JOB_ID"`
	JobImage                              string    `mapstructure:"CI_JOB_IMAGE"`
	JobJwt                                string    `mapstructure:"CI_JOB_JWT"`
	JobJwtV1                              string    `mapstructure:"CI_JOB_JWT_V1"`
	JobJwtV2                              string    `mapstructure:"CI_JOB_JWT_V2"`
	JobManual                             string    `mapstructure:"CI_JOB_MANUAL"`
	JobName                               string    `mapstructure:"CI_JOB_NAME"`
	JobNameSlug                           string    `mapstructure:"CI_JOB_NAME_SLUG"`
	JobStage                              string    `mapstructure:"CI_JOB_STAGE"`
	JobStatus                             string    `mapstructure:"CI_JOB_STATUS"`
	JobTimeout                            int32     `mapstructure:"CI_JOB_TIMEOUT"`
	JobToken                              string    `mapstructure:"CI_JOB_TOKEN"`
	JobUrl                                string    `mapstructure:"CI_JOB_URL"`
	JobStartedAt                          Timestamp `mapstructure:"CI_JOB_STARTED_AT"`
	KubernetesActive                      string    `mapstructure:"CI_KUBERNETES_ACTIVE"`
	NodeIndex                             int16     `mapstructure:"CI_NODE_INDEX"`
	NodeTotal                             int16     `mapstructure:"CI_NODE_TOTAL"`
	OpenMergeRequests                     string    `mapstructure:"CI_OPEN_MERGE_REQUESTS"`
	PagesDomain                           string    `mapstructure:"CI_PAGES_DOMAIN"`
	PagesUrl                              string    `mapstructure:"CI_PAGES_URL"`
	PipelineId                            int32     `mapstructure:"CI_PIPELINE_ID"`
	PipelineIid                           int32     `mapstructure:"CI_PIPELINE_IID"`
	PipelineSource                        string    `mapstructure:"CI_PIPELINE_SOURCE"`
	PipelineTriggered                     bool      `mapstructure:"CI_PIPELINE_TRIGGERED"`
	PipelineUrl                           string    `mapstructure:"CI_PIPELINE_URL"`
	PipelineCreatedAt                     Timestamp `mapstructure:"CI_PIPELINE_CREATED_AT"`
	ProjectDir                            string    `mapstructure:"CI_PROJECT_DIR"`
	ProjectId                             int32     `mapstructure:"CI_PROJECT_ID"`
	ProjectName                           string    `mapstructure:"CI_PROJECT_NAME"`
	ProjectNamespace                      string    `mapstructure:"CI_PROJECT_NAMESPACE"`
	ProjectNamespaceId                    int32     `mapstructure:"CI_PROJECT_NAMESPACE_ID"`
	ProjectPathSlug                       string    `mapstructure:"CI_PROJECT_PATH_SLUG"`
	ProjectPath                           string    `mapstructure:"CI_PROJECT_PATH"`
	ProjectRepositoryLanguage             string    `mapstructure:"CI_PROJECT_REPOSITORY_LANGUAGES"`
	ProjectRootNamespace                  string    `mapstructure:"CI_PROJECT_ROOT_NAMESPACE"`
	ProjectTitle                          string    `mapstructure:"CI_PROJECT_TITLE"`
	ProjectDescription                    string    `mapstructure:"CI_PROJECT_DESCRIPTION"`
	ProjectUrl                            string    `mapstructure:"CI_PROJECT_URL"`
	ProjectVisibility                     string    `mapstructure:"CI_PROJECT_VISIBILITY"`
	ProjectClassificationLabel            string    `mapstructure:"CI_PROJECT_CLASSIFICATION_LABEL"`
	RegistryImage                         string    `mapstructure:"CI_REGISTRY_IMAGE"`
	RegistryPassword                      string    `mapstructure:"CI_REGISTRY_PASSWORD"`
	RegistryUser                          string    `mapstructure:"CI_REGISTRY_USER"`
	Registry                              string    `mapstructure:"CI_REGISTRY"`
	RepositoryUrl                         string    `mapstructure:"CI_REPOSITORY_URL"`
	RunnerDescription                     string    `mapstructure:"CI_RUNNER_DESCRIPTION"`
	RunnerExecutableArch                  string    `mapstructure:"CI_RUNNER_EXECUTABLE_ARCH"`
	RunnerId                              int16     `mapstructure:"CI_RUNNER_ID"`
	RunnerRevision                        string    `mapstructure:"CI_RUNNER_REVISION"`
	RunnerShortToken                      string    `mapstructure:"CI_RUNNER_SHORT_TOKEN"`
	RunnerTags                            string    `mapstructure:"CI_RUNNER_TAGS"`
	RunnerVersion                         string    `mapstructure:"CI_RUNNER_VERSION"`
	ServerHost                            string    `mapstructure:"CI_SERVER_HOST"`
	ServerName                            string    `mapstructure:"CI_SERVER_NAME"`
	ServerPort                            int16     `mapstructure:"CI_SERVER_PORT"`
	ServerProtocol                        string    `mapstructure:"CI_SERVER_PROTOCOL"`
	ServerRevision                        string    `mapstructure:"CI_SERVER_REVISION"`
	ServerTlsCaFile                       string    `mapstructure:"CI_SERVER_TLS_CA_FILE"`
	ServerTlsCertFile                     string    `mapstructure:"CI_SERVER_TLS_CERT_FILE"`
	ServerTlsKeyFile                      string    `mapstructure:"CI_SERVER_TLS_KEY_FILE"`
	ServerUrl                             string    `mapstructure:"CI_SERVER_URL"`
	SemverVersionMajor                    int8      `mapstructure:"CI_SERVER_VERSION_MAJOR"`
	SemverVersionMinor                    int8      `mapstructure:"CI_SERVER_VERSION_MINOR"`
	SemverVersionPatch                    int8      `mapstructure:"CI_SERVER_VERSION_PATCH"`
	SemverVersion                         string    `mapstructure:"CI_SERVER_VERSION"`
	Server                                string    `mapstructure:"CI_SERVER"`
	SharedEnvironment                     bool      `mapstructure:"CI_SHARED_ENVIRONMENT"`
	TemplateRegistryHost                  string    `mapstructure:"CI_TEMPLATE_REGISTRY_HOST"`
	GitlabFeatures                        string    `mapstructure:"GITLAB_FEATURES"`
	GitlabuserEmail                       string    `mapstructure:"GITLAB_USER_EMAIL"`
	GitlabuserId                          int16     `mapstructure:"GITLAB_USER_ID"`
	GitlabuserLogin                       string    `mapstructure:"GITLAB_USER_LOGIN"`
	GitlabUsername                        string    `mapstructure:"GITLAB_USER_NAME"`
	TriggerPayload                        string    `mapstructure:"TRIGGER_PAYLOAD"`
	MergeRequestApproved                  bool      `mapstructure:"CI_MERGE_REQUEST_APPROVED"`
	MergeRequestAssignees                 string    `mapstructure:"CI_MERGE_REQUEST_ASSIGNEES"`
	MergeRequestId                        int16     `mapstructure:"CI_MERGE_REQUEST_ID"`
	MergeRequestIid                       int16     `mapstructure:"CI_MERGE_REQUEST_IID"`
	MergeRequestLabels                    string    `mapstructure:"CI_MERGE_REQUEST_LABELS"`
	MergeRequestMilestone                 string    `mapstructure:"CI_MERGE_REQUEST_MILESTONE"`
	MergeRequestProjectId                 int16     `mapstructure:"CI_MERGE_REQUEST_PROJECT_ID"`
	MergeRequestProjectPath               string    `mapstructure:"CI_MERGE_REQUEST_PROJECT_PATH"`
	MergeRequestProjectUrl                string    `mapstructure:"CI_MERGE_REQUEST_PROJECT_URL"`
	MergeRequestRefPath                   string    `mapstructure:"CI_MERGE_REQUEST_REF_PATH"`
	MergeRequestSourceBranchName          string    `mapstructure:"CI_MERGE_REQUEST_SOURCE_BRANCH_NAME"`
	MergeRequestSourceBranchSha           string    `mapstructure:"CI_MERGE_REQUEST_SOURCE_BRANCH_SHA"`
	MergeRequestSourceProjectId           int16     `mapstructure:"CI_MERGE_REQUEST_SOURCE_PROJECT_ID"`
	MergeRequestSourceProjectPath         string    `mapstructure:"CI_MERGE_REQUEST_SOURCE_PROJECT_PATH"`
	MergeRequestSourceProjectUrl          string    `mapstructure:"CI_MERGE_REQUEST_SOURCE_PROJECT_URL"`
	MergeRequestTargetBranchName          string    `mapstructure:"CI_MERGE_REQUEST_TARGET_BRANCH_NAME"`
	MergeRequestTargetBranchProtected     bool      `mapstructure:"CI_MERGE_REQUEST_TARGET_BRANCH_PROTECTED"`
	MergeRequestTargetBranchSha           string    `mapstructure:"CI_MERGE_REQUEST_TARGET_BRANCH_SHA"`
	MergeRequestTitle                     string    `mapstructure:"CI_MERGE_REQUEST_TITLE"`
	MergeRequestEventType                 string    `mapstructure:"CI_MERGE_REQUEST_EVENT_TYPE"`
	MergeRequestDiffId                    int16     `mapstructure:"CI_MERGE_REQUEST_DIFF_ID"`
	MergeRequestDiffBaseSha               string    `mapstructure:"CI_MERGE_REQUEST_DIFF_BASE_SHA"`
}