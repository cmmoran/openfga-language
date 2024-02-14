// Code generated from /app/OpenFGALexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type OpenFGALexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var OpenFGALexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  ChannelNames           []string
  ModeNames              []string
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func openfgalexerLexerInit() {
  staticData := &OpenFGALexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE", "CONDITION_DEF",
  }
  staticData.LiteralNames = []string{
    "", "':'", "','", "'<'", "'>'", "'['", "", "'('", "')'", "", "", "'#'", 
    "'and'", "'or'", "'but not'", "'from'", "'model'", "'schema'", "", "'type'", 
    "'condition'", "'relations'", "'define'", "'with'", "'=='", "'!='", 
    "'in'", "'<='", "'>='", "'&&'", "'||'", "']'", "'{'", "'}'", "'.'", 
    "'-'", "'!'", "'?'", "'+'", "'*'", "'/'", "'%'", "'true'", "'false'", 
    "'null'",
  }
  staticData.SymbolicNames = []string{
    "", "COLON", "COMMA", "LESS", "GREATER", "LBRACKET", "RBRACKET", "LPAREN", 
    "RPAREN", "WHITESPACE", "IDENTIFIER", "HASH", "AND", "OR", "BUT_NOT", 
    "FROM", "MODEL", "SCHEMA", "SCHEMA_VERSION", "TYPE", "CONDITION", "RELATIONS", 
    "DEFINE", "KEYWORD_WITH", "EQUALS", "NOT_EQUALS", "IN", "LESS_EQUALS", 
    "GREATER_EQUALS", "LOGICAL_AND", "LOGICAL_OR", "RPRACKET", "LBRACE", 
    "RBRACE", "DOT", "MINUS", "EXCLAM", "QUESTIONMARK", "PLUS", "STAR", 
    "SLASH", "PERCENT", "CEL_TRUE", "CEL_FALSE", "NUL", "CEL_COMMENT", "NUM_FLOAT", 
    "NUM_INT", "NUM_UINT", "STRING", "BYTES", "NEWLINE", "CONDITION_PARAM_CONTAINER", 
    "CONDITION_PARAM_TYPE",
  }
  staticData.RuleNames = []string{
    "HASH", "COLON", "COMMA", "AND", "OR", "BUT_NOT", "FROM", "MODEL", "SCHEMA", 
    "SCHEMA_VERSION", "TYPE", "CONDITION", "RELATIONS", "DEFINE", "KEYWORD_WITH", 
    "EQUALS", "NOT_EQUALS", "IN", "LESS", "LESS_EQUALS", "GREATER_EQUALS", 
    "GREATER", "LOGICAL_AND", "LOGICAL_OR", "LBRACKET", "RPRACKET", "LBRACE", 
    "RBRACE", "LPAREN", "RPAREN", "DOT", "MINUS", "EXCLAM", "QUESTIONMARK", 
    "PLUS", "STAR", "SLASH", "PERCENT", "CEL_TRUE", "CEL_FALSE", "NUL", 
    "BACKSLASH", "LETTER", "DIGIT", "EXPONENT", "HEXDIGIT", "RAW", "ESC_SEQ", 
    "ESC_CHAR_SEQ", "ESC_OCT_SEQ", "ESC_BYTE_SEQ", "ESC_UNI_SEQ", "WHITESPACE", 
    "CEL_COMMENT", "NUM_FLOAT", "NUM_INT", "NUM_UINT", "STRING", "BYTES", 
    "IDENTIFIER", "NEWLINE", "CONDITION_DEF_END", "CONDITION_PARAM_CONTAINER", 
    "CONDITION_PARAM_TYPE", "CONDITION_PARAM_TYPE_LESS", "CONDITION_PARAM_TYPE_GREATER", 
    "CONDITION_OPEN", "CONDITION_COLON", "CONDITION_COMMA", "CONDITION_WS", 
    "CONDITION_NAME",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 53, 665, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 
	7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 
	7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 
	14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 
	2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 
	25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 
	7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 
	35, 2, 36, 7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 
	2, 41, 7, 41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 
	46, 7, 46, 2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 
	7, 51, 2, 52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 
	56, 2, 57, 7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 
	2, 62, 7, 62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 
	67, 7, 67, 2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 1, 0, 1, 0, 1, 1, 
	1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 
	1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 
	1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 
	1, 9, 4, 9, 185, 8, 9, 11, 9, 12, 9, 186, 1, 9, 1, 9, 4, 9, 191, 8, 9, 
	11, 9, 12, 9, 192, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 
	11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 
	1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 
	13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 
	1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 
	18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 
	1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 
	27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 
	1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 
	37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 
	1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 
	43, 1, 43, 1, 44, 1, 44, 3, 44, 311, 8, 44, 1, 44, 4, 44, 314, 8, 44, 11, 
	44, 12, 44, 315, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 
	3, 47, 326, 8, 47, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 
	49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 
	1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 
	51, 1, 51, 1, 51, 3, 51, 359, 8, 51, 1, 52, 4, 52, 362, 8, 52, 11, 52, 
	12, 52, 363, 1, 53, 1, 53, 1, 53, 1, 53, 5, 53, 370, 8, 53, 10, 53, 12, 
	53, 373, 9, 53, 1, 53, 1, 53, 1, 54, 4, 54, 378, 8, 54, 11, 54, 12, 54, 
	379, 1, 54, 1, 54, 4, 54, 384, 8, 54, 11, 54, 12, 54, 385, 1, 54, 3, 54, 
	389, 8, 54, 1, 54, 4, 54, 392, 8, 54, 11, 54, 12, 54, 393, 1, 54, 1, 54, 
	1, 54, 1, 54, 4, 54, 400, 8, 54, 11, 54, 12, 54, 401, 1, 54, 3, 54, 405, 
	8, 54, 3, 54, 407, 8, 54, 1, 55, 4, 55, 410, 8, 55, 11, 55, 12, 55, 411, 
	1, 55, 1, 55, 1, 55, 1, 55, 4, 55, 418, 8, 55, 11, 55, 12, 55, 419, 3, 
	55, 422, 8, 55, 1, 56, 4, 56, 425, 8, 56, 11, 56, 12, 56, 426, 1, 56, 1, 
	56, 1, 56, 1, 56, 1, 56, 1, 56, 4, 56, 435, 8, 56, 11, 56, 12, 56, 436, 
	1, 56, 1, 56, 3, 56, 441, 8, 56, 1, 57, 1, 57, 1, 57, 5, 57, 446, 8, 57, 
	10, 57, 12, 57, 449, 9, 57, 1, 57, 1, 57, 1, 57, 1, 57, 5, 57, 455, 8, 
	57, 10, 57, 12, 57, 458, 9, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 
	1, 57, 5, 57, 467, 8, 57, 10, 57, 12, 57, 470, 9, 57, 1, 57, 1, 57, 1, 
	57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 5, 57, 481, 8, 57, 10, 57, 
	12, 57, 484, 9, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 5, 57, 492, 
	8, 57, 10, 57, 12, 57, 495, 9, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 5, 
	57, 502, 8, 57, 10, 57, 12, 57, 505, 9, 57, 1, 57, 1, 57, 1, 57, 1, 57, 
	1, 57, 1, 57, 1, 57, 1, 57, 5, 57, 515, 8, 57, 10, 57, 12, 57, 518, 9, 
	57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 
	5, 57, 530, 8, 57, 10, 57, 12, 57, 533, 9, 57, 1, 57, 1, 57, 1, 57, 1, 
	57, 3, 57, 539, 8, 57, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 3, 59, 546, 8, 
	59, 1, 59, 1, 59, 1, 59, 1, 59, 5, 59, 552, 8, 59, 10, 59, 12, 59, 555, 
	9, 59, 1, 60, 3, 60, 558, 8, 60, 1, 60, 3, 60, 561, 8, 60, 1, 60, 1, 60, 
	3, 60, 565, 8, 60, 1, 60, 3, 60, 568, 8, 60, 1, 60, 3, 60, 571, 8, 60, 
	1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 
	62, 1, 62, 3, 62, 585, 8, 62, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 
	1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 
	63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 
	1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 
	63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 
	1, 63, 3, 63, 636, 8, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 65, 1, 65, 1, 
	65, 1, 65, 1, 66, 1, 66, 1, 66, 1, 66, 1, 67, 1, 67, 1, 67, 1, 67, 1, 68, 
	1, 68, 1, 68, 1, 68, 1, 69, 1, 69, 1, 69, 1, 69, 1, 70, 1, 70, 1, 70, 1, 
	70, 4, 468, 482, 516, 531, 0, 71, 2, 11, 4, 1, 6, 2, 8, 12, 10, 13, 12, 
	14, 14, 15, 16, 16, 18, 17, 20, 18, 22, 19, 24, 20, 26, 21, 28, 22, 30, 
	23, 32, 24, 34, 25, 36, 26, 38, 3, 40, 27, 42, 28, 44, 4, 46, 29, 48, 30, 
	50, 5, 52, 31, 54, 32, 56, 33, 58, 7, 60, 8, 62, 34, 64, 35, 66, 36, 68, 
	37, 70, 38, 72, 39, 74, 40, 76, 41, 78, 42, 80, 43, 82, 44, 84, 0, 86, 
	0, 88, 0, 90, 0, 92, 0, 94, 0, 96, 0, 98, 0, 100, 0, 102, 0, 104, 0, 106, 
	9, 108, 45, 110, 46, 112, 47, 114, 48, 116, 49, 118, 50, 120, 10, 122, 
	51, 124, 0, 126, 52, 128, 53, 130, 0, 132, 0, 134, 0, 136, 0, 138, 0, 140, 
	0, 142, 0, 2, 0, 1, 16, 2, 0, 65, 90, 97, 122, 2, 0, 69, 69, 101, 101, 
	2, 0, 43, 43, 45, 45, 3, 0, 48, 57, 65, 70, 97, 102, 2, 0, 82, 82, 114, 
	114, 10, 0, 34, 34, 39, 39, 63, 63, 92, 92, 96, 98, 102, 102, 110, 110, 
	114, 114, 116, 116, 118, 118, 2, 0, 88, 88, 120, 120, 3, 0, 9, 9, 12, 12, 
	32, 32, 1, 0, 10, 10, 2, 0, 85, 85, 117, 117, 4, 0, 10, 10, 13, 13, 34, 
	34, 92, 92, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 1, 0, 92, 92, 3, 0, 10, 
	10, 13, 13, 34, 34, 3, 0, 10, 10, 13, 13, 39, 39, 2, 0, 66, 66, 98, 98, 
	713, 0, 2, 1, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 6, 1, 0, 0, 0, 0, 8, 1, 0, 
	0, 0, 0, 10, 1, 0, 0, 0, 0, 12, 1, 0, 0, 0, 0, 14, 1, 0, 0, 0, 0, 16, 1, 
	0, 0, 0, 0, 18, 1, 0, 0, 0, 0, 20, 1, 0, 0, 0, 0, 22, 1, 0, 0, 0, 0, 24, 
	1, 0, 0, 0, 0, 26, 1, 0, 0, 0, 0, 28, 1, 0, 0, 0, 0, 30, 1, 0, 0, 0, 0, 
	32, 1, 0, 0, 0, 0, 34, 1, 0, 0, 0, 0, 36, 1, 0, 0, 0, 0, 38, 1, 0, 0, 0, 
	0, 40, 1, 0, 0, 0, 0, 42, 1, 0, 0, 0, 0, 44, 1, 0, 0, 0, 0, 46, 1, 0, 0, 
	0, 0, 48, 1, 0, 0, 0, 0, 50, 1, 0, 0, 0, 0, 52, 1, 0, 0, 0, 0, 54, 1, 0, 
	0, 0, 0, 56, 1, 0, 0, 0, 0, 58, 1, 0, 0, 0, 0, 60, 1, 0, 0, 0, 0, 62, 1, 
	0, 0, 0, 0, 64, 1, 0, 0, 0, 0, 66, 1, 0, 0, 0, 0, 68, 1, 0, 0, 0, 0, 70, 
	1, 0, 0, 0, 0, 72, 1, 0, 0, 0, 0, 74, 1, 0, 0, 0, 0, 76, 1, 0, 0, 0, 0, 
	78, 1, 0, 0, 0, 0, 80, 1, 0, 0, 0, 0, 82, 1, 0, 0, 0, 0, 106, 1, 0, 0, 
	0, 0, 108, 1, 0, 0, 0, 0, 110, 1, 0, 0, 0, 0, 112, 1, 0, 0, 0, 0, 114, 
	1, 0, 0, 0, 0, 116, 1, 0, 0, 0, 0, 118, 1, 0, 0, 0, 0, 120, 1, 0, 0, 0, 
	0, 122, 1, 0, 0, 0, 1, 124, 1, 0, 0, 0, 1, 126, 1, 0, 0, 0, 1, 128, 1, 
	0, 0, 0, 1, 130, 1, 0, 0, 0, 1, 132, 1, 0, 0, 0, 1, 134, 1, 0, 0, 0, 1, 
	136, 1, 0, 0, 0, 1, 138, 1, 0, 0, 0, 1, 140, 1, 0, 0, 0, 1, 142, 1, 0, 
	0, 0, 2, 144, 1, 0, 0, 0, 4, 146, 1, 0, 0, 0, 6, 148, 1, 0, 0, 0, 8, 150, 
	1, 0, 0, 0, 10, 154, 1, 0, 0, 0, 12, 157, 1, 0, 0, 0, 14, 165, 1, 0, 0, 
	0, 16, 170, 1, 0, 0, 0, 18, 176, 1, 0, 0, 0, 20, 184, 1, 0, 0, 0, 22, 194, 
	1, 0, 0, 0, 24, 199, 1, 0, 0, 0, 26, 211, 1, 0, 0, 0, 28, 221, 1, 0, 0, 
	0, 30, 228, 1, 0, 0, 0, 32, 233, 1, 0, 0, 0, 34, 236, 1, 0, 0, 0, 36, 239, 
	1, 0, 0, 0, 38, 242, 1, 0, 0, 0, 40, 244, 1, 0, 0, 0, 42, 247, 1, 0, 0, 
	0, 44, 250, 1, 0, 0, 0, 46, 252, 1, 0, 0, 0, 48, 255, 1, 0, 0, 0, 50, 258, 
	1, 0, 0, 0, 52, 260, 1, 0, 0, 0, 54, 262, 1, 0, 0, 0, 56, 264, 1, 0, 0, 
	0, 58, 266, 1, 0, 0, 0, 60, 268, 1, 0, 0, 0, 62, 270, 1, 0, 0, 0, 64, 272, 
	1, 0, 0, 0, 66, 274, 1, 0, 0, 0, 68, 276, 1, 0, 0, 0, 70, 278, 1, 0, 0, 
	0, 72, 280, 1, 0, 0, 0, 74, 282, 1, 0, 0, 0, 76, 284, 1, 0, 0, 0, 78, 286, 
	1, 0, 0, 0, 80, 291, 1, 0, 0, 0, 82, 297, 1, 0, 0, 0, 84, 302, 1, 0, 0, 
	0, 86, 304, 1, 0, 0, 0, 88, 306, 1, 0, 0, 0, 90, 308, 1, 0, 0, 0, 92, 317, 
	1, 0, 0, 0, 94, 319, 1, 0, 0, 0, 96, 325, 1, 0, 0, 0, 98, 327, 1, 0, 0, 
	0, 100, 330, 1, 0, 0, 0, 102, 335, 1, 0, 0, 0, 104, 358, 1, 0, 0, 0, 106, 
	361, 1, 0, 0, 0, 108, 365, 1, 0, 0, 0, 110, 406, 1, 0, 0, 0, 112, 421, 
	1, 0, 0, 0, 114, 440, 1, 0, 0, 0, 116, 538, 1, 0, 0, 0, 118, 540, 1, 0, 
	0, 0, 120, 545, 1, 0, 0, 0, 122, 557, 1, 0, 0, 0, 124, 572, 1, 0, 0, 0, 
	126, 584, 1, 0, 0, 0, 128, 635, 1, 0, 0, 0, 130, 637, 1, 0, 0, 0, 132, 
	641, 1, 0, 0, 0, 134, 645, 1, 0, 0, 0, 136, 649, 1, 0, 0, 0, 138, 653, 
	1, 0, 0, 0, 140, 657, 1, 0, 0, 0, 142, 661, 1, 0, 0, 0, 144, 145, 5, 35, 
	0, 0, 145, 3, 1, 0, 0, 0, 146, 147, 5, 58, 0, 0, 147, 5, 1, 0, 0, 0, 148, 
	149, 5, 44, 0, 0, 149, 7, 1, 0, 0, 0, 150, 151, 5, 97, 0, 0, 151, 152, 
	5, 110, 0, 0, 152, 153, 5, 100, 0, 0, 153, 9, 1, 0, 0, 0, 154, 155, 5, 
	111, 0, 0, 155, 156, 5, 114, 0, 0, 156, 11, 1, 0, 0, 0, 157, 158, 5, 98, 
	0, 0, 158, 159, 5, 117, 0, 0, 159, 160, 5, 116, 0, 0, 160, 161, 5, 32, 
	0, 0, 161, 162, 5, 110, 0, 0, 162, 163, 5, 111, 0, 0, 163, 164, 5, 116, 
	0, 0, 164, 13, 1, 0, 0, 0, 165, 166, 5, 102, 0, 0, 166, 167, 5, 114, 0, 
	0, 167, 168, 5, 111, 0, 0, 168, 169, 5, 109, 0, 0, 169, 15, 1, 0, 0, 0, 
	170, 171, 5, 109, 0, 0, 171, 172, 5, 111, 0, 0, 172, 173, 5, 100, 0, 0, 
	173, 174, 5, 101, 0, 0, 174, 175, 5, 108, 0, 0, 175, 17, 1, 0, 0, 0, 176, 
	177, 5, 115, 0, 0, 177, 178, 5, 99, 0, 0, 178, 179, 5, 104, 0, 0, 179, 
	180, 5, 101, 0, 0, 180, 181, 5, 109, 0, 0, 181, 182, 5, 97, 0, 0, 182, 
	19, 1, 0, 0, 0, 183, 185, 3, 88, 43, 0, 184, 183, 1, 0, 0, 0, 185, 186, 
	1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188, 1, 0, 
	0, 0, 188, 190, 5, 46, 0, 0, 189, 191, 3, 88, 43, 0, 190, 189, 1, 0, 0, 
	0, 191, 192, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 
	21, 1, 0, 0, 0, 194, 195, 5, 116, 0, 0, 195, 196, 5, 121, 0, 0, 196, 197, 
	5, 112, 0, 0, 197, 198, 5, 101, 0, 0, 198, 23, 1, 0, 0, 0, 199, 200, 5, 
	99, 0, 0, 200, 201, 5, 111, 0, 0, 201, 202, 5, 110, 0, 0, 202, 203, 5, 
	100, 0, 0, 203, 204, 5, 105, 0, 0, 204, 205, 5, 116, 0, 0, 205, 206, 5, 
	105, 0, 0, 206, 207, 5, 111, 0, 0, 207, 208, 5, 110, 0, 0, 208, 209, 1, 
	0, 0, 0, 209, 210, 6, 11, 0, 0, 210, 25, 1, 0, 0, 0, 211, 212, 5, 114, 
	0, 0, 212, 213, 5, 101, 0, 0, 213, 214, 5, 108, 0, 0, 214, 215, 5, 97, 
	0, 0, 215, 216, 5, 116, 0, 0, 216, 217, 5, 105, 0, 0, 217, 218, 5, 111, 
	0, 0, 218, 219, 5, 110, 0, 0, 219, 220, 5, 115, 0, 0, 220, 27, 1, 0, 0, 
	0, 221, 222, 5, 100, 0, 0, 222, 223, 5, 101, 0, 0, 223, 224, 5, 102, 0, 
	0, 224, 225, 5, 105, 0, 0, 225, 226, 5, 110, 0, 0, 226, 227, 5, 101, 0, 
	0, 227, 29, 1, 0, 0, 0, 228, 229, 5, 119, 0, 0, 229, 230, 5, 105, 0, 0, 
	230, 231, 5, 116, 0, 0, 231, 232, 5, 104, 0, 0, 232, 31, 1, 0, 0, 0, 233, 
	234, 5, 61, 0, 0, 234, 235, 5, 61, 0, 0, 235, 33, 1, 0, 0, 0, 236, 237, 
	5, 33, 0, 0, 237, 238, 5, 61, 0, 0, 238, 35, 1, 0, 0, 0, 239, 240, 5, 105, 
	0, 0, 240, 241, 5, 110, 0, 0, 241, 37, 1, 0, 0, 0, 242, 243, 5, 60, 0, 
	0, 243, 39, 1, 0, 0, 0, 244, 245, 5, 60, 0, 0, 245, 246, 5, 61, 0, 0, 246, 
	41, 1, 0, 0, 0, 247, 248, 5, 62, 0, 0, 248, 249, 5, 61, 0, 0, 249, 43, 
	1, 0, 0, 0, 250, 251, 5, 62, 0, 0, 251, 45, 1, 0, 0, 0, 252, 253, 5, 38, 
	0, 0, 253, 254, 5, 38, 0, 0, 254, 47, 1, 0, 0, 0, 255, 256, 5, 124, 0, 
	0, 256, 257, 5, 124, 0, 0, 257, 49, 1, 0, 0, 0, 258, 259, 5, 91, 0, 0, 
	259, 51, 1, 0, 0, 0, 260, 261, 5, 93, 0, 0, 261, 53, 1, 0, 0, 0, 262, 263, 
	5, 123, 0, 0, 263, 55, 1, 0, 0, 0, 264, 265, 5, 125, 0, 0, 265, 57, 1, 
	0, 0, 0, 266, 267, 5, 40, 0, 0, 267, 59, 1, 0, 0, 0, 268, 269, 5, 41, 0, 
	0, 269, 61, 1, 0, 0, 0, 270, 271, 5, 46, 0, 0, 271, 63, 1, 0, 0, 0, 272, 
	273, 5, 45, 0, 0, 273, 65, 1, 0, 0, 0, 274, 275, 5, 33, 0, 0, 275, 67, 
	1, 0, 0, 0, 276, 277, 5, 63, 0, 0, 277, 69, 1, 0, 0, 0, 278, 279, 5, 43, 
	0, 0, 279, 71, 1, 0, 0, 0, 280, 281, 5, 42, 0, 0, 281, 73, 1, 0, 0, 0, 
	282, 283, 5, 47, 0, 0, 283, 75, 1, 0, 0, 0, 284, 285, 5, 37, 0, 0, 285, 
	77, 1, 0, 0, 0, 286, 287, 5, 116, 0, 0, 287, 288, 5, 114, 0, 0, 288, 289, 
	5, 117, 0, 0, 289, 290, 5, 101, 0, 0, 290, 79, 1, 0, 0, 0, 291, 292, 5, 
	102, 0, 0, 292, 293, 5, 97, 0, 0, 293, 294, 5, 108, 0, 0, 294, 295, 5, 
	115, 0, 0, 295, 296, 5, 101, 0, 0, 296, 81, 1, 0, 0, 0, 297, 298, 5, 110, 
	0, 0, 298, 299, 5, 117, 0, 0, 299, 300, 5, 108, 0, 0, 300, 301, 5, 108, 
	0, 0, 301, 83, 1, 0, 0, 0, 302, 303, 5, 92, 0, 0, 303, 85, 1, 0, 0, 0, 
	304, 305, 7, 0, 0, 0, 305, 87, 1, 0, 0, 0, 306, 307, 2, 48, 57, 0, 307, 
	89, 1, 0, 0, 0, 308, 310, 7, 1, 0, 0, 309, 311, 7, 2, 0, 0, 310, 309, 1, 
	0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 313, 1, 0, 0, 0, 312, 314, 3, 88, 43, 
	0, 313, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315, 
	316, 1, 0, 0, 0, 316, 91, 1, 0, 0, 0, 317, 318, 7, 3, 0, 0, 318, 93, 1, 
	0, 0, 0, 319, 320, 7, 4, 0, 0, 320, 95, 1, 0, 0, 0, 321, 326, 3, 98, 48, 
	0, 322, 326, 3, 102, 50, 0, 323, 326, 3, 104, 51, 0, 324, 326, 3, 100, 
	49, 0, 325, 321, 1, 0, 0, 0, 325, 322, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 
	325, 324, 1, 0, 0, 0, 326, 97, 1, 0, 0, 0, 327, 328, 3, 84, 41, 0, 328, 
	329, 7, 5, 0, 0, 329, 99, 1, 0, 0, 0, 330, 331, 3, 84, 41, 0, 331, 332, 
	2, 48, 51, 0, 332, 333, 2, 48, 55, 0, 333, 334, 2, 48, 55, 0, 334, 101, 
	1, 0, 0, 0, 335, 336, 3, 84, 41, 0, 336, 337, 7, 6, 0, 0, 337, 338, 3, 
	92, 45, 0, 338, 339, 3, 92, 45, 0, 339, 103, 1, 0, 0, 0, 340, 341, 3, 84, 
	41, 0, 341, 342, 5, 117, 0, 0, 342, 343, 3, 92, 45, 0, 343, 344, 3, 92, 
	45, 0, 344, 345, 3, 92, 45, 0, 345, 346, 3, 92, 45, 0, 346, 359, 1, 0, 
	0, 0, 347, 348, 3, 84, 41, 0, 348, 349, 5, 85, 0, 0, 349, 350, 3, 92, 45, 
	0, 350, 351, 3, 92, 45, 0, 351, 352, 3, 92, 45, 0, 352, 353, 3, 92, 45, 
	0, 353, 354, 3, 92, 45, 0, 354, 355, 3, 92, 45, 0, 355, 356, 3, 92, 45, 
	0, 356, 357, 3, 92, 45, 0, 357, 359, 1, 0, 0, 0, 358, 340, 1, 0, 0, 0, 
	358, 347, 1, 0, 0, 0, 359, 105, 1, 0, 0, 0, 360, 362, 7, 7, 0, 0, 361, 
	360, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0, 363, 364, 
	1, 0, 0, 0, 364, 107, 1, 0, 0, 0, 365, 366, 5, 47, 0, 0, 366, 367, 5, 47, 
	0, 0, 367, 371, 1, 0, 0, 0, 368, 370, 8, 8, 0, 0, 369, 368, 1, 0, 0, 0, 
	370, 373, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 
	374, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 374, 375, 6, 53, 1, 0, 375, 109, 
	1, 0, 0, 0, 376, 378, 3, 88, 43, 0, 377, 376, 1, 0, 0, 0, 378, 379, 1, 
	0, 0, 0, 379, 377, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380, 381, 1, 0, 0, 
	0, 381, 383, 5, 46, 0, 0, 382, 384, 3, 88, 43, 0, 383, 382, 1, 0, 0, 0, 
	384, 385, 1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 
	388, 1, 0, 0, 0, 387, 389, 3, 90, 44, 0, 388, 387, 1, 0, 0, 0, 388, 389, 
	1, 0, 0, 0, 389, 407, 1, 0, 0, 0, 390, 392, 3, 88, 43, 0, 391, 390, 1, 
	0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 391, 1, 0, 0, 0, 393, 394, 1, 0, 0, 
	0, 394, 395, 1, 0, 0, 0, 395, 396, 3, 90, 44, 0, 396, 407, 1, 0, 0, 0, 
	397, 399, 5, 46, 0, 0, 398, 400, 3, 88, 43, 0, 399, 398, 1, 0, 0, 0, 400, 
	401, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402, 404, 
	1, 0, 0, 0, 403, 405, 3, 90, 44, 0, 404, 403, 1, 0, 0, 0, 404, 405, 1, 
	0, 0, 0, 405, 407, 1, 0, 0, 0, 406, 377, 1, 0, 0, 0, 406, 391, 1, 0, 0, 
	0, 406, 397, 1, 0, 0, 0, 407, 111, 1, 0, 0, 0, 408, 410, 3, 88, 43, 0, 
	409, 408, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 409, 1, 0, 0, 0, 411, 
	412, 1, 0, 0, 0, 412, 422, 1, 0, 0, 0, 413, 414, 5, 48, 0, 0, 414, 415, 
	5, 120, 0, 0, 415, 417, 1, 0, 0, 0, 416, 418, 3, 92, 45, 0, 417, 416, 1, 
	0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 417, 1, 0, 0, 0, 419, 420, 1, 0, 0, 
	0, 420, 422, 1, 0, 0, 0, 421, 409, 1, 0, 0, 0, 421, 413, 1, 0, 0, 0, 422, 
	113, 1, 0, 0, 0, 423, 425, 3, 88, 43, 0, 424, 423, 1, 0, 0, 0, 425, 426, 
	1, 0, 0, 0, 426, 424, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 428, 1, 0, 
	0, 0, 428, 429, 7, 9, 0, 0, 429, 441, 1, 0, 0, 0, 430, 431, 5, 48, 0, 0, 
	431, 432, 5, 120, 0, 0, 432, 434, 1, 0, 0, 0, 433, 435, 3, 92, 45, 0, 434, 
	433, 1, 0, 0, 0, 435, 436, 1, 0, 0, 0, 436, 434, 1, 0, 0, 0, 436, 437, 
	1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 439, 7, 9, 0, 0, 439, 441, 1, 0, 
	0, 0, 440, 424, 1, 0, 0, 0, 440, 430, 1, 0, 0, 0, 441, 115, 1, 0, 0, 0, 
	442, 447, 5, 34, 0, 0, 443, 446, 3, 96, 47, 0, 444, 446, 8, 10, 0, 0, 445, 
	443, 1, 0, 0, 0, 445, 444, 1, 0, 0, 0, 446, 449, 1, 0, 0, 0, 447, 445, 
	1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 450, 1, 0, 0, 0, 449, 447, 1, 0, 
	0, 0, 450, 539, 5, 34, 0, 0, 451, 456, 5, 39, 0, 0, 452, 455, 3, 96, 47, 
	0, 453, 455, 8, 11, 0, 0, 454, 452, 1, 0, 0, 0, 454, 453, 1, 0, 0, 0, 455, 
	458, 1, 0, 0, 0, 456, 454, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 459, 
	1, 0, 0, 0, 458, 456, 1, 0, 0, 0, 459, 539, 5, 39, 0, 0, 460, 461, 5, 34, 
	0, 0, 461, 462, 5, 34, 0, 0, 462, 463, 5, 34, 0, 0, 463, 468, 1, 0, 0, 
	0, 464, 467, 3, 96, 47, 0, 465, 467, 8, 12, 0, 0, 466, 464, 1, 0, 0, 0, 
	466, 465, 1, 0, 0, 0, 467, 470, 1, 0, 0, 0, 468, 469, 1, 0, 0, 0, 468, 
	466, 1, 0, 0, 0, 469, 471, 1, 0, 0, 0, 470, 468, 1, 0, 0, 0, 471, 472, 
	5, 34, 0, 0, 472, 473, 5, 34, 0, 0, 473, 539, 5, 34, 0, 0, 474, 475, 5, 
	39, 0, 0, 475, 476, 5, 39, 0, 0, 476, 477, 5, 39, 0, 0, 477, 482, 1, 0, 
	0, 0, 478, 481, 3, 96, 47, 0, 479, 481, 8, 12, 0, 0, 480, 478, 1, 0, 0, 
	0, 480, 479, 1, 0, 0, 0, 481, 484, 1, 0, 0, 0, 482, 483, 1, 0, 0, 0, 482, 
	480, 1, 0, 0, 0, 483, 485, 1, 0, 0, 0, 484, 482, 1, 0, 0, 0, 485, 486, 
	5, 39, 0, 0, 486, 487, 5, 39, 0, 0, 487, 539, 5, 39, 0, 0, 488, 489, 3, 
	94, 46, 0, 489, 493, 5, 34, 0, 0, 490, 492, 8, 13, 0, 0, 491, 490, 1, 0, 
	0, 0, 492, 495, 1, 0, 0, 0, 493, 491, 1, 0, 0, 0, 493, 494, 1, 0, 0, 0, 
	494, 496, 1, 0, 0, 0, 495, 493, 1, 0, 0, 0, 496, 497, 5, 34, 0, 0, 497, 
	539, 1, 0, 0, 0, 498, 499, 3, 94, 46, 0, 499, 503, 5, 39, 0, 0, 500, 502, 
	8, 14, 0, 0, 501, 500, 1, 0, 0, 0, 502, 505, 1, 0, 0, 0, 503, 501, 1, 0, 
	0, 0, 503, 504, 1, 0, 0, 0, 504, 506, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 
	506, 507, 5, 39, 0, 0, 507, 539, 1, 0, 0, 0, 508, 509, 3, 94, 46, 0, 509, 
	510, 5, 34, 0, 0, 510, 511, 5, 34, 0, 0, 511, 512, 5, 34, 0, 0, 512, 516, 
	1, 0, 0, 0, 513, 515, 9, 0, 0, 0, 514, 513, 1, 0, 0, 0, 515, 518, 1, 0, 
	0, 0, 516, 517, 1, 0, 0, 0, 516, 514, 1, 0, 0, 0, 517, 519, 1, 0, 0, 0, 
	518, 516, 1, 0, 0, 0, 519, 520, 5, 34, 0, 0, 520, 521, 5, 34, 0, 0, 521, 
	522, 5, 34, 0, 0, 522, 539, 1, 0, 0, 0, 523, 524, 3, 94, 46, 0, 524, 525, 
	5, 39, 0, 0, 525, 526, 5, 39, 0, 0, 526, 527, 5, 39, 0, 0, 527, 531, 1, 
	0, 0, 0, 528, 530, 9, 0, 0, 0, 529, 528, 1, 0, 0, 0, 530, 533, 1, 0, 0, 
	0, 531, 532, 1, 0, 0, 0, 531, 529, 1, 0, 0, 0, 532, 534, 1, 0, 0, 0, 533, 
	531, 1, 0, 0, 0, 534, 535, 5, 39, 0, 0, 535, 536, 5, 39, 0, 0, 536, 537, 
	5, 39, 0, 0, 537, 539, 1, 0, 0, 0, 538, 442, 1, 0, 0, 0, 538, 451, 1, 0, 
	0, 0, 538, 460, 1, 0, 0, 0, 538, 474, 1, 0, 0, 0, 538, 488, 1, 0, 0, 0, 
	538, 498, 1, 0, 0, 0, 538, 508, 1, 0, 0, 0, 538, 523, 1, 0, 0, 0, 539, 
	117, 1, 0, 0, 0, 540, 541, 7, 15, 0, 0, 541, 542, 3, 116, 57, 0, 542, 119, 
	1, 0, 0, 0, 543, 546, 3, 86, 42, 0, 544, 546, 5, 95, 0, 0, 545, 543, 1, 
	0, 0, 0, 545, 544, 1, 0, 0, 0, 546, 553, 1, 0, 0, 0, 547, 552, 3, 86, 42, 
	0, 548, 552, 3, 88, 43, 0, 549, 552, 5, 95, 0, 0, 550, 552, 3, 64, 31, 
	0, 551, 547, 1, 0, 0, 0, 551, 548, 1, 0, 0, 0, 551, 549, 1, 0, 0, 0, 551, 
	550, 1, 0, 0, 0, 552, 555, 1, 0, 0, 0, 553, 551, 1, 0, 0, 0, 553, 554, 
	1, 0, 0, 0, 554, 121, 1, 0, 0, 0, 555, 553, 1, 0, 0, 0, 556, 558, 3, 106, 
	52, 0, 557, 556, 1, 0, 0, 0, 557, 558, 1, 0, 0, 0, 558, 564, 1, 0, 0, 0, 
	559, 561, 5, 13, 0, 0, 560, 559, 1, 0, 0, 0, 560, 561, 1, 0, 0, 0, 561, 
	562, 1, 0, 0, 0, 562, 565, 5, 10, 0, 0, 563, 565, 2, 12, 13, 0, 564, 560, 
	1, 0, 0, 0, 564, 563, 1, 0, 0, 0, 565, 567, 1, 0, 0, 0, 566, 568, 3, 106, 
	52, 0, 567, 566, 1, 0, 0, 0, 567, 568, 1, 0, 0, 0, 568, 570, 1, 0, 0, 0, 
	569, 571, 3, 122, 60, 0, 570, 569, 1, 0, 0, 0, 570, 571, 1, 0, 0, 0, 571, 
	123, 1, 0, 0, 0, 572, 573, 3, 60, 29, 0, 573, 574, 1, 0, 0, 0, 574, 575, 
	6, 61, 2, 0, 575, 576, 6, 61, 3, 0, 576, 125, 1, 0, 0, 0, 577, 578, 5, 
	109, 0, 0, 578, 579, 5, 97, 0, 0, 579, 585, 5, 112, 0, 0, 580, 581, 5, 
	108, 0, 0, 581, 582, 5, 105, 0, 0, 582, 583, 5, 115, 0, 0, 583, 585, 5, 
	116, 0, 0, 584, 577, 1, 0, 0, 0, 584, 580, 1, 0, 0, 0, 585, 127, 1, 0, 
	0, 0, 586, 587, 5, 98, 0, 0, 587, 588, 5, 111, 0, 0, 588, 589, 5, 111, 
	0, 0, 589, 636, 5, 108, 0, 0, 590, 591, 5, 115, 0, 0, 591, 592, 5, 116, 
	0, 0, 592, 593, 5, 114, 0, 0, 593, 594, 5, 105, 0, 0, 594, 595, 5, 110, 
	0, 0, 595, 636, 5, 103, 0, 0, 596, 597, 5, 105, 0, 0, 597, 598, 5, 110, 
	0, 0, 598, 636, 5, 116, 0, 0, 599, 600, 5, 117, 0, 0, 600, 601, 5, 105, 
	0, 0, 601, 602, 5, 110, 0, 0, 602, 636, 5, 116, 0, 0, 603, 604, 5, 100, 
	0, 0, 604, 605, 5, 111, 0, 0, 605, 606, 5, 117, 0, 0, 606, 607, 5, 98, 
	0, 0, 607, 608, 5, 108, 0, 0, 608, 636, 5, 101, 0, 0, 609, 610, 5, 100, 
	0, 0, 610, 611, 5, 117, 0, 0, 611, 612, 5, 114, 0, 0, 612, 613, 5, 97, 
	0, 0, 613, 614, 5, 116, 0, 0, 614, 615, 5, 105, 0, 0, 615, 616, 5, 111, 
	0, 0, 616, 636, 5, 110, 0, 0, 617, 618, 5, 116, 0, 0, 618, 619, 5, 105, 
	0, 0, 619, 620, 5, 109, 0, 0, 620, 621, 5, 101, 0, 0, 621, 622, 5, 115, 
	0, 0, 622, 623, 5, 116, 0, 0, 623, 624, 5, 97, 0, 0, 624, 625, 5, 109, 
	0, 0, 625, 636, 5, 112, 0, 0, 626, 627, 5, 105, 0, 0, 627, 628, 5, 112, 
	0, 0, 628, 629, 5, 97, 0, 0, 629, 630, 5, 100, 0, 0, 630, 631, 5, 100, 
	0, 0, 631, 632, 5, 114, 0, 0, 632, 633, 5, 101, 0, 0, 633, 634, 5, 115, 
	0, 0, 634, 636, 5, 115, 0, 0, 635, 586, 1, 0, 0, 0, 635, 590, 1, 0, 0, 
	0, 635, 596, 1, 0, 0, 0, 635, 599, 1, 0, 0, 0, 635, 603, 1, 0, 0, 0, 635, 
	609, 1, 0, 0, 0, 635, 617, 1, 0, 0, 0, 635, 626, 1, 0, 0, 0, 636, 129, 
	1, 0, 0, 0, 637, 638, 3, 38, 18, 0, 638, 639, 1, 0, 0, 0, 639, 640, 6, 
	64, 4, 0, 640, 131, 1, 0, 0, 0, 641, 642, 3, 44, 21, 0, 642, 643, 1, 0, 
	0, 0, 643, 644, 6, 65, 5, 0, 644, 133, 1, 0, 0, 0, 645, 646, 3, 58, 28, 
	0, 646, 647, 1, 0, 0, 0, 647, 648, 6, 66, 6, 0, 648, 135, 1, 0, 0, 0, 649, 
	650, 3, 4, 1, 0, 650, 651, 1, 0, 0, 0, 651, 652, 6, 67, 7, 0, 652, 137, 
	1, 0, 0, 0, 653, 654, 3, 6, 2, 0, 654, 655, 1, 0, 0, 0, 655, 656, 6, 68, 
	8, 0, 656, 139, 1, 0, 0, 0, 657, 658, 3, 106, 52, 0, 658, 659, 1, 0, 0, 
	0, 659, 660, 6, 69, 9, 0, 660, 141, 1, 0, 0, 0, 661, 662, 3, 120, 59, 0, 
	662, 663, 1, 0, 0, 0, 663, 664, 6, 70, 10, 0, 664, 143, 1, 0, 0, 0, 46, 
	0, 1, 186, 192, 310, 315, 325, 358, 363, 371, 379, 385, 388, 393, 401, 
	404, 406, 411, 419, 421, 426, 436, 440, 445, 447, 454, 456, 466, 468, 480, 
	482, 493, 503, 516, 531, 538, 545, 551, 553, 557, 560, 564, 567, 570, 584, 
	635, 11, 5, 1, 0, 0, 1, 0, 7, 8, 0, 4, 0, 0, 7, 3, 0, 7, 4, 0, 7, 7, 0, 
	7, 1, 0, 7, 2, 0, 7, 9, 0, 7, 10, 0,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// OpenFGALexerInit initializes any static state used to implement OpenFGALexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewOpenFGALexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func OpenFGALexerInit() {
  staticData := &OpenFGALexerLexerStaticData
  staticData.once.Do(openfgalexerLexerInit)
}

