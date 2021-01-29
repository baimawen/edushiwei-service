/*
 * @Author: your name
 * @Date: 2021-01-29 11:30:38
 * @LastEditTime: 2021-01-29 13:04:59
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \edushiwei-service\config\config.go
 */
package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email  Email  `mapstructure:"email" json:"email" yaml:"email"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Local  Local  `mapstructure:"local" json:"local" yaml:"local"`
}
