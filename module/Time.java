package module;
import java.util.*;
public class Time
{
	private static long times;
	private static double scale;
	
	private static long before = 0;
	private static int count = 0;
	//スケールを設定せず初期化
	public Time(){
		times = 0;
		scale = 1.0;
	}
	//スケールを設定
	public Time( double scale ){
		times = 0;
		this.scale = scale;
	}
	
	//ストップウォッチスタート、現在時刻返却
	public static void timerStart(){
		times = System.currentTimeMillis();
	}
	//ストップ、差分を返す
	public static long timerStop(){
		return (long)( ( System.currentTimeMillis() - times ) * scale );
	}
	
	//時間にスタンプ
	public static void stamp(){
		long stamp = System.currentTimeMillis();
		if(count != 0)
			System.out.println("[" + count + "] " + (stamp - before) + "ms");
		before = stamp;
		count++;
	}
	public static void stamp(int count){
		long stamp = System.currentTimeMillis();
		if(count != 0)
			System.out.println("[" + count + "] " + (stamp - before) + "ms");
		before = stamp;
	}
	
	//現在時刻を返す
	public static long now(){
		return System.currentTimeMillis();
	}
	//引数との差分を出力
	public static long now(long times){
		return System.currentTimeMillis() - times;
	}
	//プログラムを指定した時間停止させる
	public static void sleep( long interval ){
		try{
			Thread.sleep( (long)(interval / scale) );
		}catch( InterruptedException ie ){}
	}
}