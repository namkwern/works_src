package com.example.a1503.btserver.amodule;

import android.app.Activity;
import android.content.Context;
import android.hardware.Sensor;
import android.hardware.SensorEventListener;
import android.hardware.SensorManager;

import java.util.List;

/**
 * Created by a1503 on 2016/05/06.
 */
public class SimpleSensor
{
	public boolean flag = true;
	SensorManager manager;
	SensorEventListener sEvent;
	int type, speed;
	
	public SimpleSensor(Activity act){
		flag = true;
		manager = (SensorManager) act.getSystemService(Context.SENSOR_SERVICE);
	}
	
	//センサー取得
	public void set(SensorEventListener sEvent, int type, int speed){
		this.sEvent = sEvent;
		this.type = type;
		this.speed = speed;
	}
	
	//センサーの取得とリスナー登録
	public void start(){
		List<Sensor> sensorList = manager.getSensorList(type);
		manager.registerListener(sEvent, sensorList.get(0), speed);
	}
	
	//一時停止
	public void pause(){
		if(manager != null)
			manager.unregisterListener(sEvent);
	}
}
