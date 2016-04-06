package libyaml

type HostVolume struct {
	HostPath             string `yaml:"host_path" json:"host_path"`
	Owner                string `yaml:"owner" json:"owner"`                                     // TODO: not yet supported
	Permission           string `yaml:"permission" json:"permission"`                           // TODO: not yet supported
	IsExcludedFromBackup string `yaml:"is_excluded_from_backup" json:"is_excluded_from_backup"` // TODO: not yet supported
	MinDiskSpace         string `yaml:"min_disk_space" json:"min_disk_space"`
}

type ContainerVolume struct {
	HostPath             string `yaml:"host_path" json:"host_path"`
	ContainerPath        string `yaml:"container_path" json:"container_path"`
	Permission           string `yaml:"permission" json:"permission"`                           // TODO: deprecate
	Owner                string `yaml:"owner" json:"owner"`                                     // TODO: deprecate
	IsExcludedFromBackup string `yaml:"is_excluded_from_backup" json:"is_excluded_from_backup"` // TODO: deprecate
}
