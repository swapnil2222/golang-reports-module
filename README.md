# golang-reports-module

This module will fetch the data from third party service which provided details about reports
  
  User can enter any one of the option from following.
  
   1.Show all reports
   
   2.Search reports by category
   
   3.Search reports by date range

 Expected Input :
 ****** Enter your option : ******
 * 1  OR
 * 2 OR
 * 3

 Expected Output : All reports based on option
 
 ## How to run ?
 
 - Run server.go file from server folder.
 
 ## How to test code ?
 - Hit the api url mentioned on server terminal.
 - It will ask to enter user option on terminal (from 1,2 or 3).
 - Server will return the response based on option.
 - If 1 is selected server will return the all reports as a response.
 - If 2 is selected, enter the catergory name e.g Energy Crisis, Running On Empty or Other.
 - If 3 is selected, enter the start date and end date in format like 2012-1-1 and you will get the output.
  
 
 
