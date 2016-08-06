package com.example.a1503.btserver.amodule;

import android.app.Activity;
import android.graphics.Bitmap;
import android.graphics.Color;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.RadioButton;
import android.widget.RadioGroup;
import android.widget.TextView;

import java.util.ArrayList;

/**
 * Created by a1503 on 2016/05/05.
 */
//レイアウトの生成と、関連オブジェクトの生成・追加
public class LayoutConfigration {
	private LinearLayout layout;
	private Activity act;
	private ArrayList<RadioButton> radioList = new ArrayList<RadioButton>();
	
	public float defaultSize = 20.0f;
	public static final int MATCH = LinearLayout.LayoutParams.MATCH_PARENT;
	public static final int WRAP = LinearLayout.LayoutParams.WRAP_CONTENT;
	
	int id = 0;
	
	//constractor hr==true→horizontal
	public LayoutConfigration(Activity act, float size, boolean hr) {
		this(act, hr);
		defaultSize = size;
	}
	public LayoutConfigration(Activity act, boolean hr) {
		this.act = act;
		int vrhr = hr ? LinearLayout.HORIZONTAL : LinearLayout.VERTICAL;
		layout = new LinearLayout(act);
		layout.setBackgroundColor(Color.WHITE);
		layout.setOrientation(vrhr);
	}
	public LayoutConfigration(LinearLayout layout) {
		this.layout = layout;
	}
	
	//テキストビュー
	public TextView getTextView(String str, float size) {
		TextView text = new TextView(act);
		text.setTextSize(size);
		text.setText(str);
		add(text);
		return text;
	}
	
	public TextView getTextView(String str) {
		return getTextView(str, defaultSize);
	}
	
	public TextView getTextView(float size) {
		return getTextView("", size);
	}
	
	public TextView getTextView() {
		return getTextView("");
	}
	
	//ボタン
	public Button getButton(String str, float size, View.OnClickListener event) {
		Button btn = new Button(act);
		btn.setText(str);
		btn.setTextSize(size);
		btn.setOnClickListener(event);
		add(btn);
		return btn;
	}
	
	public Button getButton(String str, View.OnClickListener event) {
		return getButton(str, defaultSize, event);
	}
	
	//テキストボックス
	public EditText getEditText(float size) {
		EditText etext = new EditText(act);
		etext.setTextSize(size);
		add(etext);
		return etext;
	}
	
	public EditText getEditText() {
		return getEditText(defaultSize);
	}
	
	//ラジオボタン
	public RadioButton getRadioButton(String str) {
		return getRadioButton(str, defaultSize);
	}
	public RadioButton getRadioButton(String str, float size) {
		RadioButton radio = new RadioButton(act);
		radio.setText(str);
		radio.setTextSize(size);
		radioList.add(radio);
		return radio;
	}
	public RadioButton getRadioButton(String str, View.OnClickListener event){
		return getRadioButton(str, defaultSize, event);
	}
	public RadioButton getRadioButton(String str, float size, View.OnClickListener event){
		RadioButton radio = getRadioButton(str, size);
		radio.setOnClickListener(event);
		return radio;
	}
	
	//ラジオをまとめる
	public RadioGroup getRadioGroup(int checked){
		if(!(radioList.size() > checked) || 0 > checked)
			throw new IndexOutOfBoundsException("チェックをつけられません");
		RadioGroup radios = new RadioGroup(act);
		for(int n = 0; n < radioList.size(); n++) {
			RadioButton r = radioList.get(n);
			r.setId(id++);
			if(n == checked)
				r.setChecked(true);
			else
				r.setChecked(false);
			radios.addView(r);
		}
		add(radios);
		radioList.clear();
		return radios;
	}
	//まとめずに作っちゃう
	public RadioGroup getRadioGroup(int checked, String... str){
		for(String s: str)
			getRadioButton(s);
		return getRadioGroup(checked);
	}
	//イベントもまとめる
	public RadioGroup getRadioGroup(int checked, String[] str, View.OnClickListener event){
		for(String s: str)
			getRadioButton(s, event);
		return getRadioGroup(checked);
	}
	
	//Imageビュー
	public ImageView getImageView(Bitmap bitmap){
		ImageView image = getImageView();
		image.setImageBitmap(bitmap);
		return image;
	}
	public ImageView getImageView(){
		ImageView image = new ImageView(act);
		add(image);
		return image;
	}
	
	//パラメーターセット　
	public void setParam(View v, int i, int j){
		v.setLayoutParams(new LinearLayout.LayoutParams(i, j));
	}
	public void setParam(int i, int j){
		setParam(layout, i, j);
	}
	
	//viewの追加(これを使う)
	public void add(View... view){
		for(View v: view){
			v.setId(id++);
			layout.addView(v);
		}
	}
	//レイアウトの追加
	public void add(LayoutConfigration... layouts) {
		for(LayoutConfigration ls: layouts)
			layout.addView(ls.getLayout());
	}
	
	//レイアウトの取得
	public LinearLayout getLayout(){
		return layout;
	}
	
	//レイアウトの描画
	public void disp(){
		act.setContentView(layout);
	}
}