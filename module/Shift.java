package module;
public class Shift
{
	public static int[] up(int[] suu, int up)
	{
		for(int n = 0; n < up; n++){
			for(int m = 0; m < suu.length; m++){
				if(m >= 1 && m < suu.length){
					suu[m - 1] = suu[m];
				}
			}
			suu[suu.length - 1 - n] = 0;
		}
		return suu;
	}
	public static int[] down(int[] suu, int down)
	{
		for(int n = 0; n < down; n++){
			for(int m = suu.length - 1; m >= 0; m--){
				if(m >= 0 && m < suu.length - 1){
					suu[m + 1] = suu[m];
				}
			}
			suu[n] = 0;
		}
		return suu;
	}
	//削除　配列、要素番号
	public static int[] del(int[] suu, int iti)
	{
		int[] kekka = new int[suu.length - 1];
		int m;
		for(int n = 0; n < kekka.length; n++){
			m = n;
			if(n >= iti)m = n + 1;
			kekka[n] = suu[m];
		}
		return kekka;
	}
	public static double[] del(double[] suu, int iti)
	{
		double[] kekka = new double[suu.length - 1];
		int m;
		for(int n = 0; n < kekka.length; n++){
			m = n;
			if(n >= iti)m = n + 1;
			kekka[n] = suu[m];
		}
		return kekka;
	}
	//要素数で返す　配列、要素番号
	public static int[] delnum(int[] suu, int iti)
	{
		int[] kekka = new int[suu.length - 1];
		int m;
		for(int n = 0; n < kekka.length; n++){
			m = n;
			if(n >= iti)m = n + 1;
			kekka[n] = m;
		}
		return kekka;
	}
	public static int[] delnum(double[] suu, int iti)
	{
		int[] kekka = new int[suu.length - 1];
		int m;
		for(int n = 0; n < kekka.length; n++){
			m = n;
			if(n >= iti)m = n + 1;
			kekka[n] = m;
		}
		return kekka;
	}
	public static int[] delnum(String[] suu, int iti)
	{
		int[] kekka = new int[suu.length - 1];
		int m;
		for(int n = 0; n < kekka.length; n++){
			m = n;
			if(n >= iti)m = n + 1;
			kekka[n] = m;
		}
		return kekka;
	}
	
}