// NewOpenFGALexer produces a new lexer instance for the optional input antlr.CharStream.
func NewOpenFGALexer(input antlr.CharStream) *OpenFGALexer {
  OpenFGALexerInit()
	l := new(OpenFGALexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &OpenFGALexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "OpenFGALexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// OpenFGALexer tokens.
const (
	OpenFGALexerCOLON = 1
	OpenFGALexerCOMMA = 2
	OpenFGALexerLESS = 3
	OpenFGALexerGREATER = 4
	OpenFGALexerLBRACKET = 5
	OpenFGALexerRBRACKET = 6
	OpenFGALexerLPAREN = 7
	OpenFGALexerRPAREN = 8
	OpenFGALexerWHITESPACE = 9
	OpenFGALexerIDENTIFIER = 10
	OpenFGALexerHASH = 11
	OpenFGALexerAND = 12
	OpenFGALexerOR = 13
	OpenFGALexerBUT_NOT = 14
	OpenFGALexerFROM = 15
	OpenFGALexerMODEL = 16
	OpenFGALexerSCHEMA = 17
	OpenFGALexerSCHEMA_VERSION = 18
	OpenFGALexerTYPE = 19
	OpenFGALexerCONDITION = 20
	OpenFGALexerRELATIONS = 21
	OpenFGALexerDEFINE = 22
	OpenFGALexerKEYWORD_WITH = 23
	OpenFGALexerEQUALS = 24
	OpenFGALexerNOT_EQUALS = 25
	OpenFGALexerIN = 26
	OpenFGALexerLESS_EQUALS = 27
	OpenFGALexerGREATER_EQUALS = 28
	OpenFGALexerLOGICAL_AND = 29
	OpenFGALexerLOGICAL_OR = 30
	OpenFGALexerRPRACKET = 31
	OpenFGALexerLBRACE = 32
	OpenFGALexerRBRACE = 33
	OpenFGALexerDOT = 34
	OpenFGALexerMINUS = 35
	OpenFGALexerEXCLAM = 36
	OpenFGALexerQUESTIONMARK = 37
	OpenFGALexerPLUS = 38
	OpenFGALexerSTAR = 39
	OpenFGALexerSLASH = 40
	OpenFGALexerPERCENT = 41
	OpenFGALexerCEL_TRUE = 42
	OpenFGALexerCEL_FALSE = 43
	OpenFGALexerNUL = 44
	OpenFGALexerCEL_COMMENT = 45
	OpenFGALexerNUM_FLOAT = 46
	OpenFGALexerNUM_INT = 47
	OpenFGALexerNUM_UINT = 48
	OpenFGALexerSTRING = 49
	OpenFGALexerBYTES = 50
	OpenFGALexerNEWLINE = 51
	OpenFGALexerCONDITION_PARAM_CONTAINER = 52
	OpenFGALexerCONDITION_PARAM_TYPE = 53
)

// OpenFGALexerCONDITION_DEF is the OpenFGALexer mode.
const OpenFGALexerCONDITION_DEF = 1

