-- script boot.applescript
on run argv

set deviceName to (item 1 of argv) 
set deviceState to (item 2 of argv) 

if deviceState = "Booted"
    tell application "Simulator.app"
      activate
    end tell
    return "Device Already Booted"
end if      

if application "Simulator.app" is running then
  do shell script "xcrun simctl boot " & deviceName
  tell application "Simulator.app"
    activate
  end tell
  return "Activating Sim App"
else  
  do shell script "xcrun simctl boot " & deviceName
  tell application "Simulator.app"
    launch
  end tell  
  return "Launching Sim App"
end if  


end run