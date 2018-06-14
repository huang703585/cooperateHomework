package csInterface;

import java.io.*;
import java.net.*;

public class csInterface {
	private static int port=8080;
	private static String inetaddress="localhost";
	private StringBuffer recv,send;
	private Socket client;
	
	public csInterface(){
		recv=new StringBuffer();
		send=new StringBuffer();
	
		try
		{
			client=new Socket(inetaddress,port);
		}
		catch(UnknownHostException e){
			e.printStackTrace();
		}
		catch(IOException e){
			e.printStackTrace();
		}
	}
	
	public void sendTo(String s)
	{
		try{
			PrintWriter pw = new PrintWriter(client.getOutputStream());  
	        pw.write(s);  
	        pw.close();
		}
		catch(IOException e){
			e.printStackTrace();
		}
	}
	
	public String recvIn()
	{
		try{
			InputStream is = client.getInputStream();
			BufferedReader in = new BufferedReader(new InputStreamReader(is));
			StringBuffer buffer = new StringBuffer();
			String line = "";
			while((line = in.readLine())!=null){
				buffer.append(line);
			}
			return buffer.toString();
		}
		catch (IOException e){
			e.printStackTrace();
			return null;
		}
	}
}
