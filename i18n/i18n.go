package i18n

import (
	"encoding/json"
	"fmt"
	"os"
)

// Translator 是多语言翻译器的主结构
type Translator struct {
	translations map[string]map[string]string
	defaultLang  string
}

// New 创建一个新的翻译器实例
func New(defaultLang string) *Translator {
	return &Translator{
		translations: make(map[string]map[string]string),
		defaultLang:  defaultLang,
	}
}

// LoadTranslationFile 从JSON文件加载翻译
func (t *Translator) LoadTranslationFile(lang, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取翻译文件失败: %w", err)
	}

	translations := make(map[string]string)
	if err := json.Unmarshal(data, &translations); err != nil {
		return fmt.Errorf("解析翻译文件失败: %w", err)
	}

	t.translations[lang] = translations
	return nil
}

// T 翻译指定的key
func (t *Translator) T(lang, key string, args ...interface{}) string {
	// 如果语言不存在，使用默认语言
	translations, ok := t.translations[lang]
	if !ok {
		translations = t.translations[t.defaultLang]
	}

	// 获取翻译文本
	text, ok := translations[key]
	if !ok {
		return key
	}

	// 如果有参数，进行格式化
	if len(args) > 0 {
		text = fmt.Sprintf(text, args...)
	}

	return text
}
