package module;
import java.util.*;
public class Calculate
{
//‡Œv‚ğ‹‚ß‚é
	public static int total(int[] suu)
	{
		int goukei = 0;
		for(int n = 0; n < suu.length; n++)goukei += suu[n];
		return goukei;
	}
	public static double total(double[] suu)
	{
		double goukei = 0.0;
		for(int n = 0; n < suu.length; n++)goukei += suu[n];
		return goukei;
	}
	//•½‹Ï‚ğ‹‚ß‚é ˆø”‚ğ‘‚â‚·‚±‚Æ‚Ådouble‚Å•Ô‚·
	public static int average(int[] suu)
	{
		int goukei = 0;
		for(int n = 0; n < suu.length; n++)goukei += suu[n];
		int heikin = goukei / suu.length;
		return heikin;
	}
	public static double average(int[] suu, int aaa)
	{
		int goukei = 0;
		for(int n = 0; n < suu.length; n++)goukei += suu[n];
		double heikin = goukei * 1.0 / suu.length;
		return heikin;
	}
	public static double average(double[] suu)
	{
		double goukei = 0.0;
		for(int n = 0; n < suu.length; n++)goukei += suu[n];
		double heikin = goukei / suu.length;
		return heikin;
	}
	//‘S—v‘f‚É‘Î‚·‚él‘¥‰‰Z
	
	//‘Î”
	public static int taisuu(int ki, int si)
	{
		int tai = 0;
		for(int j = si; j >= 1; j /= ki)tai++;
		return tai - 1;
	}
	//‚×‚«æ
	public static int bekijou(int ki, int si)
	{
		int jou = 1;
		for(int j = 0; j < si; j++)jou *= ki;
		return jou;
	}
	
	//‘fˆö”•ª‰ğ
	public static ArrayList<String> factorization(int number){
		ArrayList<String>list01 = new ArrayList<String>();
		int limit = number / 2;
		for(int n = 2; n <= limit; n++){
			int count = 0;
			while(number % n == 0){
				number /= n;
				count++;
			}
			if(count != 0)list01.add(n + "^" + count);
			if(list01.size() == 0 && Math.sqrt(limit * 2) < n)break;
		}
		return list01;
	}
}