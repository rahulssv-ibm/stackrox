package env

var (
	// ProcessFilterMaxProcessPaths sets the maximum number of process filter unique paths
	ProcessFilterMaxProcessPaths = RegisterIntegerSetting("ROX_PROCESS_FILTER_MAX_PROCESS_PATHS", 5000)

	// BucketSizesConfig sets the bucket sizes for fan out filtering of process indicators.
	BucketSizesConfig = RegisterIntegerListSetting("ROX_PROCESS_FILTER_BUCKET_SIZES", []int{8, 6, 4, 2})
)
