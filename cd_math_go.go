package cd_math_go

//some math calculations

import (
	"fmt"
	"math"
	"strconv"
)

/*
//округляет до 9 цифр после запятой
    public static double round_To_9 (double in_Number) {

        return (double) Math.round(in_Number * 1_000_000_000.0) / 1_000_000_000.0;

    }

    //округляет до places цифр после запятой
    //просто обрезает не округляя ни вверх ни вниз
    public static double round_To(double in_Number, int places) {

        double scale = Math.pow(10, places);

        return (double) (long)(in_Number * scale) / scale;

    }

*/

//https://docs.oracle.com/javase/8/docs/api/java/lang/Math.html#atan2-double-double-
//https://www.geeksforgeeks.org/java-lang-math-atan2-java/
//https://en.wikipedia.org/wiki/Atan2
//стандартные atan2 в Java и Javascript возвращают результат в диапазоне
//-Pi +Pi (-180 +180) градусов
//так как это будет наша Longitude,  необходимо скорректировать
//в диапазон 0..360 градусов

// тесты у стандартной atan2 выиграл, потому что не округляет
// на сегодня самая корректная
func Atn2RAD(y float64, x float64) float64 {

	var phi float64

	// азимут вектора
	if x == 0 && y == 0 {
		return 0
	}

	var absY float64 = math.Abs(y)
	var absX float64 = math.Abs(x)

	//console.log(`abs_y =  ${abs_y}, abs_x =  ${abs_x},`);

	//https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/atan
	// The Math.atan() method returns a numeric value between - π/2  and  π/2  radians. (+90 and -90 degrees)
	// The angle that the line [(0,0);(x,y)] forms with the x-axis in a Cartesian coordinate system - Math.atan(y / x);
	if absX > absY {
		//arctan
		phi = math.Atan(absY / absX)
	} else {
		//arctan
		phi = math.Pi/2 - math.Atan(absX/absY)
	}

	if x < 0 {
		phi = math.Pi - phi
	}
	if y < 0 {
		phi = -phi
	}

	//console.log(`return value from atn2 =  ${phi}`);
	return phi
}

// на сегодня самая корректная
func Atn2RADWith360Check(y float64, x float64) float64 {

	phi := Atn2RAD(y, x)

	//если меньше 0 или больше 360 градусов, корректируем
	if phi < 0 {
		phi = phi + 2*math.Pi
	}
	if phi > 2*math.Pi {
		phi = phi - 2*math.Pi
	}

	return phi
}

func Atn2RADWith90Check(y float64, x float64) float64 {
	theta := Atn2RAD(y, x)

	//не может быть больше 90 или меньше -90 градусов
	if theta > math.Pi/2 {
		fmt.Printf("\n!!!atn2_RAD_with_90_check has Problems with +-90degrees %05.5f", theta)
	}
	if theta < -math.Pi/2 {
		fmt.Printf("\n!!!atn2_RAD_with_90_check has Problems with +-90degrees %05.5f", theta)
	}

	return theta

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// убирает минус или значения больше 360
// если угол должен быть от 0 до 360
// все в Радианах
func Convert_to_0_360_RAD(longitude float64) float64 {

	coeff := Abs(int(longitude / (2 * math.Pi)))

	if longitude < 0 {
		return longitude + float64(coeff)*2*math.Pi + 2*math.Pi
	} else {
		return longitude - float64(coeff)*2*math.Pi
	}

}

// убирает минус или значения больше 360
// если угол должен быть от 0 до 360
// все в Градусах
func Convert_to_0_360_DEG(longitude float64) float64 {

	coeff := Abs(int(longitude / 360))

	if longitude < 0 {
		return longitude + float64(coeff*360+360)
	} else {
		return longitude - float64(coeff*360)
	}

}

// prec controls the number of digits (excluding the exponent)
//
//	prec of -1 uses the smallest number of digits
func TruncFloat(f float64, prec int) float64 {
	floatBits := 64

	if math.IsNaN(f) || math.IsInf(f, 1) || math.IsInf(f, -1) {
		fmt.Println("error in TruncFloat")
		return 0
	}

	fTruncStr := strconv.FormatFloat(f, 'f', prec+1, floatBits)
	fTruncStr = fTruncStr[:len(fTruncStr)-1]
	fTrunc, err := strconv.ParseFloat(fTruncStr, floatBits)
	if err != nil {
		fmt.Println("error in TruncFloat")
		return 0

	}

	return fTrunc
}
