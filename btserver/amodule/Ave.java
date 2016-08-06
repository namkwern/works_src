package com.example.a1503.btserver.amodule;

/**
 * Created by a1503 on 2016/06/17.
 */
public class Ave {
	private int count = 0;
	private double ave = 0.0;
	
	public void add(double d){
		count++;
		ave = ave * (1 - 1/count) + d * (1/count);
	}
	public double get(){
		double rv = ave;
		reset();
		return rv;
	}
	public int getCount(){
		return count;
	}
	public void reset(){
		ave = 0.0;
		count = 0;
	}
}
