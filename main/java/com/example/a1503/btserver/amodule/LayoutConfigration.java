package com.example.a1503.btserver.amodule;

import android.app.Activity;
import android.graphics.Color;
import android.view.View;
import android.widget.Button;
import android.widget.LinearLayout;
import android.widget.TextView;

/**
 * Created by a1503 on 2016/05/05.
 */
//レイアウトの生成と、関連オブジェクトの生成・追加
public class LayoutConfigration
{
	private LinearLayout layout;
	private Activity act;
	public float defaultSize;
	
	//constractor
	public LayoutConfigration(Activity act){
		this.act = act;
		layout = new LinearLayout(act);
		layout.setBackgroundColor(Color.WHITE);
		layout.setOrientation(LinearLayout.VERTICAL);
		act.setContentView(layout);
	}
	public LayoutConfigration(Activity act, float size){
		this(act);
		defaultSize = size;
	}
	
	public TextView getTextView(String str, float size){
		TextView text = new TextView(act);
		text.setTextSize(size);
		text.setText(str);
		layout.addView(text);
		return text;
	}
	public TextView getTextView(String str){
		return getTextView(str, defaultSize);
	}
	public TextView getTextView(float size){
		return getTextView("", size);
	}
	public TextView getTextView(){
		return getTextView("");
	}
	
	public Button getButton(String str, float size, View.OnClickListener event){
		Button btn = new Button(act);
		btn.setText(str);
		btn.setTextSize(size);
		btn.setOnClickListener(event);
		layout.addView(btn);
		return btn;
	}
	public Button getButton(String str, View.OnClickListener event){
		return getButton(str, defaultSize, event);
	}
}