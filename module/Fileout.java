package module;
import java.io.*;
public class Fileout
{

	private PrintWriter pw;
	//コンストラクタでファイル名入力ファイルを開く
	//なければファイルを作成
	public Fileout(){}
	public Fileout(String filename)throws IOException
	{
		try{
			File newfile = new File(filename);
			if(!(newfile.exists())){
				newfile.createNewFile();
				System.out.println(filename + "を作成しました");
			}
			pw = new PrintWriter(new BufferedWriter(new FileWriter(filename)));
		}catch(IOException e){
			
		}
	}
	public Fileout(String filename, boolean post)throws IOException
	{
		this(filename);
		pw = new PrintWriter(new BufferedWriter(new FileWriter(filename, post)));
	}
	//ファイルを閉じる
	public void close(){
		pw.close();
	}
	
	//ファイルに書き込む
	public void out(String str){
		pw.println(str);
	}
//■■配列で操作■■

	//■■出力　配名■■
	public void out(int[] osuu)
	{
		for(int n = 0; n < osuu.length; n++)
		{
			pw.println(osuu[n]);
		}
	} 
	public void out(double[] osuu)
	{
		for(int n = 0; n < osuu.length; n++)
		{
			pw.println(osuu[n]);
		}
			
	} 
	public void out(String[] osuu)
	{
		for(int n = 0; n < osuu.length; n++)
		{
				pw.println(osuu[n]);
		}
	} 
	//■■出力　二次元配列■■
	public void out(int[][] osuu)
	{
		for(int n = 0; n < osuu.length; n++){
			for(int m = 0; m < osuu[n].length; m++){
				pw.println(osuu[n][m]);
			}
		}
	}
	public void out(double[][] osuu)
	{
		for(int n = 0; n < osuu.length; n++){
			for(int m = 0; m < osuu[n].length; m++){
				pw.println(osuu[n][m]);
			}
		}
	}
	public void out(String[][] osuu)
	{
		for(int n = 0; n < osuu.length; n++){
			for(int m = 0; m < osuu[n].length; m++){
					pw.println(osuu[n][m]);
			}
		}
		
	}
	//■■出力　三次元配列、■■
	public void out(int[][][] osuu)
	{
		
			
		for(int n = 0; n < osuu.length; n++){
	 		for(int m = 0; m < osuu[n].length; m++){
				for(int k = 0; k < osuu[n][m].length; k++){
					pw.println(osuu[n][m][k]);
				}
			}
		}
		
	}
	public void out(double[][][] osuu)
	{
		for(int n = 0; n < osuu.length; n++){
			for(int m = 0; m < osuu[n].length; m++){
				for(int k = 0; k < osuu[n][m].length; k++){
					pw.println(osuu[n][m][k]);
				}
			}
		}
	}
	public void out(String[][][] osuu)
	{
		for(int n = 0; n < osuu.length; n++){
			for(int m = 0; m < osuu[n].length; m++){
				for(int k = 0; k < osuu[n][m].length; k++){
					pw.println(osuu[n][m]);
				}
			}
		}
	}
}