# drawlab 

drawlab is a cool programming environment written in go where you can 
draw pictures via functions (similar to processing or p5.js) in the console 

# What does it can?
## supported shapes 
* Triangle -> triangle(x1, y1, x2, y2, x3, y3)
* Circle -> circle(x1, y2, radius) 
* Line -> line(x1, y2, x2, y2)
* Point -> point(x1, y2)
* Text -> text(x1, y2, "content")

## changes to environment 
* change character -> char(new character)
* change color -> color(new color)
    * supported colors include: 
        * red 
        * green 
        * cyan 
        * yellow 
        * blue 
        * white 
        * purple

## preview
Here you can see how drawn triangle, some text, lines and circles look like 
<pre><code>
                                 test                                 
========================================================================
|                                                                      |
|                  ----                ----                            |
|               ---                        ---                         |
|              --                            --             ---        |
|            --                                --          -   --      |
|    ------------                               --        -      --    |
|  --      --    --                              --      -         --  |
| -        -       --          omg it's so cool   -     -------------- |
|          -        -                             -                    |
|          -        ---                           -                    |
|          -        --------------------------------                   |
|          -        -      ----                   -                    |
| -        -       --          -----              -                    |
|  --      --    --                 ----         --                    |
|    ------------                       ----    --                     |
|          - --                             -----                      |
|              --                            --  ---                   |
|               ---                        ---                         |
|                  ----                ----                            |
|                                                                      |
========================================================================
</code></pre>

And here's it's code (golang)
<pre><code>
init(70, 20, "hello")

point(30, 10)
circle(10, 10, 5)
circle(30, 10, 10)
line(20, 10, 50, 10)
line(20, 10, 50, 17)
triangle(55, 7, 60, 3, 68, 7)
text(30, 7, "omg it's so cool")
</code></pre>