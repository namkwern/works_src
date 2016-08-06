package com.example.a1503.btserver.amodule;

import android.app.Activity;
import android.widget.Toast;

/**
 * Created by a1503 on 2016/05/31.
 */
public class Debug {
	Activity act;
	boolean flag;
	
	public Debug(Activity act, boolean flag){
		this.flag = flag;
		this.act = act;
	}
	
	public void alert(String str){
		if(flag)Toast.makeText(act, str, Toast.LENGTH_SHORT).show();
	}
	
	//assert(flaseがあったらエラー)
	public void ass(String str, boolean... b){
		if(!trueAll(b) && flag)throw new AssertionError(str);
	}
	public void ass(boolean... b){
		ass("", b);
	}
	
	public boolean trueAll(boolean... b){
		boolean allflag = true;
		for(boolean bb: b)
			if(!bb)allflag = false;
		return allflag;
	}
	
	public void nullCheck(Object... o){
		if(flag){
			for(int n = 0; n < o.length; n++){
				if (o[n] == null)
					System.out.println(n);
			}
		}
	}
}
