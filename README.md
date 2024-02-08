# Sujet Projet Groupie-Tracker

## Objectives

Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a application, displaying the information.

- It will be given an [API](https://groupietrackers.herokuapp.com/api), that consists in four parts:

  - The first one, `artists`, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

  - The second one, `locations`, consists in their last and/or upcoming concert locations.

  - The third one, `dates`, consists in their last and/or upcoming concert dates.

  - And the last one, `relation`, does the link between all the other parts, `artists`, `dates` and `locations`.

- Given all this you should a user friendly application where you can display the bands info through several data visualizations (examples : blocks, cards, tables, list, pages, graphics, etc). It is up to you to decide how you will display it.

- This project also focuses on the creation of events/actions and on their visualization.

  - The event/action we want you to do is known as a client call to the server (client-server). We can say it is a feature of your choice that needs to trigger an action. This action must communicate with the server in order to recieve information, ([request-response])(https://en.wikipedia.org/wiki/Request%E2%80%93response)
  - An event consists in a system that responds to some kind of action triggered by the client, time, or any other factor.

### Groupie Tracker GUI geolocalization

#### Objectives

the Groupie Tracker GUI geolocalization consists on mapping the different concerts locations of a certain artist/band given by the Client.

- You must use a process of converting addresses (ex: Germany Mainz) into geographic coordinates (ex: 49,59380 8,15052) which you must use to place markers for the concerts locations of a certain artist/band on a map.

- You are free to use the [Map API](https://rapidapi.com/blog/top-map-apis/) you found more appropriate.

### Groupie Tracker GUI search bar

#### Objectives

Groupie Tracker GUI search bar consists of creating a functional program that searches, inside your application, for a specific text input.
So the focus of this project is to create a way for the client to search a member or artist or any other attribute in the data system you made.

- The program should handle at least these search cases :
  - artist/band name
  - members
  - locations
  - first album date
  - creation date
- The program must handle search input as case-insensitive.
- The search bar must have typing suggestions as you write.
  - The search bar must identify and display in each suggestion the individual type of the search cases. (ex: Freddie Mercury -> member)
  - For example if you start writing `"phil"` it should appear as suggestions `Phil Collins - member` and `Phil Collins - artist/band`. This is just an example of a display.

#### Example

Lets imagine you have created a card system to display the band data. The user can directly search for the band he needs. Here is an example:

- While the user is typing for the member he desires to see, the search bar gives the suggestion of all the possible options.

![searchExample](https://gist.github.com/assets/53236467/ff5efd44-2a07-4da3-820f-a039f3a32700)


### Groupie Tracker GUI Filters

#### Objectives

Groupie Tracker GUI Filters consists on letting the user filter the artists/bands that will be shown.

- Your project must incorporate at least these four filters:

  - filter by creation date
  - filter by first album date
  - filter by number of members
  - filter by locations of concerts

- Your filters must be of at least these two types:
  - a range filter (filters the results between two values)
  - a check box filter (filters the results by one or multiple selection)

#### Example

Here is an example of both types of filters:

![filters_example](https://gist.github.com/assets/53236467/2c749b8e-1774-4d3a-a95a-b9e1676cfab1)



#### Hints

- You have to pay attention to the locations. For example Seattle, Washington, Washington **is part of** USA.

### Groupie Tracker GUI visualizations

#### Objectives

Groupie tracker visualizations consists of manipulating the data coming from the API and displaying it in the most presentable way possible to you. The [_Schneiderman's 8 Golden Rules of Interface Design_](https://www.interaction-design.org/literature/article/shneiderman-s-eight-golden-rules-will-help-you-design-better-interfaces) must be followed:

- Strive for consistency
- Enable frequent users to use shortcuts
- Offer informative feedback
- Design dialogue to yield closure
- Offer simple error handling
- Permit easy reversal of actions
- Support internal locus of control
- Reduce short-term memory load

And you must add **shortcuts** to your application.

## Instructions

- The application must be written in **Go**.
- The application cannot crash at any time.
- The code must respect the [**good practices**](../good-practices/README.md).
- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).

## Allowed packages

- Only the [standard Go](https://golang.org/pkg/) packages are allowed expected [Fyne](https://github.com/fyne-io/fyne)

## Help

- You can see an example of a RESTful API [here](https://rickandmortyapi.com/)

### Fyne

- You can see an example of a GUI with Fyne [here](https://apps.fyne.io/)
- General tutorial: [https://m.youtube.com/watch?v=-v1vz_NcWng](https://m.youtube.com/watch?v=-v1vz_NcWng)
- Documentation: [https://developer.fyne.io/](https://developer.fyne.io/)
- Tutorials on various features: [https://www.youtube.com/playlist?list=PL5vZ49dm2gshlo1EIxFNcQFBUfLFoHPfp](https://www.youtube.com/playlist?list=PL5vZ49dm2gshlo1EIxFNcQFBUfLFoHPfp)

### GCC

To install GCC for Fyne:

1. Install the .exe from https://jmeubank.github.io/tdm-gcc/download/

2. Use tdm64-gcc-10.3.0-2.exe

3. Add C:\TDM-GCC-64\bin to the PATH in User Variables and System Variables.

> Note: To edit the Path variable, press the Windows key, type 'path', choose 'Edit the system environment variables', click 'Environment Variables', find the Path variable in System variables and User variables, then edit it

## Bonus

- integrate spotify embed in your application (+1) (ex: [spotify](https://developer.spotify.com/documentation/widgets/generate/embed/))
- make a favorites system (+1)
