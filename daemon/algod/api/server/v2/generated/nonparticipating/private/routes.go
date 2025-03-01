// Package private provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameterWithLocation("simple", false, "catchpoint", runtime.ParamLocationPath, ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameterWithLocation("simple", false, "catchpoint", runtime.ParamLocationPath, ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE(baseURL+"/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST(baseURL+"/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.POST(baseURL+"/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+x9/XMbN7Lgv4LivirHPo7kz+xaV1vvFDvJ6uLELkvJ3nu2LwFnmiRWQ2ACYCQyPv3v",
	"V2gAM5gZgBxKjJ3Uez/Z4uCj0Wg0uhv98XGSi1UlOHCtJicfJxWVdAUaJP5F81zUXGesMH8VoHLJKs0E",
	"n5z4b0RpyfhiMp0w82tF9XIynXC6graN6T+dSPi1ZhKKyYmWNUwnKl/CipqB9aYyrZuR1tlCZG6IUzvE",
	"2cvJzZYPtCgkKDWE8jUvN4TxvKwLIFpSrmhuPilyzfSS6CVTxHUmjBPBgYg50ctOYzJnUBbqyC/y1xrk",
	"Jlilmzy9pJsWxEyKEoZwvhCrGePgoYIGqGZDiBakgDk2WlJNzAwGVt9QC6KAynxJ5kLuANUCEcILvF5N",
	"Tt5NFPACJO5WDuwK/zuXAL9BpqlcgJ58mMYWN9cgM81WkaWdOexLUHWpFcG2uMYFuwJOTK8j8n2tNJkB",
	"oZy8/eYFefLkyXOzkBXVGgpHZMlVtbOHa7LdJyeTgmrwn4e0RsuFkJQXWdP+7TcvcP5zt8CxrahSED8s",
	"p+YLOXuZWoDvGCEhxjUscB861G96RA5F+/MM5kLCyD2xjQ+6KeH8n3VXcqrzZSUY15F9IfiV2M9RHhZ0",
	"38bDGgA67SuDKWkGffcwe/7h46Ppo4c3f3l3mv2n+/PZk5uRy3/RjLsDA9GGeS0l8HyTLSRQPC1Lyof4",
	"eOvoQS1FXRZkSa9w8+kKWb3rS0xfyzqvaFkbOmG5FKflQihCHRkVMKd1qYmfmNS8NGzKjOaonTBFKimu",
	"WAHF1HDf6yXLlySnyg6B7cg1K0tDg7WCIkVr8dVtOUw3IUoMXLfCBy7oj4uMdl07MAFr5AZZXgoFmRY7",
	"rid/41BekPBCae8qtd9lRS6WQHBy88Fetog7bmi6LDdE474WhCpCib+apoTNyUbU5Bo3p2SX2N+txmBt",
	"RQzScHM696g5vCn0DZARQd5MiBIoR+T5czdEGZ+zRS1Bkesl6KW78ySoSnAFRMz+Bbk22/6/z1//QIQk",
	"34NSdAFvaH5JgOeigOKInM0JFzogDUdLiEPTM7UOB1fskv+XEoYmVmpR0fwyfqOXbMUiq/qertmqXhFe",
	"r2YgzZb6K0QLIkHXkqcAsiPuIMUVXQ8nvZA1z3H/22k7spyhNqaqkm4QYSu6/vvDqQNHEVqWpAJeML4g",
	"es2TcpyZezd4mRQ1L0aIOdrsaXCxqgpyNmdQkGaULZC4aXbBw/h+8LTCVwCOHyQJTjPLDnA4rCM0Y063",
	"+UIquoCAZI7Ij4654VctLoE3hE5mG/xUSbhiolZNpwSMOPV2CZwLDVklYc4iNHbu0GEYjG3jOPDKyUC5",
	"4JoyDoVhzgi00GCZVRKmYMLt+s7wFp9RBV8+Td3x7deRuz8X/V3fuuOjdhsbZfZIRq5O89Ud2Lhk1ek/",
	"Qj8M51ZskdmfBxvJFhfmtpmzEm+if5n982ioFTKBDiL83aTYglNdSzh5zx+Yv0hGzjXlBZWF+WVlf/q+",
	"LjU7ZwvzU2l/eiUWLD9niwQyG1ijChd2W9l/zHhxdqzXUb3ilRCXdRUuKO8orrMNOXuZ2mQ75r6Eedpo",
	"u6HicbH2ysi+PfS62cgEkEncVdQ0vISNBAMtzef4z3qO9ETn8jfzT1WVpreu5jHUGjp2VzKaD5xZ4bSq",
	"SpZTg8S37rP5apgAWEWCti2O8UI9+RiAWElRgdTMDkqrKitFTstMaapxpH+TMJ+cTP5y3Npfjm13dRxM",
	"/sr0OsdORmS1YlBGq2qPMd4Y0UdtYRaGQeMnZBOW7aHQxLjdRENKzLDgEq4o10etytLhB80BfudmavFt",
	"pR2L754KlkQ4sQ1noKwEbBveUyRAPUG0EkQrCqSLUsyaH744raoWg/j9tKosPlB6BIaCGayZ0uo+Lp+2",
	"Jymc5+zlEfk2HBtFccHLjbkcrKhh7oa5u7XcLdbYltwa2hHvKYLbKeSR2RqPBiPmH4LiUK1YitJIPTtp",
	"xTT+h2sbkpn5fVTnPweJhbhNExcqWg5zVsfBXwLl5ose5QwJx5l7jshpv+/tyMaMEieYW9HK1v20427B",
	"Y4PCa0krC6D7Yu9SxlFJs40srHfkpiMZXRTm4AwHtIZQ3fqs7TwPUUiQFHowfFWK/PIfVC0PcOZnfqzh",
	"8cNpyBJoAZIsqVoeTWJSRni82tHGHDHTEBV8MgumOmqWeKjl7VhaQTUNlubgjYslFvXYD5keyIju8hr/",
	"Q0tiPpuzbVi/HfaIXCADU/Y4u0eGwmj7VkGwM5kGaIUQZGUVfGK07r2gfNFOHt+nUXv0tbUpuB1yi2h2",
	"6GLNCnWobcLBUnsVCqhnL61Gp2GlIlpbsyoqJd3E127nGoOAC1GREq6g7INgWRaOZhEi1gfnC1+JdQym",
	"r8R6wBPEGg6yE2YclKs9dnfA99JBJuRuzOPYY5BuFmhkeYXsgYcikJmltVafzoS8HTvu8VlOWhs8oWbU",
	"4Daa9pCETesqc2czYsezDXoDtc+e27lof/gYxjpYONf0d8CCMqMeAgvdgQ6NBbGqWAkHIP1l9BacUQVP",
	"HpPzf5w+e/T458fPvjQkWUmxkHRFZhsNinzhlFWi9KaE+8OVobpYlzo++pdPveW2O25sHCVqmcOKVsOh",
	"rEXYyoS2GTHthljrohlX3QA4iiOCudos2ol97DCgvWTKiJyr2UE2I4Wwop2lIA6SAnYS077La6fZhEuU",
	"G1kfQrcHKYWMXl2VFFrkosyuQComIs9Lb1wL4lp4eb/q/26hJddUETM32sJrjhJWhLL0mo/n+3boizVv",
	"cbOV89v1Rlbn5h2zL13ke9OqIhXITK85KWBWLzqq4VyKFaGkwI54R38L2sotbAXnmq6q1/P5YXRngQNF",
	"dFi2AmVmIraFkRoU5IJb15Ad6qobdQx6+ojxNkudBsBh5HzDczS8HuLYpjX5FeP4CqQ2PA/UegNjCcWi",
	"Q5Z3V99T6LBT3VMRcAw6XuFntPy8hFLTb4S8aMW+b6Woq4MLef05xy6HusU421Jh+nqjAuOLsuuOtDCw",
	"H8XW+FkW9MIfX7cGhB4p8hVbLHWgZ72RQswPD2Nslhig+MFqqaXpM9RVfxCFYSa6VgcQwdrBWg5n6Dbk",
	"a3Qmak0o4aIA3PxaxYWzhAMLvpzjg78O5T29tIrnDAx15bQ2q60rgs/Zg/ui7ZjR3J7QDFGjEo95zSus",
	"bWWns84RpQRabMgMgBMxcy9m7i0PF0nxLV578caJhhF+0YGrkiIHpaDInKVuJ2i+nb069BY8IeAIcDML",
	"UYLMqbwzsJdXO+G8hE2GniOKfPHdT+r+Z4BXC03LHYjFNjH0NnYP9yw6hHrc9NsIrj95SHZUAvH3CtEC",
	"pdkSNKRQuBdOkvvXh2iwi3dHyxVIfKD8XSneT3I3AmpA/Z3p/a7Q1lXCH9Kpt0bCMxvGKRdesIoNVlKl",
	"s11s2TTq6OBmBQEnjHFiHDgheL2iSttHdcYLtAXa6wTnsUKYmSINcFINMSP/5DWQ4di5uQe5qlWjjqi6",
	"qoTUUMTWwGG9Za4fYN3MJebB2I3OowWpFewaOYWlYHyHLLsSiyCqm7cn53UyXBy+0Jh7fhNFZQeIFhHb",
	"ADn3rQLshj5hCUCYahFtCYepHuU0jmjTidKiqgy30FnNm34pNJ3b1qf6x7btkLiobu/tQoBCVzTX3kF+",
	"bTFrvQGXVBEHB1nRSyN7oBnEvv4PYTaHMVOM55Bto3xU8Uyr8AjsPKR1tZC0gKyAkm6Gg/5oPxP7edsA",
	"uOOtuis0ZNatK77pLSV7L5otQwscT8WER4JfSG6OoFEFWgJxvXeMXACOHWNOjo7uNUPhXNEt8uPhsu1W",
	"R0bE2/BKaLPjjh4QZMfRxwCcwEMz9O1RgZ2zVvfsT/EfoNwEjRyx/yQbUKkltOPvtYCEDdV5zAfnpcfe",
	"exw4yjaTbGwHH0kd2YRB9w2VmuWsQl3nO9gcXPXrTxB9dyUFaMpKKEjwwaqBVdifWIek/pi3UwVH2d6G",
	"4A+Mb5HllEyhyNMF/hI2qHO/sZ6uganjELpsZFRzP1FOEFDvP2dE8LAJrGmuy40R1PQSNuQaJBBVz1ZM",
	"a+vB3lV1taiycIDou8aWGd2rZvRNcesz6zkOFSxvuBXTidUJtsN30VMMOuhwukAlRDnCQjZARhSCUQ4w",
	"pBJm15lzpvfu1J6SOkA6po1P2s31f0910IwrIP8hapJTjipXraGRaYREQQEFSDODEcGaOZ2rS4shKGEF",
	"VpPELw8e9Bf+4IHbc6bIHK59BIpp2EfHgwdox3kjlO4crgPYQ81xO4tcH/jgYy4+p4X0ecpuVws38pid",
	"fNMbvHklMmdKKUe4Zvl3ZgC9k7kes/aQRsa5meC4o95yOk/2w3Xjvp+zVV1SfYhXK7iiZSauQEpWwE5O",
	"7iZmgn99RcvXTTeMroHc0GgOWY4xISPHggvTx4aR7NINW/c6tlpBwaiGckMqCTnYsAcj8qkGxiNiHSLz",
	"JeULlPSlqBfOI8+Og5y6VtamIms+GCIqDek1z9A6HePczgvbR74YOQio0cX6pm2reVzTZj4X7DTmSg2Q",
	"1zf1R1+3ppOkqmqQetWqqhY53fCdEVy8I6gF+GknHvkGgqgzQssQX+G2mFNgNvf3sbW3Q8egHE4c+Ai2",
	"H1NugkZPLjcHkFbsQERCJUHh3RLal5T9KuZhqJ67fNRGaVgNTfC268+J4/c2qegJXjIO2Upw2ESj0xmH",
	"7/Fj9Djh/ZbojJJGqm9feejA3wOrO88YarwrfnG3+ye0/9SkvhHyUG+ZdsDRcvmIp8Od7+Ruyts+cNKy",
	"jLwJukCePgNQ0yZxAJOEKiVyhsLWWaGm9qC5Z0QX9dNF/5vGPfkAZ68/bu/xK4wRReMulBWhJC8Zmn4F",
	"V1rWuX7PKRqXgqVGvJa8Fp02N77wTeL2zYj50Q31nlP0WGtMTlFPizlE7CvfAHiro6oXC1C6p6TMAd5z",
	"14pxUnOmca6VOS6ZPS8VSHQdOrItV3RD5oYmtCC/gRRkVuuu2I5xakqzsnQvcWYaIubvOdWkBKo0+Z7x",
	"izUO51/r/ZHloK+FvGywEL/dF8BBMZXFvau+tV/RE9gtf+m8gjGvgP3svSzbwNmJWWYnVv7/fvHvJ+9O",
	"s/+k2W8Ps+f/4/jDx6c39x8Mfnx88/e//7/uT09u/n7/3/8ttlMe9lgUlYP87KVTac9eot7SPt4MYP9k",
	"hvsV41mUyEI3jB5tkS8wYtgR0P2uVUsv4T3Xa24I6YqWrDC85Tbk0L9hBmfRno4e1XQ2omfF8mvdUxu4",
	"A5chESbTY423lqKGDonxeEV8TXQhiHhe5jW3W+mlbxuO4x3DxHzaxKTadDUnBAMWl9R7Nbo/Hz/7cjJt",
	"Aw2b75PpxH39EKFkVqxj4aQFrGNKnjsgeDDuKVLRjQId5x4Ie9QHzjplhMOuYDUDqZas+vScQmk2i3M4",
	"H+TgjEVrfsatR7s5P/g2uXFPHmL+6eHWEqCASi9jaSw6ghq2ancToOcvUklxBXxK2BEc9Y01hdEXnTde",
	"CXSO6RRQ+xRjtKHmHFhC81QRYD1cyCiLSIx+ev787vJXB1eH3MAxuPpzNg+R/m8tyL1vv74gx45hqns2",
	"stkOHcSiRlRpF27V8SQy3Mwm77FC3nv+nr+EOePMfD95zwuq6fGMKpar41qB/IqWlOdwtBDkxEdwvaSa",
	"vucDSSuZXyuInSNVPStZTi5DhaQlT5szZTjC+/fvaLkQ799/GDhVDNUHN1WUv9gJMiMIi1pnLuNDJuGa",
	"ytijlWoi/nFkm9Jl26xWyBa1tWz6jBJu/DjPo1Wl+pG/w+VXVWmWH5ChcnGtZsuI0kJ6WcQIKBYa3N8f",
	"hLsYJL32dpVagSK/rGj1jnH9gWTv64cPnwDphML+4q58Q5ObCkZbV5KRyX2jCi7cqpWw1pJmFV3E3sbe",
	"v3+ngVa4+ygvr9DGUZYEu3VCcL1HPQ7VLsDjI70BFo69wwlxcee2l8/uFV8CfsItxDZG3Ghf7G+7X0FQ",
	"7q23qxfYO9ilWi8zc7ajq1KGxP3ONEl/FkbI8m4Uii1QW3X5kWZA8iXkly5xDawqvZl2untPHSdoetbB",
	"lE1pZEPqMKkGvizMgNRVQZ0oTvmmn91AgdbeH/gtXMLmQrQ5OfZJZ9CNrlepg4qUGkiXhljDY+vG6G++",
	"cwdDxb6qfJA6Rit6sjhp6ML3SR9kK/Ie4BDHiKIT/Z1CBJURRFjiT6DgFgs1492J9GPLM1rGzN58kfRG",
	"nvcT16RVnpznVrgatLrb7yvA/GjiWpEZNXK7cKm9bAR5wMVqRReQkJDDx52RcdqdByEcZNe9F73pxLx/",
	"oQ3umyjItnFm1hylFDBfDKmgMtPz1/Mz2fdD9zKBGTsdwmYlikmNY6NlOlR2HtlsCsIUaHECBslbgcOD",
	"0cVIKNksqfJZxzA5mz/Lo2SA3zEjwrY8OGeBq1mQga3JcuN5bv+cDrRLlw3Hp8DxeW9C1XJEDhsj4aN3",
	"e2w7BEcBqIASFnbhtrEnlDY7Q7tBBo7X83nJOJAs5rUWmEGDa8bNAUY+fkCItcCT0SPEyDgAG9/FcWDy",
	"gwjPJl/sAyR32SWoHxtf1IO/IR73Zf24jcgjKsPCWeJVK/ccgDpXx+b+6jnc4jCE8SkxbO6KlobNOY2v",
	"HWSQjgXF1l7yFeeZcT8lzm55ALEXy15rslfRbVYTykwe6LhAtwXimVhnNvAzKvHO1jND71HXdgxDjR1M",
	"m/jmniIzsUZvH7xarCv1DljScHgwAg1/zRTSK/ZL3eYWmG3TbpemYlSokGScOa8hl5Q4MWbqhASTIpcv",
	"glw2twKgZ+xoE0M75XenktoVT4aXeXurTdscbT5qKHb8U0couksJ/A2tME32mTd9iSVqp+g6rXQT7wQi",
	"ZIzoDZsYPtIMn4IUlIBKQdYRorLL2Mup0W0Ab5xz3y0wXmB6H8o39wNPKAkLpjS0RnTvJ/E5zJMUswoK",
	"MU+vTldybtb3VojmmrLPiNixs8xPvgJ0JZ4zqXSGLxDRJZhG3yhUqr8xTeOyUtfXyubgZUWcN+C0l7DJ",
	"ClbWcXp183730kz7Q8MSVT1Dfsu4dViZYc7oqAfmlqmtk+7WBb+yC35FD7becafBNDUTS0Mu3Tn+JOei",
	"x3m3sYMIAcaIY7hrSZRuYZBB5OyQOwZyU/DGf7TN+jo4TIUfe6fXjo/fTd1RdqToWgKDwdZVMHwmMmIJ",
	"00HK5WFIa+IM0KpixbpnC7WjJjVmupfBwyeq62EBd9cNtgMDKNK+hTlIiJoQmk/WO7oRl8JEhRjZ3UmF",
	"E9n0pPG/a0rzF2VTOSKY6BZGMJdaMr3Hre9lJ/VidymR2gXDWWvG9ZdPhxTZ2PgNLGN24zxuWj83ikYX",
	"8YG6ZVOZ79gEllDcQ/IM2HM4FVO+EMeQbJsYyF2UewG0/A42P5m2uJzJzXRyN0N2jPLdiDtw/aY5bFE8",
	"o6OENWx23qX2RDmtKimuaJk5c3+KUUhx5RgFNvevA5/44olT9sXXp6/eOPBvppO8BCqzRnBLrgrbVX+a",
	"VdlklIkD4hP9Gw3ca1BWsA82v8mgFz4RXC/BZUwPdINBatf2+Sc4iu7JYB7319rJ+9xLlV3ilhcrqJoH",
	"q9aYat+rum9U9Iqy0lsxPbQJ3ypc3Lj8wFGuEA5w57eu4MkyOyi7GZzu+OloqWsHT8K5XmNKpLh0wl3C",
	"JGRF7u2qy4LuKUdZx7jq45lYt7fnyDv5GyE7zN851kffvvyF3WeMB7m7HR4Trka+Ckdf8DwiSEvkl8Uv",
	"5jQ+eBAetQcPpuSX0n0IAMTfZ+53NBY9eBA1S0a1DsMkUKngdAX3GyfB5EZ8WhWVw/W4C/r0aoWoQ1/v",
	"NBk2FGofsTy6rx32riVz+CzcLwWUYH7aHUDT23SL7hCYMSfoPOVI3/hIrGzhD0UE77sEYQyHIS1k9iuK",
	"qY2tlXd4hHi9QstopkqWx9+M+EwZ9sqtL4BpTLBxQrk2I9Ys4VrCaxaMZZqNydXVAzKYI4pMFU0X1uJu",
	"Jtzxrjn7tQbCCuDafJJ4r/WuOq8c4KgDgdToQsO53MD2xbEd/i46U5jWuy8zIhDbFabQ82AA7svGBOgX",
	"2ljYW51pXwemcMYB497ifOTow1GzdcZedj0IxukxYwrAeUbn8osn5ogWdGMqm0vxG8TtVmjuiwRg+kTm",
	"DL32foNQPQvLGHVYSmOtbuvStbPv2u7xunFq4++sC/tFN7nTb3OZxk/1fht5G6VXxdMEOiSnlLDw6aLr",
	"2ZZgLXi8Al8OTFvtnzUpt+fJRh92HKTjpzIMRTi247en0sE8CN8o6fWMxnJ6G13IwBRsb+cBVgviO/sN",
	"UE2Inp2dBA5ITVtmM5hUINsA9GE2tFvqNXba0RpNq8AgRYWqy9Q6jZRKRIap+TXlthaa6Wf5leutwL6Y",
	"mF7XQmL+IRV/Ky4gZytaxhWcIh++CxZswWyZr1pBUEfKDWRLKFoqcrW4msBTh5qzOXk4DYrZud0o2BVT",
	"bFYCtnhkW8yowuuyeb1oupjlAddLhc0fj2i+rHkhodBLZRGrBGl0TxTyGo+HGehrAE4eYrtHz8kX6Ouh",
	"2BXcN1h0QtDk5NFzfKmzfzyM3bKuTNs2ll0gz/6n49lxOkZnFzuGYZJu1KNoqhZbpzV9O2w5TbbrmLOE",
	"Ld2FsvssrSinC4i7F652wGT74m7i60sPL7ywRQaVlmJDmI7PD5oa/pQIWTLsz4JBcrFaMb1yHgFKrAw9",
	"tUWi7KR+OFux0OX393D5j+hYU3m/gp6t6xOrMXSVcDlG96cf6Aq6aJ0SapNOlax1efNVR8iZz2mHBQ+a",
	"OgcWN2Yus3SUJdEDbk4qybhG+0et59nfjFosaW7Y31EK3Gz25dNI4YBubm2+H+CfHO8SFMirOOplguy9",
	"zOL6ki+44NnKcJTifhsiGJzKpAdQ3Ncj5XCyfeixkq8ZJUuSW90hNxpw6jsRHt8y4B1JsVnPXvS498o+",
	"OWXWMk4etDY79OPbV07KWAkZS1TbHncncUjQksEVOnzHN8mMece9kOWoXbgL9J/3udqLnIFY5s9yVBHw",
	"RqdtgV5GhP/pe1eUeCB7J5zTrPdZ0+cTB7BFjZZWQuuYzR79QqTRJFEaffAAgX7wYOqEuV8edz9bJvXg",
	"QTx9W9RwZH5tsXAXvQ77xvbwKxEx4/haKc0TugtSi5jRUqzWfDBHeeaGmpJuXYpPfxcexv057uISPwXv",
	"37/DLx4P+EcfEZ/5yOMGtk58diUJQgnq8kRJpmi+B851lHwl1mMJp8dJPfH8AVCUQMlIIxOuZFB3KPro",
	"vNPrIaBRM+oMSmFUpTClemiV/vPg2Sx+ugXbNSuLn9oEG72LRFKeL6OuSTPT8ee2PnCzRMsqo1mal5Rz",
	"KKPDWQ3tZ6/JRXTNf4mx86wYH9m2X/fKLre3uBbwLpgeKD+hQS/TpZkgxGo3d0ETG1cuREFwnjYlcMsc",
	"hwXkgqo2v9agdOxo4Afrn49PNob52qIqBHiBNpwj8i1GERtYOvke0XbiE3J1k9PUVSloMcVEYRdfn74i",
	"dlbbx1a5tEVdFmg66K4iausdn6ynKVgZj0IdP872sDizaqWzpgZLLM+HadFWiWE9BwA0KoTYOSIvg2L+",
	"NiWIGYJgnji5giIo+WI1CqQJ8x+tab5EQ0nnIkuT/PhqRJ4qVVASvSlt2qQAx3Nn4HYFiWw9oikRegny",
	"minAuCO4gm5qkSbPjjPU+VQj3eXJmnNLKUd7yBRNwu990e6BswKJf+GMQtZD/J5qsi3mtW9xpnPsFc1I",
	"2q/0NKiFbhNVNCUrv/fV7CkXnOWYDzQmEGEahHFvJiNSp8YfO9TEndDI4YrWl2oiHhwWkxWnPCN0iBu+",
	"PwZfzaZa6rB/ali7ugML0MpxNiimvkyas84zrsCldDdEFPJJISMeFjGRI2tec/ckI4xwTphbvjHffnDG",
	"OAz9u2Qc1W6HNidmW/s5VrDXRldnmiwEKLeebpoX9c70OcKMJwWsPxz5ivc4hvXpMcu2DmzDoU69O5tz",
	"HzNtX5i2Lg9l83PHN8VOelpVbtJ0Eb145dA1TyI45kThX7UD5Dbjh6NtIbetfqh4nxpCgyt0oYEK7+EB",
	"YTQF5XrVW42KYCkKWxDrjR9NRsV4BIxXjPv3nPgFkUevBNwYPK+JfiqXVFsRcBRPuwBaNj4zfYamtHsQ",
	"vOtQ/SycBiW4Rj9HehvbWngJxtE0aAU3yjfEHwpD3YEw8YKWjR9npLIdSlVOiCowOLRX6y7GOAzj9tU0",
	"uxfAjgK607Y7pqTd9yZK5fuY1cUCdEaLIpZh/yv8SvArKWqUHGANed1kYq8qkmN6u26+vyG1uYlywVW9",
	"2jKXb3DH6YLikRFqCAtY+h3GeOLZBv/dp7Rx48G5d0SHd9cs9ktyOYxQiUm9hqYzxRbZeEzgnXJ3dLRT",
	"347Q2/4HpfRSLLqAfA4jaYLLhXsU429fm4sjTII1cJa1V0uTowodU4WvgY5qY5NdpcuV8CobJNvHJ9im",
	"pPB2M0S6OPAUL79EFFVo8rb3qzUDp2Kp8mToH9UuCYGmZCsLSgZ2W8fFnhF9+J6Rcla0voqHMz67tW5F",
	"qPcjHwL0nQ9SIRVlzmGlZRZDzDo332G45xg/2naD+4twIXtJ++h3V6nwOp/zFr/3i4degstMVEm4YqL2",
	"riDeIdOrhPbXTinOJsAxuv6om/PnNj4nTeUXroiTXabTyb/7ybrvEuBabv4AhvPBpg/Kkg6lXWueapuQ",
	"pv7HqHognVtxTD7oWOphJxt2CqPuKOs6IKuXY8SBYZnW6eSs2OvCjKWvnthRYscuXnQ1nd2zzeiJR6wS",
	"irVleGLVWEd6Pl9gQdUgO+lwLO8RdwW5xtpLraePBNgnV6mZLKjv/t9ZPhPqdOMg7pJ7bsvoOSy4tOOO",
	"HwTdB4kjbLGao/H5K08bf04bjnJNFWZ7tiXWuwGco8PI5nPINbvakeTgn0vgQQD91NtlEJZ5kPOANUEV",
	"mCNvf6tjC9C2HARb4QlyVd8ZnFRQ7SVs7inSoYZo9Zwmoug26dEQA8gdMkMiQsX8pawh2bmwMNVQBmLB",
	"+yfa7tAmmk0W3gxSdtxyLk+S5uJo03hsmTJe+W/UXKbrXsltMD4glQdhWDgsrX+8xDptqimK7dOrhVo6",
	"ORsmob526dkwJUXzduITtYHyv/n8M3aWkl1CWBoUX6quqSx8i6jpxVt1si330SB5gS961Qd63szMWm/y",
	"4Vt1JK0pBmbkpTBiRJaKbuk6cDfeT/eUdVOzVXbQNd3ANQfpSiij/FsKBZkW3vt8GxzbUGF98W6FBJVM",
	"JW6BSyb4e9tmMMSSChQT+lHnghcukEhYUQOdDPIMpufchuwX9ruPCPYp9XdamBp63V3byccRMDVAYkj1",
	"c+Juy92RxrcxNjHOQWb+5amfdJCD7L6GVFIUdW4v6PBgNAa50Sk9t7CSqJ0mH66ypyMEEbuXsDm2SpAv",
	"iuV3MATaSk4W9CBZVW+TD2p+UzG4FwcB73NarqaTSogySzx2nA0zJfYp/pLll1AQc1N4f9tEoULyBdrY",
	"m9fs6+XGZwasKuBQ3D8i5JTbCAf/sN0t1dGbnN/T2+Zf46xFbZOXOqPa0XsedxXHtKLyjtzMD7Odhykw",
	"rO6OU9lBduThWyeyNEp6HSnbeTRWKx8+NfdLKbZEZaGIySTn9sXqBR70mOEI47GDxAH4kEmJe+kiqhQx",
	"l8zbxIyboeKYCidDgDTwMaHLDRRu8CgCmjKJOxyFGh+htsJc6yc0FI/KUlxneIyyJs9sTOky7VT3mvCp",
	"9dt+ht5mEHgcUeVEiA1Z0oLkQkrIwx7xsCgL1UpIyEqBDkixt9G5NhLhCmMhOCnFgojKKPo2X7N/RYrW",
	"PxzMVXNO8UKHwN8jigKa56h9CuL6kKbP2CkPVV7SJj+xi87sK1vCJRKUS3biMGQbD+HdUuFx/+qRF8uI",
	"sQwx5wlk7xKRjsj3ruwWgDnicO02FJ7GKmB219WvxZqqjKzFiuVxdP+5XISSjj0x6o1mfbHFFWycLjZD",
	"nhLyseZFGE/PEM3A6ayM3g/u+LmXMaRz818UG/rjkjk4fpbgocMj7Vh/licvqB4ACKkNHtO1tBUZwuuj",
	"qfMqFjbYFN/1+oCOZDjoPnE32MwIBwdKw52AGrhsNQB+YTWmqc3OY92/ZmLtv99v0/fcCvib7VQeq2Ib",
	"OcUNabkiuz7UP8ERol4l2504bGXz2VhXjqZ6zkjmHwCQdu7owDDKxWNfMOaUlVBkNILks0axngbqgQsL",
	"6NdEY8px8pxaw9oSiBm7luBCz21J814N1YoaUhJN86H5ixewBoVx4bYQJFXWWOuNxq6eel+DEVVWwhV0",
	"fF5cPHyNUgi7grAWu+1MCoAKn1D6in3MmSO8y3vanlt7FrgDjMFuVP2ziLU7RXbodlFNdM0ze0zU2KNk",
	"ILpiRU07+FN3qEqdLkg9EB8zKybaAzFmmh/tCG/9AKe+f0yU8Zj4MI4P7c2C4qjbxoB2OnfhiYqeeh73",
	"7QqTPTRWYZytaF6PLIm3fENV9JqnrShDkm8l8fHV4gPEfr2GHKWarvPS3XFCcDCieolckiK4bHb49ta4",
	"z0LDW0k4OV5M1VCADLZVxlpbuV9HQxdhyXqsgsWN2GukZqw84fi/439TLNxrBzIqoC2EEVbmfwn+2QNz",
	"yzYWXyfQsuZC805aU5darK8/ssA9dUU3REj8hwtNfq1pyeYbPKEWfN+NqCU1JOTeWewDoHP6MhNvF0ym",
	"HjCvwgo/lV03GztmMNzGjBIAba5AIqQz2a/oJYTbgG+blvPk2rAcVc9WTCm87HrbOcSCW7wPD1/RAoJY",
	"EkxS1a1A5tMWmt7/sw19CafyuWWqkuZtRWFFVz2roi1t5IlLL2G1PTZqqB57EmjKJbVEK31MZGFTl1j8",
	"NXkKUBLB/8yYllRutnhq7nz+jjkco+S8C+xBGRkUww+2jH3qGrbhpVuiykYt5dC7MPaRfQA0vtT5BD87",
	"wLeJ2XwyoE+B/2j+uNQyxoD/R8F7ovpOCK8ttPMJsNyJm47Aak2AM7HOJMzVrvdkawM0irBsI669EwHj",
	"uQSq7AP72WunsrXp0Rg3KqR1AWueMJpRCpgz3jJLxqtutXvHrjFLGt8ECAstqYjWhMU8JSUYMeyKlq+v",
	"QEpWpDbOnA5b/SNMT+2tx65vRPlv7tThAEy12g+GY0Eb7hM0Mxd4weZzkNY7S2nKCyqLsDnjJAdp7n1y",
	"TTfq9mZ6A62sjXyxw1BPA2mmGyQcmOyRtC0g5ca9Ad3RiN4ASA9oTR9hBUc3wIgF3BpFtEgYvYcwxGPT",
	"6TorxQKDdBIE6PLQ4TOFVVYER4OtlYf2m0ex32D7NJiC1x18LXDWMVNsP2evEXWo8PzImd560qw1rR81",
	"Zd3a7EHw9M8XrW+t3Zwh/ccC3S5scf0w2K1fq9bvtX1jt/NBovZO14Kb2EV8ZXRRkqG5Vo1/yeg8ZMbC",
	"6awOm6Fuq7Z4z4IKqvvnzvthaPQZKMUWKVMXjLinTchakv09kADPFrhzZ6s7bfMibcYZL2sEz69xiCpR",
	"ZfkYlyqbpbtwBm0HaRfGBH0E5urEupvX57bmcic7RCeBvZWUbyPu9hLo73qXqfJtSnbKoJHgoF1juZgj",
	"L8MjbM046CjfGC+m/RCOrsGmYRKEEgl5LdGgeU03u0uMJLJDnv/j9Nmjxz8/fvYlMQ1IwRag2gyjvRId",
	"rdsN4307y6d1tBksT8c3wQf3WsT5lzIfs9BsijtrlttayY1HC5TsYwmNXACxUtTD0hC32iscp/Wc/WNt",
	"V2yRB9+xGAp+nz1z7oHxBZxyp7+IOdnOM9qHEX/cI/zCCP+RS8pv7S0WmLLHpoNLb0OPrUH2D0OFkWjZ",
	"g9Fes9zfg+KiUubtqu6NAm0YORkhDwQgERLVCWYJi3K2Sf+kte2iFdg/mPUvse/bh7SdvrsIie+wA7ww",
	"xqlt17ibOnA+c/a87xukBEv5kKKEzvJ3hU25BbYvj8EWOVVXa7Alkm0OoO6+BDFx6kUTapaQbQcRaViB",
	"0+g3ZRmJZLPaN56pkHCMYCmvaPnpuQaWZj1FfEDxNu2/HoYzhUi2qFS3S6b0io6aOwhdOtzU/A1Gz/0T",
	"zB5F7zk3lHt0HNxmaDuhpfU0nLtIZDMkucYxrVPJoy/JzKVnriTkTPUfM+2Lk4vFwugdkGzuQuFgrXeE",
	"C+1a509C34GM597zgPwQPEoINP60ELZH9DMzlcTJjVJ5jPoGZBHBX4xHheXcdlwXl52Y/FYWD240IeHA",
	"sflBlp09Y/OHherGLs/Gn5tLp1YwXOfo27qD28hF3a5tbGKJ0bmUscD+mHwQ8bzHpjsmpDhIAuS90h//",
	"DqkoLI7cGG7eGMX8lEpOaBPwJfJg9vajZuVON4NOVtOb6WQBHBRTmLfzZ5dt/NPepR4CGx47PKoW1rvE",
	"9FvERNbamTyYKshXOiJVqesWSUyKoSd5LZneYKU5b4ZhP0eTZnzbBGC7AP7mBcTdfVpcQlPtsw3XrpW/",
	"Xb8VtMT7yD7McHMLifKIfL2mq6p0RkXy93uzv8KTvz0tHj559NfZ3x4+e5jD02fPHz6kz5/SR8+fPILH",
	"f3v29CE8mn/5fPa4ePz08ezp46dfPnueP3n6aPb0y+d/vWf4kAHZAurT6J5M/k92Wi5EdvrmLLswwLY4",
	"oRX7DszeoK48F1gJySA1x5MIK8rKyYn/6X/5E3aUi1U7vP914jL6T5ZaV+rk+Pj6+voo7HK8wPjMTIs6",
	"Xx77ebA+TUdeeXPW+CRb7wnc0dYGiZvqSOEUv739+vyCnL45O2oJZnIyeXj08OiRK4bIacUmJ5Mn+BOe",
	"niXu+7EjtsnJx5vp5HgJtMR0BuaPFWjJcv9JAi027v/qmi4WII/Q7dz+dPX42IsVxx9dnOrNtm/H4cP8",
	"8cdOOG+xoyc+Kh9/9CXRtrfulMNy/jxBh5FQbGt2PMME8mObggoap5eCyoY6/ojicvL3Y2fziH9EtcWe",
	"h2Mf8x5v2cHSR702sO7osWZFsJKc6nxZV8cf8T9IvTeWnZQQi3+3iY0paZtPCdOEzoTEIlo6XxoO4qv3",
	"MBW0DGtqnhXmGJheLywEvhiirXZ/8m7ogI4DET8S8gxzINoj3Zmp5dr4wBmU/G7upE779mZ69zB7/uHj",
	"o+mjhzd/MTeP+/PZk5uRsRovmnHJeXOtjGz4AUvfoFcanvTHDx969uaUh4A0j91JDhY3UKLaRdpNapze",
	"hre+o4W0g7Hbqt5ApEHGjhIdveGHwgty9Kd7rnirpamTrQ2H72eTL4gP4cO5H326uc+4dbUzN4e94W6m",
	"k2efcvVn3JA8LQm2DGquDbf+R37JxTX3LY04Uq9WVG78MVYdpkDcZuOlRxcKH74ku6IoBXLBgxQ0fDH5",
	"gMHMsTDKBL9Rmt6C35ybXv/Nbz4Vv8FNOgS/6Q50YH7zeM8z/+df8X9tDvv04d8+HQQ+CvyCrUDU+s/K",
	"4c8tu70Th3cCp02xe6zX/Bhduo4/dsRn93kgPnd/b7uHLa5WogAv74r53FYf3vb5+KP9N5gI1hVItgJu",
	"ywC6X236wWMsQrcZ/rzhefTH4To6qdcSPx9/7PzZ1S/UstaFuLZ1ZKJXJlZNp6Wr/onG5EYx1YL4Adpc",
	"b+S1S09bbtCCzgogFOtmiFq3lgPrlOqC2pq3HTMCUUtnRF8wjhOgkR5nsWVuaeDyoyAXvEB9uHc9O8h+",
	"EAUMr2e8gH+tQW7aG9jBOJl2+LMj8EhR2Ttfd0N2erMf+eNjgn0JGxKH+Vir/t/H15Rpc4m7pGuI0WFn",
	"DbQ8dhUWer+2SY0HXzBTc/BjGJkX/fWYdqm9q6f70trRj30lPvbVKbGJRt4t1n9uDXqhgQzJpTGNvftg",
	"dh0rgjpKau09J8fHGCexFEofT26mH3u2oPDjh2ajfQmwZsNvPtz8/wAAAP//Jj/x5wbwAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
