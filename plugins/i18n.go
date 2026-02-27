package plugins

import (
	"sort"
	"strings"
)

// Strings holds every user-facing message the bot can send.
// Format-string fields use standard fmt verbs (%s, %d, etc.).
type Strings struct {
	// ping
	Pong        string
	PongLatency string // fmt: latency ms

	// meta
	MetaUsage string

	// dispatch
	GroupOnly string
	SudoOnly  string

	// setprefix
	SetPrefixUsage   string
	SetPrefixUpdated string // fmt: prefix list
	SaveFailed       string // fmt: error string

	// setsudo
	SetSudoUsage  string
	SudoAdded     string // fmt: phone
	SudoRemoved   string // fmt: phone
	SudoNotFound  string // fmt: phone
	UnknownAction string

	// setmode
	SetModeUsage   string
	ModePublicSet  string
	ModePrivateSet string

	// lang
	LangCurrent string // fmt: language name
	LangSet     string // fmt: language name
	LangUnknown string // fmt: attempted code, available codes
	LangUsage   string // fmt: available codes

	// menu
	MenuGreeting string // fmt: push name
}

// LangNames maps ISO language codes to their native display names.
var LangNames = map[string]string{
	"en": "English",
	"es": "Español",
	"pt": "Português",
	"ar": "العربية",
	"hi": "हिन्दी",
	"fr": "Français",
	"de": "Deutsch",
	"ru": "Русский",
	"tr": "Türkçe",
	"sw": "Kiswahili",
}

