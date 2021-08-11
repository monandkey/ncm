module github.com/monandkey/ncm

go 1.16

replace local.packages/ncm => ./pkg/ncm

replace local.packages/cmd => ./cmd

require (
	local.packages/cmd v0.0.0-00010101000000-000000000000
	local.packages/ncm v0.0.0-00010101000000-000000000000 // indirect
)
