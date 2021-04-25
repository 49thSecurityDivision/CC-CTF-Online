#!/usr/bin/python3
import falcon, bjoern, json, sqlite3, uuid

app = falcon.API()
conn = sqlite3.connect(':memory:')

class login():
    def on_get(self, req, resp):
        print(req.params)
        un = req.params["username"].lower()
        pw = req.params["password"]
        c = conn.cursor()
        c.execute('SELECT * FROM users WHERE (username = ? and password = ?)', (un,pw,))
        if (c.fetchone()):
            cookie = str(uuid.uuid4())
            c.execute('UPDATE users SET cookie = ? WHERE username = ?', (cookie,un,))
            resp.body = "{'message':'success!'}"
            resp.set_cookie("X-Login-Token", cookie)
            resp.status = falcon.HTTP_200
        else:
            resp.body = "{'message':'bad username or password'}"
            resp.set_cookie("X-Login-Token", "NULL")
            resp.status = falcon.HTTP_401

class passwordReset():
    def on_get(self, req, resp):
        un = req.params["username"].lower()
        c = conn.cursor()
        command = "SELECT username FROM users WHERE username = '" + un+ "'"
        #print(command)

        try:#make it easier
            c.execute(command)
        except sqlite3.Error as er:
            resp.body = " ".join(er.args)
            resp.status = falcon.HTTP_500
            return
        
        user = c.fetchall()
        if (len(user) > 0):
            resp.body = "Email sent to " + json.dumps(user)
            resp.status = falcon.HTTP_200
        else:
            resp.body = "{'message':'username not in users table'}"
            resp.status = falcon.HTTP_401

class verifyCookie():
    def on_get(self, req, resp):
        c = conn.cursor()
        cookie = req.get_cookie_values("X-Login-Token")
        if ( not (cookie)):
            cookie = ""

        c.execute('SELECT * FROM users WHERE (cookie = ?)', (cookie,))
        if (c.fetchone()):
            resp.body = "{'message':'valid'}"
            resp.status = falcon.HTTP_200
        else:
            resp.body = "{'message':'cookie is not in users table'}"
            resp.status = falcon.HTTP_401

#The API def
app.add_route('/api/login', login())
app.add_route('/api/login/reset', passwordReset())#check for username. Respond "Email Sent!"
app.add_route('/api/session/verify', verifyCookie())

if __name__ == '__main__':
    c = conn.cursor()
    c.execute("CREATE TABLE users (username varchar(100) PRIMARY KEY, password varchar(500), cookie varchar(500))")
    c.execute("INSERT INTO users(username, password) VALUES (?,?)",["janet", "kittenkabal"])
    c.execute("INSERT INTO users(username, password) VALUES (?,?)",["bigq", "gimiBBQ87!"])
    c.execute("INSERT INTO users(username, password) VALUES (?,?)",["clarance", "flag{NiceSQLSkillzCzjTC3lPgCKWxVPhygRhq5wTiUC3xxnzPb}"])
    c.execute("INSERT INTO users(username, password) VALUES (?,?)",["zero", "lilBasedkid"])
    conn.commit()

    bjoern.listen(app, '0.0.0.0', 8080)
    bjoern.run()
