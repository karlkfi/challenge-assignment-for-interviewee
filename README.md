

Snurtle Flopping(TM) Competition - coding challenge
=========

Snurlte Flopping is the newest craze taking over all of North Central Tasmania!  These year's rounds were some of the best
we have ever seen. There are 30 amazing Flopper's in this year's prelimns.

In the prelimary rounds the competitors were given 60 minutes to score as many flops as possible.  But, some weren't able
to finish as many rounds as they needed to qualify for the finals because as we all know, the snurtle's are hard to catch 
and you can't score flops without catching a SNURTLE!


* Your task is to select the 3 finalists out of the 30 compeitiors.  To be eligible the competior MUST of accumulated enough
flops to accrue a score.  In order to qualify for the finals, they must have at least **7 scorable floppings.**

* The final score for each compeitor is the avg of the scores, **EXCLUDING the best and worst score.**

e.x.
```
person_1:
 - 90 # highest score is removed
 - 86
 - 65 
 - 87
 - 57 # lowest score is removed
 - 90
 - 77
 - 82
person_2: #disqualified, didn't score the min number of floppings
 - 99
 - 88
 - 99
 - 77
 - 85
 - 98
```

* Select the top 3 based on the above criteria, don't worry about rounding, but the averages should be floating point numbers

* output should be of the form
```
winners:
  - name: Cal Ripken
    avg: 76.3333333
  - name: Boy George
    avg: 73.2
  - name: Dwight Schrute
    avg: 70
disqualifications:
  - Dougie Fresh
  - Hubert Blaine Wolfeschlegelsteinhausenbergerdorff Sr.
```

* The scores are in the `scores.yaml` file
* Please output your resutls in yaml, into a file called `results.yaml`

