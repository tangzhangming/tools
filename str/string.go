package str


//返回字符串中指定值之后的所有内容
func After(s string, f string) string {}

//返回字符串中指定值最后一次出现后的所有内容
func AfterLast(s string, f string) string {}

//字符串转换为 ASCII 值
func Ascii(s string) string {}

//返回字符串中指定值最后一次出现前的所有内容
func Before(s string, f string) string {}

//返回字符串中指定值最后一次出现前的所有内容
func BeforeLast(s string, f string) string {}

//返回字符串中指定值最后一次出现前的所有内容
func Between(s string, start string, end string) string {}

//将指定字符串转换为 驼峰式 表示方法
func Camel(s string) string {}

//将字符串转换为 烤串式（ kebab-case ） 表示方法：
// fooBar => foo-bar, foo_bar => foo-bar, 
func Kebab(s string) string {}

// 用于判断指定字符串是否包含指定数组中的所有值
func ContainsAll(s string, ...value string) string {}

// 将指定的字符串修改为以指定的值结尾的形式
// str.Finish('www.baidu.com', 'http://') == str.Finish('http://www.baidu.com', 'http://') == http://www.baidu.com
func Start(s string, f string) string {}

// 将指定的字符串修改为以指定的值结尾的形式
// str.Finish('this/string', '/') == str.Finish('this/string/', '/') == this/string/
func Finish(s string, f string) string {}

// 用来判断字符串是否与指定模式匹配。星号 * 可用于表示通配符
// str.Is('foo*', 'foobar') == true
func Is(s string, f string) string {}

// 方法会使用重复的字符掩盖字符串的一部分，并可用于混淆字符串段，例如电子邮件地址和电话号码
// taylor@example.com => tay***@example.com
func Mask(s string, start string, end string) string {}


// 在指定字符串的左侧填充上另一字符串
// PadLeft("abcdd", 10, "01") => 01010abcdd
func PadLeft(s string, len string, tcw string) string {}

// 在指定字符串的右侧填充上另一字符串
// PadRight("abcdd", 10, "01") => abcdd01010
func PadRight(s string, len string, tcw string) string {}

