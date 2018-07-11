
import java.io.*;
import java.net.*;
import java.lang.Integer;
import java.util.ArrayList;
import java.util.List;
import java.lang.Boolean;
import java.lang.Float;

public class csInterface {
	private static int port=8080;
	private static String inetaddress="localhost";
	private Socket client;
	
	public csInterface(){
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
	        pw.flush();
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
		        int singlechar;
			while((singlechar=in.read())!=(int)'\n'){
			    buffer.append((char)singlechar);
			}
			return buffer.toString();
		}
		catch (IOException e){
			e.printStackTrace();
			return null;
		}
	}
    public void close()
    {
        try {
	    client.close();
	}
	catch(IOException e){
	    e.printStackTrace();
	}
    }
    
    public boolean SignIn(String username, String password)
    {
	sendTo("SignIn "+username+" "+password+"\n");
	String temp = recvIn();
	return Boolean.parseBoolean(temp.substring(0,4));
    }

    public boolean SignUp(String username, String password)
    {
	sendTo("SignUp "+username+" "+password+"\n");
	String temp = recvIn();
	return Boolean.parseBoolean(temp.substring(0,4));
    }

    public String[] getAnnouncement(String lasttime)
    {
	sendTo("Announcement "+lasttime+"\n");
	List<String> recv = new ArrayList<String>();
	boolean flag = true ;
	
	while(flag)
	    {
		String temp = recvIn();
		if(temp.equals("end"))
		    flag = false;
		recv.add(temp);
	    }
	String[] ret = (String[])recv.toArray(new String[recv.size()]);
	return ret;
    }

    public float getAccount()
    {
	sendTo("Account\n");
	String recv = recvIn();
	return Float.parseFloat(recv);
    }

    public String getMaintain()
    {
	sendTo("Maintain\n");
	String recv = recvIn();
	return recv;
    }

    public String getContact()
    {
	sendTo("Contact\n");
	String recv = recvIn();
	return recv;
    }
}
