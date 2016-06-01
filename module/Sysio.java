package module;
import java.io.*;
//１一文字、２配列、３配列番号指定
public class Sysio
{
	//■■引数と同じ型で返す
	public static int in(int suu) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str = br.readLine();
		suu = Integer.parseInt(str);
		return suu;
	}
	public static double in(double suu) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str = br.readLine();
		suu = Double.parseDouble(str);
		return suu;
	}
	String in(String suu) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str = br.readLine();
		return str;
	}
	//入力要求時に文字列を表示する
	public static int in(int suu, String moji) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		System.out.print(moji);
		String str = br.readLine();
		suu = Integer.parseInt(str);
		return suu;
	}
	public static double in(double suu, String moji) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		System.out.print(moji);
		String str = br.readLine();
		suu = Double.parseDouble(str);
		return suu;
	}
	public static String in(String suu, String moji) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		System.out.print(moji);
		String str = br.readLine();
		return str;
	}
	//■■配列
	public static int[] in(int[] suu) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str;
		for(int n = 0; n < suu.length; n++){
			str = br.readLine();
			suu[n] = Integer.parseInt(str);
		}
		return suu;
	}
	public static double[] in(double[] suu) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str;
		for(int n = 0; n < suu.length; n++){
			str = br.readLine();
			suu[n] = Double.parseDouble(str);
		}
		return suu;
	}
	public static String[] in(String[] suu) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		for(int n = 0; n < suu.length; n++){
			suu[n] = br.readLine();
		}
		return suu;
	}
	//入力要求時に文字列を表示する(一度)
	public static int[] in(int[] suu, String moji) throws IOException
	{
		System.out.print(moji);
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str;
		for(int n = 0; n < suu.length; n++){
			str = br.readLine();
			suu[n] = Integer.parseInt(str);
		}
		return suu;
	}
	public static double[] in(double[] suu, String moji) throws IOException
	{
		System.out.print(moji);
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str;
		for(int n = 0; n < suu.length; n++){
			str = br.readLine();
			suu[n] = Double.parseDouble(str);
		}
		return suu;
	}
	public static String[] in(String[] suu, String moji) throws IOException
	{
		System.out.print(moji);
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		for(int n = 0; n < suu.length; n++){
			suu[n] = br.readLine();
		}
		return suu;
	}
	//入力要求時に文字列を表示する(毎度)
	public static int[] in(int[] suu, String[] moji) throws IOException
	{
		
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str;
		for(int n = 0; n < suu.length; n++){
			System.out.print(moji[n]);
			str = br.readLine();
			suu[n] = Integer.parseInt(str);
		}
		return suu;
	}
	public static double[] in(double[] suu, String[] moji) throws IOException
	{
		
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str;
		for(int n = 0; n < suu.length; n++){
			System.out.print(moji[n]);
			str = br.readLine();
			suu[n] = Double.parseDouble(str);
		}
		return suu;
	}
	public static String[] in(String[] suu, String[] moji) throws IOException
	{
		
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		for(int n = 0; n < suu.length; n++){
			System.out.print(moji[n]);
			suu[n] = br.readLine();
		}
		return suu;
	}
	
	//■■配列、入力要求回数
	public static int[] in(int[] suu, int kai) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str;
		for(int n = 0; n < kai; n++){
			str = br.readLine();
			suu[n] = Integer.parseInt(str);
		}
		return suu;
	}
	public static double[] in(double[] suu, int kai) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str;
		for(int n = 0; n < kai; n++){
			str = br.readLine();
			suu[n] = Double.parseDouble(str);
		}
		return suu;
	}
	public static String[] in(String[] suu, int kai) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		for(int n = 0; n < kai; n++){
			suu[n] = br.readLine();
		}
		return suu;
	}
	//入力要求時に文字列を表示する(一度)
	public static int[] in(int suu[], int kai, String moji) throws IOException
	{
		System.out.print(moji);
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str;
		for(int n = 0; n < kai; n++){
			
			str = br.readLine();
			suu[n] = Integer.parseInt(str);
		}
		return suu;
	}
	public static double[] in(double[] suu, int kai, String moji) throws IOException
	{
		System.out.print(moji);
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str;
		for(int n = 0; n < kai; n++){
			str = br.readLine();
			suu[n] = Double.parseDouble(str);
		}
		return suu;
	}
	public static String[] in(String[] suu, int kai, String moji) throws IOException
	{
		System.out.print(moji);
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		for(int n = 0; n < kai; n++){
			suu[n] = br.readLine();
		}
		return suu;
	}
	
	
	//入力要求時に文字列を表示する(毎度)
	public static int[] in(int suu[], int kai, String[] moji) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str;
		for(int n = 0; n < kai; n++){
			System.out.print(moji[n]);
			str = br.readLine();
			suu[n] = Integer.parseInt(str);
		}
		return suu;
	}
	public static double[] in(double[] suu, int kai, String[] moji) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str;
		for(int n = 0; n < kai; n++){
			System.out.print(moji[n]);
			str = br.readLine();
			suu[n] = Double.parseDouble(str);
		}
		return suu;
	}
	public static String[] in(String[] suu, int kai, String[] moji) throws IOException
	{
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		for(int n = 0; n < kai; n++){
			System.out.print(moji[n]);
			suu[n] = br.readLine();
		}
		return suu;
	}
	//2次元配列の文字列を表で表示する
	public static void out(String[][] str){
		for(int n = 0; n < str.length; n++){
			for(int m = 0; m < str[0].length; m++){
				System.out.print(str[n][m] + "\t");
			}
			System.out.print("\n");
		}
	}
}