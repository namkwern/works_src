package com.example.a1503.btserver.amodule;

import android.app.Activity;
import android.widget.Toast;

/**
 * Created by a1503 on 2016/05/31.
 */
public class Debug {
	boolean flag;
	public Debug(boolean flag){
		this.flag = flag;
	}
	
	public void alert(Activity act, String str){
		if(flag)Toast.makeText(act, str, Toast.LENGTH_SHORT).show();
	}
	
	//assert
	public void ass(boolean b, String str){
		if(!b && flag)throw new AssertionError(str);
	}
	public void ass(boolean b){
		ass(b, "");
	}
}
