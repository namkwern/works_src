package com.example.a1503.car_logger.amodule;

/**
 * Created by a1503 on 2016/05/06.
 */
public class Formatting 
{
	//小数点以下の文字数を第二引数の値の数に合わせる
	public static String toString(float f, int i) {
		return toString((double)f, i);
	}
	public static String toString(double d, int i){
		String s = String.valueOf(d);
		return toString(s, i);
	}
	public static String toString(String s, int i){
		int index = s.indexOf(".");
		if(index != -1 && s.length() > index + i)
			return s.substring(0, index + i + 1);
		else
			return s;
	}
}
