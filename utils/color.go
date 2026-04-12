package utils

import "fmt"

type ColorBuilder struct {
	color     int
	bold      bool
	underline bool
}

func Red() *ColorBuilder    { return &ColorBuilder{color: 31} }
func Green() *ColorBuilder  { return &ColorBuilder{color: 32} }
func Yellow() *ColorBuilder { return &ColorBuilder{color: 33} }
func Blue() *ColorBuilder   { return &ColorBuilder{color: 34} }
func Cyan() *ColorBuilder   { return &ColorBuilder{color: 36} }

func (c *ColorBuilder) Bold() *ColorBuilder {
	c.bold = true
	return c
}

func (c *ColorBuilder) Underline() *ColorBuilder {
	c.underline = true
	return c
}

func (c *ColorBuilder) buildCode() string {
	if c.bold && c.underline {
		return fmt.Sprintf("\033[1;4;%dm", c.color)
	}
	if c.bold {
		return fmt.Sprintf("\033[1;%dm", c.color)
	}
	if c.underline {
		return fmt.Sprintf("\033[4;%dm", c.color)
	}
	return fmt.Sprintf("\033[%dm", c.color)
}

func (c *ColorBuilder) Sprint(args ...any) string {
	msg := fmt.Sprint(args...)
	return fmt.Sprintf("%s%s\033[0m", c.buildCode(), msg)
}

func (c *ColorBuilder) Sprintf(format string, args ...any) string {
	msg := fmt.Sprintf(format, args...)
	return fmt.Sprintf("%s%s\033[0m", c.buildCode(), msg)
}

func (c *ColorBuilder) Print(args ...any) {
	fmt.Print(c.Sprint(args...))
}

func (c *ColorBuilder) Printf(format string, args ...any) {
	fmt.Print(c.Sprintf(format, args...))
}

func (c *ColorBuilder) Println(args ...any) {
	fmt.Println(c.Sprint(args...))
}
