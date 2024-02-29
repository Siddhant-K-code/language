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
    "'and'", "'or'", "'but not'", "'from'", "'model'", "'schema'", "'1.1'", 
    "'type'", "'condition'", "'relations'", "'relation'", "'define'", "'with'", 
    "'=='", "'!='", "'in'", "'<='", "'>='", "'&&'", "'||'", "']'", "'{'", 
    "'}'", "'.'", "'-'", "'!'", "'?'", "'+'", "'*'", "'/'", "'%'", "'true'", 
    "'false'", "'null'",
  }
  staticData.SymbolicNames = []string{
    "", "COLON", "COMMA", "LESS", "GREATER", "LBRACKET", "RBRACKET", "LPAREN", 
    "RPAREN", "WHITESPACE", "IDENTIFIER", "HASH", "AND", "OR", "BUT_NOT", 
    "FROM", "MODEL", "SCHEMA", "SCHEMA_VERSION", "TYPE", "CONDITION", "RELATIONS", 
    "RELATION", "DEFINE", "KEYWORD_WITH", "EQUALS", "NOT_EQUALS", "IN", 
    "LESS_EQUALS", "GREATER_EQUALS", "LOGICAL_AND", "LOGICAL_OR", "RPRACKET", 
    "LBRACE", "RBRACE", "DOT", "MINUS", "EXCLAM", "QUESTIONMARK", "PLUS", 
    "STAR", "SLASH", "PERCENT", "CEL_TRUE", "CEL_FALSE", "NUL", "CEL_COMMENT", 
    "NUM_FLOAT", "NUM_INT", "NUM_UINT", "STRING", "BYTES", "NEWLINE", "CONDITION_PARAM_CONTAINER", 
    "CONDITION_PARAM_TYPE",
  }
  staticData.RuleNames = []string{
    "HASH", "COLON", "COMMA", "AND", "OR", "BUT_NOT", "FROM", "MODEL", "SCHEMA", 
    "SCHEMA_VERSION", "TYPE", "CONDITION", "RELATIONS", "RELATION", "DEFINE", 
    "KEYWORD_WITH", "EQUALS", "NOT_EQUALS", "IN", "LESS", "LESS_EQUALS", 
    "GREATER_EQUALS", "GREATER", "LOGICAL_AND", "LOGICAL_OR", "LBRACKET", 
    "RPRACKET", "LBRACE", "RBRACE", "LPAREN", "RPAREN", "DOT", "MINUS", 
    "EXCLAM", "QUESTIONMARK", "PLUS", "STAR", "SLASH", "PERCENT", "CEL_TRUE", 
    "CEL_FALSE", "NUL", "BACKSLASH", "LETTER", "DIGIT", "EXPONENT", "HEXDIGIT", 
    "RAW", "ESC_SEQ", "ESC_CHAR_SEQ", "ESC_OCT_SEQ", "ESC_BYTE_SEQ", "ESC_UNI_SEQ", 
    "WHITESPACE", "CEL_COMMENT", "NUM_FLOAT", "NUM_INT", "NUM_UINT", "STRING", 
    "BYTES", "IDENTIFIER", "NEWLINE", "CONDITION_DEF_END", "CONDITION_PARAM_CONTAINER", 
    "CONDITION_PARAM_TYPE", "CONDITION_PARAM_TYPE_LESS", "CONDITION_PARAM_TYPE_GREATER", 
    "CONDITION_OPEN", "CONDITION_COLON", "CONDITION_COMMA", "CONDITION_WS", 
    "CONDITION_NAME",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 54, 669, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 
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
	67, 7, 67, 2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 1, 0, 
	1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 
	1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 
	1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 
	1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 
	1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 
	11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 
	1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 
	14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 
	1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 
	19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 
	1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 
	27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 
	1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 
	38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 
	1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 
	43, 1, 44, 1, 44, 1, 45, 1, 45, 3, 45, 315, 8, 45, 1, 45, 4, 45, 318, 8, 
	45, 11, 45, 12, 45, 319, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 
	1, 48, 3, 48, 330, 8, 48, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 
	50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 
	1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 
	52, 1, 52, 1, 52, 1, 52, 3, 52, 363, 8, 52, 1, 53, 4, 53, 366, 8, 53, 11, 
	53, 12, 53, 367, 1, 54, 1, 54, 1, 54, 1, 54, 5, 54, 374, 8, 54, 10, 54, 
	12, 54, 377, 9, 54, 1, 54, 1, 54, 1, 55, 4, 55, 382, 8, 55, 11, 55, 12, 
	55, 383, 1, 55, 1, 55, 4, 55, 388, 8, 55, 11, 55, 12, 55, 389, 1, 55, 3, 
	55, 393, 8, 55, 1, 55, 4, 55, 396, 8, 55, 11, 55, 12, 55, 397, 1, 55, 1, 
	55, 1, 55, 1, 55, 4, 55, 404, 8, 55, 11, 55, 12, 55, 405, 1, 55, 3, 55, 
	409, 8, 55, 3, 55, 411, 8, 55, 1, 56, 4, 56, 414, 8, 56, 11, 56, 12, 56, 
	415, 1, 56, 1, 56, 1, 56, 1, 56, 4, 56, 422, 8, 56, 11, 56, 12, 56, 423, 
	3, 56, 426, 8, 56, 1, 57, 4, 57, 429, 8, 57, 11, 57, 12, 57, 430, 1, 57, 
	1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 4, 57, 439, 8, 57, 11, 57, 12, 57, 440, 
	1, 57, 1, 57, 3, 57, 445, 8, 57, 1, 58, 1, 58, 1, 58, 5, 58, 450, 8, 58, 
	10, 58, 12, 58, 453, 9, 58, 1, 58, 1, 58, 1, 58, 1, 58, 5, 58, 459, 8, 
	58, 10, 58, 12, 58, 462, 9, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 
	1, 58, 5, 58, 471, 8, 58, 10, 58, 12, 58, 474, 9, 58, 1, 58, 1, 58, 1, 
	58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 5, 58, 485, 8, 58, 10, 58, 
	12, 58, 488, 9, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 5, 58, 496, 
	8, 58, 10, 58, 12, 58, 499, 9, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 5, 
	58, 506, 8, 58, 10, 58, 12, 58, 509, 9, 58, 1, 58, 1, 58, 1, 58, 1, 58, 
	1, 58, 1, 58, 1, 58, 1, 58, 5, 58, 519, 8, 58, 10, 58, 12, 58, 522, 9, 
	58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 
	5, 58, 534, 8, 58, 10, 58, 12, 58, 537, 9, 58, 1, 58, 1, 58, 1, 58, 1, 
	58, 3, 58, 543, 8, 58, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 3, 60, 550, 8, 
	60, 1, 60, 1, 60, 1, 60, 1, 60, 5, 60, 556, 8, 60, 10, 60, 12, 60, 559, 
	9, 60, 1, 61, 3, 61, 562, 8, 61, 1, 61, 3, 61, 565, 8, 61, 1, 61, 1, 61, 
	3, 61, 569, 8, 61, 1, 61, 3, 61, 572, 8, 61, 1, 61, 3, 61, 575, 8, 61, 
	1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 
	63, 1, 63, 3, 63, 589, 8, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 
	1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 
	64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 
	1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 
	64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 
	1, 64, 3, 64, 640, 8, 64, 1, 65, 1, 65, 1, 65, 1, 65, 1, 66, 1, 66, 1, 
	66, 1, 66, 1, 67, 1, 67, 1, 67, 1, 67, 1, 68, 1, 68, 1, 68, 1, 68, 1, 69, 
	1, 69, 1, 69, 1, 69, 1, 70, 1, 70, 1, 70, 1, 70, 1, 71, 1, 71, 1, 71, 1, 
	71, 4, 472, 486, 520, 535, 0, 72, 2, 11, 4, 1, 6, 2, 8, 12, 10, 13, 12, 
	14, 14, 15, 16, 16, 18, 17, 20, 18, 22, 19, 24, 20, 26, 21, 28, 22, 30, 
	23, 32, 24, 34, 25, 36, 26, 38, 27, 40, 3, 42, 28, 44, 29, 46, 4, 48, 30, 
	50, 31, 52, 5, 54, 32, 56, 33, 58, 34, 60, 7, 62, 8, 64, 35, 66, 36, 68, 
	37, 70, 38, 72, 39, 74, 40, 76, 41, 78, 42, 80, 43, 82, 44, 84, 45, 86, 
	0, 88, 0, 90, 0, 92, 0, 94, 0, 96, 0, 98, 0, 100, 0, 102, 0, 104, 0, 106, 
	0, 108, 9, 110, 46, 112, 47, 114, 48, 116, 49, 118, 50, 120, 51, 122, 10, 
	124, 52, 126, 0, 128, 53, 130, 54, 132, 0, 134, 0, 136, 0, 138, 0, 140, 
	0, 142, 0, 144, 0, 2, 0, 1, 16, 2, 0, 65, 90, 97, 122, 2, 0, 69, 69, 101, 
	101, 2, 0, 43, 43, 45, 45, 3, 0, 48, 57, 65, 70, 97, 102, 2, 0, 82, 82, 
	114, 114, 10, 0, 34, 34, 39, 39, 63, 63, 92, 92, 96, 98, 102, 102, 110, 
	110, 114, 114, 116, 116, 118, 118, 2, 0, 88, 88, 120, 120, 3, 0, 9, 9, 
	12, 12, 32, 32, 1, 0, 10, 10, 2, 0, 85, 85, 117, 117, 4, 0, 10, 10, 13, 
	13, 34, 34, 92, 92, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 1, 0, 92, 92, 
	3, 0, 10, 10, 13, 13, 34, 34, 3, 0, 10, 10, 13, 13, 39, 39, 2, 0, 66, 66, 
	98, 98, 715, 0, 2, 1, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 6, 1, 0, 0, 0, 0, 8, 
	1, 0, 0, 0, 0, 10, 1, 0, 0, 0, 0, 12, 1, 0, 0, 0, 0, 14, 1, 0, 0, 0, 0, 
	16, 1, 0, 0, 0, 0, 18, 1, 0, 0, 0, 0, 20, 1, 0, 0, 0, 0, 22, 1, 0, 0, 0, 
	0, 24, 1, 0, 0, 0, 0, 26, 1, 0, 0, 0, 0, 28, 1, 0, 0, 0, 0, 30, 1, 0, 0, 
	0, 0, 32, 1, 0, 0, 0, 0, 34, 1, 0, 0, 0, 0, 36, 1, 0, 0, 0, 0, 38, 1, 0, 
	0, 0, 0, 40, 1, 0, 0, 0, 0, 42, 1, 0, 0, 0, 0, 44, 1, 0, 0, 0, 0, 46, 1, 
	0, 0, 0, 0, 48, 1, 0, 0, 0, 0, 50, 1, 0, 0, 0, 0, 52, 1, 0, 0, 0, 0, 54, 
	1, 0, 0, 0, 0, 56, 1, 0, 0, 0, 0, 58, 1, 0, 0, 0, 0, 60, 1, 0, 0, 0, 0, 
	62, 1, 0, 0, 0, 0, 64, 1, 0, 0, 0, 0, 66, 1, 0, 0, 0, 0, 68, 1, 0, 0, 0, 
	0, 70, 1, 0, 0, 0, 0, 72, 1, 0, 0, 0, 0, 74, 1, 0, 0, 0, 0, 76, 1, 0, 0, 
	0, 0, 78, 1, 0, 0, 0, 0, 80, 1, 0, 0, 0, 0, 82, 1, 0, 0, 0, 0, 84, 1, 0, 
	0, 0, 0, 108, 1, 0, 0, 0, 0, 110, 1, 0, 0, 0, 0, 112, 1, 0, 0, 0, 0, 114, 
	1, 0, 0, 0, 0, 116, 1, 0, 0, 0, 0, 118, 1, 0, 0, 0, 0, 120, 1, 0, 0, 0, 
	0, 122, 1, 0, 0, 0, 0, 124, 1, 0, 0, 0, 1, 126, 1, 0, 0, 0, 1, 128, 1, 
	0, 0, 0, 1, 130, 1, 0, 0, 0, 1, 132, 1, 0, 0, 0, 1, 134, 1, 0, 0, 0, 1, 
	136, 1, 0, 0, 0, 1, 138, 1, 0, 0, 0, 1, 140, 1, 0, 0, 0, 1, 142, 1, 0, 
	0, 0, 1, 144, 1, 0, 0, 0, 2, 146, 1, 0, 0, 0, 4, 148, 1, 0, 0, 0, 6, 150, 
	1, 0, 0, 0, 8, 152, 1, 0, 0, 0, 10, 156, 1, 0, 0, 0, 12, 159, 1, 0, 0, 
	0, 14, 167, 1, 0, 0, 0, 16, 172, 1, 0, 0, 0, 18, 178, 1, 0, 0, 0, 20, 185, 
	1, 0, 0, 0, 22, 189, 1, 0, 0, 0, 24, 194, 1, 0, 0, 0, 26, 206, 1, 0, 0, 
	0, 28, 216, 1, 0, 0, 0, 30, 225, 1, 0, 0, 0, 32, 232, 1, 0, 0, 0, 34, 237, 
	1, 0, 0, 0, 36, 240, 1, 0, 0, 0, 38, 243, 1, 0, 0, 0, 40, 246, 1, 0, 0, 
	0, 42, 248, 1, 0, 0, 0, 44, 251, 1, 0, 0, 0, 46, 254, 1, 0, 0, 0, 48, 256, 
	1, 0, 0, 0, 50, 259, 1, 0, 0, 0, 52, 262, 1, 0, 0, 0, 54, 264, 1, 0, 0, 
	0, 56, 266, 1, 0, 0, 0, 58, 268, 1, 0, 0, 0, 60, 270, 1, 0, 0, 0, 62, 272, 
	1, 0, 0, 0, 64, 274, 1, 0, 0, 0, 66, 276, 1, 0, 0, 0, 68, 278, 1, 0, 0, 
	0, 70, 280, 1, 0, 0, 0, 72, 282, 1, 0, 0, 0, 74, 284, 1, 0, 0, 0, 76, 286, 
	1, 0, 0, 0, 78, 288, 1, 0, 0, 0, 80, 290, 1, 0, 0, 0, 82, 295, 1, 0, 0, 
	0, 84, 301, 1, 0, 0, 0, 86, 306, 1, 0, 0, 0, 88, 308, 1, 0, 0, 0, 90, 310, 
	1, 0, 0, 0, 92, 312, 1, 0, 0, 0, 94, 321, 1, 0, 0, 0, 96, 323, 1, 0, 0, 
	0, 98, 329, 1, 0, 0, 0, 100, 331, 1, 0, 0, 0, 102, 334, 1, 0, 0, 0, 104, 
	339, 1, 0, 0, 0, 106, 362, 1, 0, 0, 0, 108, 365, 1, 0, 0, 0, 110, 369, 
	1, 0, 0, 0, 112, 410, 1, 0, 0, 0, 114, 425, 1, 0, 0, 0, 116, 444, 1, 0, 
	0, 0, 118, 542, 1, 0, 0, 0, 120, 544, 1, 0, 0, 0, 122, 549, 1, 0, 0, 0, 
	124, 561, 1, 0, 0, 0, 126, 576, 1, 0, 0, 0, 128, 588, 1, 0, 0, 0, 130, 
	639, 1, 0, 0, 0, 132, 641, 1, 0, 0, 0, 134, 645, 1, 0, 0, 0, 136, 649, 
	1, 0, 0, 0, 138, 653, 1, 0, 0, 0, 140, 657, 1, 0, 0, 0, 142, 661, 1, 0, 
	0, 0, 144, 665, 1, 0, 0, 0, 146, 147, 5, 35, 0, 0, 147, 3, 1, 0, 0, 0, 
	148, 149, 5, 58, 0, 0, 149, 5, 1, 0, 0, 0, 150, 151, 5, 44, 0, 0, 151, 
	7, 1, 0, 0, 0, 152, 153, 5, 97, 0, 0, 153, 154, 5, 110, 0, 0, 154, 155, 
	5, 100, 0, 0, 155, 9, 1, 0, 0, 0, 156, 157, 5, 111, 0, 0, 157, 158, 5, 
	114, 0, 0, 158, 11, 1, 0, 0, 0, 159, 160, 5, 98, 0, 0, 160, 161, 5, 117, 
	0, 0, 161, 162, 5, 116, 0, 0, 162, 163, 5, 32, 0, 0, 163, 164, 5, 110, 
	0, 0, 164, 165, 5, 111, 0, 0, 165, 166, 5, 116, 0, 0, 166, 13, 1, 0, 0, 
	0, 167, 168, 5, 102, 0, 0, 168, 169, 5, 114, 0, 0, 169, 170, 5, 111, 0, 
	0, 170, 171, 5, 109, 0, 0, 171, 15, 1, 0, 0, 0, 172, 173, 5, 109, 0, 0, 
	173, 174, 5, 111, 0, 0, 174, 175, 5, 100, 0, 0, 175, 176, 5, 101, 0, 0, 
	176, 177, 5, 108, 0, 0, 177, 17, 1, 0, 0, 0, 178, 179, 5, 115, 0, 0, 179, 
	180, 5, 99, 0, 0, 180, 181, 5, 104, 0, 0, 181, 182, 5, 101, 0, 0, 182, 
	183, 5, 109, 0, 0, 183, 184, 5, 97, 0, 0, 184, 19, 1, 0, 0, 0, 185, 186, 
	5, 49, 0, 0, 186, 187, 5, 46, 0, 0, 187, 188, 5, 49, 0, 0, 188, 21, 1, 
	0, 0, 0, 189, 190, 5, 116, 0, 0, 190, 191, 5, 121, 0, 0, 191, 192, 5, 112, 
	0, 0, 192, 193, 5, 101, 0, 0, 193, 23, 1, 0, 0, 0, 194, 195, 5, 99, 0, 
	0, 195, 196, 5, 111, 0, 0, 196, 197, 5, 110, 0, 0, 197, 198, 5, 100, 0, 
	0, 198, 199, 5, 105, 0, 0, 199, 200, 5, 116, 0, 0, 200, 201, 5, 105, 0, 
	0, 201, 202, 5, 111, 0, 0, 202, 203, 5, 110, 0, 0, 203, 204, 1, 0, 0, 0, 
	204, 205, 6, 11, 0, 0, 205, 25, 1, 0, 0, 0, 206, 207, 5, 114, 0, 0, 207, 
	208, 5, 101, 0, 0, 208, 209, 5, 108, 0, 0, 209, 210, 5, 97, 0, 0, 210, 
	211, 5, 116, 0, 0, 211, 212, 5, 105, 0, 0, 212, 213, 5, 111, 0, 0, 213, 
	214, 5, 110, 0, 0, 214, 215, 5, 115, 0, 0, 215, 27, 1, 0, 0, 0, 216, 217, 
	5, 114, 0, 0, 217, 218, 5, 101, 0, 0, 218, 219, 5, 108, 0, 0, 219, 220, 
	5, 97, 0, 0, 220, 221, 5, 116, 0, 0, 221, 222, 5, 105, 0, 0, 222, 223, 
	5, 111, 0, 0, 223, 224, 5, 110, 0, 0, 224, 29, 1, 0, 0, 0, 225, 226, 5, 
	100, 0, 0, 226, 227, 5, 101, 0, 0, 227, 228, 5, 102, 0, 0, 228, 229, 5, 
	105, 0, 0, 229, 230, 5, 110, 0, 0, 230, 231, 5, 101, 0, 0, 231, 31, 1, 
	0, 0, 0, 232, 233, 5, 119, 0, 0, 233, 234, 5, 105, 0, 0, 234, 235, 5, 116, 
	0, 0, 235, 236, 5, 104, 0, 0, 236, 33, 1, 0, 0, 0, 237, 238, 5, 61, 0, 
	0, 238, 239, 5, 61, 0, 0, 239, 35, 1, 0, 0, 0, 240, 241, 5, 33, 0, 0, 241, 
	242, 5, 61, 0, 0, 242, 37, 1, 0, 0, 0, 243, 244, 5, 105, 0, 0, 244, 245, 
	5, 110, 0, 0, 245, 39, 1, 0, 0, 0, 246, 247, 5, 60, 0, 0, 247, 41, 1, 0, 
	0, 0, 248, 249, 5, 60, 0, 0, 249, 250, 5, 61, 0, 0, 250, 43, 1, 0, 0, 0, 
	251, 252, 5, 62, 0, 0, 252, 253, 5, 61, 0, 0, 253, 45, 1, 0, 0, 0, 254, 
	255, 5, 62, 0, 0, 255, 47, 1, 0, 0, 0, 256, 257, 5, 38, 0, 0, 257, 258, 
	5, 38, 0, 0, 258, 49, 1, 0, 0, 0, 259, 260, 5, 124, 0, 0, 260, 261, 5, 
	124, 0, 0, 261, 51, 1, 0, 0, 0, 262, 263, 5, 91, 0, 0, 263, 53, 1, 0, 0, 
	0, 264, 265, 5, 93, 0, 0, 265, 55, 1, 0, 0, 0, 266, 267, 5, 123, 0, 0, 
	267, 57, 1, 0, 0, 0, 268, 269, 5, 125, 0, 0, 269, 59, 1, 0, 0, 0, 270, 
	271, 5, 40, 0, 0, 271, 61, 1, 0, 0, 0, 272, 273, 5, 41, 0, 0, 273, 63, 
	1, 0, 0, 0, 274, 275, 5, 46, 0, 0, 275, 65, 1, 0, 0, 0, 276, 277, 5, 45, 
	0, 0, 277, 67, 1, 0, 0, 0, 278, 279, 5, 33, 0, 0, 279, 69, 1, 0, 0, 0, 
	280, 281, 5, 63, 0, 0, 281, 71, 1, 0, 0, 0, 282, 283, 5, 43, 0, 0, 283, 
	73, 1, 0, 0, 0, 284, 285, 5, 42, 0, 0, 285, 75, 1, 0, 0, 0, 286, 287, 5, 
	47, 0, 0, 287, 77, 1, 0, 0, 0, 288, 289, 5, 37, 0, 0, 289, 79, 1, 0, 0, 
	0, 290, 291, 5, 116, 0, 0, 291, 292, 5, 114, 0, 0, 292, 293, 5, 117, 0, 
	0, 293, 294, 5, 101, 0, 0, 294, 81, 1, 0, 0, 0, 295, 296, 5, 102, 0, 0, 
	296, 297, 5, 97, 0, 0, 297, 298, 5, 108, 0, 0, 298, 299, 5, 115, 0, 0, 
	299, 300, 5, 101, 0, 0, 300, 83, 1, 0, 0, 0, 301, 302, 5, 110, 0, 0, 302, 
	303, 5, 117, 0, 0, 303, 304, 5, 108, 0, 0, 304, 305, 5, 108, 0, 0, 305, 
	85, 1, 0, 0, 0, 306, 307, 5, 92, 0, 0, 307, 87, 1, 0, 0, 0, 308, 309, 7, 
	0, 0, 0, 309, 89, 1, 0, 0, 0, 310, 311, 2, 48, 57, 0, 311, 91, 1, 0, 0, 
	0, 312, 314, 7, 1, 0, 0, 313, 315, 7, 2, 0, 0, 314, 313, 1, 0, 0, 0, 314, 
	315, 1, 0, 0, 0, 315, 317, 1, 0, 0, 0, 316, 318, 3, 90, 44, 0, 317, 316, 
	1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 319, 320, 1, 0, 
	0, 0, 320, 93, 1, 0, 0, 0, 321, 322, 7, 3, 0, 0, 322, 95, 1, 0, 0, 0, 323, 
	324, 7, 4, 0, 0, 324, 97, 1, 0, 0, 0, 325, 330, 3, 100, 49, 0, 326, 330, 
	3, 104, 51, 0, 327, 330, 3, 106, 52, 0, 328, 330, 3, 102, 50, 0, 329, 325, 
	1, 0, 0, 0, 329, 326, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 329, 328, 1, 0, 
	0, 0, 330, 99, 1, 0, 0, 0, 331, 332, 3, 86, 42, 0, 332, 333, 7, 5, 0, 0, 
	333, 101, 1, 0, 0, 0, 334, 335, 3, 86, 42, 0, 335, 336, 2, 48, 51, 0, 336, 
	337, 2, 48, 55, 0, 337, 338, 2, 48, 55, 0, 338, 103, 1, 0, 0, 0, 339, 340, 
	3, 86, 42, 0, 340, 341, 7, 6, 0, 0, 341, 342, 3, 94, 46, 0, 342, 343, 3, 
	94, 46, 0, 343, 105, 1, 0, 0, 0, 344, 345, 3, 86, 42, 0, 345, 346, 5, 117, 
	0, 0, 346, 347, 3, 94, 46, 0, 347, 348, 3, 94, 46, 0, 348, 349, 3, 94, 
	46, 0, 349, 350, 3, 94, 46, 0, 350, 363, 1, 0, 0, 0, 351, 352, 3, 86, 42, 
	0, 352, 353, 5, 85, 0, 0, 353, 354, 3, 94, 46, 0, 354, 355, 3, 94, 46, 
	0, 355, 356, 3, 94, 46, 0, 356, 357, 3, 94, 46, 0, 357, 358, 3, 94, 46, 
	0, 358, 359, 3, 94, 46, 0, 359, 360, 3, 94, 46, 0, 360, 361, 3, 94, 46, 
	0, 361, 363, 1, 0, 0, 0, 362, 344, 1, 0, 0, 0, 362, 351, 1, 0, 0, 0, 363, 
	107, 1, 0, 0, 0, 364, 366, 7, 7, 0, 0, 365, 364, 1, 0, 0, 0, 366, 367, 
	1, 0, 0, 0, 367, 365, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 109, 1, 0, 
	0, 0, 369, 370, 5, 47, 0, 0, 370, 371, 5, 47, 0, 0, 371, 375, 1, 0, 0, 
	0, 372, 374, 8, 8, 0, 0, 373, 372, 1, 0, 0, 0, 374, 377, 1, 0, 0, 0, 375, 
	373, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376, 378, 1, 0, 0, 0, 377, 375, 
	1, 0, 0, 0, 378, 379, 6, 54, 1, 0, 379, 111, 1, 0, 0, 0, 380, 382, 3, 90, 
	44, 0, 381, 380, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 
	383, 384, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 387, 5, 46, 0, 0, 386, 
	388, 3, 90, 44, 0, 387, 386, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 387, 
	1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 392, 1, 0, 0, 0, 391, 393, 3, 92, 
	45, 0, 392, 391, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 411, 1, 0, 0, 0, 
	394, 396, 3, 90, 44, 0, 395, 394, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 
	395, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 400, 
	3, 92, 45, 0, 400, 411, 1, 0, 0, 0, 401, 403, 5, 46, 0, 0, 402, 404, 3, 
	90, 44, 0, 403, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 403, 1, 0, 
	0, 0, 405, 406, 1, 0, 0, 0, 406, 408, 1, 0, 0, 0, 407, 409, 3, 92, 45, 
	0, 408, 407, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 411, 1, 0, 0, 0, 410, 
	381, 1, 0, 0, 0, 410, 395, 1, 0, 0, 0, 410, 401, 1, 0, 0, 0, 411, 113, 
	1, 0, 0, 0, 412, 414, 3, 90, 44, 0, 413, 412, 1, 0, 0, 0, 414, 415, 1, 
	0, 0, 0, 415, 413, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416, 426, 1, 0, 0, 
	0, 417, 418, 5, 48, 0, 0, 418, 419, 5, 120, 0, 0, 419, 421, 1, 0, 0, 0, 
	420, 422, 3, 94, 46, 0, 421, 420, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 
	421, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 426, 1, 0, 0, 0, 425, 413, 
	1, 0, 0, 0, 425, 417, 1, 0, 0, 0, 426, 115, 1, 0, 0, 0, 427, 429, 3, 90, 
	44, 0, 428, 427, 1, 0, 0, 0, 429, 430, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 
	430, 431, 1, 0, 0, 0, 431, 432, 1, 0, 0, 0, 432, 433, 7, 9, 0, 0, 433, 
	445, 1, 0, 0, 0, 434, 435, 5, 48, 0, 0, 435, 436, 5, 120, 0, 0, 436, 438, 
	1, 0, 0, 0, 437, 439, 3, 94, 46, 0, 438, 437, 1, 0, 0, 0, 439, 440, 1, 
	0, 0, 0, 440, 438, 1, 0, 0, 0, 440, 441, 1, 0, 0, 0, 441, 442, 1, 0, 0, 
	0, 442, 443, 7, 9, 0, 0, 443, 445, 1, 0, 0, 0, 444, 428, 1, 0, 0, 0, 444, 
	434, 1, 0, 0, 0, 445, 117, 1, 0, 0, 0, 446, 451, 5, 34, 0, 0, 447, 450, 
	3, 98, 48, 0, 448, 450, 8, 10, 0, 0, 449, 447, 1, 0, 0, 0, 449, 448, 1, 
	0, 0, 0, 450, 453, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 451, 452, 1, 0, 0, 
	0, 452, 454, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 454, 543, 5, 34, 0, 0, 455, 
	460, 5, 39, 0, 0, 456, 459, 3, 98, 48, 0, 457, 459, 8, 11, 0, 0, 458, 456, 
	1, 0, 0, 0, 458, 457, 1, 0, 0, 0, 459, 462, 1, 0, 0, 0, 460, 458, 1, 0, 
	0, 0, 460, 461, 1, 0, 0, 0, 461, 463, 1, 0, 0, 0, 462, 460, 1, 0, 0, 0, 
	463, 543, 5, 39, 0, 0, 464, 465, 5, 34, 0, 0, 465, 466, 5, 34, 0, 0, 466, 
	467, 5, 34, 0, 0, 467, 472, 1, 0, 0, 0, 468, 471, 3, 98, 48, 0, 469, 471, 
	8, 12, 0, 0, 470, 468, 1, 0, 0, 0, 470, 469, 1, 0, 0, 0, 471, 474, 1, 0, 
	0, 0, 472, 473, 1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 473, 475, 1, 0, 0, 0, 
	474, 472, 1, 0, 0, 0, 475, 476, 5, 34, 0, 0, 476, 477, 5, 34, 0, 0, 477, 
	543, 5, 34, 0, 0, 478, 479, 5, 39, 0, 0, 479, 480, 5, 39, 0, 0, 480, 481, 
	5, 39, 0, 0, 481, 486, 1, 0, 0, 0, 482, 485, 3, 98, 48, 0, 483, 485, 8, 
	12, 0, 0, 484, 482, 1, 0, 0, 0, 484, 483, 1, 0, 0, 0, 485, 488, 1, 0, 0, 
	0, 486, 487, 1, 0, 0, 0, 486, 484, 1, 0, 0, 0, 487, 489, 1, 0, 0, 0, 488, 
	486, 1, 0, 0, 0, 489, 490, 5, 39, 0, 0, 490, 491, 5, 39, 0, 0, 491, 543, 
	5, 39, 0, 0, 492, 493, 3, 96, 47, 0, 493, 497, 5, 34, 0, 0, 494, 496, 8, 
	13, 0, 0, 495, 494, 1, 0, 0, 0, 496, 499, 1, 0, 0, 0, 497, 495, 1, 0, 0, 
	0, 497, 498, 1, 0, 0, 0, 498, 500, 1, 0, 0, 0, 499, 497, 1, 0, 0, 0, 500, 
	501, 5, 34, 0, 0, 501, 543, 1, 0, 0, 0, 502, 503, 3, 96, 47, 0, 503, 507, 
	5, 39, 0, 0, 504, 506, 8, 14, 0, 0, 505, 504, 1, 0, 0, 0, 506, 509, 1, 
	0, 0, 0, 507, 505, 1, 0, 0, 0, 507, 508, 1, 0, 0, 0, 508, 510, 1, 0, 0, 
	0, 509, 507, 1, 0, 0, 0, 510, 511, 5, 39, 0, 0, 511, 543, 1, 0, 0, 0, 512, 
	513, 3, 96, 47, 0, 513, 514, 5, 34, 0, 0, 514, 515, 5, 34, 0, 0, 515, 516, 
	5, 34, 0, 0, 516, 520, 1, 0, 0, 0, 517, 519, 9, 0, 0, 0, 518, 517, 1, 0, 
	0, 0, 519, 522, 1, 0, 0, 0, 520, 521, 1, 0, 0, 0, 520, 518, 1, 0, 0, 0, 
	521, 523, 1, 0, 0, 0, 522, 520, 1, 0, 0, 0, 523, 524, 5, 34, 0, 0, 524, 
	525, 5, 34, 0, 0, 525, 526, 5, 34, 0, 0, 526, 543, 1, 0, 0, 0, 527, 528, 
	3, 96, 47, 0, 528, 529, 5, 39, 0, 0, 529, 530, 5, 39, 0, 0, 530, 531, 5, 
	39, 0, 0, 531, 535, 1, 0, 0, 0, 532, 534, 9, 0, 0, 0, 533, 532, 1, 0, 0, 
	0, 534, 537, 1, 0, 0, 0, 535, 536, 1, 0, 0, 0, 535, 533, 1, 0, 0, 0, 536, 
	538, 1, 0, 0, 0, 537, 535, 1, 0, 0, 0, 538, 539, 5, 39, 0, 0, 539, 540, 
	5, 39, 0, 0, 540, 541, 5, 39, 0, 0, 541, 543, 1, 0, 0, 0, 542, 446, 1, 
	0, 0, 0, 542, 455, 1, 0, 0, 0, 542, 464, 1, 0, 0, 0, 542, 478, 1, 0, 0, 
	0, 542, 492, 1, 0, 0, 0, 542, 502, 1, 0, 0, 0, 542, 512, 1, 0, 0, 0, 542, 
	527, 1, 0, 0, 0, 543, 119, 1, 0, 0, 0, 544, 545, 7, 15, 0, 0, 545, 546, 
	3, 118, 58, 0, 546, 121, 1, 0, 0, 0, 547, 550, 3, 88, 43, 0, 548, 550, 
	5, 95, 0, 0, 549, 547, 1, 0, 0, 0, 549, 548, 1, 0, 0, 0, 550, 557, 1, 0, 
	0, 0, 551, 556, 3, 88, 43, 0, 552, 556, 3, 90, 44, 0, 553, 556, 5, 95, 
	0, 0, 554, 556, 3, 66, 32, 0, 555, 551, 1, 0, 0, 0, 555, 552, 1, 0, 0, 
	0, 555, 553, 1, 0, 0, 0, 555, 554, 1, 0, 0, 0, 556, 559, 1, 0, 0, 0, 557, 
	555, 1, 0, 0, 0, 557, 558, 1, 0, 0, 0, 558, 123, 1, 0, 0, 0, 559, 557, 
	1, 0, 0, 0, 560, 562, 3, 108, 53, 0, 561, 560, 1, 0, 0, 0, 561, 562, 1, 
	0, 0, 0, 562, 568, 1, 0, 0, 0, 563, 565, 5, 13, 0, 0, 564, 563, 1, 0, 0, 
	0, 564, 565, 1, 0, 0, 0, 565, 566, 1, 0, 0, 0, 566, 569, 5, 10, 0, 0, 567, 
	569, 2, 12, 13, 0, 568, 564, 1, 0, 0, 0, 568, 567, 1, 0, 0, 0, 569, 571, 
	1, 0, 0, 0, 570, 572, 3, 108, 53, 0, 571, 570, 1, 0, 0, 0, 571, 572, 1, 
	0, 0, 0, 572, 574, 1, 0, 0, 0, 573, 575, 3, 124, 61, 0, 574, 573, 1, 0, 
	0, 0, 574, 575, 1, 0, 0, 0, 575, 125, 1, 0, 0, 0, 576, 577, 3, 62, 30, 
	0, 577, 578, 1, 0, 0, 0, 578, 579, 6, 62, 2, 0, 579, 580, 6, 62, 3, 0, 
	580, 127, 1, 0, 0, 0, 581, 582, 5, 109, 0, 0, 582, 583, 5, 97, 0, 0, 583, 
	589, 5, 112, 0, 0, 584, 585, 5, 108, 0, 0, 585, 586, 5, 105, 0, 0, 586, 
	587, 5, 115, 0, 0, 587, 589, 5, 116, 0, 0, 588, 581, 1, 0, 0, 0, 588, 584, 
	1, 0, 0, 0, 589, 129, 1, 0, 0, 0, 590, 591, 5, 98, 0, 0, 591, 592, 5, 111, 
	0, 0, 592, 593, 5, 111, 0, 0, 593, 640, 5, 108, 0, 0, 594, 595, 5, 115, 
	0, 0, 595, 596, 5, 116, 0, 0, 596, 597, 5, 114, 0, 0, 597, 598, 5, 105, 
	0, 0, 598, 599, 5, 110, 0, 0, 599, 640, 5, 103, 0, 0, 600, 601, 5, 105, 
	0, 0, 601, 602, 5, 110, 0, 0, 602, 640, 5, 116, 0, 0, 603, 604, 5, 117, 
	0, 0, 604, 605, 5, 105, 0, 0, 605, 606, 5, 110, 0, 0, 606, 640, 5, 116, 
	0, 0, 607, 608, 5, 100, 0, 0, 608, 609, 5, 111, 0, 0, 609, 610, 5, 117, 
	0, 0, 610, 611, 5, 98, 0, 0, 611, 612, 5, 108, 0, 0, 612, 640, 5, 101, 
	0, 0, 613, 614, 5, 100, 0, 0, 614, 615, 5, 117, 0, 0, 615, 616, 5, 114, 
	0, 0, 616, 617, 5, 97, 0, 0, 617, 618, 5, 116, 0, 0, 618, 619, 5, 105, 
	0, 0, 619, 620, 5, 111, 0, 0, 620, 640, 5, 110, 0, 0, 621, 622, 5, 116, 
	0, 0, 622, 623, 5, 105, 0, 0, 623, 624, 5, 109, 0, 0, 624, 625, 5, 101, 
	0, 0, 625, 626, 5, 115, 0, 0, 626, 627, 5, 116, 0, 0, 627, 628, 5, 97, 
	0, 0, 628, 629, 5, 109, 0, 0, 629, 640, 5, 112, 0, 0, 630, 631, 5, 105, 
	0, 0, 631, 632, 5, 112, 0, 0, 632, 633, 5, 97, 0, 0, 633, 634, 5, 100, 
	0, 0, 634, 635, 5, 100, 0, 0, 635, 636, 5, 114, 0, 0, 636, 637, 5, 101, 
	0, 0, 637, 638, 5, 115, 0, 0, 638, 640, 5, 115, 0, 0, 639, 590, 1, 0, 0, 
	0, 639, 594, 1, 0, 0, 0, 639, 600, 1, 0, 0, 0, 639, 603, 1, 0, 0, 0, 639, 
	607, 1, 0, 0, 0, 639, 613, 1, 0, 0, 0, 639, 621, 1, 0, 0, 0, 639, 630, 
	1, 0, 0, 0, 640, 131, 1, 0, 0, 0, 641, 642, 3, 40, 19, 0, 642, 643, 1, 
	0, 0, 0, 643, 644, 6, 65, 4, 0, 644, 133, 1, 0, 0, 0, 645, 646, 3, 46, 
	22, 0, 646, 647, 1, 0, 0, 0, 647, 648, 6, 66, 5, 0, 648, 135, 1, 0, 0, 
	0, 649, 650, 3, 60, 29, 0, 650, 651, 1, 0, 0, 0, 651, 652, 6, 67, 6, 0, 
	652, 137, 1, 0, 0, 0, 653, 654, 3, 4, 1, 0, 654, 655, 1, 0, 0, 0, 655, 
	656, 6, 68, 7, 0, 656, 139, 1, 0, 0, 0, 657, 658, 3, 6, 2, 0, 658, 659, 
	1, 0, 0, 0, 659, 660, 6, 69, 8, 0, 660, 141, 1, 0, 0, 0, 661, 662, 3, 108, 
	53, 0, 662, 663, 1, 0, 0, 0, 663, 664, 6, 70, 9, 0, 664, 143, 1, 0, 0, 
	0, 665, 666, 3, 122, 60, 0, 666, 667, 1, 0, 0, 0, 667, 668, 6, 71, 10, 
	0, 668, 145, 1, 0, 0, 0, 44, 0, 1, 314, 319, 329, 362, 367, 375, 383, 389, 
	392, 397, 405, 408, 410, 415, 423, 425, 430, 440, 444, 449, 451, 458, 460, 
	470, 472, 484, 486, 497, 507, 520, 535, 542, 549, 555, 557, 561, 564, 568, 
	571, 574, 588, 639, 11, 5, 1, 0, 0, 1, 0, 7, 8, 0, 4, 0, 0, 7, 3, 0, 7, 
	4, 0, 7, 7, 0, 7, 1, 0, 7, 2, 0, 7, 9, 0, 7, 10, 0,
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
	OpenFGALexerRELATION = 22
	OpenFGALexerDEFINE = 23
	OpenFGALexerKEYWORD_WITH = 24
	OpenFGALexerEQUALS = 25
	OpenFGALexerNOT_EQUALS = 26
	OpenFGALexerIN = 27
	OpenFGALexerLESS_EQUALS = 28
	OpenFGALexerGREATER_EQUALS = 29
	OpenFGALexerLOGICAL_AND = 30
	OpenFGALexerLOGICAL_OR = 31
	OpenFGALexerRPRACKET = 32
	OpenFGALexerLBRACE = 33
	OpenFGALexerRBRACE = 34
	OpenFGALexerDOT = 35
	OpenFGALexerMINUS = 36
	OpenFGALexerEXCLAM = 37
	OpenFGALexerQUESTIONMARK = 38
	OpenFGALexerPLUS = 39
	OpenFGALexerSTAR = 40
	OpenFGALexerSLASH = 41
	OpenFGALexerPERCENT = 42
	OpenFGALexerCEL_TRUE = 43
	OpenFGALexerCEL_FALSE = 44
	OpenFGALexerNUL = 45
	OpenFGALexerCEL_COMMENT = 46
	OpenFGALexerNUM_FLOAT = 47
	OpenFGALexerNUM_INT = 48
	OpenFGALexerNUM_UINT = 49
	OpenFGALexerSTRING = 50
	OpenFGALexerBYTES = 51
	OpenFGALexerNEWLINE = 52
	OpenFGALexerCONDITION_PARAM_CONTAINER = 53
	OpenFGALexerCONDITION_PARAM_TYPE = 54
)

// OpenFGALexerCONDITION_DEF is the OpenFGALexer mode.
const OpenFGALexerCONDITION_DEF = 1

