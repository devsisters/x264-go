//go:build required

package x264

import (
	_ "github.com/devsisters/x264-go/x264c/external"
	_ "github.com/devsisters/x264-go/x264c/external/x264"
	_ "github.com/devsisters/x264-go/x264c/external/x264/common"
	_ "github.com/devsisters/x264-go/x264c/external/x264/common/aarch64"
	_ "github.com/devsisters/x264-go/x264c/external/x264/common/arm"
	_ "github.com/devsisters/x264-go/x264c/external/x264/common/mips"
	_ "github.com/devsisters/x264-go/x264c/external/x264/common/opencl"
	_ "github.com/devsisters/x264-go/x264c/external/x264/common/ppc"
	_ "github.com/devsisters/x264-go/x264c/external/x264/common/x86"
	_ "github.com/devsisters/x264-go/x264c/external/x264/encoder"
)
