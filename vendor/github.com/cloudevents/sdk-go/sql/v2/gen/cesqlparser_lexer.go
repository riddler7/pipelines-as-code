// Code generated from CESQLParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package gen

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CESQLParserLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var cesqlparserlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cesqlparserlexerLexerInit() {
	staticData := &cesqlparserlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "'('", "')'", "','", "'''", "'\"'", "'AND'", "'OR'", "'XOR'",
		"'NOT'", "'*'", "'/'", "'%'", "'+'", "'-'", "'='", "'!='", "'>'", "'>='",
		"'<'", "'<>'", "'<='", "'LIKE'", "'EXISTS'", "'IN'", "'TRUE'", "'FALSE'",
	}
	staticData.symbolicNames = []string{
		"", "SPACE", "LR_BRACKET", "RR_BRACKET", "COMMA", "SINGLE_QUOTE_SYMB",
		"DOUBLE_QUOTE_SYMB", "AND", "OR", "XOR", "NOT", "STAR", "DIVIDE", "MODULE",
		"PLUS", "MINUS", "EQUAL", "NOT_EQUAL", "GREATER", "GREATER_OR_EQUAL",
		"LESS", "LESS_GREATER", "LESS_OR_EQUAL", "LIKE", "EXISTS", "IN", "TRUE",
		"FALSE", "DQUOTED_STRING_LITERAL", "SQUOTED_STRING_LITERAL", "INTEGER_LITERAL",
		"IDENTIFIER", "IDENTIFIER_WITH_NUMBER", "FUNCTION_IDENTIFIER_WITH_UNDERSCORE",
	}
	staticData.ruleNames = []string{
		"SPACE", "ID_LITERAL", "DQUOTA_STRING", "SQUOTA_STRING", "INT_DIGIT",
		"FN_LITERAL", "LR_BRACKET", "RR_BRACKET", "COMMA", "SINGLE_QUOTE_SYMB",
		"DOUBLE_QUOTE_SYMB", "QUOTE_SYMB", "AND", "OR", "XOR", "NOT", "STAR",
		"DIVIDE", "MODULE", "PLUS", "MINUS", "EQUAL", "NOT_EQUAL", "GREATER",
		"GREATER_OR_EQUAL", "LESS", "LESS_GREATER", "LESS_OR_EQUAL", "LIKE",
		"EXISTS", "IN", "TRUE", "FALSE", "DQUOTED_STRING_LITERAL", "SQUOTED_STRING_LITERAL",
		"INTEGER_LITERAL", "IDENTIFIER", "IDENTIFIER_WITH_NUMBER", "FUNCTION_IDENTIFIER_WITH_UNDERSCORE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 33, 238, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 4, 0, 81, 8, 0, 11, 0, 12, 0,
		82, 1, 0, 1, 0, 1, 1, 4, 1, 88, 8, 1, 11, 1, 12, 1, 89, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 5, 2, 98, 8, 2, 10, 2, 12, 2, 101, 9, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 111, 8, 3, 10, 3, 12, 3, 114,
		9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 122, 8, 5, 10, 5, 12, 5,
		125, 9, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 3, 11, 139, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 3, 35, 215, 8, 35, 1, 35, 4, 35,
		218, 8, 35, 11, 35, 12, 35, 219, 1, 36, 4, 36, 223, 8, 36, 11, 36, 12,
		36, 224, 1, 37, 4, 37, 228, 8, 37, 11, 37, 12, 37, 229, 1, 38, 1, 38, 5,
		38, 234, 8, 38, 10, 38, 12, 38, 237, 9, 38, 0, 0, 39, 1, 1, 3, 0, 5, 0,
		7, 0, 9, 0, 11, 0, 13, 2, 15, 3, 17, 4, 19, 5, 21, 6, 23, 0, 25, 7, 27,
		8, 29, 9, 31, 10, 33, 11, 35, 12, 37, 13, 39, 14, 41, 15, 43, 16, 45, 17,
		47, 18, 49, 19, 51, 20, 53, 21, 55, 22, 57, 23, 59, 24, 61, 25, 63, 26,
		65, 27, 67, 28, 69, 29, 71, 30, 73, 31, 75, 32, 77, 33, 1, 0, 9, 3, 0,
		9, 10, 13, 13, 32, 32, 3, 0, 48, 57, 65, 90, 97, 122, 2, 0, 34, 34, 92,
		92, 2, 0, 39, 39, 92, 92, 1, 0, 48, 57, 1, 0, 65, 90, 2, 0, 65, 90, 95,
		95, 2, 0, 43, 43, 45, 45, 2, 0, 65, 90, 97, 122, 246, 0, 1, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0,
		0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1,
		0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45,
		1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0,
		53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0,
		0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0,
		0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0,
		0, 0, 0, 77, 1, 0, 0, 0, 1, 80, 1, 0, 0, 0, 3, 87, 1, 0, 0, 0, 5, 91, 1,
		0, 0, 0, 7, 104, 1, 0, 0, 0, 9, 117, 1, 0, 0, 0, 11, 119, 1, 0, 0, 0, 13,
		126, 1, 0, 0, 0, 15, 128, 1, 0, 0, 0, 17, 130, 1, 0, 0, 0, 19, 132, 1,
		0, 0, 0, 21, 134, 1, 0, 0, 0, 23, 138, 1, 0, 0, 0, 25, 140, 1, 0, 0, 0,
		27, 144, 1, 0, 0, 0, 29, 147, 1, 0, 0, 0, 31, 151, 1, 0, 0, 0, 33, 155,
		1, 0, 0, 0, 35, 157, 1, 0, 0, 0, 37, 159, 1, 0, 0, 0, 39, 161, 1, 0, 0,
		0, 41, 163, 1, 0, 0, 0, 43, 165, 1, 0, 0, 0, 45, 167, 1, 0, 0, 0, 47, 170,
		1, 0, 0, 0, 49, 172, 1, 0, 0, 0, 51, 175, 1, 0, 0, 0, 53, 177, 1, 0, 0,
		0, 55, 180, 1, 0, 0, 0, 57, 183, 1, 0, 0, 0, 59, 188, 1, 0, 0, 0, 61, 195,
		1, 0, 0, 0, 63, 198, 1, 0, 0, 0, 65, 203, 1, 0, 0, 0, 67, 209, 1, 0, 0,
		0, 69, 211, 1, 0, 0, 0, 71, 214, 1, 0, 0, 0, 73, 222, 1, 0, 0, 0, 75, 227,
		1, 0, 0, 0, 77, 231, 1, 0, 0, 0, 79, 81, 7, 0, 0, 0, 80, 79, 1, 0, 0, 0,
		81, 82, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 1,
		0, 0, 0, 84, 85, 6, 0, 0, 0, 85, 2, 1, 0, 0, 0, 86, 88, 7, 1, 0, 0, 87,
		86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0,
		0, 90, 4, 1, 0, 0, 0, 91, 99, 5, 34, 0, 0, 92, 93, 5, 92, 0, 0, 93, 98,
		9, 0, 0, 0, 94, 95, 5, 34, 0, 0, 95, 98, 5, 34, 0, 0, 96, 98, 8, 2, 0,
		0, 97, 92, 1, 0, 0, 0, 97, 94, 1, 0, 0, 0, 97, 96, 1, 0, 0, 0, 98, 101,
		1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 102, 1, 0, 0,
		0, 101, 99, 1, 0, 0, 0, 102, 103, 5, 34, 0, 0, 103, 6, 1, 0, 0, 0, 104,
		112, 5, 39, 0, 0, 105, 106, 5, 92, 0, 0, 106, 111, 9, 0, 0, 0, 107, 108,
		5, 39, 0, 0, 108, 111, 5, 39, 0, 0, 109, 111, 8, 3, 0, 0, 110, 105, 1,
		0, 0, 0, 110, 107, 1, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111, 114, 1, 0, 0,
		0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 115, 1, 0, 0, 0, 114,
		112, 1, 0, 0, 0, 115, 116, 5, 39, 0, 0, 116, 8, 1, 0, 0, 0, 117, 118, 7,
		4, 0, 0, 118, 10, 1, 0, 0, 0, 119, 123, 7, 5, 0, 0, 120, 122, 7, 6, 0,
		0, 121, 120, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123,
		124, 1, 0, 0, 0, 124, 12, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 127, 5,
		40, 0, 0, 127, 14, 1, 0, 0, 0, 128, 129, 5, 41, 0, 0, 129, 16, 1, 0, 0,
		0, 130, 131, 5, 44, 0, 0, 131, 18, 1, 0, 0, 0, 132, 133, 5, 39, 0, 0, 133,
		20, 1, 0, 0, 0, 134, 135, 5, 34, 0, 0, 135, 22, 1, 0, 0, 0, 136, 139, 3,
		19, 9, 0, 137, 139, 3, 21, 10, 0, 138, 136, 1, 0, 0, 0, 138, 137, 1, 0,
		0, 0, 139, 24, 1, 0, 0, 0, 140, 141, 5, 65, 0, 0, 141, 142, 5, 78, 0, 0,
		142, 143, 5, 68, 0, 0, 143, 26, 1, 0, 0, 0, 144, 145, 5, 79, 0, 0, 145,
		146, 5, 82, 0, 0, 146, 28, 1, 0, 0, 0, 147, 148, 5, 88, 0, 0, 148, 149,
		5, 79, 0, 0, 149, 150, 5, 82, 0, 0, 150, 30, 1, 0, 0, 0, 151, 152, 5, 78,
		0, 0, 152, 153, 5, 79, 0, 0, 153, 154, 5, 84, 0, 0, 154, 32, 1, 0, 0, 0,
		155, 156, 5, 42, 0, 0, 156, 34, 1, 0, 0, 0, 157, 158, 5, 47, 0, 0, 158,
		36, 1, 0, 0, 0, 159, 160, 5, 37, 0, 0, 160, 38, 1, 0, 0, 0, 161, 162, 5,
		43, 0, 0, 162, 40, 1, 0, 0, 0, 163, 164, 5, 45, 0, 0, 164, 42, 1, 0, 0,
		0, 165, 166, 5, 61, 0, 0, 166, 44, 1, 0, 0, 0, 167, 168, 5, 33, 0, 0, 168,
		169, 5, 61, 0, 0, 169, 46, 1, 0, 0, 0, 170, 171, 5, 62, 0, 0, 171, 48,
		1, 0, 0, 0, 172, 173, 5, 62, 0, 0, 173, 174, 5, 61, 0, 0, 174, 50, 1, 0,
		0, 0, 175, 176, 5, 60, 0, 0, 176, 52, 1, 0, 0, 0, 177, 178, 5, 60, 0, 0,
		178, 179, 5, 62, 0, 0, 179, 54, 1, 0, 0, 0, 180, 181, 5, 60, 0, 0, 181,
		182, 5, 61, 0, 0, 182, 56, 1, 0, 0, 0, 183, 184, 5, 76, 0, 0, 184, 185,
		5, 73, 0, 0, 185, 186, 5, 75, 0, 0, 186, 187, 5, 69, 0, 0, 187, 58, 1,
		0, 0, 0, 188, 189, 5, 69, 0, 0, 189, 190, 5, 88, 0, 0, 190, 191, 5, 73,
		0, 0, 191, 192, 5, 83, 0, 0, 192, 193, 5, 84, 0, 0, 193, 194, 5, 83, 0,
		0, 194, 60, 1, 0, 0, 0, 195, 196, 5, 73, 0, 0, 196, 197, 5, 78, 0, 0, 197,
		62, 1, 0, 0, 0, 198, 199, 5, 84, 0, 0, 199, 200, 5, 82, 0, 0, 200, 201,
		5, 85, 0, 0, 201, 202, 5, 69, 0, 0, 202, 64, 1, 0, 0, 0, 203, 204, 5, 70,
		0, 0, 204, 205, 5, 65, 0, 0, 205, 206, 5, 76, 0, 0, 206, 207, 5, 83, 0,
		0, 207, 208, 5, 69, 0, 0, 208, 66, 1, 0, 0, 0, 209, 210, 3, 5, 2, 0, 210,
		68, 1, 0, 0, 0, 211, 212, 3, 7, 3, 0, 212, 70, 1, 0, 0, 0, 213, 215, 7,
		7, 0, 0, 214, 213, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 217, 1, 0, 0,
		0, 216, 218, 3, 9, 4, 0, 217, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219,
		217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 72, 1, 0, 0, 0, 221, 223, 7,
		8, 0, 0, 222, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 222, 1, 0, 0,
		0, 224, 225, 1, 0, 0, 0, 225, 74, 1, 0, 0, 0, 226, 228, 7, 1, 0, 0, 227,
		226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230,
		1, 0, 0, 0, 230, 76, 1, 0, 0, 0, 231, 235, 7, 5, 0, 0, 232, 234, 7, 6,
		0, 0, 233, 232, 1, 0, 0, 0, 234, 237, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0,
		235, 236, 1, 0, 0, 0, 236, 78, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 14, 0,
		82, 89, 97, 99, 110, 112, 123, 138, 214, 219, 224, 229, 235, 1, 6, 0, 0,
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

// CESQLParserLexerInit initializes any static state used to implement CESQLParserLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCESQLParserLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CESQLParserLexerInit() {
	staticData := &cesqlparserlexerLexerStaticData
	staticData.once.Do(cesqlparserlexerLexerInit)
}

// NewCESQLParserLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCESQLParserLexer(input antlr.CharStream) *CESQLParserLexer {
	CESQLParserLexerInit()
	l := new(CESQLParserLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &cesqlparserlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "CESQLParser.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CESQLParserLexer tokens.
const (
	CESQLParserLexerSPACE                               = 1
	CESQLParserLexerLR_BRACKET                          = 2
	CESQLParserLexerRR_BRACKET                          = 3
	CESQLParserLexerCOMMA                               = 4
	CESQLParserLexerSINGLE_QUOTE_SYMB                   = 5
	CESQLParserLexerDOUBLE_QUOTE_SYMB                   = 6
	CESQLParserLexerAND                                 = 7
	CESQLParserLexerOR                                  = 8
	CESQLParserLexerXOR                                 = 9
	CESQLParserLexerNOT                                 = 10
	CESQLParserLexerSTAR                                = 11
	CESQLParserLexerDIVIDE                              = 12
	CESQLParserLexerMODULE                              = 13
	CESQLParserLexerPLUS                                = 14
	CESQLParserLexerMINUS                               = 15
	CESQLParserLexerEQUAL                               = 16
	CESQLParserLexerNOT_EQUAL                           = 17
	CESQLParserLexerGREATER                             = 18
	CESQLParserLexerGREATER_OR_EQUAL                    = 19
	CESQLParserLexerLESS                                = 20
	CESQLParserLexerLESS_GREATER                        = 21
	CESQLParserLexerLESS_OR_EQUAL                       = 22
	CESQLParserLexerLIKE                                = 23
	CESQLParserLexerEXISTS                              = 24
	CESQLParserLexerIN                                  = 25
	CESQLParserLexerTRUE                                = 26
	CESQLParserLexerFALSE                               = 27
	CESQLParserLexerDQUOTED_STRING_LITERAL              = 28
	CESQLParserLexerSQUOTED_STRING_LITERAL              = 29
	CESQLParserLexerINTEGER_LITERAL                     = 30
	CESQLParserLexerIDENTIFIER                          = 31
	CESQLParserLexerIDENTIFIER_WITH_NUMBER              = 32
	CESQLParserLexerFUNCTION_IDENTIFIER_WITH_UNDERSCORE = 33
)
