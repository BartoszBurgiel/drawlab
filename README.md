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
can := canvas.New(70, 20, "test")
can.Point(30, 10)
can.Circle(10, 10, 5)
can.Circle(30, 10, 10)
can.Line(20, 10, 50, 10)
can.Line(20, 10, 50, 17)
can.Triangle(55, 7, 60, 3, 68, 7)
can.Text(30, 7, "omg it's so cool")
can.Draw()
</code></pre>