package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Asos struct {
	tm      string
	stn     int
	help    int
	authKey string
}

// Functional Option Pattern
// The ellipsis `...` before the type name `func(*Asos)` indicates that this function, `New` can take any number of arguments of that type.
// Because of the cariadic parameter, `options` inside the function will be a slice of functions (`[]func(*Asos)`).
func New(options ...func(*Asos)) *Asos {
	asos := &Asos{}
	for _, o := range options {
		o(asos)
	}
	return asos
}

func WithTm(tm string) func(*Asos) {
	return func(asos *Asos) {
		asos.tm = tm
	}
}

func WithStn(stn int) func(*Asos) {
	return func(asos *Asos) {
		asos.stn = stn
	}
}

func WithHelp(help int) func(*Asos) {
	return func(asos *Asos) {
		asos.help = help
	}
}

func WithAuthKey() func(*Asos) {
	authKey := os.Getenv("SECRET_APIKEY")
	// path := os.Getenv("SECRET_APIKEY_PATH")
	// if path == "" {
	// 	fmt.Println("The SECRET_APIKEY_PATH variable is not set.")
	// 	return nil
	// }

	// file, err := os.Open(path)
	// if err != nil {
	// 	fmt.Println("Error opening The SECRET_APIKEY_PATH secret file:", err)
	// 	return nil
	// }

	// Reader := bufio.NewReader(file)
	// fmt.Fscan(Reader, &authKey)
	// defer file.Close()

	return func(asos *Asos) {
		asos.authKey = authKey
	}
}

func AsosGetRequest(asos *Asos) ([]byte, error) {
	request := fmt.Sprintf(
		"https://apihub.kma.go.kr/api/typ01/url/kma_sfctm2.php?tm=%v&stn=%v&help=%v&authKey=%v",
		asos.tm,
		asos.stn,
		asos.help,
		asos.authKey,
	)

	resp, err := http.Get(request)

	if err != nil {
		fmt.Printf("Error sending request to API endpoint. %+v\n", err)
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response. %+v\n", err)
		return nil, err
	}

	return responseBody, nil
}

// #START7777
// #--------------------------------------------------------------------------------------------------
// #  기상청 지상관측 시간자료 [입력인수형태][예] ?tm=201007151200&stn=0&help=1
// #--------------------------------------------------------------------------------------------------
// #  1. TM     : 관측시각 (KST)
// #  2. STN    : 국내 지점번호
// #  3. WD     : 풍향 (16방위)
// #  4. WS     : 풍속 (m/s)
// #  5. GST_WD : 돌풍향 (16방위)
// #  6. GST_WS : 돌풍속 (m/s)
// #  7. GST_TM : 돌풍속이 관측된 시각 (시분)
// #  8. PA     : 현지기압 (hPa)
// #  9. PS     : 해면기압 (hPa)
// # 10. PT     : 기압변화경향 (Code 0200)
// # 11. PR     : 기압변화량 (hPa)
// # 12. TA     : 기온 (C)
// # 13. TD     : 이슬점온도 (C)
// # 14. HM     : 상대습도 (%)
// # 15. PV     : 수증기압 (hPa)
// # 16. RN     : 강수량 (mm) : 여름철에는 1시간강수량, 겨울철에는 3시간강수량
// # 17. RN_DAY : 일강수량 (mm) : 해당시간까지 관측된 양(통계표)
// # 18. RN_JUN : 일강수량 (mm) : 해당시간까지 관측된 양을 전문으로 입력한 값(전문)
// # 19. RN_INT : 강수강도 (mm/h) : 관측하는 곳이 별로 없음
// # 20. SD_HR3 : 3시간 신적설 (cm) : 3시간 동안 내린 신적설의 높이
// # 21. SD_DAY : 일 신적설 (cm) : 00시00분부터 위 관측시간까지 내린 신적설의 높이
// # 22. SD_TOT : 적설 (cm) : 치우지 않고 그냥 계속 쌓이도록 놔눈 경우의 적설의 높이
// # 23. WC     : GTS 현재일기 (Code 4677)
// # 24. WP     : GTS 과거일기 (Code 4561) .. 3(황사),4(안개),5(가랑비),6(비),7(눈),8(소나기),9(뇌전)
// # 25. WW     : 국내식 일기코드 (문자열 22개) : 2자리씩 11개까지 기록 가능 (코드는 기상자원과 문의)
// # 26. CA_TOT : 전운량 (1/10)
// # 27. CA_MID : 중하층운량 (1/10)
// # 28. CH_MIN : 최저운고 (100m)
// # 29. CT     : 운형 (문자열 8개) : 2자리 코드로 4개까지 기록 가능
// # 30. CT_TOP : GTS 상층운형 (Code 0509)
// # 31. CT_MID : GTS 중층운형 (Code 0515)
// # 32. CT_LOW : GTS 하층운형 (Code 0513)
// # 33. VS     : 시정 (10m)
// # 34. SS     : 일조 (hr)
// # 35. SI     : 일사 (MJ/m2)
// # 36. ST_GD  : 지면상태 코드 (코드는 기상자원과 문의)
// # 37. TS     : 지면온도 (C)
// # 38. TE_005 : 5cm 지중온도 (C)
// # 39. TE_01  : 10cm 지중온도 (C)
// # 40. TE_02  : 20cm 지중온도 (C)
// # 41. TE_03  : 30cm 지중온도 (C)
// # 42. ST_SEA : 해면상태 코드 (코드는 기상자원과 문의)
// # 43. WH     : 파고 (m) : 해안관측소에서 목측한 값
// # 44. BF     : Beaufart 최대풍력(GTS코드)
// # 45. IR     : 강수자료 유무 (Code 1819) .. 1(Sec1에 포함), 2(Sec3에 포함), 3(무강수), 4(결측)
// # 46. IX     : 유인관측/무인관측 및 일기 포함여부 (code 1860) .. 1,2,3(유인) 4,5,6(무인) / 1,4(포함), 2,5(생략), 3,6(결측)
