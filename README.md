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
</pre>
