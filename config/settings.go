package config

import "context"

type SettingsContext string

const settings = SettingsContext("settings")

func FromContext(ctx context.Context) Settings {
	return ctx.Value(settings).(Settings)
}

func ToContext(ctx context.Context, conf Settings) context.Context {
	return context.WithValue(ctx, settings, conf)
}
