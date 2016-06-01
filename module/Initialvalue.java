package module;
public class Initialvalue
{
	//初期化したい配列、代入する値
	public static int[] number(int[] suu3822, int syokiti2938)
	{
		for(int n = 0; n < suu3822.length; n++){
			suu3822[n] = syokiti2938;
		}
		return suu3822;
	}
	public static double[] number(double[] suu3822, double syokiti2938)
	{
		for(int n = 0; n < suu3822.length; n++){
			suu3822[n] = syokiti2938;
		}
		return suu3822;
	}
	public static String[] number(String[] suu3822, String syokiti2938)
	{
		for(int n = 0; n < suu3822.length; n++){
			suu3822[n] = syokiti2938;
		}
		return suu3822;
	}
	
	//初期化したい配列、代入する最初の数字、上昇または下降する間隔
	public static int[] number(int[] suu3822, int syokiti2938, int henka6570)
	{
		for(int n = 0; n < suu3822.length; n++){
			suu3822[n] = syokiti2938 + (henka6570 * n);
		}
		return suu3822;
	}
	public static double[] number(double[] suu3822, double syokiti2938, double henka6570)
	{
		for(int n = 0; n < suu3822.length; n++){
			suu3822[n] = syokiti2938 + (henka6570 * n);
		}
		return suu3822;
	}
}