// availableLangs returns an alphabetically sorted, comma-separated list of
// all registered language codes.
func availableLangs() string {
	codes := make([]string, 0, len(translations))
	for code := range translations {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	return strings.Join(codes, ", ")
}

// langList returns a formatted list of all supported languages, one per line,
// in the form "English - en".
func langList() string {
	codes := make([]string, 0, len(LangNames))
	for code := range LangNames {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	var b strings.Builder
	for _, code := range codes {
		b.WriteString(LangNames[code] + " - " + code + "\n")
	}
	return strings.TrimRight(b.String(), "\n")
}

// T returns the Strings set for the current bot language, falling back to
// English when the stored code has no registered translation.
func T() *Strings {
	lang := BotSettings.GetLanguage()
	if s, ok := translations[lang]; ok {
		return s
	}
	return translations["en"]
}

// translations is the registry of all supported languages.
var translations = map[string]*Strings{
	"en": {
		Pong:             "Pong",
		PongLatency:      "Pong (%dms)",
		MetaUsage:        "Provide a query to send to Meta AI.\nUsage: meta <question>",
		GroupOnly:        "This command is restricted to group chats.",
		SudoOnly:         "This command is restricted to authorised users.",
		SetPrefixUsage:   "Specify one or more command prefixes.\nUsage: setprefix <p1> <p2> ...\nUse the token \"empty\" for a prefix-free entry.\nExample: setprefix . / #",
		SetPrefixUpdated: "Command prefix updated: %s",
		SaveFailed:       "Failed to persist settings: %s",
		SetSudoUsage:     "Usage: setsudo add|remove <phone>\nExample: setsudo add 1234567890",
		SudoAdded:        "%s has been granted sudo access.",
		SudoRemoved:      "Sudo access has been revoked from %s.",
		SudoNotFound:     "%s is not a sudo user.",
		UnknownAction:    "Unknown action. Accepted values: add, remove",
		SetModeUsage:     "Usage: setmode public|private",
		ModePublicSet:    "Mode set to public. All users may issue commands.",
		ModePrivateSet:   "Mode set to private. Commands are restricted to sudo users.",
		LangCurrent:      "Current language: %s",
		LangSet:          "Language set to: %s",
		LangUnknown:      "Unrecognised language code: %s\nAvailable: %s",
		LangUsage:        "Usage: lang <code>\nAvailable: %s",
		MenuGreeting:     "ʜᴇʏ, %s!",
	},
	"es": {
		Pong:             "Pong",
		PongLatency:      "Pong (%dms)",
		MetaUsage:        "Proporcione una consulta para enviar a Meta AI.\nUso: meta <pregunta>",
		GroupOnly:        "Este comando está restringido a conversaciones de grupo.",
		SudoOnly:         "Este comando está restringido a usuarios autorizados.",
		SetPrefixUsage:   "Especifique uno o más prefijos de comando.\nUso: setprefix <p1> <p2> ...\nUtilice el símbolo \"empty\" para una entrada sin prefijo.\nEjemplo: setprefix . / #",
		SetPrefixUpdated: "Prefijo de comando actualizado: %s",
		SaveFailed:       "Error al guardar la configuración: %s",
		SetSudoUsage:     "Uso: setsudo add|remove <teléfono>\nEjemplo: setsudo add 1234567890",
		SudoAdded:        "Se ha concedido acceso sudo a %s.",
		SudoRemoved:      "El acceso sudo de %s ha sido revocado.",
		SudoNotFound:     "%s no es un usuario sudo.",
		UnknownAction:    "Acción desconocida. Valores aceptados: add, remove",
		SetModeUsage:     "Uso: setmode public|private",
		ModePublicSet:    "Modo establecido en público. Todos los usuarios pueden ejecutar comandos.",
		ModePrivateSet:   "Modo establecido en privado. Los comandos están restringidos a usuarios sudo.",
		LangCurrent:      "Idioma actual: %s",
		LangSet:          "Idioma establecido en: %s",
		LangUnknown:      "Código de idioma no reconocido: %s\nDisponibles: %s",
		LangUsage:        "Uso: lang <código>\nDisponibles: %s",
		MenuGreeting:     "ʜᴏʟᴀ, %s!",
	},
	"pt": {
		Pong:             "Pong",
		PongLatency:      "Pong (%dms)",
		MetaUsage:        "Forneça uma consulta para enviar ao Meta AI.\nUtilização: meta <pergunta>",
		GroupOnly:        "Este comando está restrito a conversas em grupo.",
		SudoOnly:         "Este comando está restrito a utilizadores autorizados.",
		SetPrefixUsage:   "Especifique um ou mais prefixos de comando.\nUtilização: setprefix <p1> <p2> ...\nUtilize o símbolo \"empty\" para uma entrada sem prefixo.\nExemplo: setprefix . / #",
		SetPrefixUpdated: "Prefixo de comando atualizado: %s",
		SaveFailed:       "Falha ao guardar as definições: %s",
		SetSudoUsage:     "Utilização: setsudo add|remove <telefone>\nExemplo: setsudo add 1234567890",
		SudoAdded:        "%s recebeu acesso sudo.",
		SudoRemoved:      "O acesso sudo de %s foi revogado.",
		SudoNotFound:     "%s não é um utilizador sudo.",
		UnknownAction:    "Ação desconhecida. Valores aceites: add, remove",
		SetModeUsage:     "Utilização: setmode public|private",
		ModePublicSet:    "Modo definido como público. Todos os utilizadores podem executar comandos.",
		ModePrivateSet:   "Modo definido como privado. Os comandos estão restritos a utilizadores sudo.",
		LangCurrent:      "Idioma atual: %s",
		LangSet:          "Idioma definido para: %s",
		LangUnknown:      "Código de idioma não reconhecido: %s\nDisponíveis: %s",
		LangUsage:        "Utilização: lang <código>\nDisponíveis: %s",
		MenuGreeting:     "ᴏʟᴀ, %s!",
	},
	"ar": {
		Pong:             "بونغ",
		PongLatency:      "بونغ (%dms)",
		MetaUsage:        "أدخل استفساراً لإرساله إلى Meta AI.\nالاستخدام: meta <سؤال>",
		GroupOnly:        "هذا الأمر مقتصر على المجموعات.",
		SudoOnly:         "هذا الأمر مقتصر على المستخدمين المخوَّلين.",
		SetPrefixUsage:   "حدِّد بادئة واحدة أو أكثر.\nالاستخدام: setprefix <ب1> <ب2> ...\nاستخدم الرمز \"empty\" لإدخال بدون بادئة.\nمثال: setprefix . / #",
		SetPrefixUpdated: "تم تحديث بادئة الأمر: %s",
		SaveFailed:       "فشل حفظ الإعدادات: %s",
		SetSudoUsage:     "الاستخدام: setsudo add|remove <رقم الهاتف>\nمثال: setsudo add 1234567890",
		SudoAdded:        "تم منح %s صلاحيات المشرف.",
		SudoRemoved:      "تم سحب صلاحيات المشرف من %s.",
		SudoNotFound:     "%s ليس مستخدماً مشرفاً.",
		UnknownAction:    "إجراء غير معروف. القيم المقبولة: add، remove",
		SetModeUsage:     "الاستخدام: setmode public|private",
		ModePublicSet:    "تم تعيين الوضع على عام. يمكن لجميع المستخدمين تنفيذ الأوامر.",
		ModePrivateSet:   "تم تعيين الوضع على خاص. الأوامر مقتصرة على المشرفين.",
		LangCurrent:      "اللغة الحالية: %s",
		LangSet:          "تم تعيين اللغة إلى: %s",
		LangUnknown:      "رمز اللغة غير معروف: %s\nالمتاحة: %s",
		LangUsage:        "الاستخدام: lang <رمز>\nالمتاحة: %s",
		MenuGreeting:     "أهلاً، %s!",
	},
	"hi": {
		Pong:             "पोंग",
		PongLatency:      "पोंग (%dms)",
		MetaUsage:        "Meta AI को भेजने के लिए एक प्रश्न दर्ज करें।\nउपयोग: meta <प्रश्न>",
		GroupOnly:        "यह आदेश केवल समूह वार्तालापों में उपयोग किया जा सकता है।",
		SudoOnly:         "यह आदेश केवल अधिकृत उपयोगकर्ताओं के लिए है।",
		SetPrefixUsage:   "एक या अधिक आदेश उपसर्ग निर्दिष्ट करें।\nउपयोग: setprefix <उ1> <उ2> ...\nबिना उपसर्ग की प्रविष्टि के लिए \"empty\" का उपयोग करें।\nउदाहरण: setprefix . / #",
		SetPrefixUpdated: "आदेश उपसर्ग अद्यतन किया गया: %s",
		SaveFailed:       "सेटिंग्स सहेजने में विफल: %s",
		SetSudoUsage:     "उपयोग: setsudo add|remove <फ़ोन>\nउदाहरण: setsudo add 1234567890",
		SudoAdded:        "%s को sudo अधिकार प्रदान किए गए।",
		SudoRemoved:      "%s के sudo अधिकार रद्द किए गए।",
		SudoNotFound:     "%s एक sudo उपयोगकर्ता नहीं है।",
		UnknownAction:    "अज्ञात क्रिया। स्वीकृत मान: add, remove",
		SetModeUsage:     "उपयोग: setmode public|private",
		ModePublicSet:    "मोड सार्वजनिक पर सेट किया गया। सभी उपयोगकर्ता आदेश दे सकते हैं।",
		ModePrivateSet:   "मोड निजी पर सेट किया गया। आदेश केवल sudo उपयोगकर्ताओं तक सीमित हैं।",
		LangCurrent:      "वर्तमान भाषा: %s",
		LangSet:          "भाषा निर्धारित की गई: %s",
		LangUnknown:      "अपरिचित भाषा कोड: %s\nउपलब्ध: %s",
		LangUsage:        "उपयोग: lang <कोड>\nउपलब्ध: %s",
		MenuGreeting:     "नमस्ते, %s!",
	},
	"fr": {
		Pong:             "Pong",
		PongLatency:      "Pong (%dms)",
		MetaUsage:        "Veuillez saisir une requête à envoyer à Meta AI.\nUtilisation : meta <question>",
		GroupOnly:        "Cette commande est réservée aux conversations de groupe.",
		SudoOnly:         "Cette commande est réservée aux utilisateurs autorisés.",
		SetPrefixUsage:   "Indiquez un ou plusieurs préfixes de commande.\nUtilisation : setprefix <p1> <p2> ...\nUtilisez le symbole \"empty\" pour une entrée sans préfixe.\nExemple : setprefix . / #",
		SetPrefixUpdated: "Préfixe de commande mis à jour : %s",
		SaveFailed:       "Échec de l'enregistrement des paramètres : %s",
		SetSudoUsage:     "Utilisation : setsudo add|remove <téléphone>\nExemple : setsudo add 1234567890",
		SudoAdded:        "Les droits sudo ont été accordés à %s.",
		SudoRemoved:      "Les droits sudo de %s ont été révoqués.",
		SudoNotFound:     "%s n'est pas un utilisateur sudo.",
		UnknownAction:    "Action inconnue. Valeurs acceptées : add, remove",
		SetModeUsage:     "Utilisation : setmode public|private",
		ModePublicSet:    "Mode défini sur public. Tous les utilisateurs peuvent exécuter des commandes.",
		ModePrivateSet:   "Mode défini sur privé. Les commandes sont réservées aux utilisateurs sudo.",
		LangCurrent:      "Langue actuelle : %s",
		LangSet:          "Langue définie sur : %s",
		LangUnknown:      "Code de langue non reconnu : %s\nDisponibles : %s",
		LangUsage:        "Utilisation : lang <code>\nDisponibles : %s",
		MenuGreeting:     "ʙᴏɴᴊᴏᴜʀ, %s!",
	},
	"de": {
		Pong:             "Pong",
		PongLatency:      "Pong (%dms)",
		MetaUsage:        "Geben Sie eine Anfrage für Meta AI ein.\nVerwendung: meta <Frage>",
		GroupOnly:        "Dieser Befehl ist auf Gruppenunterhaltungen beschränkt.",
		SudoOnly:         "Dieser Befehl ist auf autorisierte Benutzer beschränkt.",
		SetPrefixUsage:   "Geben Sie ein oder mehrere Befehlspräfixe an.\nVerwendung: setprefix <p1> <p2> ...\nVerwenden Sie das Symbol \"empty\" für einen Eintrag ohne Präfix.\nBeispiel: setprefix . / #",
		SetPrefixUpdated: "Befehlspräfix aktualisiert: %s",
		SaveFailed:       "Einstellungen konnten nicht gespeichert werden: %s",
		SetSudoUsage:     "Verwendung: setsudo add|remove <Telefon>\nBeispiel: setsudo add 1234567890",
		SudoAdded:        "%s wurde sudo-Zugriff gewährt.",
		SudoRemoved:      "%s wurde der sudo-Zugriff entzogen.",
		SudoNotFound:     "%s ist kein sudo-Benutzer.",
		UnknownAction:    "Unbekannte Aktion. Gültige Werte: add, remove",
		SetModeUsage:     "Verwendung: setmode public|private",
		ModePublicSet:    "Modus auf öffentlich gesetzt. Alle Benutzer können Befehle ausführen.",
		ModePrivateSet:   "Modus auf privat gesetzt. Befehle sind auf sudo-Benutzer beschränkt.",
		LangCurrent:      "Aktuelle Sprache: %s",
		LangSet:          "Sprache gesetzt auf: %s",
		LangUnknown:      "Unbekannter Sprachcode: %s\nVerfügbar: %s",
		LangUsage:        "Verwendung: lang <Code>\nVerfügbar: %s",
		MenuGreeting:     "ʜᴀʟʟᴏ, %s!",
	},
	"ru": {
		Pong:             "Понг",
		PongLatency:      "Понг (%dms)",
		MetaUsage:        "Введите запрос для отправки в Meta AI.\nИспользование: meta <вопрос>",
		GroupOnly:        "Эта команда доступна только в групповых чатах.",
		SudoOnly:         "Эта команда доступна только авторизованным пользователям.",
		SetPrefixUsage:   "Укажите один или несколько префиксов команды.\nИспользование: setprefix <п1> <п2> ...\nИспользуйте символ \"empty\" для записи без префикса.\nПример: setprefix . / #",
		SetPrefixUpdated: "Префикс команды обновлён: %s",
		SaveFailed:       "Не удалось сохранить настройки: %s",
		SetSudoUsage:     "Использование: setsudo add|remove <телефон>\nПример: setsudo add 1234567890",
		SudoAdded:        "Пользователю %s предоставлены права sudo.",
		SudoRemoved:      "Права sudo пользователя %s отозваны.",
		SudoNotFound:     "%s не является пользователем sudo.",
		UnknownAction:    "Неизвестное действие. Допустимые значения: add, remove",
		SetModeUsage:     "Использование: setmode public|private",
		ModePublicSet:    "Режим установлен на публичный. Все пользователи могут выполнять команды.",
		ModePrivateSet:   "Режим установлен на приватный. Команды доступны только пользователям sudo.",
		LangCurrent:      "Текущий язык: %s",
		LangSet:          "Язык установлен: %s",
		LangUnknown:      "Неизвестный код языка: %s\nДоступные: %s",
		LangUsage:        "Использование: lang <код>\nДоступные: %s",
		MenuGreeting:     "Привет, %s!",
	},
	"tr": {
		Pong:             "Pong",
		PongLatency:      "Pong (%dms)",
		MetaUsage:        "Meta AI'ya göndermek için bir sorgu girin.\nKullanım: meta <soru>",
		GroupOnly:        "Bu komut yalnızca grup sohbetleriyle sınırlıdır.",
		SudoOnly:         "Bu komut yalnızca yetkili kullanıcılarla sınırlıdır.",
		SetPrefixUsage:   "Bir veya daha fazla komut öneki belirtin.\nKullanım: setprefix <ö1> <ö2> ...\nÖneksiz giriş için \"empty\" sembolünü kullanın.\nÖrnek: setprefix . / #",
		SetPrefixUpdated: "Komut ön eki güncellendi: %s",
		SaveFailed:       "Ayarlar kaydedilemedi: %s",
		SetSudoUsage:     "Kullanım: setsudo add|remove <telefon>\nÖrnek: setsudo add 1234567890",
		SudoAdded:        "%s kullanıcısına sudo erişimi verildi.",
		SudoRemoved:      "%s kullanıcısının sudo erişimi iptal edildi.",
		SudoNotFound:     "%s bir sudo kullanıcısı değil.",
		UnknownAction:    "Bilinmeyen işlem. Kabul edilen değerler: add, remove",
		SetModeUsage:     "Kullanım: setmode public|private",
		ModePublicSet:    "Mod herkese açık olarak ayarlandı. Tüm kullanıcılar komut çalıştırabilir.",
		ModePrivateSet:   "Mod özel olarak ayarlandı. Komutlar sudo kullanıcılarıyla sınırlı.",
		LangCurrent:      "Mevcut dil: %s",
		LangSet:          "Dil ayarlandı: %s",
		LangUnknown:      "Tanınmayan dil kodu: %s\nMevcut: %s",
		LangUsage:        "Kullanım: lang <kod>\nMevcut: %s",
		MenuGreeting:     "ᴍᴇʀʜᴀʙᴀ, %s!",
	},
	"sw": {
		Pong:             "Pong",
		PongLatency:      "Pong (%dms)",
		MetaUsage:        "Ingiza swali la kutuma kwa Meta AI.\nMatumizi: meta <swali>",
		GroupOnly:        "Amri hii imezuiliwa kwa mazungumzo ya kikundi.",
		SudoOnly:         "Amri hii imezuiliwa kwa watumiaji walioidhinishwa.",
		SetPrefixUsage:   "Taja kiambishi awali kimoja au zaidi cha amri.\nMatumizi: setprefix <k1> <k2> ...\nTumia neno \"empty\" kwa ingizo lisilo na kiambishi awali.\nMfano: setprefix . / #",
		SetPrefixUpdated: "Kiambishi awali cha amri kimesasishwa: %s",
		SaveFailed:       "Imeshindwa kuhifadhi mipangilio: %s",
		SetSudoUsage:     "Matumizi: setsudo add|remove <simu>\nMfano: setsudo add 1234567890",
		SudoAdded:        "%s amepewa mamlaka ya sudo.",
		SudoRemoved:      "Mamlaka ya sudo ya %s imefutwa.",
		SudoNotFound:     "%s si mtumiaji wa sudo.",
		UnknownAction:    "Hatua isiyojulikana. Maadili yanayokubaliwa: add, remove",
		SetModeUsage:     "Matumizi: setmode public|private",
		ModePublicSet:    "Hali imewekwa kuwa ya umma. Watumiaji wote wanaweza kutoa amri.",
		ModePrivateSet:   "Hali imewekwa kuwa ya faragha. Amri zimezuiliwa kwa watumiaji wa sudo.",
		LangCurrent:      "Lugha ya sasa: %s",
		LangSet:          "Lugha imewekwa kuwa: %s",
		LangUnknown:      "Msimbo wa lugha usiojulikana: %s\nZinazopatikana: %s",
		LangUsage:        "Matumizi: lang <msimbo>\nZinazopatikana: %s",
		MenuGreeting:     "ʜᴜᴊᴀᴍʙᴏ, %s!",
	},
}
