package module;
import java.io.*;
public class Filein
{
//■■変数で操作■■
	private static BufferedReader br;
	public Filein(){}
	//コンストラクタでファイル名入力ファイルを開く
	public Filein(String filename)throws IOException
	{
		br = new BufferedReader(new FileReader(filename));
	}
	public Filein(String filename, int mark)throws IOException
	{
		br = new BufferedReader(new FileReader(filename));
		br.mark(mark);
	}
	//ファイルを閉じる
	public void close()throws IOException
	{
			br.close();
	}
	//■■入力　すべて一行で取得■■
	public static String in()throws IOException
	{
		String line, str = "";
		while((line = br.readLine()) != null)
			str += line + "\n";
		return str;
	}
	//■■入力　指標(行番号指定)■■
	public static String in(int index)throws IOException
	{
		String str = "";
		for(int n = 0; n <= index; n++)
		{
			str = br.readLine();
			if(str == null)break;
		}
		try{br.reset();}catch(Exception e){}
		return str;
	}

//■■配列で操作■■
	//■■入力　配列■■
	public static int[] in(int[] isuu)throws IOException
	{
		String str;
		for(int n = 0; n < isuu.length; n++)
		{
			str = br.readLine();
			if(str == null)break;
			isuu[n] = Integer.parseInt(str);
		}
		try{br.reset();}catch(Exception e){}
		return isuu;
	}
	public static double[] in(double[] isuu)throws IOException
	{
		String str;
		for(int n = 0; n < isuu.length; n++)
		{
			str = br.readLine();
			if(str == null)break;
			isuu[n] = Double.parseDouble(str);
		}
		try{br.reset();}catch(Exception e){}
		return isuu;
	}
	public static String[] in(String[] isuu)throws IOException
	{
		for(int n = 0; n < isuu.length; n++)
		{
			isuu[n] = br.readLine();
			if(isuu[n] == null)break;
		}
		try{br.reset();}catch(Exception e){}
		return isuu;
	}
	//■■入力　配列、、読み取り開始、読み取り終了■■
	public static int[] in(int[] isuu,int min3425,int max5425)throws IOException
	{
		String str;
		for(int n = 0; n <= max5425; n++)
		{
			str = br.readLine();
			if(str == null)break;
			if(n <= min3425)
				isuu[n] = Integer.parseInt(str);
		}
		try{br.reset();}catch(Exception e){}
		return isuu;
	}
	public static double[] in(double[] isuu,int min3425,int max5425)throws IOException
	{
		String str;
		for(int n = 0; n <= max5425; n++)
		{
			str = br.readLine();
			if(str == null)break;
			if(n <= min3425)
				isuu[n] = Double.parseDouble(str);
		}
		try{br.reset();}catch(Exception e){}
		return isuu;
	}
	public static String[] in(String[] isuu,int min3425,int max5425)throws IOException
	{
		String str;
		int count = 0;
		for(int n = 0; n <= max5425; n++)
		{
			str = br.readLine();
			if(str == null)
				break;
			if(n >= min3425){
				isuu[count] = str;
				count++;
			}
		}
		try{br.reset();}catch(Exception e){}
		return isuu;
	}
	//■■入力　二次元配列、■■
	public static int[][] in(int[][] isuu)throws IOException
	{
		String str = "0";
		for(int n = 0; n < isuu.length; n++){
			for(int m = 0; m < isuu[n].length; m++){
				str = br.readLine();
				if(str == null)break;
				isuu[n][m] = Integer.parseInt(str);
			}
			if(str == null)break;
		}
			try{br.reset();}catch(Exception e){}
		return isuu;
	}
	public static double[][] in(double[][] isuu)throws IOException
	{
		String str = "0";
		for(int n = 0; n < isuu.length; n++){
			for(int m = 0; m < isuu[n].length; m++){
				str = br.readLine();
				if(str == null)break;
				isuu[n][m] = Double.parseDouble(str);
			}
			if(str == null)break;
		}
			try{br.reset();}catch(Exception e){}
		return isuu;
	}
	public static String[][] in(String[][] isuu)throws IOException
	{
		String str = "0";
		for(int n = 0; n < isuu.length; n++){
			for(int m = 0; m < isuu[n].length; m++){
				str = br.readLine();
				if(str == null)break;
				isuu[n][m] = str;
			}
			if(str == null)break;
		}
		try{br.reset();}catch(Exception e){}
		return isuu;
	}
	//■■入力　三次元配列、■■
	public static int[][][] in(int[][][] isuu)throws IOException
	{
		String str = "0";
		for(int n = 0; n < isuu.length; n++){
			for(int m = 0; m < isuu[n].length; m++){
				for(int k = 0; k < isuu[n][m].length; k++){
					str = br.readLine();
					if(str == null)break;
					isuu[n][m][k] = Integer.parseInt(str);
				}
				if(str == null)break;
			}
			if(str == null)break;
		}
		try{br.reset();}catch(Exception e){}
		
		
		return isuu;
	}
	public static double[][][] in(double[][][] isuu)throws IOException
	{
		String str = "0";
		for(int n = 0; n < isuu.length; n++){
			for(int m = 0; m < isuu[n].length; m++){
				for(int k = 0; k < isuu[n][m].length; k++){
					str = br.readLine();
					if(str == null)break;
					isuu[n][m][k] = Double.parseDouble(str);
				}
				if(str == null)break;
			}
			if(str == null)break;
		}
		try{br.reset();}catch(Exception e){}
		return isuu;
	}
	public static String[][][] in(String[][][] isuu)throws IOException
	{
		String str = "0";
		for(int n = 0; n < isuu.length; n++){
			for(int m = 0; m < isuu[n].length; m++){
				for(int k = 0; k < isuu[n][m].length; k++){
					str = br.readLine();
					if(str == null)break;
					isuu[n][m][k] = str;
				}
				if(str == null)break;
			}
			if(str == null)break;
		}
		try{br.reset();}catch(Exception e){}
		return isuu;
	}
//■■データ数出力■■
	public static int size()throws IOException
	{
		int k = 0;
		String str;
		for(k = 0; ; k++)
		{
			str = br.readLine();
			if(str == null){
			k--;
			break;
			}
		}
		k++;
		try{br.reset();}catch(Exception e){}
		return k;
	}
}