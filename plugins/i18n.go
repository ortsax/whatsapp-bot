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
	PongLatency string // fmt: latency in ms (float, e.g. "Pong (1.23ms)")

	// meta
	MetaUsage string

	// dispatch
	GroupOnly string
	SudoOnly  string

	// setprefix
	SetPrefixUsage   string
	SetPrefixUpdated string // fmt: prefix list
	SaveFailed       string // fmt: error string

	// setsudo / delsudo / getsudo
	SetSudoUsage  string
	DelSudoUsage  string
	SudoAdded     string // fmt: identifier
	SudoRemoved   string // fmt: identifier
	SudoNotFound  string // fmt: identifier
	SudoList      string // fmt: newline-separated list
	SudoListEmpty string
	UnknownAction string

	// setmode
	SetModeUsage   string
	ModePublicSet  string
	ModePrivateSet string

	// enablecmd / disablecmd
	EnableCmdUsage string
	DisableCmdUsage string
	CmdEnabled     string // fmt: command name
	CmdDisabledOK  string // fmt: command name
	CmdIsDisabled  string // shown when a disabled command is called
	CmdNotFound    string // fmt: command name

	// ban / delban / getban
	BanUsage    string
	DelBanUsage string
	UserBanned  string // fmt: identifier
	UserUnbanned string // fmt: identifier
	UserNotBanned string // fmt: identifier
	BanList     string // fmt: newline-separated list
	BanListEmpty string

	// disablegc / enablegc
	GCDisabledSet   string
	GCEnabledSet    string
	GCAlreadyDisabled string
	GCAlreadyEnabled  string

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
		Pong:            "Pong",
		PongLatency:     "Pong (%.2fms)",
		MetaUsage:       "Provide a query to send to Meta AI.\nUsage: meta <question>",
		GroupOnly:       "This command is restricted to group chats.",
		SudoOnly:        "This command is restricted to authorised users.",
		SetPrefixUsage:  "Specify one or more command prefixes.\nUsage: setprefix <p1> <p2> ...\nUse the token \"empty\" for a prefix-free entry.\nExample: setprefix . / #",
		SetPrefixUpdated: "Command prefix updated: %s",
		SaveFailed:      "Failed to persist settings: %s",
		SetSudoUsage:    "Reply to a message or pass a phone number.\nUsage: setsudo [phone|@mention]",
		DelSudoUsage:    "Reply to a message or pass a phone number.\nUsage: delsudo [phone|@mention]",
		SudoAdded:       "%s has been granted sudo access.",
		SudoRemoved:     "Sudo access has been revoked from %s.",
		SudoNotFound:    "%s is not a sudo user.",
		SudoList:        "Sudo users:\n%s",
		SudoListEmpty:   "No sudo users have been configured.",
		UnknownAction:   "Unknown action.",
		SetModeUsage:    "Usage: setmode public|private",
		ModePublicSet:   "Mode set to public. All users may issue commands.",
		ModePrivateSet:  "Mode set to private. Commands are restricted to sudo users.",
		EnableCmdUsage:  "Usage: enablecmd <command>",
		DisableCmdUsage: "Usage: disablecmd <command>",
		CmdEnabled:      "Command '%s' has been enabled.",
		CmdDisabledOK:   "Command '%s' has been disabled.",
		CmdIsDisabled:   "This command has been disabled by the owner.",
		CmdNotFound:     "Unknown command: %s",
		GCDisabledSet:     "The bot will no longer respond to group messages.",
		GCEnabledSet:      "The bot will now respond to group messages.",
		GCAlreadyDisabled: "Group chat responses are already disabled.",
		GCAlreadyEnabled:  "Group chat responses are already enabled.",
		BanUsage:        "Reply to a message or pass a phone number.\nUsage: ban [phone|@mention]",
		DelBanUsage:     "Reply to a message or pass a phone number.\nUsage: delban [phone|@mention]",
		UserBanned:      "%s has been banned from using the bot.",
		UserUnbanned:    "%s has been unbanned.",
		UserNotBanned:   "%s is not on the ban list.",
		BanList:         "Banned users:\n%s",
		BanListEmpty:    "No users are currently banned.",
		LangCurrent:     "Current language: %s",
		LangSet:         "Language set to: %s",
		LangUnknown:     "Unrecognised language code: %s\nAvailable: %s",
		LangUsage:       "Usage: lang <code>\nAvailable: %s",
		MenuGreeting:    "ʜᴇʏ, %s!",
	},
	"es": {
		Pong:            "Pong",
		PongLatency:     "Pong (%.2fms)",
		MetaUsage:       "Proporcione una consulta para enviar a Meta AI.\nUso: meta <pregunta>",
		GroupOnly:       "Este comando está restringido a conversaciones de grupo.",
		SudoOnly:        "Este comando está restringido a usuarios autorizados.",
		SetPrefixUsage:  "Especifique uno o más prefijos de comando.\nUso: setprefix <p1> <p2> ...\nUtilice el símbolo \"empty\" para una entrada sin prefijo.\nEjemplo: setprefix . / #",
		SetPrefixUpdated: "Prefijo de comando actualizado: %s",
		SaveFailed:      "Error al guardar la configuración: %s",
		SetSudoUsage:    "Responda a un mensaje o indique un número de teléfono.\nUso: setsudo [teléfono|@mención]",
		DelSudoUsage:    "Responda a un mensaje o indique un número de teléfono.\nUso: delsudo [teléfono|@mención]",
		SudoAdded:       "Se ha concedido acceso sudo a %s.",
		SudoRemoved:     "El acceso sudo de %s ha sido revocado.",
		SudoNotFound:    "%s no es un usuario sudo.",
		SudoList:        "Usuarios sudo:\n%s",
		SudoListEmpty:   "No hay usuarios sudo configurados.",
		UnknownAction:   "Acción desconocida.",
		SetModeUsage:    "Uso: setmode public|private",
		ModePublicSet:   "Modo establecido en público. Todos los usuarios pueden ejecutar comandos.",
		ModePrivateSet:  "Modo establecido en privado. Los comandos están restringidos a usuarios sudo.",
		EnableCmdUsage:  "Uso: enablecmd <comando>",
		DisableCmdUsage: "Uso: disablecmd <comando>",
		CmdEnabled:      "El comando '%s' ha sido habilitado.",
		CmdDisabledOK:   "El comando '%s' ha sido deshabilitado.",
		CmdIsDisabled:   "Este comando ha sido deshabilitado por el propietario.",
		CmdNotFound:     "Comando desconocido: %s",
		GCDisabledSet:     "El bot ya no responderá a mensajes de grupo.",
		GCEnabledSet:      "El bot ahora responderá a mensajes de grupo.",
		GCAlreadyDisabled: "Las respuestas en grupos ya están desactivadas.",
		GCAlreadyEnabled:  "Las respuestas en grupos ya están activadas.",
		BanUsage:        "Responda a un mensaje o indique un número de teléfono.\nUso: ban [teléfono|@mención]",
		DelBanUsage:     "Responda a un mensaje o indique un número de teléfono.\nUso: delban [teléfono|@mención]",
		UserBanned:      "%s ha sido bloqueado del uso del bot.",
		UserUnbanned:    "%s ha sido desbloqueado.",
		UserNotBanned:   "%s no está en la lista de bloqueados.",
		BanList:         "Usuarios bloqueados:\n%s",
		BanListEmpty:    "No hay usuarios bloqueados actualmente.",
		LangCurrent:     "Idioma actual: %s",
		LangSet:         "Idioma establecido en: %s",
		LangUnknown:     "Código de idioma no reconocido: %s\nDisponibles: %s",
		LangUsage:       "Uso: lang <código>\nDisponibles: %s",
		MenuGreeting:    "ʜᴏʟᴀ, %s!",
	},
	"pt": {
		Pong:            "Pong",
		PongLatency:     "Pong (%.2fms)",
		MetaUsage:       "Forneça uma consulta para enviar ao Meta AI.\nUtilização: meta <pergunta>",
		GroupOnly:       "Este comando está restrito a conversas em grupo.",
		SudoOnly:        "Este comando está restrito a utilizadores autorizados.",
		SetPrefixUsage:  "Especifique um ou mais prefixos de comando.\nUtilização: setprefix <p1> <p2> ...\nUtilize o símbolo \"empty\" para uma entrada sem prefixo.\nExemplo: setprefix . / #",
		SetPrefixUpdated: "Prefixo de comando atualizado: %s",
		SaveFailed:      "Falha ao guardar as definições: %s",
		SetSudoUsage:    "Responda a uma mensagem ou indique um número de telefone.\nUtilização: setsudo [telefone|@menção]",
		DelSudoUsage:    "Responda a uma mensagem ou indique um número de telefone.\nUtilização: delsudo [telefone|@menção]",
		SudoAdded:       "%s recebeu acesso sudo.",
		SudoRemoved:     "O acesso sudo de %s foi revogado.",
		SudoNotFound:    "%s não é um utilizador sudo.",
		SudoList:        "Utilizadores sudo:\n%s",
		SudoListEmpty:   "Nenhum utilizador sudo foi configurado.",
		UnknownAction:   "Ação desconhecida.",
		SetModeUsage:    "Utilização: setmode public|private",
		ModePublicSet:   "Modo definido como público. Todos os utilizadores podem executar comandos.",
		ModePrivateSet:  "Modo definido como privado. Os comandos estão restritos a utilizadores sudo.",
		EnableCmdUsage:  "Utilização: enablecmd <comando>",
		DisableCmdUsage: "Utilização: disablecmd <comando>",
		CmdEnabled:      "O comando '%s' foi ativado.",
		CmdDisabledOK:   "O comando '%s' foi desativado.",
		CmdIsDisabled:   "Este comando foi desativado pelo proprietário.",
		CmdNotFound:     "Comando desconhecido: %s",
		GCDisabledSet:     "O bot deixará de responder a mensagens de grupo.",
		GCEnabledSet:      "O bot voltará a responder a mensagens de grupo.",
		GCAlreadyDisabled: "As respostas em grupos já estão desativadas.",
		GCAlreadyEnabled:  "As respostas em grupos já estão ativadas.",
		BanUsage:        "Responda a uma mensagem ou indique um número de telefone.\nUtilização: ban [telefone|@menção]",
		DelBanUsage:     "Responda a uma mensagem ou indique um número de telefone.\nUtilização: delban [telefone|@menção]",
		UserBanned:      "%s foi banido de usar o bot.",
		UserUnbanned:    "%s foi desbanido.",
		UserNotBanned:   "%s não está na lista de banidos.",
		BanList:         "Utilizadores banidos:\n%s",
		BanListEmpty:    "Nenhum utilizador está atualmente banido.",
		LangCurrent:     "Idioma atual: %s",
		LangSet:         "Idioma definido para: %s",
		LangUnknown:     "Código de idioma não reconhecido: %s\nDisponíveis: %s",
		LangUsage:       "Utilização: lang <código>\nDisponíveis: %s",
		MenuGreeting:    "ᴏʟᴀ, %s!",
	},
	"ar": {
		Pong:            "بونغ",
		PongLatency:     "بونغ (%.2fms)",
		MetaUsage:       "أدخل استفساراً لإرساله إلى Meta AI.\nالاستخدام: meta <سؤال>",
		GroupOnly:       "هذا الأمر مقتصر على المجموعات.",
		SudoOnly:        "هذا الأمر مقتصر على المستخدمين المخوَّلين.",
		SetPrefixUsage:  "حدِّد بادئة واحدة أو أكثر.\nالاستخدام: setprefix <ب1> <ب2> ...\nاستخدم الرمز \"empty\" لإدخال بدون بادئة.\nمثال: setprefix . / #",
		SetPrefixUpdated: "تم تحديث بادئة الأمر: %s",
		SaveFailed:      "فشل حفظ الإعدادات: %s",
		SetSudoUsage:    "ردَّ على رسالة أو أدخل رقم هاتف.\nالاستخدام: setsudo [رقم|@ذكر]",
		DelSudoUsage:    "ردَّ على رسالة أو أدخل رقم هاتف.\nالاستخدام: delsudo [رقم|@ذكر]",
		SudoAdded:       "تم منح %s صلاحيات المشرف.",
		SudoRemoved:     "تم سحب صلاحيات المشرف من %s.",
		SudoNotFound:    "%s ليس مستخدماً مشرفاً.",
		SudoList:        "المستخدمون المشرفون:\n%s",
		SudoListEmpty:   "لا يوجد مستخدمون مشرفون حتى الآن.",
		UnknownAction:   "إجراء غير معروف.",
		SetModeUsage:    "الاستخدام: setmode public|private",
		ModePublicSet:   "تم تعيين الوضع على عام. يمكن لجميع المستخدمين تنفيذ الأوامر.",
		ModePrivateSet:  "تم تعيين الوضع على خاص. الأوامر مقتصرة على المشرفين.",
		EnableCmdUsage:  "الاستخدام: enablecmd <أمر>",
		DisableCmdUsage: "الاستخدام: disablecmd <أمر>",
		CmdEnabled:      "تم تفعيل الأمر '%s'.",
		CmdDisabledOK:   "تم تعطيل الأمر '%s'.",
		CmdIsDisabled:   "هذا الأمر معطَّل من قبل المالك.",
		CmdNotFound:     "أمر غير معروف: %s",
		GCDisabledSet:     "لن يستجيب البوت بعد الآن لرسائل المجموعات.",
		GCEnabledSet:      "سيستجيب البوت من الآن لرسائل المجموعات.",
		GCAlreadyDisabled: "ردود المجموعات معطَّلة بالفعل.",
		GCAlreadyEnabled:  "ردود المجموعات مفعَّلة بالفعل.",
		BanUsage:        "ردَّ على رسالة أو أدخل رقم هاتف.\nالاستخدام: ban [رقم|@ذكر]",
		DelBanUsage:     "ردَّ على رسالة أو أدخل رقم هاتف.\nالاستخدام: delban [رقم|@ذكر]",
		UserBanned:      "تم حظر %s من استخدام البوت.",
		UserUnbanned:    "تم رفع الحظر عن %s.",
		UserNotBanned:   "%s ليس في قائمة المحظورين.",
		BanList:         "المستخدمون المحظورون:\n%s",
		BanListEmpty:    "لا يوجد مستخدمون محظورون حاليًا.",
		LangCurrent:     "اللغة الحالية: %s",
		LangSet:         "تم تعيين اللغة إلى: %s",
		LangUnknown:     "رمز اللغة غير معروف: %s\nالمتاحة: %s",
		LangUsage:       "الاستخدام: lang <رمز>\nالمتاحة: %s",
		MenuGreeting:    "أهلاً، %s!",
	},
	"hi": {
		Pong:            "पोंग",
		PongLatency:     "पोंग (%.2fms)",
		MetaUsage:       "Meta AI को भेजने के लिए एक प्रश्न दर्ज करें।\nउपयोग: meta <प्रश्न>",
		GroupOnly:       "यह आदेश केवल समूह वार्तालापों में उपयोग किया जा सकता है।",
		SudoOnly:        "यह आदेश केवल अधिकृत उपयोगकर्ताओं के लिए है।",
		SetPrefixUsage:  "एक या अधिक आदेश उपसर्ग निर्दिष्ट करें।\nउपयोग: setprefix <उ1> <उ2> ...\nबिना उपसर्ग की प्रविष्टि के लिए \"empty\" का उपयोग करें।\nउदाहरण: setprefix . / #",
		SetPrefixUpdated: "आदेश उपसर्ग अद्यतन किया गया: %s",
		SaveFailed:      "सेटिंग्स सहेजने में विफल: %s",
		SetSudoUsage:    "किसी संदेश का उत्तर दें या फ़ोन नंबर दर्ज करें।\nउपयोग: setsudo [फ़ोन|@उल्लेख]",
		DelSudoUsage:    "किसी संदेश का उत्तर दें या फ़ोन नंबर दर्ज करें।\nउपयोग: delsudo [फ़ोन|@उल्लेख]",
		SudoAdded:       "%s को sudo अधिकार प्रदान किए गए।",
		SudoRemoved:     "%s के sudo अधिकार रद्द किए गए।",
		SudoNotFound:    "%s एक sudo उपयोगकर्ता नहीं है।",
		SudoList:        "Sudo उपयोगकर्ता:\n%s",
		SudoListEmpty:   "कोई sudo उपयोगकर्ता कॉन्फ़िगर नहीं किया गया है।",
		UnknownAction:   "अज्ञात क्रिया।",
		SetModeUsage:    "उपयोग: setmode public|private",
		ModePublicSet:   "मोड सार्वजनिक पर सेट किया गया। सभी उपयोगकर्ता आदेश दे सकते हैं।",
		ModePrivateSet:  "मोड निजी पर सेट किया गया। आदेश केवल sudo उपयोगकर्ताओं तक सीमित हैं।",
		EnableCmdUsage:  "उपयोग: enablecmd <आदेश>",
		DisableCmdUsage: "उपयोग: disablecmd <आदेश>",
		CmdEnabled:      "आदेश '%s' सक्षम किया गया।",
		CmdDisabledOK:   "आदेश '%s' अक्षम किया गया।",
		CmdIsDisabled:   "यह आदेश स्वामी द्वारा अक्षम किया गया है।",
		CmdNotFound:     "अज्ञात आदेश: %s",
		GCDisabledSet:     "बोट अब समूह संदेशों का उत्तर नहीं देगा।",
		GCEnabledSet:      "बोट अब समूह संदेशों का उत्तर देगा।",
		GCAlreadyDisabled: "समूह प्रतिक्रियाएँ पहले से ही अक्षम हैं।",
		GCAlreadyEnabled:  "समूह प्रतिक्रियाएँ पहले से ही सक्षम हैं।",
		BanUsage:        "किसी संदेश का उत्तर दें या फ़ोन नंबर दर्ज करें।\nउपयोग: ban [फ़ोन|@उल्लेख]",
		DelBanUsage:     "किसी संदेश का उत्तर दें या फ़ोन नंबर दर्ज करें।\nउपयोग: delban [फ़ोन|@उल्लेख]",
		UserBanned:      "%s को बोट उपयोग से प्रतिबंधित किया गया।",
		UserUnbanned:    "%s का प्रतिबंध हटा दिया गया।",
		UserNotBanned:   "%s प्रतिबंध सूची में नहीं है।",
		BanList:         "प्रतिबंधित उपयोगकर्ता:\n%s",
		BanListEmpty:    "वर्तमान में कोई उपयोगकर्ता प्रतिबंधित नहीं है।",
		LangCurrent:     "वर्तमान भाषा: %s",
		LangSet:         "भाषा निर्धारित की गई: %s",
		LangUnknown:     "अपरिचित भाषा कोड: %s\nउपलब्ध: %s",
		LangUsage:       "उपयोग: lang <कोड>\nउपलब्ध: %s",
		MenuGreeting:    "नमस्ते, %s!",
	},
	"fr": {
		Pong:            "Pong",
		PongLatency:     "Pong (%.2fms)",
		MetaUsage:       "Veuillez saisir une requête à envoyer à Meta AI.\nUtilisation : meta <question>",
		GroupOnly:       "Cette commande est réservée aux conversations de groupe.",
		SudoOnly:        "Cette commande est réservée aux utilisateurs autorisés.",
		SetPrefixUsage:  "Indiquez un ou plusieurs préfixes de commande.\nUtilisation : setprefix <p1> <p2> ...\nUtilisez le symbole \"empty\" pour une entrée sans préfixe.\nExemple : setprefix . / #",
		SetPrefixUpdated: "Préfixe de commande mis à jour : %s",
		SaveFailed:      "Échec de l'enregistrement des paramètres : %s",
		SetSudoUsage:    "Répondez à un message ou indiquez un numéro de téléphone.\nUtilisation : setsudo [téléphone|@mention]",
		DelSudoUsage:    "Répondez à un message ou indiquez un numéro de téléphone.\nUtilisation : delsudo [téléphone|@mention]",
		SudoAdded:       "Les droits sudo ont été accordés à %s.",
		SudoRemoved:     "Les droits sudo de %s ont été révoqués.",
		SudoNotFound:    "%s n'est pas un utilisateur sudo.",
		SudoList:        "Utilisateurs sudo :\n%s",
		SudoListEmpty:   "Aucun utilisateur sudo n'a été configuré.",
		UnknownAction:   "Action inconnue.",
		SetModeUsage:    "Utilisation : setmode public|private",
		ModePublicSet:   "Mode défini sur public. Tous les utilisateurs peuvent exécuter des commandes.",
		ModePrivateSet:  "Mode défini sur privé. Les commandes sont réservées aux utilisateurs sudo.",
		EnableCmdUsage:  "Utilisation : enablecmd <commande>",
		DisableCmdUsage: "Utilisation : disablecmd <commande>",
		CmdEnabled:      "La commande '%s' a été activée.",
		CmdDisabledOK:   "La commande '%s' a été désactivée.",
		CmdIsDisabled:   "Cette commande a été désactivée par le propriétaire.",
		CmdNotFound:     "Commande inconnue : %s",
		GCDisabledSet:     "Le bot ne répondra plus aux messages de groupe.",
		GCEnabledSet:      "Le bot répondra désormais aux messages de groupe.",
		GCAlreadyDisabled: "Les réponses aux groupes sont déjà désactivées.",
		GCAlreadyEnabled:  "Les réponses aux groupes sont déjà activées.",
		BanUsage:        "Répondez à un message ou indiquez un numéro de téléphone.\nUtilisation : ban [téléphone|@mention]",
		DelBanUsage:     "Répondez à un message ou indiquez un numéro de téléphone.\nUtilisation : delban [téléphone|@mention]",
		UserBanned:      "%s a été banni de l'utilisation du bot.",
		UserUnbanned:    "%s a été débanni.",
		UserNotBanned:   "%s ne figure pas dans la liste des bannis.",
		BanList:         "Utilisateurs bannis :\n%s",
		BanListEmpty:    "Aucun utilisateur n'est actuellement banni.",
		LangCurrent:     "Langue actuelle : %s",
		LangSet:         "Langue définie sur : %s",
		LangUnknown:     "Code de langue non reconnu : %s\nDisponibles : %s",
		LangUsage:       "Utilisation : lang <code>\nDisponibles : %s",
		MenuGreeting:    "ʙᴏɴᴊᴏᴜʀ, %s!",
	},
	"de": {
		Pong:            "Pong",
		PongLatency:     "Pong (%.2fms)",
		MetaUsage:       "Geben Sie eine Anfrage für Meta AI ein.\nVerwendung: meta <Frage>",
		GroupOnly:       "Dieser Befehl ist auf Gruppenunterhaltungen beschränkt.",
		SudoOnly:        "Dieser Befehl ist auf autorisierte Benutzer beschränkt.",
		SetPrefixUsage:  "Geben Sie ein oder mehrere Befehlspräfixe an.\nVerwendung: setprefix <p1> <p2> ...\nVerwenden Sie das Symbol \"empty\" für einen Eintrag ohne Präfix.\nBeispiel: setprefix . / #",
		SetPrefixUpdated: "Befehlspräfix aktualisiert: %s",
		SaveFailed:      "Einstellungen konnten nicht gespeichert werden: %s",
		SetSudoUsage:    "Antworten Sie auf eine Nachricht oder geben Sie eine Telefonnummer an.\nVerwendung: setsudo [Telefon|@Erwähnung]",
		DelSudoUsage:    "Antworten Sie auf eine Nachricht oder geben Sie eine Telefonnummer an.\nVerwendung: delsudo [Telefon|@Erwähnung]",
		SudoAdded:       "%s wurde sudo-Zugriff gewährt.",
		SudoRemoved:     "%s wurde der sudo-Zugriff entzogen.",
		SudoNotFound:    "%s ist kein sudo-Benutzer.",
		SudoList:        "Sudo-Benutzer:\n%s",
		SudoListEmpty:   "Es wurden keine sudo-Benutzer konfiguriert.",
		UnknownAction:   "Unbekannte Aktion.",
		SetModeUsage:    "Verwendung: setmode public|private",
		ModePublicSet:   "Modus auf öffentlich gesetzt. Alle Benutzer können Befehle ausführen.",
		ModePrivateSet:  "Modus auf privat gesetzt. Befehle sind auf sudo-Benutzer beschränkt.",
		EnableCmdUsage:  "Verwendung: enablecmd <Befehl>",
		DisableCmdUsage: "Verwendung: disablecmd <Befehl>",
		CmdEnabled:      "Befehl '%s' wurde aktiviert.",
		CmdDisabledOK:   "Befehl '%s' wurde deaktiviert.",
		CmdIsDisabled:   "Dieser Befehl wurde vom Eigentümer deaktiviert.",
		CmdNotFound:     "Unbekannter Befehl: %s",
		GCDisabledSet:     "Der Bot antwortet nicht mehr auf Gruppennachrichten.",
		GCEnabledSet:      "Der Bot antwortet jetzt wieder auf Gruppennachrichten.",
		GCAlreadyDisabled: "Gruppenantworten sind bereits deaktiviert.",
		GCAlreadyEnabled:  "Gruppenantworten sind bereits aktiviert.",
		BanUsage:        "Antworten Sie auf eine Nachricht oder geben Sie eine Telefonnummer an.\nVerwendung: ban [Telefon|@Erwähnung]",
		DelBanUsage:     "Antworten Sie auf eine Nachricht oder geben Sie eine Telefonnummer an.\nVerwendung: delban [Telefon|@Erwähnung]",
		UserBanned:      "%s wurde von der Bot-Nutzung gesperrt.",
		UserUnbanned:    "%s wurde entsperrt.",
		UserNotBanned:   "%s steht nicht auf der Sperrliste.",
		BanList:         "Gesperrte Benutzer:\n%s",
		BanListEmpty:    "Derzeit sind keine Benutzer gesperrt.",
		LangCurrent:     "Aktuelle Sprache: %s",
		LangSet:         "Sprache gesetzt auf: %s",
		LangUnknown:     "Unbekannter Sprachcode: %s\nVerfügbar: %s",
		LangUsage:       "Verwendung: lang <Code>\nVerfügbar: %s",
		MenuGreeting:    "ʜᴀʟʟᴏ, %s!",
	},
	"ru": {
		Pong:            "Понг",
		PongLatency:     "Понг (%.2fms)",
		MetaUsage:       "Введите запрос для отправки в Meta AI.\nИспользование: meta <вопрос>",
		GroupOnly:       "Эта команда доступна только в групповых чатах.",
		SudoOnly:        "Эта команда доступна только авторизованным пользователям.",
		SetPrefixUsage:  "Укажите один или несколько префиксов команды.\nИспользование: setprefix <п1> <п2> ...\nИспользуйте символ \"empty\" для записи без префикса.\nПример: setprefix . / #",
		SetPrefixUpdated: "Префикс команды обновлён: %s",
		SaveFailed:      "Не удалось сохранить настройки: %s",
		SetSudoUsage:    "Ответьте на сообщение или укажите номер телефона.\nИспользование: setsudo [телефон|@упоминание]",
		DelSudoUsage:    "Ответьте на сообщение или укажите номер телефона.\nИспользование: delsudo [телефон|@упоминание]",
		SudoAdded:       "Пользователю %s предоставлены права sudo.",
		SudoRemoved:     "Права sudo пользователя %s отозваны.",
		SudoNotFound:    "%s не является пользователем sudo.",
		SudoList:        "Пользователи sudo:\n%s",
		SudoListEmpty:   "Пользователи sudo не настроены.",
		UnknownAction:   "Неизвестное действие.",
		SetModeUsage:    "Использование: setmode public|private",
		ModePublicSet:   "Режим установлен на публичный. Все пользователи могут выполнять команды.",
		ModePrivateSet:  "Режим установлен на приватный. Команды доступны только пользователям sudo.",
		EnableCmdUsage:  "Использование: enablecmd <команда>",
		DisableCmdUsage: "Использование: disablecmd <команда>",
		CmdEnabled:      "Команда '%s' включена.",
		CmdDisabledOK:   "Команда '%s' отключена.",
		CmdIsDisabled:   "Эта команда отключена владельцем.",
		CmdNotFound:     "Неизвестная команда: %s",
		GCDisabledSet:     "Бот больше не будет отвечать на сообщения в группах.",
		GCEnabledSet:      "Бот снова будет отвечать на сообщения в группах.",
		GCAlreadyDisabled: "Ответы в группах уже отключены.",
		GCAlreadyEnabled:  "Ответы в группах уже включены.",
		BanUsage:        "Ответьте на сообщение или укажите номер телефона.\nИспользование: ban [телефон|@упоминание]",
		DelBanUsage:     "Ответьте на сообщение или укажите номер телефона.\nИспользование: delban [телефон|@упоминание]",
		UserBanned:      "%s заблокирован от использования бота.",
		UserUnbanned:    "Блокировка %s снята.",
		UserNotBanned:   "%s не находится в списке блокировок.",
		BanList:         "Заблокированные пользователи:\n%s",
		BanListEmpty:    "В настоящее время нет заблокированных пользователей.",
		LangCurrent:     "Текущий язык: %s",
		LangSet:         "Язык установлен: %s",
		LangUnknown:     "Неизвестный код языка: %s\nДоступные: %s",
		LangUsage:       "Использование: lang <код>\nДоступные: %s",
		MenuGreeting:    "Привет, %s!",
	},
	"tr": {
		Pong:            "Pong",
		PongLatency:     "Pong (%.2fms)",
		MetaUsage:       "Meta AI'ya göndermek için bir sorgu girin.\nKullanım: meta <soru>",
		GroupOnly:       "Bu komut yalnızca grup sohbetleriyle sınırlıdır.",
		SudoOnly:        "Bu komut yalnızca yetkili kullanıcılarla sınırlıdır.",
		SetPrefixUsage:  "Bir veya daha fazla komut öneki belirtin.\nKullanım: setprefix <ö1> <ö2> ...\nÖneksiz giriş için \"empty\" sembolünü kullanın.\nÖrnek: setprefix . / #",
		SetPrefixUpdated: "Komut ön eki güncellendi: %s",
		SaveFailed:      "Ayarlar kaydedilemedi: %s",
		SetSudoUsage:    "Bir mesajı yanıtlayın veya telefon numarası girin.\nKullanım: setsudo [telefon|@bahsetme]",
		DelSudoUsage:    "Bir mesajı yanıtlayın veya telefon numarası girin.\nKullanım: delsudo [telefon|@bahsetme]",
		SudoAdded:       "%s kullanıcısına sudo erişimi verildi.",
		SudoRemoved:     "%s kullanıcısının sudo erişimi iptal edildi.",
		SudoNotFound:    "%s bir sudo kullanıcısı değil.",
		SudoList:        "Sudo kullanıcıları:\n%s",
		SudoListEmpty:   "Hiçbir sudo kullanıcısı yapılandırılmadı.",
		UnknownAction:   "Bilinmeyen işlem.",
		SetModeUsage:    "Kullanım: setmode public|private",
		ModePublicSet:   "Mod herkese açık olarak ayarlandı. Tüm kullanıcılar komut çalıştırabilir.",
		ModePrivateSet:  "Mod özel olarak ayarlandı. Komutlar sudo kullanıcılarıyla sınırlı.",
		EnableCmdUsage:  "Kullanım: enablecmd <komut>",
		DisableCmdUsage: "Kullanım: disablecmd <komut>",
		CmdEnabled:      "'%s' komutu etkinleştirildi.",
		CmdDisabledOK:   "'%s' komutu devre dışı bırakıldı.",
		CmdIsDisabled:   "Bu komut, sahip tarafından devre dışı bırakıldı.",
		CmdNotFound:     "Bilinmeyen komut: %s",
		GCDisabledSet:     "Bot artık grup mesajlarına yanıt vermeyecek.",
		GCEnabledSet:      "Bot artık grup mesajlarına yanıt verecek.",
		GCAlreadyDisabled: "Grup yanıtları zaten devre dışı.",
		GCAlreadyEnabled:  "Grup yanıtları zaten etkin.",
		BanUsage:        "Bir mesajı yanıtlayın veya telefon numarası girin.\nKullanım: ban [telefon|@bahsetme]",
		DelBanUsage:     "Bir mesajı yanıtlayın veya telefon numarası girin.\nKullanım: delban [telefon|@bahsetme]",
		UserBanned:      "%s botu kullanmaktan yasaklandı.",
		UserUnbanned:    "%s kullanıcısının yasağı kaldırıldı.",
		UserNotBanned:   "%s yasak listesinde değil.",
		BanList:         "Yasaklı kullanıcılar:\n%s",
		BanListEmpty:    "Şu anda yasaklı kullanıcı bulunmamaktadır.",
		LangCurrent:     "Mevcut dil: %s",
		LangSet:         "Dil ayarlandı: %s",
		LangUnknown:     "Tanınmayan dil kodu: %s\nMevcut: %s",
		LangUsage:       "Kullanım: lang <kod>\nMevcut: %s",
		MenuGreeting:    "ᴍᴇʀʜᴀʙᴀ, %s!",
	},
	"sw": {
		Pong:            "Pong",
		PongLatency:     "Pong (%.2fms)",
		MetaUsage:       "Ingiza swali la kutuma kwa Meta AI.\nMatumizi: meta <swali>",
		GroupOnly:       "Amri hii imezuiliwa kwa mazungumzo ya kikundi.",
		SudoOnly:        "Amri hii imezuiliwa kwa watumiaji walioidhinishwa.",
		SetPrefixUsage:  "Taja kiambishi awali kimoja au zaidi cha amri.\nMatumizi: setprefix <k1> <k2> ...\nTumia neno \"empty\" kwa ingizo lisilo na kiambishi awali.\nMfano: setprefix . / #",
		SetPrefixUpdated: "Kiambishi awali cha amri kimesasishwa: %s",
		SaveFailed:      "Imeshindwa kuhifadhi mipangilio: %s",
		SetSudoUsage:    "Jibu ujumbe au weka nambari ya simu.\nMatumizi: setsudo [simu|@kutajwa]",
		DelSudoUsage:    "Jibu ujumbe au weka nambari ya simu.\nMatumizi: delsudo [simu|@kutajwa]",
		SudoAdded:       "%s amepewa mamlaka ya sudo.",
		SudoRemoved:     "Mamlaka ya sudo ya %s imefutwa.",
		SudoNotFound:    "%s si mtumiaji wa sudo.",
		SudoList:        "Watumiaji wa sudo:\n%s",
		SudoListEmpty:   "Hakuna watumiaji wa sudo waliowekwa.",
		UnknownAction:   "Hatua isiyojulikana.",
		SetModeUsage:    "Matumizi: setmode public|private",
		ModePublicSet:   "Hali imewekwa kuwa ya umma. Watumiaji wote wanaweza kutoa amri.",
		ModePrivateSet:  "Hali imewekwa kuwa ya faragha. Amri zimezuiliwa kwa watumiaji wa sudo.",
		EnableCmdUsage:  "Matumizi: enablecmd <amri>",
		DisableCmdUsage: "Matumizi: disablecmd <amri>",
		CmdEnabled:      "Amri '%s' imewezeshwa.",
		CmdDisabledOK:   "Amri '%s' imezimwa.",
		CmdIsDisabled:   "Amri hii imezimwa na mmiliki.",
		CmdNotFound:     "Amri isiyojulikana: %s",
		GCDisabledSet:     "Boti haitajibu tena ujumbe wa vikundi.",
		GCEnabledSet:      "Boti itajibu ujumbe wa vikundi tena.",
		GCAlreadyDisabled: "Majibu ya vikundi yamezimwa tayari.",
		GCAlreadyEnabled:  "Majibu ya vikundi yamewashwa tayari.",
		BanUsage:        "Jibu ujumbe au weka nambari ya simu.\nMatumizi: ban [simu|@kutajwa]",
		DelBanUsage:     "Jibu ujumbe au weka nambari ya simu.\nMatumizi: delban [simu|@kutajwa]",
		UserBanned:      "%s amepigwa marufuku kutumia boti.",
		UserUnbanned:    "%s ameondolewa marufuku.",
		UserNotBanned:   "%s hayupo kwenye orodha ya marufuku.",
		BanList:         "Watumiaji waliozuiliwa:\n%s",
		BanListEmpty:    "Hakuna watumiaji waliozuiliwa kwa sasa.",
		LangCurrent:     "Lugha ya sasa: %s",
		LangSet:         "Lugha imewekwa kuwa: %s",
		LangUnknown:     "Msimbo wa lugha usiojulikana: %s\nZinazopatikana: %s",
		LangUsage:       "Matumizi: lang <msimbo>\nZinazopatikana: %s",
		MenuGreeting:    "ʜᴜᴊᴀᴍʙᴏ, %s!",
	},
}
