package googleSheets

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
  "os"
  "os/user"
  "path/filepath"
  
  "golang.org/x/net/context"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
  "google.golang.org/api/sheets/v4"
)

type TableData struct{
  TableName string
  InOut string
  InOutComment string
  Analaysis string
  SQL string
  Reconciliation string
  Mapping string
  UDI string
  Testing string
  Comment string
}
// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
  cacheFile, err := tokenCacheFile()
  if err != nil {
    log.Fatalf("Unable to get path to cached credential file. %v", err)
  }
  tok, err := tokenFromFile(cacheFile)
  if err != nil {
    tok = getTokenFromWeb(config)
    saveToken(cacheFile, tok)
  }
  return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
  authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
  fmt.Printf("Go to the following link in your browser then type the "+
    "authorization code: \n%v\n", authURL)

  var code string
  if _, err := fmt.Scan(&code); err != nil {
    log.Fatalf("Unable to read authorization code %v", err)
  }

  tok, err := config.Exchange(oauth2.NoContext, code)
  if err != nil {
    log.Fatalf("Unable to retrieve token from web %v", err)
  }
  return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
  usr, err := user.Current()
  if err != nil {
    return "", err
  }
  tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
  os.MkdirAll(tokenCacheDir, 0700)
  return filepath.Join(tokenCacheDir,
    url.QueryEscape("sheets.googleapis.com-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
  f, err := os.Open(file)
  if err != nil {
    return nil, err
  }
  t := &oauth2.Token{}
  err = json.NewDecoder(f).Decode(t)
  defer f.Close()
  return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
  fmt.Printf("Saving credential file to: %s\n", file)
  f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
  if err != nil {
    log.Fatalf("Unable to cache oauth token: %v", err)
  }
  defer f.Close()
  json.NewEncoder(f).Encode(token)
}

func Run() {
  ctx := context.Background()

  b, err := ioutil.ReadFile("/home/cabox/workspace/src/googleSheets/client_secret.json")
  if err != nil {
    log.Fatalf("Unable to read client secret file: %v", err)
  }

  // If modifying these scopes, delete your previously saved credentials
  // at ~/.credentials/sheets.googleapis.com-go-quickstart.json
  config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
  if err != nil {
    log.Fatalf("Unable to parse client secret file to config: %v", err)
  }
  client := getClient(ctx, config)

  srv, err := sheets.New(client)
  if err != nil {
    log.Fatalf("Unable to retrieve Sheets Client %v", err)
  }

  // Prints the names and majors of students in a sample spreadsheet:
  // https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
  //https://docs.google.com/spreadsheets/d/1R6j7FFSQrfS0NjV4UsiBAmcYjU8--fcd6C_tCZVJDxY/edit
  spreadsheetId := "1R6j7FFSQrfS0NjV4UsiBAmcYjU8--fcd6C_tCZVJDxY"
  readRange := "A2:J"
  resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
  if err != nil {
    log.Fatalf("Unable to retrieve data from sheet. %v", err)
  }

  if len(resp.Values) > 0 {
    fmt.Println("Name, Major:")
    for _, row := range resp.Values {
      // Print columns A and E, which correspond to indices 0 and 4.
      a := row[0]
      b := row[1]
      c := row[2]
      d := row[3]
      e:= row[4]
      f:= row[5]
      g:= row[6]
      h:= row[7]
      i:= row[8]
      j:= row[9]
       if a== nil{
        a =""  
      } 
       if b== nil{
        b =""  
      } 
       if c== nil{
        c =""  
      } 
       if d== nil{
        d =""  
      } 
       if e== nil{
        e =""  
      } 
       if f== nil{
        f =""  
      } 
       if g== nil{
        g =""  
      } 
       if h== nil{
        h =""  
      } 
       if i== nil{
        i =""  
      } 
       if j== nil{
        j =""  
      } 
      
      fmt.Print(a ," ", b ," ", c," ",d," ",e," ",f," ",g," ",h," ",i," ",j,"\n ")
    }
  } else {
    fmt.Print("No data found.")
  }
}
  //////////////////////////////////////////////////////////////////////////
  //////////////////////////////////////////////////////////////////////////
  //////////////////////////////////////////////////////////////////////////
  //////////////////////////////////////////////////////////////////////////
  //////////////////////////////////////////////////////////////////////////
  func GetTables() []TableData{
       ctx := context.Background()

    b, err := ioutil.ReadFile("/home/cabox/workspace/src/googleSheets/client_secret.json")
    if err != nil {
      log.Fatalf("Unable to read client secret file: %v", err)
    }

    // If modifying these scopes, delete your previously saved credentials
    // at ~/.credentials/sheets.googleapis.com-go-quickstart.json
    config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
    if err != nil {
      log.Fatalf("Unable to parse client secret file to config: %v", err)
    }
    client := getClient(ctx, config)

    srv, err := sheets.New(client)
    if err != nil {
      log.Fatalf("Unable to retrieve Sheets Client %v", err)
    }

    // Prints the names and majors of students in a sample spreadsheet:
    // https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
    //https://docs.google.com/spreadsheets/d/1R6j7FFSQrfS0NjV4UsiBAmcYjU8--fcd6C_tCZVJDxY/edit
    spreadsheetId := "1R6j7FFSQrfS0NjV4UsiBAmcYjU8--fcd6C_tCZVJDxY"
    readRange := "A2:J"
    resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
    if err != nil {
      log.Fatalf("Unable to retrieve data from sheet. %v", err)
    }
    
    var tables []TableData
    
    if len(resp.Values) > 0 {
      for _, row := range resp.Values {
        table := TableData{
          TableName : (row[0].(string)),
            InOut : (row[1].(string)),
            InOutComment : (row[2].(string)),
            Analaysis : (row[3].(string)),
            SQL : (row[4].(string)),
            Reconciliation : (row[5].(string)),
            Mapping : (row[6].(string)),
            UDI : (row[7].(string)),
            Testing : (row[8].(string)),
            Comment : (row[9].(string)),
        }
        tables = append(tables, table)
        
        }
      }else {
        fmt.Print("No data found.")
    }
   
//    fmt.Printf("%+v \n", tables)
    return tables
  }

