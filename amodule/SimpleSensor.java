package com.example.a1503.car_logger.amodule;

import android.app.Activity;
import android.content.Context;
import android.hardware.Sensor;
import android.hardware.SensorEvent;
import android.hardware.SensorEventListener;
import android.hardware.SensorManager;
import android.widget.TextView;

import com.example.a1503.car_logger.amodule.*;

import java.util.List;

/**
 * Created by a1503 on 2016/05/06.
 */
public class SimpleSensor
{
	public boolean flag = true;
	SensorManager manager;
	SensorEventListener sEvent;
	
	public SimpleSensor(Activity act){
		flag = true;
		manager = (SensorManager) act.getSystemService(Context.SENSOR_SERVICE);
	}
	
	//センサーの取得とリスナー登録
	public void start(SensorEventListener sEvent,int type, int speed){
		this.sEvent = sEvent;
		List<Sensor> sensorList = manager.getSensorList(type);
		manager.registerListener(sEvent, sensorList.get(0), speed);
	}
	
	public void pause(){
		if(manager != null)
			manager.unregisterListener(sEvent);
	}
}
