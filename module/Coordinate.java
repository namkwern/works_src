package module;
import java.util.*;
//座標からcin,cos,tanや角度、最短距離を計算する
public class Coordinate{
	double criteria = 0;
	final static String TAB = "↑方向を0°、→方向を90°とする";
	private int x, y;
	private double shotest;
	
	//座標指定
	public Coordinate(int xa, int ya, int xb, int yb){
		this(xb - xa, yb - ya);
	}
	public Coordinate(int x, int y){
		this.x = x;
		this.y = y;
		shotest = Math.sqrt(Math.pow(x, 2) + Math.pow(y, 2));
	}
	
	//角度の基準
	public void setCriteria(double criteria){
		this.criteria = criteria;
	}
	
	//スケール設定、値を再計算して設定
	public void setScale(double xscale, double yscale){
		x *= xscale;
		y *= yscale;
		shotest = Math.sqrt(Math.pow(x, 2) + Math.pow(y, 2));
	}
	
	//sin cos tan
	public double sin(){
		return y / shotest;
	}
	
	public double cos(){
		return x / shotest;
	}
	
	public double tan(){
		return (double)y / x;
	}
	
	
	//角度を判定
	public double getAngle(){
		return Math.asin(sin()) / Math.PI * 180;
	}
	
	//
	public double getShotest(){
		return shotest;
	}
}