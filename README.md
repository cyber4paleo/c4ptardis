## C4P TARDIS

# details to come


# RDF scratch pad



 ```
 <http://host.geolink.org/sesar/id/dataset/ffaf89d9-8e7f-4770-b049-224d87eb2850>
     a <http://schema.geolink.org/1.0/base/main#Dataset> ;
     rdfs:label "RC1307, Coring>PistonCorer"
     geolink:hasCruise <http://lod.bco-dmo.org/geolink/id/deployment/57413>
     geolink:hasTitle
     geolink:hasAbstract

 ```


# Encoded CSIRO call

 http://resource.geosciml.org/sparql/isc2014?query=PREFIX+isc%3A+%3Chttp%3A%2F%2Fresource.geosciml.org%2Fclassifier%2Fics%2Fischart%2F%3E%0D%0APREFIX+gts%3A+%3Chttp%3A%2F%2Fresource.geosciml.org%2Fontology%2Ftimescale%2Fgts%23%3E%0D%0APREFIX+thors%3A+%3Chttp%3A%2F%2Fresource.geosciml.org%2Fontology%2Ftimescale%2Fthors%23%3E%0D%0APREFIX+tm%3A+%3Chttp%3A%2F%2Fdef.seegrid.csiro.au%2Fisotc211%2Fiso19108%2F2002%2Ftemporal%23%3E%0D%0APREFIX+rdfs%3A+%3Chttp%3A%2F%2Fwww.w3.org%2F2000%2F01%2Frdf-schema%23%3E%0D%0A%0D%0ASELECT+%2A%0D%0AWHERE+%7B%0D%0A++++++++++++++++%3Fera+gts%3Arank+%3Frank+.%0D%0A+++++++++++++++++%3Fera+thors%3Abegin%2Ftm%3AtemporalPosition%2Ftm%3Avalue+%3Fbegin+.%0D%0A+++++++++++++++++%3Fera+thors%3Abegin%2Ftm%3AtemporalPosition%2Ftm%3Aframe+%3Chttp%3A%2F%2Fresource.geosciml.org%2Fclassifier%2Fcgi%2Fgeologicage%2Fma%3E+.%0D%0A+++++++++++++++++%3Fera+thors%3Aend%2Ftm%3AtemporalPosition%2Ftm%3Avalue+%3Fend+.%0D%0A+++++++++++++++++%3Fera+thors%3Aend%2Ftm%3AtemporalPosition%2Ftm%3Aframe+%3Chttp%3A%2F%2Fresource.geosciml.org%2Fclassifier%2Fcgi%2Fgeologicage%2Fma%3E+.%0D%0A++++++++++++++++%3Fera+rdfs%3Alabel+%3Fname+.%0D%0A+++++++++++++++++BIND+%28+%22439.%22%5E%5Exsd%3Adecimal+AS+%3FtargetAge+%29%0D%0A+++++++++++++++++FILTER+%28+%3FtargetAge+%3E+xsd%3Adecimal%28%3Fend%29+%29%0D%0A+++++++++++++++++FILTER+%28+%3FtargetAge+%3C+xsd%3Adecimal%28%3Fbegin%29+%29%0D%0A%7D%0D%0A



