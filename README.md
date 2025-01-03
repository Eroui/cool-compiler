# cool-compiler

## Gramar
Will be adding things to the grammar as things progress to match specification of the Cool language 

<pre>
Program      : [[Class;]]+
Class        : class TYPE { [[feature;]]* } 
Feature      : Method | Attribute
Method       : ID(Formal [[,Formal]]*): TYPE { Expression }
             | ID(): TYPE { Expression }
Attribute    : ID:TYPE
             | ID:TYPE <- Expression
Formal       : ID:TYPE
Expression   : ID<-Expression 
             | ID(Expression[[,Expression]*)
             | ID()
             | if Expression then Expression else Expression fi
             | while Expression loop Expression pool
             | {[[Expression;]]+}
             | Expression + Expression
             | Expression - Expression 
             | Expression * Expression
             | Expression / Expression
             | Expression < Expression
             | Expression <= Expression
             | Expression = Expression
             | ID  
             | Integer 
             | String 
             | true | false 
< /pre>
