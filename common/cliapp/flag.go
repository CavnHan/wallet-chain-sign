package cliapp

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// CloneableGeneric 接口嵌入了 cli.Generic 接口，并添加了 Clone 方法。
// 这个接口用于表示可以克隆的通用标志值。
type CloneableGeneric interface {
	cli.Generic
	Clone() any
}

// ProtectFlags 函数接受一个 cli.Flag 切片，并返回一个新的切片，其中包含了所有标志的克隆副本。
// 这个函数通过调用 cloneFlag 函数来克隆每个标志，并在克隆失败时抛出一个恐慌（panic）。
func ProtectFlags(flags []cli.Flag) []cli.Flag {
	// 创建一个新的切片，用于存储克隆的标志
	out := make([]cli.Flag, 0, len(flags))
	// 遍历每个标志
	for _, f := range flags {
		// 克隆标志
		fCopy, err := cloneFlag(f)
		if err != nil {
			// 如果克隆失败，抛出一个恐慌（panic）
			panic(fmt.Errorf("failed to clone flag %q: %w", f.Names()[0], err))
		}
		// 将克隆的标志添加到新的切片中
		out = append(out, fCopy)
	}
	// 返回包含克隆标志的新切片
	return out
}

// cloneFlag 函数接受一个 cli.Flag，并返回它的克隆副本。
// 如果标志是 *cli.GenericFlag 类型，并且其值实现了 CloneableGeneric 接口，则克隆该值并返回新的标志副本。
// 对于其他类型的标志，直接返回原始标志。
func cloneFlag(f cli.Flag) (cli.Flag, error) {
	// 使用类型断言检查标志的具体类型
	switch typedFlag := f.(type) {
	case *cli.GenericFlag:
		// 如果标志是 *cli.GenericFlag 类型，并且其值实现了 CloneableGeneric 接口
		if genValue, ok := typedFlag.Value.(CloneableGeneric); ok {
			// 创建 *cli.GenericFlag 的副本
			cpy := *typedFlag
			// 克隆标志的值
			cpyVal, ok := genValue.Clone().(cli.Generic)
			if !ok {
				// 如果克隆的值不是 cli.Generic 类型，返回错误
				return nil, fmt.Errorf("cloned Generic value is not Generic: %T", typedFlag)
			}
			// 将克隆的值赋给副本
			cpy.Value = cpyVal
			// 返回标志副本
			return &cpy, nil
		} else {
			// 如果值不能被克隆，返回错误
			return nil, fmt.Errorf("cannot clone Generic value: %T", typedFlag)
		}
	default:
		// 对于其他类型的标志，直接返回原始标志
		return f, nil
	}
}
