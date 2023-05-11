It is a basic Movie crud project without using databases.
Array of Movie struct is used for storing data.
Threre are three operations:-
    ROUTES      FUNCTION        ENDPOINTS     METHODS
1.  GETALL      getmovies       /movies       get
2.  get by id   getmovie        /movies/id    get
3.  create      createmovie     /movies/id    post
4.  update      updatemovie     /movies/id    put
5.  delete      deletemovie     /movies/id    delete

* Main function has routes by gotilla/mux, which 
  is used to create router and added to http.ListenAndServe.
