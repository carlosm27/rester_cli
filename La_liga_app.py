import http.client
import json
from tkinter import *
from tkinter import messagebox


connection=http.client.HTTPConnection('api.football-data.org')
headers={'X-Auth-Token':'295f60cecff7454b82d66a7d80f0edd1'}
connection.request('GET','https://api.football-data.org/v2/competitions/PD/matches',None,headers)
response=json.loads(connection.getresponse().read().decode())


# explicit function to get the team

def get_matches(team):
    result=response

    if result:
        json=result.json()
        matches=json['season']['380']['name']

        return matches

    else:
        print("No content Found")


def search():
    matches=matches_text.get()
    match=get_matches()

    

#Create a object

app=Tk()

#add title

app.title("La Liga App")

#adjust window size
app.geometry("500x500")

#add labels,buttons and text

team_text=StringVar()
team_entry=Entry(app,textvariable=team_text)
team_entry.pack()

Search_btn=Button(app,text="Search team", width=12,command=search)
Search_btn.pack()

app.mainloop()
