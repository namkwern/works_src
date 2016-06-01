package module;
public class Rank
{
	public static int[] up(int[] suu, int syoki)
	{
		int[] jun = new int[suu.length];
		for(int n = 0; n < suu.length; n++)jun[n] = syoki;
		for(int n = 0; n < suu.length; n++){
			for(int m = n; m < suu.length; m++){
				if(suu[n] < suu[m])jun[m]++;
				if(suu[n] > suu[m])jun[n]++;
			}
		}
		return jun;
	}
	public static int[] up(double[] suu, int syoki)
	{
		int[] jun = new int[suu.length];
		for(int n = 0; n < suu.length; n++)jun[n] = syoki;
		for(int n = 0; n < suu.length; n++){
			for(int m = n; m < suu.length; m++){
				if(suu[n] < suu[m])jun[m]++;
				if(suu[n] > suu[m])jun[n]++;
			}
		}
		return jun;
	}
	public static int[] up(String[] suu, int syoki)
	{
		int[] num = new int[suu.length];
		int[] jun = new int[suu.length];
		for(int n = 0; n < suu.length; n++){
			jun[n] = syoki;
			num[n] = Integer.parseInt(suu[n]);
		}
		for(int n = 0; n < suu.length; n++){
			for(int m = n; m < suu.length; m++){
				if(num[n] < num[m])jun[m]++;
				if(num[n] > num[m])jun[n]++;
			}
		}
		return jun;
	}
	//ダウン
	public static int[] down(int[] suu, int syoki)
	{
		int[] jun = new int[suu.length];
		for(int n = 0; n < suu.length; n++)jun[n] = syoki;
		for(int n = 0; n < suu.length; n++){
			for(int m = n; m < suu.length; m++){
				if(suu[n] > suu[m])jun[m]++;
				if(suu[n] < suu[m])jun[n]++;
			}
		}
		return jun;
	}
	public static int[] down(double[] suu, int syoki)
	{
		int[] jun = new int[suu.length];
		for(int n = 0; n < suu.length; n++)jun[n] = syoki;
		for(int n = 0; n < suu.length; n++){
			for(int m = n; m < suu.length; m++){
				if(suu[n] > suu[m])jun[m]++;
				if(suu[n] < suu[m])jun[n]++;
			}
		}
		return jun;
	}
	public static int[] down(String[] suu, int syoki)
	{
		double[] num = new double[suu.length];
		int[] jun = new int[suu.length];
		for(int n = 0; n < suu.length; n++){
			jun[n] = syoki;
			num[n] = Double.parseDouble(suu[n]);
		}
		for(int n = 0; n < suu.length; n++){
			for(int m = n; m < suu.length; m++){
				if(num[n] > num[m])jun[m]++;
				if(num[n] < num[m])jun[n]++;
			}
		}
		return jun;
	}
	
	
	
	//ランクソート
	public static int[] sort(int[] suu, int[] jun){
		int[] data = new int[suu.length];
		for(int n = 0; n < suu.length; n++){
			data[n] = suu[jun[n]];
		}
		return data;
	}
	public static double[] sort(double[] suu, int[] jun){
		double[] data = new double[suu.length];
		for(int n = 0; n < suu.length; n++){
			data[n] = suu[jun[n]];
		}
		return data;
	}
	public static String[] sort(String[] suu, int[] jun){
		String[] data = new String[suu.length];
		for(int n = 0; n < suu.length; n++){
			data[n] = suu[jun[n]];
		}
		return data;
	}
	//ランクソート二次元配列
	public static int[][] sort(int[][] suu, int[] jun){
		int[][] data = new int[suu.length][suu[0].length];
		for(int n = 0; n < suu.length; n++){
			for(int m = 0; m < suu[0].length; m++){
				data[n][jun[m]] = suu[n][m];
			}
		}
		return data;
	}
	public static double[][] sort(double[][] suu, int[] jun){
		double[][] data = new double[suu.length][suu[0].length];
		for(int n = 0; n < suu.length; n++){
			for(int m = 0; m < suu[0].length; m++){
				data[n][jun[m]] = suu[n][m];
			}
		}
		return data;
	}
	public static String[][] sort(String[][] suu, int[] jun){
		String[][] data = new String[suu.length][suu[0].length];
		for(int n = 0; n < suu.length; n++){
			for(int m = 0; m < suu[0].length; m++){
				data[n][jun[m]] = suu[n][m];
			}
		}
		return data;
	}
}
