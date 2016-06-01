package module;

import java.util.*;

public class Shellsort
{
	//シェルソート
	public static int[] up(int[] suu)
	{
		int jou, h, work, a;
		//間隔作成
		for(a = 0; suu.length - 1 > ((int)Math.pow(3, a) - 1) / 2; a++);
		while(a > 0){
			//間隔作成
			jou = (int)Math.pow(3, --a);
			h = (jou - 1) / 2;
			//ソート
			for(int n = 0; n < h; n++){
				for(int k = h + n; k < suu.length; k += h){
					for(int m = k; m >= h ; m -= h){
						if(suu[m] < suu[m - h]){
							work = suu[m];
							suu[m] = suu[m - h];
							suu[m - h] = work;
						}else{
							break;
						}
					}
				}
			}
		}
	return suu;
	}
	public static int[] down(int[] suu)
	{
		int jou, h, work, a;
		//間隔作成
		for(a = 0; suu.length - 1 > ((int)Math.pow(3, a) - 1) / 2; a++);
		while(a > 0){
			//間隔作成
			jou = (int)Math.pow(3, --a);
			h = (jou - 1) / 2;
			//ソート
			for(int n = 0; n < h; n++){
				for(int k = h + n; k < suu.length; k += h){
					for(int m = k; m >= h ; m -= h){
						if(suu[m] > suu[m - h]){
							work = suu[m];
							suu[m] = suu[m - h];
							suu[m - h] = work;
						}else{
							break;
						}
					}
				}
			}
		}
	return suu;
	}
}
