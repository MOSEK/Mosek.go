/* MOSEK package.

   Thin wrapper for MOSEK solver API.
*/
package mosek

// /*<comment>*/
// // SEE https://github.com/golang/go/wiki/cgo#function-variables
//
// #include <stdlib.h>
// #include <stdint.h>
//
// typedef void * MSKuserhandle_t;
// typedef const int8_t * string_t;
// typedef void * MSKtask_t;
// typedef void * MSKenv_t;
//
// extern int MSK_analyzenames(MSKtask_t,int32_t,int32_t);
// extern int MSK_analyzeproblem(MSKtask_t,int32_t);
// extern int MSK_analyzesolution(MSKtask_t,int32_t,int32_t);
// extern int MSK_appendacc(MSKtask_t,int64_t,int64_t,int64_t *,double *);
// extern int MSK_appendaccs(MSKtask_t,int64_t,int64_t *,int64_t,int64_t *,double *);
// extern int MSK_appendaccseq(MSKtask_t,int64_t,int64_t,int64_t,double *);
// extern int MSK_appendaccsseq(MSKtask_t,int64_t,int64_t *,int64_t,int64_t,double *);
// extern int MSK_appendafes(MSKtask_t,int64_t);
// extern int MSK_appendbarvars(MSKtask_t,int32_t,int32_t *);
// extern int MSK_appendcone(MSKtask_t,int32_t,double,int32_t,int32_t *);
// extern int MSK_appendconeseq(MSKtask_t,int32_t,double,int32_t,int32_t);
// extern int MSK_appendconesseq(MSKtask_t,int32_t,int32_t *,double *,int32_t *,int32_t);
// extern int MSK_appendcons(MSKtask_t,int32_t);
// extern int MSK_appenddjcs(MSKtask_t,int64_t);
// extern int MSK_appenddualexpconedomain(MSKtask_t,int64_t *);
// extern int MSK_appenddualgeomeanconedomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appenddualpowerconedomain(MSKtask_t,int64_t,int64_t,double *,int64_t *);
// extern int MSK_appendprimalexpconedomain(MSKtask_t,int64_t *);
// extern int MSK_appendprimalgeomeanconedomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendprimalpowerconedomain(MSKtask_t,int64_t,int64_t,double *,int64_t *);
// extern int MSK_appendquadraticconedomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendrdomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendrminusdomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendrplusdomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendrquadraticconedomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendrzerodomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendsparsesymmat(MSKtask_t,int32_t,int64_t,int32_t *,int32_t *,double *,int64_t *);
// extern int MSK_appendsparsesymmatlist(MSKtask_t,int32_t,int32_t *,int64_t *,int32_t *,int32_t *,double *,int64_t *);
// extern int MSK_appendsvecpsdconedomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendvars(MSKtask_t,int32_t);
// extern int MSK_basiscond(MSKtask_t,double *,double *);
// extern int MSK_checkmemtask(MSKtask_t,const char *,int32_t);
// extern int MSK_chgconbound(MSKtask_t,int32_t,int32_t,int32_t,double);
// extern int MSK_chgvarbound(MSKtask_t,int32_t,int32_t,int32_t,double);
// extern int MSK_commitchanges(MSKtask_t);
// extern int MSK_deletesolution(MSKtask_t,int32_t);
// extern int MSK_dualsensitivity(MSKtask_t,int32_t,int32_t *,double *,double *,double *,double *);
// extern int MSK_emptyafebarfrow(MSKtask_t,int64_t);
// extern int MSK_emptyafebarfrowlist(MSKtask_t,int64_t,int64_t *);
// extern int MSK_emptyafefcol(MSKtask_t,int32_t);
// extern int MSK_emptyafefcollist(MSKtask_t,int64_t,int32_t *);
// extern int MSK_emptyafefrow(MSKtask_t,int64_t);
// extern int MSK_emptyafefrowlist(MSKtask_t,int64_t,int64_t *);
// extern int MSK_evaluateacc(MSKtask_t,int32_t,int64_t,double *);
// extern int MSK_evaluateaccs(MSKtask_t,int32_t,double *);
// extern int MSK_generateaccnames(MSKtask_t,int64_t,int64_t *,const char *,int32_t,int32_t *,int64_t *,int32_t,int32_t *,int64_t,const char **);
// extern int MSK_generatebarvarnames(MSKtask_t,int32_t,int32_t *,const char *,int32_t,int32_t *,int64_t *,int32_t,int32_t *,int64_t,const char **);
// extern int MSK_generateconenames(MSKtask_t,int32_t,int32_t *,const char *,int32_t,int32_t *,int64_t *,int32_t,int32_t *,int64_t,const char **);
// extern int MSK_generateconnames(MSKtask_t,int32_t,int32_t *,const char *,int32_t,int32_t *,int64_t *,int32_t,int32_t *,int64_t,const char **);
// extern int MSK_generatedjcnames(MSKtask_t,int64_t,int64_t *,const char *,int32_t,int32_t *,int64_t *,int32_t,int32_t *,int64_t,const char **);
// extern int MSK_generatevarnames(MSKtask_t,int32_t,int32_t *,const char *,int32_t,int32_t *,int64_t *,int32_t,int32_t *,int64_t,const char **);
// extern int MSK_getaccafeidxlist(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getaccb(MSKtask_t,int64_t,double *);
// extern int MSK_getaccbarfblocktriplet(MSKtask_t,int64_t,int64_t *,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_getaccbarfnumblocktriplets(MSKtask_t,int64_t *);
// extern int MSK_getaccdomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getaccdoty(MSKtask_t,int32_t,int64_t,double *);
// extern int MSK_getaccdotys(MSKtask_t,int32_t,double *);
// extern int MSK_getaccfnumnz(MSKtask_t,int64_t *);
// extern int MSK_getaccftrip(MSKtask_t,int64_t *,int32_t *,double *);
// extern int MSK_getaccgvector(MSKtask_t,double *);
// extern int MSK_getaccn(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getaccname(MSKtask_t,int64_t,int32_t,unsigned char *);
// extern int MSK_getaccnamelen(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getaccntot(MSKtask_t,int64_t *);
// extern int MSK_getaccs(MSKtask_t,int64_t *,int64_t *,double *);
// extern int MSK_getacol(MSKtask_t,int32_t,int32_t *,int32_t *,double *);
// extern int MSK_getacolnumnz(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getacolslice64(MSKtask_t,int32_t,int32_t,int64_t,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_getacolslicenumnz64(MSKtask_t,int32_t,int32_t,int64_t *);
// extern int MSK_getacolslicetrip(MSKtask_t,int32_t,int32_t,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_getafebarfblocktriplet(MSKtask_t,int64_t,int64_t *,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_getafebarfnumblocktriplets(MSKtask_t,int64_t *);
// extern int MSK_getafebarfnumrowentries(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getafebarfrow(MSKtask_t,int64_t,int32_t *,int64_t *,int64_t *,int64_t *,double *);
// extern int MSK_getafebarfrowinfo(MSKtask_t,int64_t,int32_t *,int64_t *);
// extern int MSK_getafefnumnz(MSKtask_t,int64_t *);
// extern int MSK_getafefrow(MSKtask_t,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_getafefrownumnz(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getafeftrip(MSKtask_t,int64_t *,int32_t *,double *);
// extern int MSK_getafeg(MSKtask_t,int64_t,double *);
// extern int MSK_getafegslice(MSKtask_t,int64_t,int64_t,double *);
// extern int MSK_getaij(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_getapiecenumnz(MSKtask_t,int32_t,int32_t,int32_t,int32_t,int32_t *);
// extern int MSK_getarow(MSKtask_t,int32_t,int32_t *,int32_t *,double *);
// extern int MSK_getarownumnz(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getarowslice64(MSKtask_t,int32_t,int32_t,int64_t,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_getarowslicenumnz64(MSKtask_t,int32_t,int32_t,int64_t *);
// extern int MSK_getarowslicetrip(MSKtask_t,int32_t,int32_t,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_getatrip(MSKtask_t,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_getatruncatetol(MSKtask_t,double *);
// extern int MSK_getbarablocktriplet(MSKtask_t,int64_t,int64_t *,int32_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_getbaraidx(MSKtask_t,int64_t,int64_t,int32_t *,int32_t *,int64_t *,int64_t *,double *);
// extern int MSK_getbaraidxij(MSKtask_t,int64_t,int32_t *,int32_t *);
// extern int MSK_getbaraidxinfo(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getbarasparsity(MSKtask_t,int64_t,int64_t *,int64_t *);
// extern int MSK_getbarcblocktriplet(MSKtask_t,int64_t,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_getbarcidx(MSKtask_t,int64_t,int64_t,int32_t *,int64_t *,int64_t *,double *);
// extern int MSK_getbarcidxinfo(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getbarcidxj(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getbarcsparsity(MSKtask_t,int64_t,int64_t *,int64_t *);
// extern int MSK_getbarsj(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_getbarsslice(MSKtask_t,int32_t,int32_t,int32_t,int64_t,double *);
// extern int MSK_getbarvarname(MSKtask_t,int32_t,int32_t,unsigned char *);
// extern int MSK_getbarvarnameindex(MSKtask_t,const char *,int32_t *,int32_t *);
// extern int MSK_getbarvarnamelen(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getbarxj(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_getbarxslice(MSKtask_t,int32_t,int32_t,int32_t,int64_t,double *);
// extern int MSK_getc(MSKtask_t,double *);
// extern int MSK_getcfix(MSKtask_t,double *);
// extern int MSK_getcj(MSKtask_t,int32_t,double *);
// extern int MSK_getclist(MSKtask_t,int32_t,int32_t *,double *);
// extern int MSK_getconbound(MSKtask_t,int32_t,int32_t *,double *,double *);
// extern int MSK_getconboundslice(MSKtask_t,int32_t,int32_t,int32_t *,double *,double *);
// extern int MSK_getcone(MSKtask_t,int32_t,int32_t *,double *,int32_t *,int32_t *);
// extern int MSK_getconeinfo(MSKtask_t,int32_t,int32_t *,double *,int32_t *);
// extern int MSK_getconename(MSKtask_t,int32_t,int32_t,unsigned char *);
// extern int MSK_getconenameindex(MSKtask_t,const char *,int32_t *,int32_t *);
// extern int MSK_getconenamelen(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getconname(MSKtask_t,int32_t,int32_t,unsigned char *);
// extern int MSK_getconnameindex(MSKtask_t,const char *,int32_t *,int32_t *);
// extern int MSK_getconnamelen(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getcslice(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_getdimbarvarj(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getdjcafeidxlist(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdjcb(MSKtask_t,int64_t,double *);
// extern int MSK_getdjcdomainidxlist(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdjcname(MSKtask_t,int64_t,int32_t,unsigned char *);
// extern int MSK_getdjcnamelen(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getdjcnumafe(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdjcnumafetot(MSKtask_t,int64_t *);
// extern int MSK_getdjcnumdomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdjcnumdomaintot(MSKtask_t,int64_t *);
// extern int MSK_getdjcnumterm(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdjcnumtermtot(MSKtask_t,int64_t *);
// extern int MSK_getdjcs(MSKtask_t,int64_t *,int64_t *,double *,int64_t *,int64_t *);
// extern int MSK_getdjctermsizelist(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdomainn(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdomainname(MSKtask_t,int64_t,int32_t,unsigned char *);
// extern int MSK_getdomainnamelen(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getdomaintype(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getdouinf(MSKtask_t,int32_t,double *);
// extern int MSK_getdouparam(MSKtask_t,int32_t,double *);
// extern int MSK_getdualobj(MSKtask_t,int32_t,double *);
// extern int MSK_getdualsolutionnorms(MSKtask_t,int32_t,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getdviolacc(MSKtask_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_getdviolbarvar(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getdviolcon(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getdviolcones(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getdviolvar(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getinfeasiblesubproblem(MSKtask_t,int32_t,MSKtask_t*);
// extern int MSK_getinfindex(MSKtask_t,int32_t,const char *,int32_t *);
// extern int MSK_getinfmax(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getinfname(MSKtask_t,int32_t,int32_t,unsigned char *);
// extern int MSK_getintinf(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getintparam(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getlenbarvarj(MSKtask_t,int32_t,int64_t *);
// extern int MSK_getlintinf(MSKtask_t,int32_t,int64_t *);
// extern int MSK_getmaxnumanz64(MSKtask_t,int64_t *);
// extern int MSK_getmaxnumbarvar(MSKtask_t,int32_t *);
// extern int MSK_getmaxnumcon(MSKtask_t,int32_t *);
// extern int MSK_getmaxnumcone(MSKtask_t,int32_t *);
// extern int MSK_getmaxnumqnz64(MSKtask_t,int64_t *);
// extern int MSK_getmaxnumvar(MSKtask_t,int32_t *);
// extern int MSK_getmemusagetask(MSKtask_t,int64_t *,int64_t *);
// extern int MSK_getnumacc(MSKtask_t,int64_t *);
// extern int MSK_getnumafe(MSKtask_t,int64_t *);
// extern int MSK_getnumanz(MSKtask_t,int32_t *);
// extern int MSK_getnumanz64(MSKtask_t,int64_t *);
// extern int MSK_getnumbarablocktriplets(MSKtask_t,int64_t *);
// extern int MSK_getnumbaranz(MSKtask_t,int64_t *);
// extern int MSK_getnumbarcblocktriplets(MSKtask_t,int64_t *);
// extern int MSK_getnumbarcnz(MSKtask_t,int64_t *);
// extern int MSK_getnumbarvar(MSKtask_t,int32_t *);
// extern int MSK_getnumcon(MSKtask_t,int32_t *);
// extern int MSK_getnumcone(MSKtask_t,int32_t *);
// extern int MSK_getnumconemem(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getnumdjc(MSKtask_t,int64_t *);
// extern int MSK_getnumdomain(MSKtask_t,int64_t *);
// extern int MSK_getnumintvar(MSKtask_t,int32_t *);
// extern int MSK_getnumparam(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getnumqconknz64(MSKtask_t,int32_t,int64_t *);
// extern int MSK_getnumqobjnz64(MSKtask_t,int64_t *);
// extern int MSK_getnumsymmat(MSKtask_t,int64_t *);
// extern int MSK_getnumvar(MSKtask_t,int32_t *);
// extern int MSK_getobjname(MSKtask_t,int32_t,unsigned char *);
// extern int MSK_getobjnamelen(MSKtask_t,int32_t *);
// extern int MSK_getobjsense(MSKtask_t,int32_t *);
// extern int MSK_getparammax(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getparamname(MSKtask_t,int32_t,int32_t,unsigned char *);
// extern int MSK_getpowerdomainalpha(MSKtask_t,int64_t,double *);
// extern int MSK_getpowerdomaininfo(MSKtask_t,int64_t,int64_t *,int64_t *);
// extern int MSK_getprimalobj(MSKtask_t,int32_t,double *);
// extern int MSK_getprimalsolutionnorms(MSKtask_t,int32_t,double *,double *,double *);
// extern int MSK_getprobtype(MSKtask_t,int32_t *);
// extern int MSK_getprosta(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getpviolacc(MSKtask_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_getpviolbarvar(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getpviolcon(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getpviolcones(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getpvioldjc(MSKtask_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_getpviolvar(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getqconk64(MSKtask_t,int32_t,int64_t,int64_t *,int32_t *,int32_t *,double *);
// extern int MSK_getqobj64(MSKtask_t,int64_t,int64_t *,int32_t *,int32_t *,double *);
// extern int MSK_getqobjij(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_getreducedcosts(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getskc(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getskcslice(MSKtask_t,int32_t,int32_t,int32_t,int32_t *);
// extern int MSK_getskn(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getskx(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getskxslice(MSKtask_t,int32_t,int32_t,int32_t,int32_t *);
// extern int MSK_getslc(MSKtask_t,int32_t,double *);
// extern int MSK_getslcslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getslx(MSKtask_t,int32_t,double *);
// extern int MSK_getslxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getsnx(MSKtask_t,int32_t,double *);
// extern int MSK_getsnxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getsolsta(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getsolution(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t *,int32_t *,int32_t *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getsolutioninfo(MSKtask_t,int32_t,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getsolutioninfonew(MSKtask_t,int32_t,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getsolutionnew(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t *,int32_t *,int32_t *,double *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getsolutionslice(MSKtask_t,int32_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getsparsesymmat(MSKtask_t,int64_t,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_getstrparam(MSKtask_t,int32_t,int32_t,int32_t *,unsigned char *);
// extern int MSK_getstrparamlen(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getsuc(MSKtask_t,int32_t,double *);
// extern int MSK_getsucslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getsux(MSKtask_t,int32_t,double *);
// extern int MSK_getsuxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getsymmatinfo(MSKtask_t,int64_t,int32_t *,int64_t *,int32_t *);
// extern int MSK_gettaskname(MSKtask_t,int32_t,unsigned char *);
// extern int MSK_gettasknamelen(MSKtask_t,int32_t *);
// extern int MSK_getvarbound(MSKtask_t,int32_t,int32_t *,double *,double *);
// extern int MSK_getvarboundslice(MSKtask_t,int32_t,int32_t,int32_t *,double *,double *);
// extern int MSK_getvarname(MSKtask_t,int32_t,int32_t,unsigned char *);
// extern int MSK_getvarnameindex(MSKtask_t,const char *,int32_t *,int32_t *);
// extern int MSK_getvarnamelen(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getvartype(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getvartypelist(MSKtask_t,int32_t,int32_t *,int32_t *);
// extern int MSK_getxc(MSKtask_t,int32_t,double *);
// extern int MSK_getxcslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getxx(MSKtask_t,int32_t,double *);
// extern int MSK_getxxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_gety(MSKtask_t,int32_t,double *);
// extern int MSK_getyslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_infeasibilityreport(MSKtask_t,int32_t,int32_t);
// extern int MSK_initbasissolve(MSKtask_t,int32_t *);
// extern int MSK_inputdata64(MSKtask_t,int32_t,int32_t,int32_t,int32_t,double *,double,int64_t *,int64_t *,int32_t *,double *,int32_t *,double *,double *,int32_t *,double *,double *);
// extern int MSK_isdouparname(MSKtask_t,const char *,int32_t *);
// extern int MSK_isintparname(MSKtask_t,const char *,int32_t *);
// extern int MSK_isstrparname(MSKtask_t,const char *,int32_t *);
// extern int MSK_linkfiletotaskstream(MSKtask_t,int32_t,const char *,int32_t);
// extern int MSK_onesolutionsummary(MSKtask_t,int32_t,int32_t);
// extern int MSK_optimizersummary(MSKtask_t,int32_t);
// extern int MSK_optimizetrm(MSKtask_t,int32_t *);
// extern int MSK_primalrepair(MSKtask_t,double *,double *,double *,double *);
// extern int MSK_primalsensitivity(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t,int32_t *,int32_t *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_probtypetostr(MSKtask_t,int32_t,unsigned char *);
// extern int MSK_prostatostr(MSKtask_t,int32_t,unsigned char *);
// extern int MSK_putacc(MSKtask_t,int64_t,int64_t,int64_t,int64_t *,double *);
// extern int MSK_putaccb(MSKtask_t,int64_t,int64_t,double *);
// extern int MSK_putaccbj(MSKtask_t,int64_t,int64_t,double);
// extern int MSK_putaccdoty(MSKtask_t,int32_t,int64_t,double *);
// extern int MSK_putacclist(MSKtask_t,int64_t,int64_t *,int64_t *,int64_t,int64_t *,double *);
// extern int MSK_putaccname(MSKtask_t,int64_t,const char *);
// extern int MSK_putacol(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_putacollist(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putacolslice64(MSKtask_t,int32_t,int32_t,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_putafebarfblocktriplet(MSKtask_t,int64_t,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putafebarfentry(MSKtask_t,int64_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_putafebarfentrylist(MSKtask_t,int64_t,int64_t *,int32_t *,int64_t *,int64_t *,int64_t,int64_t *,double *);
// extern int MSK_putafebarfrow(MSKtask_t,int64_t,int32_t,int32_t *,int64_t *,int64_t *,int64_t,int64_t *,double *);
// extern int MSK_putafefcol(MSKtask_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_putafefentry(MSKtask_t,int64_t,int32_t,double);
// extern int MSK_putafefentrylist(MSKtask_t,int64_t,int64_t *,int32_t *,double *);
// extern int MSK_putafefrow(MSKtask_t,int64_t,int32_t,int32_t *,double *);
// extern int MSK_putafefrowlist(MSKtask_t,int64_t,int64_t *,int32_t *,int64_t *,int64_t,int32_t *,double *);
// extern int MSK_putafeg(MSKtask_t,int64_t,double);
// extern int MSK_putafeglist(MSKtask_t,int64_t,int64_t *,double *);
// extern int MSK_putafegslice(MSKtask_t,int64_t,int64_t,double *);
// extern int MSK_putaij(MSKtask_t,int32_t,int32_t,double);
// extern int MSK_putaijlist64(MSKtask_t,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_putarow(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_putarowlist64(MSKtask_t,int32_t,int32_t *,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_putarowslice64(MSKtask_t,int32_t,int32_t,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_putatruncatetol(MSKtask_t,double);
// extern int MSK_putbarablocktriplet(MSKtask_t,int64_t,int32_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putbaraij(MSKtask_t,int32_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_putbaraijlist(MSKtask_t,int32_t,int32_t *,int32_t *,int64_t *,int64_t *,int64_t *,double *);
// extern int MSK_putbararowlist(MSKtask_t,int32_t,int32_t *,int64_t *,int64_t *,int32_t *,int64_t *,int64_t *,double *);
// extern int MSK_putbarcblocktriplet(MSKtask_t,int64_t,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putbarcj(MSKtask_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_putbarsj(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_putbarvarname(MSKtask_t,int32_t,const char *);
// extern int MSK_putbarxj(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_putcfix(MSKtask_t,double);
// extern int MSK_putcj(MSKtask_t,int32_t,double);
// extern int MSK_putclist(MSKtask_t,int32_t,int32_t *,double *);
// extern int MSK_putconbound(MSKtask_t,int32_t,int32_t,double,double);
// extern int MSK_putconboundlist(MSKtask_t,int32_t,int32_t *,int32_t *,double *,double *);
// extern int MSK_putconboundlistconst(MSKtask_t,int32_t,int32_t *,int32_t,double,double);
// extern int MSK_putconboundslice(MSKtask_t,int32_t,int32_t,int32_t *,double *,double *);
// extern int MSK_putconboundsliceconst(MSKtask_t,int32_t,int32_t,int32_t,double,double);
// extern int MSK_putcone(MSKtask_t,int32_t,int32_t,double,int32_t,int32_t *);
// extern int MSK_putconename(MSKtask_t,int32_t,const char *);
// extern int MSK_putconname(MSKtask_t,int32_t,const char *);
// extern int MSK_putconsolutioni(MSKtask_t,int32_t,int32_t,int32_t,double,double,double);
// extern int MSK_putcslice(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_putdjc(MSKtask_t,int64_t,int64_t,int64_t *,int64_t,int64_t *,double *,int64_t,int64_t *);
// extern int MSK_putdjcname(MSKtask_t,int64_t,const char *);
// extern int MSK_putdjcslice(MSKtask_t,int64_t,int64_t,int64_t,int64_t *,int64_t,int64_t *,double *,int64_t,int64_t *,int64_t *);
// extern int MSK_putdomainname(MSKtask_t,int64_t,const char *);
// extern int MSK_putdouparam(MSKtask_t,int32_t,double);
// extern int MSK_putintparam(MSKtask_t,int32_t,int32_t);
// extern int MSK_putmaxnumacc(MSKtask_t,int64_t);
// extern int MSK_putmaxnumafe(MSKtask_t,int64_t);
// extern int MSK_putmaxnumanz(MSKtask_t,int64_t);
// extern int MSK_putmaxnumbarvar(MSKtask_t,int32_t);
// extern int MSK_putmaxnumcon(MSKtask_t,int32_t);
// extern int MSK_putmaxnumcone(MSKtask_t,int32_t);
// extern int MSK_putmaxnumdjc(MSKtask_t,int64_t);
// extern int MSK_putmaxnumdomain(MSKtask_t,int64_t);
// extern int MSK_putmaxnumqnz(MSKtask_t,int64_t);
// extern int MSK_putmaxnumvar(MSKtask_t,int32_t);
// extern int MSK_putnadouparam(MSKtask_t,const char *,double);
// extern int MSK_putnaintparam(MSKtask_t,const char *,int32_t);
// extern int MSK_putnastrparam(MSKtask_t,const char *,const char *);
// extern int MSK_putobjname(MSKtask_t,const char *);
// extern int MSK_putobjsense(MSKtask_t,int32_t);
// extern int MSK_putoptserverhost(MSKtask_t,const char *);
// extern int MSK_putparam(MSKtask_t,const char *,const char *);
// extern int MSK_putqcon(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putqconk(MSKtask_t,int32_t,int32_t,int32_t *,int32_t *,double *);
// extern int MSK_putqobj(MSKtask_t,int32_t,int32_t *,int32_t *,double *);
// extern int MSK_putqobjij(MSKtask_t,int32_t,int32_t,double);
// extern int MSK_putskc(MSKtask_t,int32_t,int32_t *);
// extern int MSK_putskcslice(MSKtask_t,int32_t,int32_t,int32_t,int32_t *);
// extern int MSK_putskx(MSKtask_t,int32_t,int32_t *);
// extern int MSK_putskxslice(MSKtask_t,int32_t,int32_t,int32_t,int32_t *);
// extern int MSK_putslc(MSKtask_t,int32_t,double *);
// extern int MSK_putslcslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_putslx(MSKtask_t,int32_t,double *);
// extern int MSK_putslxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_putsnx(MSKtask_t,int32_t,double *);
// extern int MSK_putsnxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_putsolution(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_putsolutionnew(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t *,double *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_putsolutionyi(MSKtask_t,int32_t,int32_t,double);
// extern int MSK_putstrparam(MSKtask_t,int32_t,const char *);
// extern int MSK_putsuc(MSKtask_t,int32_t,double *);
// extern int MSK_putsucslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_putsux(MSKtask_t,int32_t,double *);
// extern int MSK_putsuxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_puttaskname(MSKtask_t,const char *);
// extern int MSK_putvarbound(MSKtask_t,int32_t,int32_t,double,double);
// extern int MSK_putvarboundlist(MSKtask_t,int32_t,int32_t *,int32_t *,double *,double *);
// extern int MSK_putvarboundlistconst(MSKtask_t,int32_t,int32_t *,int32_t,double,double);
// extern int MSK_putvarboundslice(MSKtask_t,int32_t,int32_t,int32_t *,double *,double *);
// extern int MSK_putvarboundsliceconst(MSKtask_t,int32_t,int32_t,int32_t,double,double);
// extern int MSK_putvarname(MSKtask_t,int32_t,const char *);
// extern int MSK_putvarsolutionj(MSKtask_t,int32_t,int32_t,int32_t,double,double,double,double);
// extern int MSK_putvartype(MSKtask_t,int32_t,int32_t);
// extern int MSK_putvartypelist(MSKtask_t,int32_t,int32_t *,int32_t *);
// extern int MSK_putxc(MSKtask_t,int32_t,double *);
// extern int MSK_putxcslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_putxx(MSKtask_t,int32_t,double *);
// extern int MSK_putxxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_puty(MSKtask_t,int32_t,double *);
// extern int MSK_putyslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_readbsolution(MSKtask_t,const char *,int32_t);
// extern int MSK_readdataautoformat(MSKtask_t,const char *);
// extern int MSK_readdataformat(MSKtask_t,const char *,int32_t,int32_t);
// extern int MSK_readjsonsol(MSKtask_t,const char *);
// extern int MSK_readjsonstring(MSKtask_t,const char *);
// extern int MSK_readlpstring(MSKtask_t,const char *);
// extern int MSK_readopfstring(MSKtask_t,const char *);
// extern int MSK_readparamfile(MSKtask_t,const char *);
// extern int MSK_readptfstring(MSKtask_t,const char *);
// extern int MSK_readsolution(MSKtask_t,int32_t,const char *);
// extern int MSK_readsolutionfile(MSKtask_t,const char *);
// extern int MSK_readsummary(MSKtask_t,int32_t);
// extern int MSK_readtask(MSKtask_t,const char *);
// extern int MSK_removebarvars(MSKtask_t,int32_t,int32_t *);
// extern int MSK_removecones(MSKtask_t,int32_t,int32_t *);
// extern int MSK_removecons(MSKtask_t,int32_t,int32_t *);
// extern int MSK_removevars(MSKtask_t,int32_t,int32_t *);
// extern int MSK_resizetask(MSKtask_t,int32_t,int32_t,int32_t,int64_t,int64_t);
// extern int MSK_sensitivityreport(MSKtask_t,int32_t);
// extern int MSK_setdefaults(MSKtask_t);
// extern int MSK_sktostr(MSKtask_t,int32_t,unsigned char *);
// extern int MSK_solstatostr(MSKtask_t,int32_t,unsigned char *);
// extern int MSK_solutiondef(MSKtask_t,int32_t,int *);
// extern int MSK_solutionsummary(MSKtask_t,int32_t);
// extern int MSK_solvewithbasis(MSKtask_t,int,int32_t,int32_t *,double *,int32_t *);
// extern int MSK_strtoconetype(MSKtask_t,const char *,int32_t *);
// extern int MSK_strtosk(MSKtask_t,const char *,int32_t *);
// extern int MSK_toconic(MSKtask_t);
// extern int MSK_updatesolutioninfo(MSKtask_t,int32_t);
// extern int MSK_writebsolution(MSKtask_t,const char *,int32_t);
// extern int MSK_writedata(MSKtask_t,const char *);
// extern int MSK_writejsonsol(MSKtask_t,const char *);
// extern int MSK_writeparamfile(MSKtask_t,const char *);
// extern int MSK_writesolution(MSKtask_t,int32_t,const char *);
// extern int MSK_writesolutionfile(MSKtask_t,const char *);
// extern int MSK_writetask(MSKtask_t,const char *);
// extern int MSK_axpy(MSKenv_t,int32_t,double,double *,double *);
// extern int MSK_checkinall(MSKenv_t);
// extern int MSK_checkinlicense(MSKenv_t,int32_t);
// extern int MSK_checkoutlicense(MSKenv_t,int32_t);
// extern int MSK_dot(MSKenv_t,int32_t,double *,double *,double *);
// extern int MSK_echointro(MSKenv_t,int32_t);
// extern int MSK_expirylicenses(MSKenv_t,int64_t *);
// extern int MSK_gemm(MSKenv_t,int32_t,int32_t,int32_t,int32_t,int32_t,double,double *,double *,double,double *);
// extern int MSK_gemv(MSKenv_t,int32_t,int32_t,int32_t,double,double *,double *,double,double *);
// extern int MSK_getcodedesc(int32_t,unsigned char *,unsigned char *);
// extern int MSK_getversion(int32_t *,int32_t *,int32_t *);
// extern int MSK_licensecleanup();
// extern int MSK_linkfiletoenvstream(MSKenv_t,int32_t,const char *,int32_t);
// extern int MSK_potrf(MSKenv_t,int32_t,int32_t,double *);
// extern int MSK_putlicensecode(MSKenv_t,int32_t *);
// extern int MSK_putlicensedebug(MSKenv_t,int32_t);
// extern int MSK_putlicensepath(MSKenv_t,const char *);
// extern int MSK_putlicensewait(MSKenv_t,int32_t);
// extern int MSK_resetexpirylicenses(MSKenv_t);
// extern int MSK_syeig(MSKenv_t,int32_t,int32_t,double *,double *);
// extern int MSK_syevd(MSKenv_t,int32_t,int32_t,double *,double *);
// extern int MSK_syrk(MSKenv_t,int32_t,int32_t,int32_t,int32_t,double,double *,double,double *);
import "C"

import (
	"bytes"
)

// Constants

type Basindtype int32
// Basis identification
const (
    MSK_BI_NEVER Basindtype = 0 // Never do basis identification.
    MSK_BI_ALWAYS Basindtype = 1 // Basis identification is always performed even if the interior-point optimizer terminates abnormally.
    MSK_BI_NO_ERROR Basindtype = 2 // Basis identification is performed if the interior-point optimizer terminates without an error.
    MSK_BI_IF_FEASIBLE Basindtype = 3 // Basis identification is not performed if the interior-point optimizer terminates with a problem status saying that the problem is primal or dual infeasible.
    MSK_BI_RESERVERED Basindtype = 4 // Not currently in use.
    MSK_BI_END Basindtype = 4
)
func (self Basindtype) String() string {
  switch self {
    case MSK_BI_NEVER: return "MSK_BI_NEVER"
    case MSK_BI_ALWAYS: return "MSK_BI_ALWAYS"
    case MSK_BI_NO_ERROR: return "MSK_BI_NO_ERROR"
    case MSK_BI_IF_FEASIBLE: return "MSK_BI_IF_FEASIBLE"
    case MSK_BI_RESERVERED: return "MSK_BI_RESERVERED"
    default: return "<unknown>"
  }
} // Basindtype.String()

type Boundkey int32
// Bound keys
const (
    MSK_BK_LO Boundkey = 0 // The constraint or variable has a finite lower bound and an infinite upper bound.
    MSK_BK_UP Boundkey = 1 // The constraint or variable has an infinite lower bound and an finite upper bound.
    MSK_BK_FX Boundkey = 2 // The constraint or variable is fixed.
    MSK_BK_FR Boundkey = 3 // The constraint or variable is free.
    MSK_BK_RA Boundkey = 4 // The constraint or variable is ranged.
    MSK_BK_END Boundkey = 4
)
func (self Boundkey) String() string {
  switch self {
    case MSK_BK_LO: return "MSK_BK_LO"
    case MSK_BK_UP: return "MSK_BK_UP"
    case MSK_BK_FX: return "MSK_BK_FX"
    case MSK_BK_FR: return "MSK_BK_FR"
    case MSK_BK_RA: return "MSK_BK_RA"
    default: return "<unknown>"
  }
} // Boundkey.String()

type Mark int32
// Mark
const (
    MSK_MARK_LO Mark = 0 // The lower bound is selected for sensitivity analysis.
    MSK_MARK_UP Mark = 1 // The upper bound is selected for sensitivity analysis.
    MSK_MARK_END Mark = 1
)
func (self Mark) String() string {
  switch self {
    case MSK_MARK_LO: return "MSK_MARK_LO"
    case MSK_MARK_UP: return "MSK_MARK_UP"
    default: return "<unknown>"
  }
} // Mark.String()

type Simdegen int32
// Degeneracy strategies
const (
    MSK_SIM_DEGEN_NONE Simdegen = 0 // The simplex optimizer should use no degeneration strategy.
    MSK_SIM_DEGEN_FREE Simdegen = 1 // The simplex optimizer chooses the degeneration strategy.
    MSK_SIM_DEGEN_AGGRESSIVE Simdegen = 2 // The simplex optimizer should use an aggressive degeneration strategy.
    MSK_SIM_DEGEN_MODERATE Simdegen = 3 // The simplex optimizer should use a moderate degeneration strategy.
    MSK_SIM_DEGEN_MINIMUM Simdegen = 4 // The simplex optimizer should use a minimum degeneration strategy.
    MSK_SIM_DEGEN_END Simdegen = 4
)
func (self Simdegen) String() string {
  switch self {
    case MSK_SIM_DEGEN_NONE: return "MSK_SIM_DEGEN_NONE"
    case MSK_SIM_DEGEN_FREE: return "MSK_SIM_DEGEN_FREE"
    case MSK_SIM_DEGEN_AGGRESSIVE: return "MSK_SIM_DEGEN_AGGRESSIVE"
    case MSK_SIM_DEGEN_MODERATE: return "MSK_SIM_DEGEN_MODERATE"
    case MSK_SIM_DEGEN_MINIMUM: return "MSK_SIM_DEGEN_MINIMUM"
    default: return "<unknown>"
  }
} // Simdegen.String()

type Transpose int32
// Transposed matrix.
const (
    MSK_TRANSPOSE_NO Transpose = 0 // No transpose is applied.
    MSK_TRANSPOSE_YES Transpose = 1 // A transpose is applied.
    MSK_TRANSPOSE_END Transpose = 1
)
func (self Transpose) String() string {
  switch self {
    case MSK_TRANSPOSE_NO: return "MSK_TRANSPOSE_NO"
    case MSK_TRANSPOSE_YES: return "MSK_TRANSPOSE_YES"
    default: return "<unknown>"
  }
} // Transpose.String()

type Uplo int32
// Triangular part of a symmetric matrix.
const (
    MSK_UPLO_LO Uplo = 0 // Lower part.
    MSK_UPLO_UP Uplo = 1 // Upper part.
    MSK_UPLO_END Uplo = 1
)
func (self Uplo) String() string {
  switch self {
    case MSK_UPLO_LO: return "MSK_UPLO_LO"
    case MSK_UPLO_UP: return "MSK_UPLO_UP"
    default: return "<unknown>"
  }
} // Uplo.String()

type Simreform int32
// Problem reformulation.
const (
    MSK_SIM_REFORMULATION_OFF Simreform = 0 // Disallow the simplex optimizer to reformulate the problem.
    MSK_SIM_REFORMULATION_ON Simreform = 1 // Allow the simplex optimizer to reformulate the problem.
    MSK_SIM_REFORMULATION_FREE Simreform = 2 // The simplex optimizer can choose freely.
    MSK_SIM_REFORMULATION_AGGRESSIVE Simreform = 3 // The simplex optimizer should use an aggressive reformulation strategy.
    MSK_SIM_REFORMULATION_END Simreform = 3
)
func (self Simreform) String() string {
  switch self {
    case MSK_SIM_REFORMULATION_OFF: return "MSK_SIM_REFORMULATION_OFF"
    case MSK_SIM_REFORMULATION_ON: return "MSK_SIM_REFORMULATION_ON"
    case MSK_SIM_REFORMULATION_FREE: return "MSK_SIM_REFORMULATION_FREE"
    case MSK_SIM_REFORMULATION_AGGRESSIVE: return "MSK_SIM_REFORMULATION_AGGRESSIVE"
    default: return "<unknown>"
  }
} // Simreform.String()

type Simdupvec int32
// Exploit duplicate columns.
const (
    MSK_SIM_EXPLOIT_DUPVEC_OFF Simdupvec = 0 // Disallow the simplex optimizer to exploit duplicated columns.
    MSK_SIM_EXPLOIT_DUPVEC_ON Simdupvec = 1 // Allow the simplex optimizer to exploit duplicated columns.
    MSK_SIM_EXPLOIT_DUPVEC_FREE Simdupvec = 2 // The simplex optimizer can choose freely.
    MSK_SIM_EXPLOIT_DUPVEC_END Simdupvec = 2
)
func (self Simdupvec) String() string {
  switch self {
    case MSK_SIM_EXPLOIT_DUPVEC_OFF: return "MSK_SIM_EXPLOIT_DUPVEC_OFF"
    case MSK_SIM_EXPLOIT_DUPVEC_ON: return "MSK_SIM_EXPLOIT_DUPVEC_ON"
    case MSK_SIM_EXPLOIT_DUPVEC_FREE: return "MSK_SIM_EXPLOIT_DUPVEC_FREE"
    default: return "<unknown>"
  }
} // Simdupvec.String()

type Simhotstart int32
// Hot-start type employed by the simplex optimizer
const (
    MSK_SIM_HOTSTART_NONE Simhotstart = 0 // The simplex optimizer performs a coldstart.
    MSK_SIM_HOTSTART_FREE Simhotstart = 1 // The simplex optimize chooses the hot-start type.
    MSK_SIM_HOTSTART_STATUS_KEYS Simhotstart = 2 // Only the status keys of the constraints and variables are used to choose the type of hot-start.
    MSK_SIM_HOTSTART_END Simhotstart = 2
)
func (self Simhotstart) String() string {
  switch self {
    case MSK_SIM_HOTSTART_NONE: return "MSK_SIM_HOTSTART_NONE"
    case MSK_SIM_HOTSTART_FREE: return "MSK_SIM_HOTSTART_FREE"
    case MSK_SIM_HOTSTART_STATUS_KEYS: return "MSK_SIM_HOTSTART_STATUS_KEYS"
    default: return "<unknown>"
  }
} // Simhotstart.String()

type Intpnthotstart int32
// Hot-start type employed by the interior-point optimizers.
const (
    MSK_INTPNT_HOTSTART_NONE Intpnthotstart = 0 // The interior-point optimizer performs a coldstart.
    MSK_INTPNT_HOTSTART_PRIMAL Intpnthotstart = 1 // The interior-point optimizer exploits the primal solution only.
    MSK_INTPNT_HOTSTART_DUAL Intpnthotstart = 2 // The interior-point optimizer exploits the dual solution only.
    MSK_INTPNT_HOTSTART_PRIMAL_DUAL Intpnthotstart = 3 // The interior-point optimizer exploits both the primal and dual solution.
    MSK_INTPNT_HOTSTART_END Intpnthotstart = 3
)
func (self Intpnthotstart) String() string {
  switch self {
    case MSK_INTPNT_HOTSTART_NONE: return "MSK_INTPNT_HOTSTART_NONE"
    case MSK_INTPNT_HOTSTART_PRIMAL: return "MSK_INTPNT_HOTSTART_PRIMAL"
    case MSK_INTPNT_HOTSTART_DUAL: return "MSK_INTPNT_HOTSTART_DUAL"
    case MSK_INTPNT_HOTSTART_PRIMAL_DUAL: return "MSK_INTPNT_HOTSTART_PRIMAL_DUAL"
    default: return "<unknown>"
  }
} // Intpnthotstart.String()

type Purify int32
// Solution purification employed optimizer.
const (
    MSK_PURIFY_NONE Purify = 0 // The optimizer performs no solution purification.
    MSK_PURIFY_PRIMAL Purify = 1 // The optimizer purifies the primal solution.
    MSK_PURIFY_DUAL Purify = 2 // The optimizer purifies the dual solution.
    MSK_PURIFY_PRIMAL_DUAL Purify = 3 // The optimizer purifies both the primal and dual solution.
    MSK_PURIFY_AUTO Purify = 4 // TBD
    MSK_PURIFY_END Purify = 4
)
func (self Purify) String() string {
  switch self {
    case MSK_PURIFY_NONE: return "MSK_PURIFY_NONE"
    case MSK_PURIFY_PRIMAL: return "MSK_PURIFY_PRIMAL"
    case MSK_PURIFY_DUAL: return "MSK_PURIFY_DUAL"
    case MSK_PURIFY_PRIMAL_DUAL: return "MSK_PURIFY_PRIMAL_DUAL"
    case MSK_PURIFY_AUTO: return "MSK_PURIFY_AUTO"
    default: return "<unknown>"
  }
} // Purify.String()

type Callbackcode int32
// Progress callback codes
const (
    MSK_CALLBACK_BEGIN_BI Callbackcode = 0 // The basis identification procedure has been started.
    MSK_CALLBACK_BEGIN_CONIC Callbackcode = 1 // The callback function is called when the conic optimizer is started.
    MSK_CALLBACK_BEGIN_DUAL_BI Callbackcode = 2 // The callback function is called from within the basis identification procedure when the dual phase is started.
    MSK_CALLBACK_BEGIN_DUAL_SENSITIVITY Callbackcode = 3 // Dual sensitivity analysis is started.
    MSK_CALLBACK_BEGIN_DUAL_SETUP_BI Callbackcode = 4 // The callback function is called when the dual BI phase is started.
    MSK_CALLBACK_BEGIN_DUAL_SIMPLEX Callbackcode = 5 // The callback function is called when the dual simplex optimizer started.
    MSK_CALLBACK_BEGIN_DUAL_SIMPLEX_BI Callbackcode = 6 // The callback function is called from within the basis identification procedure when the dual simplex clean-up phase is started.
    MSK_CALLBACK_BEGIN_INFEAS_ANA Callbackcode = 7 // The callback function is called when the infeasibility analyzer is started.
    MSK_CALLBACK_BEGIN_INTPNT Callbackcode = 8 // The callback function is called when the interior-point optimizer is started.
    MSK_CALLBACK_BEGIN_LICENSE_WAIT Callbackcode = 9 // Begin waiting for license.
    MSK_CALLBACK_BEGIN_MIO Callbackcode = 10 // The callback function is called when the mixed-integer optimizer is started.
    MSK_CALLBACK_BEGIN_OPTIMIZER Callbackcode = 11 // The callback function is called when the optimizer is started.
    MSK_CALLBACK_BEGIN_PRESOLVE Callbackcode = 12 // The callback function is called when the presolve is started.
    MSK_CALLBACK_BEGIN_PRIMAL_BI Callbackcode = 13 // The callback function is called from within the basis identification procedure when the primal phase is started.
    MSK_CALLBACK_BEGIN_PRIMAL_REPAIR Callbackcode = 14 // Begin primal feasibility repair.
    MSK_CALLBACK_BEGIN_PRIMAL_SENSITIVITY Callbackcode = 15 // Primal sensitivity analysis is started.
    MSK_CALLBACK_BEGIN_PRIMAL_SETUP_BI Callbackcode = 16 // The callback function is called when the primal BI setup is started.
    MSK_CALLBACK_BEGIN_PRIMAL_SIMPLEX Callbackcode = 17 // The callback function is called when the primal simplex optimizer is started.
    MSK_CALLBACK_BEGIN_PRIMAL_SIMPLEX_BI Callbackcode = 18 // The callback function is called from within the basis identification procedure when the primal simplex clean-up phase is started.
    MSK_CALLBACK_BEGIN_QCQO_REFORMULATE Callbackcode = 19 // Begin QCQO reformulation.
    MSK_CALLBACK_BEGIN_READ Callbackcode = 20 // MOSEK has started reading a problem file.
    MSK_CALLBACK_BEGIN_ROOT_CUTGEN Callbackcode = 21 // The callback function is called when root cut generation is started.
    MSK_CALLBACK_BEGIN_SIMPLEX Callbackcode = 22 // The callback function is called when the simplex optimizer is started.
    MSK_CALLBACK_BEGIN_SIMPLEX_BI Callbackcode = 23 // The callback function is called from within the basis identification procedure when the simplex clean-up phase is started.
    MSK_CALLBACK_BEGIN_SOLVE_ROOT_RELAX Callbackcode = 24 // The callback function is called when solution of root relaxation is started.
    MSK_CALLBACK_BEGIN_TO_CONIC Callbackcode = 25 // Begin conic reformulation.
    MSK_CALLBACK_BEGIN_WRITE Callbackcode = 26 // MOSEK has started writing a problem file.
    MSK_CALLBACK_CONIC Callbackcode = 27 // The callback function is called from within the conic optimizer after the information database has been updated.
    MSK_CALLBACK_DUAL_SIMPLEX Callbackcode = 28 // The callback function is called from within the dual simplex optimizer.
    MSK_CALLBACK_END_BI Callbackcode = 29 // The callback function is called when the basis identification procedure is terminated.
    MSK_CALLBACK_END_CONIC Callbackcode = 30 // The callback function is called when the conic optimizer is terminated.
    MSK_CALLBACK_END_DUAL_BI Callbackcode = 31 // The callback function is called from within the basis identification procedure when the dual phase is terminated.
    MSK_CALLBACK_END_DUAL_SENSITIVITY Callbackcode = 32 // Dual sensitivity analysis is terminated.
    MSK_CALLBACK_END_DUAL_SETUP_BI Callbackcode = 33 // The callback function is called when the dual BI phase is terminated.
    MSK_CALLBACK_END_DUAL_SIMPLEX Callbackcode = 34 // The callback function is called when the dual simplex optimizer is terminated.
    MSK_CALLBACK_END_DUAL_SIMPLEX_BI Callbackcode = 35 // The callback function is called from within the basis identification procedure when the dual clean-up phase is terminated.
    MSK_CALLBACK_END_INFEAS_ANA Callbackcode = 36 // The callback function is called when the infeasibility analyzer is terminated.
    MSK_CALLBACK_END_INTPNT Callbackcode = 37 // The callback function is called when the interior-point optimizer is terminated.
    MSK_CALLBACK_END_LICENSE_WAIT Callbackcode = 38 // End waiting for license.
    MSK_CALLBACK_END_MIO Callbackcode = 39 // The callback function is called when the mixed-integer optimizer is terminated.
    MSK_CALLBACK_END_OPTIMIZER Callbackcode = 40 // The callback function is called when the optimizer is terminated.
    MSK_CALLBACK_END_PRESOLVE Callbackcode = 41 // The callback function is called when the presolve is completed.
    MSK_CALLBACK_END_PRIMAL_BI Callbackcode = 42 // The callback function is called from within the basis identification procedure when the primal phase is terminated.
    MSK_CALLBACK_END_PRIMAL_REPAIR Callbackcode = 43 // End primal feasibility repair.
    MSK_CALLBACK_END_PRIMAL_SENSITIVITY Callbackcode = 44 // Primal sensitivity analysis is terminated.
    MSK_CALLBACK_END_PRIMAL_SETUP_BI Callbackcode = 45 // The callback function is called when the primal BI setup is terminated.
    MSK_CALLBACK_END_PRIMAL_SIMPLEX Callbackcode = 46 // The callback function is called when the primal simplex optimizer is terminated.
    MSK_CALLBACK_END_PRIMAL_SIMPLEX_BI Callbackcode = 47 // The callback function is called from within the basis identification procedure when the primal clean-up phase is terminated.
    MSK_CALLBACK_END_QCQO_REFORMULATE Callbackcode = 48 // End QCQO reformulation.
    MSK_CALLBACK_END_READ Callbackcode = 49 // MOSEK has finished reading a problem file.
    MSK_CALLBACK_END_ROOT_CUTGEN Callbackcode = 50 // The callback function is called when root cut generation is terminated.
    MSK_CALLBACK_END_SIMPLEX Callbackcode = 51 // The callback function is called when the simplex optimizer is terminated.
    MSK_CALLBACK_END_SIMPLEX_BI Callbackcode = 52 // The callback function is called from within the basis identification procedure when the simplex clean-up phase is terminated.
    MSK_CALLBACK_END_SOLVE_ROOT_RELAX Callbackcode = 53 // The callback function is called when solution of root relaxation is terminated.
    MSK_CALLBACK_END_TO_CONIC Callbackcode = 54 // End conic reformulation.
    MSK_CALLBACK_END_WRITE Callbackcode = 55 // MOSEK has finished writing a problem file.
    MSK_CALLBACK_IM_BI Callbackcode = 56 // The callback function is called from within the basis identification procedure at an intermediate point.
    MSK_CALLBACK_IM_CONIC Callbackcode = 57 // The callback function is called at an intermediate stage within the conic optimizer where the information database has not been updated.
    MSK_CALLBACK_IM_DUAL_BI Callbackcode = 58 // The callback function is called from within the basis identification procedure at an intermediate point in the dual phase.
    MSK_CALLBACK_IM_DUAL_SENSIVITY Callbackcode = 59 // The callback function is called at an intermediate stage of the dual sensitivity analysis.
    MSK_CALLBACK_IM_DUAL_SIMPLEX Callbackcode = 60 // The callback function is called at an intermediate point in the dual simplex optimizer.
    MSK_CALLBACK_IM_INTPNT Callbackcode = 61 // The callback function is called at an intermediate stage within the interior-point optimizer where the information database has not been updated.
    MSK_CALLBACK_IM_LICENSE_WAIT Callbackcode = 62 // MOSEK is waiting for a license.
    MSK_CALLBACK_IM_LU Callbackcode = 63 // The callback function is called from within the LU factorization procedure at an intermediate point.
    MSK_CALLBACK_IM_MIO Callbackcode = 64 // The callback function is called at an intermediate point in the mixed-integer optimizer.
    MSK_CALLBACK_IM_MIO_DUAL_SIMPLEX Callbackcode = 65 // The callback function is called at an intermediate point in the mixed-integer optimizer while running the dual simplex optimizer.
    MSK_CALLBACK_IM_MIO_INTPNT Callbackcode = 66 // The callback function is called at an intermediate point in the mixed-integer optimizer while running the interior-point optimizer.
    MSK_CALLBACK_IM_MIO_PRIMAL_SIMPLEX Callbackcode = 67 // The callback function is called at an intermediate point in the mixed-integer optimizer while running the primal simplex optimizer.
    MSK_CALLBACK_IM_ORDER Callbackcode = 68 // The callback function is called from within the matrix ordering procedure at an intermediate point.
    MSK_CALLBACK_IM_PRESOLVE Callbackcode = 69 // The callback function is called from within the presolve procedure at an intermediate stage.
    MSK_CALLBACK_IM_PRIMAL_BI Callbackcode = 70 // The callback function is called from within the basis identification procedure at an intermediate point in the primal phase.
    MSK_CALLBACK_IM_PRIMAL_SENSIVITY Callbackcode = 71 // The callback function is called at an intermediate stage of the primal sensitivity analysis.
    MSK_CALLBACK_IM_PRIMAL_SIMPLEX Callbackcode = 72 // The callback function is called at an intermediate point in the primal simplex optimizer.
    MSK_CALLBACK_IM_QO_REFORMULATE Callbackcode = 73 // The callback function is called at an intermediate stage of the conic quadratic reformulation.
    MSK_CALLBACK_IM_READ Callbackcode = 74 // Intermediate stage in reading.
    MSK_CALLBACK_IM_ROOT_CUTGEN Callbackcode = 75 // The callback is called from within root cut generation at an intermediate stage.
    MSK_CALLBACK_IM_SIMPLEX Callbackcode = 76 // The callback function is called from within the simplex optimizer at an intermediate point.
    MSK_CALLBACK_IM_SIMPLEX_BI Callbackcode = 77 // The callback function is called from within the basis identification procedure at an intermediate point in the simplex clean-up phase.
    MSK_CALLBACK_INTPNT Callbackcode = 78 // The callback function is called from within the interior-point optimizer after the information database has been updated.
    MSK_CALLBACK_NEW_INT_MIO Callbackcode = 79 // The callback function is called after a new integer solution has been located by the mixed-integer optimizer.
    MSK_CALLBACK_PRIMAL_SIMPLEX Callbackcode = 80 // The callback function is called from within the primal simplex optimizer.
    MSK_CALLBACK_READ_OPF Callbackcode = 81 // The callback function is called from the OPF reader.
    MSK_CALLBACK_READ_OPF_SECTION Callbackcode = 82 // A chunk of Q non-zeros has been read from a problem file.
    MSK_CALLBACK_RESTART_MIO Callbackcode = 83 // The callback function is called when the mixed-integer optimizer is restarted.
    MSK_CALLBACK_SOLVING_REMOTE Callbackcode = 84 // The callback function is called while the task is being solved on a remote server.
    MSK_CALLBACK_UPDATE_DUAL_BI Callbackcode = 85 // The callback function is called from within the basis identification procedure at an intermediate point in the dual phase.
    MSK_CALLBACK_UPDATE_DUAL_SIMPLEX Callbackcode = 86 // The callback function is called in the dual simplex optimizer.
    MSK_CALLBACK_UPDATE_DUAL_SIMPLEX_BI Callbackcode = 87 // The callback function is called from within the basis identification procedure at an intermediate point in the dual simplex clean-up phase.
    MSK_CALLBACK_UPDATE_PRESOLVE Callbackcode = 88 // The callback function is called from within the presolve procedure.
    MSK_CALLBACK_UPDATE_PRIMAL_BI Callbackcode = 89 // The callback function is called from within the basis identification procedure at an intermediate point in the primal phase.
    MSK_CALLBACK_UPDATE_PRIMAL_SIMPLEX Callbackcode = 90 // The callback function is called  in the primal simplex optimizer.
    MSK_CALLBACK_UPDATE_PRIMAL_SIMPLEX_BI Callbackcode = 91 // The callback function is called from within the basis identification procedure at an intermediate point in the primal simplex clean-up phase.
    MSK_CALLBACK_UPDATE_SIMPLEX Callbackcode = 92 // The callback function is called from simplex optimizer.
    MSK_CALLBACK_WRITE_OPF Callbackcode = 93 // The callback function is called from the OPF writer.
    MSK_CALLBACK_END Callbackcode = 93
)
func (self Callbackcode) String() string {
  switch self {
    case MSK_CALLBACK_BEGIN_BI: return "MSK_CALLBACK_BEGIN_BI"
    case MSK_CALLBACK_BEGIN_CONIC: return "MSK_CALLBACK_BEGIN_CONIC"
    case MSK_CALLBACK_BEGIN_DUAL_BI: return "MSK_CALLBACK_BEGIN_DUAL_BI"
    case MSK_CALLBACK_BEGIN_DUAL_SENSITIVITY: return "MSK_CALLBACK_BEGIN_DUAL_SENSITIVITY"
    case MSK_CALLBACK_BEGIN_DUAL_SETUP_BI: return "MSK_CALLBACK_BEGIN_DUAL_SETUP_BI"
    case MSK_CALLBACK_BEGIN_DUAL_SIMPLEX: return "MSK_CALLBACK_BEGIN_DUAL_SIMPLEX"
    case MSK_CALLBACK_BEGIN_DUAL_SIMPLEX_BI: return "MSK_CALLBACK_BEGIN_DUAL_SIMPLEX_BI"
    case MSK_CALLBACK_BEGIN_INFEAS_ANA: return "MSK_CALLBACK_BEGIN_INFEAS_ANA"
    case MSK_CALLBACK_BEGIN_INTPNT: return "MSK_CALLBACK_BEGIN_INTPNT"
    case MSK_CALLBACK_BEGIN_LICENSE_WAIT: return "MSK_CALLBACK_BEGIN_LICENSE_WAIT"
    case MSK_CALLBACK_BEGIN_MIO: return "MSK_CALLBACK_BEGIN_MIO"
    case MSK_CALLBACK_BEGIN_OPTIMIZER: return "MSK_CALLBACK_BEGIN_OPTIMIZER"
    case MSK_CALLBACK_BEGIN_PRESOLVE: return "MSK_CALLBACK_BEGIN_PRESOLVE"
    case MSK_CALLBACK_BEGIN_PRIMAL_BI: return "MSK_CALLBACK_BEGIN_PRIMAL_BI"
    case MSK_CALLBACK_BEGIN_PRIMAL_REPAIR: return "MSK_CALLBACK_BEGIN_PRIMAL_REPAIR"
    case MSK_CALLBACK_BEGIN_PRIMAL_SENSITIVITY: return "MSK_CALLBACK_BEGIN_PRIMAL_SENSITIVITY"
    case MSK_CALLBACK_BEGIN_PRIMAL_SETUP_BI: return "MSK_CALLBACK_BEGIN_PRIMAL_SETUP_BI"
    case MSK_CALLBACK_BEGIN_PRIMAL_SIMPLEX: return "MSK_CALLBACK_BEGIN_PRIMAL_SIMPLEX"
    case MSK_CALLBACK_BEGIN_PRIMAL_SIMPLEX_BI: return "MSK_CALLBACK_BEGIN_PRIMAL_SIMPLEX_BI"
    case MSK_CALLBACK_BEGIN_QCQO_REFORMULATE: return "MSK_CALLBACK_BEGIN_QCQO_REFORMULATE"
    case MSK_CALLBACK_BEGIN_READ: return "MSK_CALLBACK_BEGIN_READ"
    case MSK_CALLBACK_BEGIN_ROOT_CUTGEN: return "MSK_CALLBACK_BEGIN_ROOT_CUTGEN"
    case MSK_CALLBACK_BEGIN_SIMPLEX: return "MSK_CALLBACK_BEGIN_SIMPLEX"
    case MSK_CALLBACK_BEGIN_SIMPLEX_BI: return "MSK_CALLBACK_BEGIN_SIMPLEX_BI"
    case MSK_CALLBACK_BEGIN_SOLVE_ROOT_RELAX: return "MSK_CALLBACK_BEGIN_SOLVE_ROOT_RELAX"
    case MSK_CALLBACK_BEGIN_TO_CONIC: return "MSK_CALLBACK_BEGIN_TO_CONIC"
    case MSK_CALLBACK_BEGIN_WRITE: return "MSK_CALLBACK_BEGIN_WRITE"
    case MSK_CALLBACK_CONIC: return "MSK_CALLBACK_CONIC"
    case MSK_CALLBACK_DUAL_SIMPLEX: return "MSK_CALLBACK_DUAL_SIMPLEX"
    case MSK_CALLBACK_END_BI: return "MSK_CALLBACK_END_BI"
    case MSK_CALLBACK_END_CONIC: return "MSK_CALLBACK_END_CONIC"
    case MSK_CALLBACK_END_DUAL_BI: return "MSK_CALLBACK_END_DUAL_BI"
    case MSK_CALLBACK_END_DUAL_SENSITIVITY: return "MSK_CALLBACK_END_DUAL_SENSITIVITY"
    case MSK_CALLBACK_END_DUAL_SETUP_BI: return "MSK_CALLBACK_END_DUAL_SETUP_BI"
    case MSK_CALLBACK_END_DUAL_SIMPLEX: return "MSK_CALLBACK_END_DUAL_SIMPLEX"
    case MSK_CALLBACK_END_DUAL_SIMPLEX_BI: return "MSK_CALLBACK_END_DUAL_SIMPLEX_BI"
    case MSK_CALLBACK_END_INFEAS_ANA: return "MSK_CALLBACK_END_INFEAS_ANA"
    case MSK_CALLBACK_END_INTPNT: return "MSK_CALLBACK_END_INTPNT"
    case MSK_CALLBACK_END_LICENSE_WAIT: return "MSK_CALLBACK_END_LICENSE_WAIT"
    case MSK_CALLBACK_END_MIO: return "MSK_CALLBACK_END_MIO"
    case MSK_CALLBACK_END_OPTIMIZER: return "MSK_CALLBACK_END_OPTIMIZER"
    case MSK_CALLBACK_END_PRESOLVE: return "MSK_CALLBACK_END_PRESOLVE"
    case MSK_CALLBACK_END_PRIMAL_BI: return "MSK_CALLBACK_END_PRIMAL_BI"
    case MSK_CALLBACK_END_PRIMAL_REPAIR: return "MSK_CALLBACK_END_PRIMAL_REPAIR"
    case MSK_CALLBACK_END_PRIMAL_SENSITIVITY: return "MSK_CALLBACK_END_PRIMAL_SENSITIVITY"
    case MSK_CALLBACK_END_PRIMAL_SETUP_BI: return "MSK_CALLBACK_END_PRIMAL_SETUP_BI"
    case MSK_CALLBACK_END_PRIMAL_SIMPLEX: return "MSK_CALLBACK_END_PRIMAL_SIMPLEX"
    case MSK_CALLBACK_END_PRIMAL_SIMPLEX_BI: return "MSK_CALLBACK_END_PRIMAL_SIMPLEX_BI"
    case MSK_CALLBACK_END_QCQO_REFORMULATE: return "MSK_CALLBACK_END_QCQO_REFORMULATE"
    case MSK_CALLBACK_END_READ: return "MSK_CALLBACK_END_READ"
    case MSK_CALLBACK_END_ROOT_CUTGEN: return "MSK_CALLBACK_END_ROOT_CUTGEN"
    case MSK_CALLBACK_END_SIMPLEX: return "MSK_CALLBACK_END_SIMPLEX"
    case MSK_CALLBACK_END_SIMPLEX_BI: return "MSK_CALLBACK_END_SIMPLEX_BI"
    case MSK_CALLBACK_END_SOLVE_ROOT_RELAX: return "MSK_CALLBACK_END_SOLVE_ROOT_RELAX"
    case MSK_CALLBACK_END_TO_CONIC: return "MSK_CALLBACK_END_TO_CONIC"
    case MSK_CALLBACK_END_WRITE: return "MSK_CALLBACK_END_WRITE"
    case MSK_CALLBACK_IM_BI: return "MSK_CALLBACK_IM_BI"
    case MSK_CALLBACK_IM_CONIC: return "MSK_CALLBACK_IM_CONIC"
    case MSK_CALLBACK_IM_DUAL_BI: return "MSK_CALLBACK_IM_DUAL_BI"
    case MSK_CALLBACK_IM_DUAL_SENSIVITY: return "MSK_CALLBACK_IM_DUAL_SENSIVITY"
    case MSK_CALLBACK_IM_DUAL_SIMPLEX: return "MSK_CALLBACK_IM_DUAL_SIMPLEX"
    case MSK_CALLBACK_IM_INTPNT: return "MSK_CALLBACK_IM_INTPNT"
    case MSK_CALLBACK_IM_LICENSE_WAIT: return "MSK_CALLBACK_IM_LICENSE_WAIT"
    case MSK_CALLBACK_IM_LU: return "MSK_CALLBACK_IM_LU"
    case MSK_CALLBACK_IM_MIO: return "MSK_CALLBACK_IM_MIO"
    case MSK_CALLBACK_IM_MIO_DUAL_SIMPLEX: return "MSK_CALLBACK_IM_MIO_DUAL_SIMPLEX"
    case MSK_CALLBACK_IM_MIO_INTPNT: return "MSK_CALLBACK_IM_MIO_INTPNT"
    case MSK_CALLBACK_IM_MIO_PRIMAL_SIMPLEX: return "MSK_CALLBACK_IM_MIO_PRIMAL_SIMPLEX"
    case MSK_CALLBACK_IM_ORDER: return "MSK_CALLBACK_IM_ORDER"
    case MSK_CALLBACK_IM_PRESOLVE: return "MSK_CALLBACK_IM_PRESOLVE"
    case MSK_CALLBACK_IM_PRIMAL_BI: return "MSK_CALLBACK_IM_PRIMAL_BI"
    case MSK_CALLBACK_IM_PRIMAL_SENSIVITY: return "MSK_CALLBACK_IM_PRIMAL_SENSIVITY"
    case MSK_CALLBACK_IM_PRIMAL_SIMPLEX: return "MSK_CALLBACK_IM_PRIMAL_SIMPLEX"
    case MSK_CALLBACK_IM_QO_REFORMULATE: return "MSK_CALLBACK_IM_QO_REFORMULATE"
    case MSK_CALLBACK_IM_READ: return "MSK_CALLBACK_IM_READ"
    case MSK_CALLBACK_IM_ROOT_CUTGEN: return "MSK_CALLBACK_IM_ROOT_CUTGEN"
    case MSK_CALLBACK_IM_SIMPLEX: return "MSK_CALLBACK_IM_SIMPLEX"
    case MSK_CALLBACK_IM_SIMPLEX_BI: return "MSK_CALLBACK_IM_SIMPLEX_BI"
    case MSK_CALLBACK_INTPNT: return "MSK_CALLBACK_INTPNT"
    case MSK_CALLBACK_NEW_INT_MIO: return "MSK_CALLBACK_NEW_INT_MIO"
    case MSK_CALLBACK_PRIMAL_SIMPLEX: return "MSK_CALLBACK_PRIMAL_SIMPLEX"
    case MSK_CALLBACK_READ_OPF: return "MSK_CALLBACK_READ_OPF"
    case MSK_CALLBACK_READ_OPF_SECTION: return "MSK_CALLBACK_READ_OPF_SECTION"
    case MSK_CALLBACK_RESTART_MIO: return "MSK_CALLBACK_RESTART_MIO"
    case MSK_CALLBACK_SOLVING_REMOTE: return "MSK_CALLBACK_SOLVING_REMOTE"
    case MSK_CALLBACK_UPDATE_DUAL_BI: return "MSK_CALLBACK_UPDATE_DUAL_BI"
    case MSK_CALLBACK_UPDATE_DUAL_SIMPLEX: return "MSK_CALLBACK_UPDATE_DUAL_SIMPLEX"
    case MSK_CALLBACK_UPDATE_DUAL_SIMPLEX_BI: return "MSK_CALLBACK_UPDATE_DUAL_SIMPLEX_BI"
    case MSK_CALLBACK_UPDATE_PRESOLVE: return "MSK_CALLBACK_UPDATE_PRESOLVE"
    case MSK_CALLBACK_UPDATE_PRIMAL_BI: return "MSK_CALLBACK_UPDATE_PRIMAL_BI"
    case MSK_CALLBACK_UPDATE_PRIMAL_SIMPLEX: return "MSK_CALLBACK_UPDATE_PRIMAL_SIMPLEX"
    case MSK_CALLBACK_UPDATE_PRIMAL_SIMPLEX_BI: return "MSK_CALLBACK_UPDATE_PRIMAL_SIMPLEX_BI"
    case MSK_CALLBACK_UPDATE_SIMPLEX: return "MSK_CALLBACK_UPDATE_SIMPLEX"
    case MSK_CALLBACK_WRITE_OPF: return "MSK_CALLBACK_WRITE_OPF"
    default: return "<unknown>"
  }
} // Callbackcode.String()

type Compresstype int32
// Compression types
const (
    MSK_COMPRESS_NONE Compresstype = 0 // No compression is used.
    MSK_COMPRESS_FREE Compresstype = 1 // The type of compression used is chosen automatically.
    MSK_COMPRESS_GZIP Compresstype = 2 // The type of compression used is gzip compatible.
    MSK_COMPRESS_ZSTD Compresstype = 3 // The type of compression used is zstd compatible.
    MSK_COMPRESS_END Compresstype = 3
)
func (self Compresstype) String() string {
  switch self {
    case MSK_COMPRESS_NONE: return "MSK_COMPRESS_NONE"
    case MSK_COMPRESS_FREE: return "MSK_COMPRESS_FREE"
    case MSK_COMPRESS_GZIP: return "MSK_COMPRESS_GZIP"
    case MSK_COMPRESS_ZSTD: return "MSK_COMPRESS_ZSTD"
    default: return "<unknown>"
  }
} // Compresstype.String()

type Conetype int32
// Cone types
const (
    MSK_CT_QUAD Conetype = 0 // The cone is a quadratic cone.
    MSK_CT_RQUAD Conetype = 1 // The cone is a rotated quadratic cone.
    MSK_CT_PEXP Conetype = 2 // A primal exponential cone.
    MSK_CT_DEXP Conetype = 3 // A dual exponential cone.
    MSK_CT_PPOW Conetype = 4 // A primal power cone.
    MSK_CT_DPOW Conetype = 5 // A dual power cone.
    MSK_CT_ZERO Conetype = 6 // The zero cone.
    MSK_CT_END Conetype = 6
)
func (self Conetype) String() string {
  switch self {
    case MSK_CT_QUAD: return "MSK_CT_QUAD"
    case MSK_CT_RQUAD: return "MSK_CT_RQUAD"
    case MSK_CT_PEXP: return "MSK_CT_PEXP"
    case MSK_CT_DEXP: return "MSK_CT_DEXP"
    case MSK_CT_PPOW: return "MSK_CT_PPOW"
    case MSK_CT_DPOW: return "MSK_CT_DPOW"
    case MSK_CT_ZERO: return "MSK_CT_ZERO"
    default: return "<unknown>"
  }
} // Conetype.String()

type Domaintype int32
// Cone types
const (
    MSK_DOMAIN_R Domaintype = 0 // R.
    MSK_DOMAIN_RZERO Domaintype = 1 // The zero vector.
    MSK_DOMAIN_RPLUS Domaintype = 2 // The positive orthant.
    MSK_DOMAIN_RMINUS Domaintype = 3 // The negative orthant.
    MSK_DOMAIN_QUADRATIC_CONE Domaintype = 4 // The quadratic cone.
    MSK_DOMAIN_RQUADRATIC_CONE Domaintype = 5 // The rotated quadratic cone.
    MSK_DOMAIN_PRIMAL_EXP_CONE Domaintype = 6 // The primal exponential cone.
    MSK_DOMAIN_DUAL_EXP_CONE Domaintype = 7 // The dual exponential cone.
    MSK_DOMAIN_PRIMAL_POWER_CONE Domaintype = 8 // The primal power cone.
    MSK_DOMAIN_DUAL_POWER_CONE Domaintype = 9 // The dual power cone.
    MSK_DOMAIN_PRIMAL_GEO_MEAN_CONE Domaintype = 10 // The primal geometric mean cone.
    MSK_DOMAIN_DUAL_GEO_MEAN_CONE Domaintype = 11 // The dual geometric mean cone.
    MSK_DOMAIN_SVEC_PSD_CONE Domaintype = 12 // The vectorized positive semidefinite cone.
    MSK_DOMAIN_END Domaintype = 12
)
func (self Domaintype) String() string {
  switch self {
    case MSK_DOMAIN_R: return "MSK_DOMAIN_R"
    case MSK_DOMAIN_RZERO: return "MSK_DOMAIN_RZERO"
    case MSK_DOMAIN_RPLUS: return "MSK_DOMAIN_RPLUS"
    case MSK_DOMAIN_RMINUS: return "MSK_DOMAIN_RMINUS"
    case MSK_DOMAIN_QUADRATIC_CONE: return "MSK_DOMAIN_QUADRATIC_CONE"
    case MSK_DOMAIN_RQUADRATIC_CONE: return "MSK_DOMAIN_RQUADRATIC_CONE"
    case MSK_DOMAIN_PRIMAL_EXP_CONE: return "MSK_DOMAIN_PRIMAL_EXP_CONE"
    case MSK_DOMAIN_DUAL_EXP_CONE: return "MSK_DOMAIN_DUAL_EXP_CONE"
    case MSK_DOMAIN_PRIMAL_POWER_CONE: return "MSK_DOMAIN_PRIMAL_POWER_CONE"
    case MSK_DOMAIN_DUAL_POWER_CONE: return "MSK_DOMAIN_DUAL_POWER_CONE"
    case MSK_DOMAIN_PRIMAL_GEO_MEAN_CONE: return "MSK_DOMAIN_PRIMAL_GEO_MEAN_CONE"
    case MSK_DOMAIN_DUAL_GEO_MEAN_CONE: return "MSK_DOMAIN_DUAL_GEO_MEAN_CONE"
    case MSK_DOMAIN_SVEC_PSD_CONE: return "MSK_DOMAIN_SVEC_PSD_CONE"
    default: return "<unknown>"
  }
} // Domaintype.String()

type Nametype int32
// Name types
const (
    MSK_NAME_TYPE_GEN Nametype = 0 // General names. However, no duplicate and blank names are allowed.
    MSK_NAME_TYPE_MPS Nametype = 1 // MPS type names.
    MSK_NAME_TYPE_LP Nametype = 2 // LP type names.
    MSK_NAME_TYPE_END Nametype = 2
)
func (self Nametype) String() string {
  switch self {
    case MSK_NAME_TYPE_GEN: return "MSK_NAME_TYPE_GEN"
    case MSK_NAME_TYPE_MPS: return "MSK_NAME_TYPE_MPS"
    case MSK_NAME_TYPE_LP: return "MSK_NAME_TYPE_LP"
    default: return "<unknown>"
  }
} // Nametype.String()

type Symmattype int32
// Cone types
const (
    MSK_SYMMAT_TYPE_SPARSE Symmattype = 0 // Sparse symmetric matrix.
    MSK_SYMMAT_TYPE_END Symmattype = 0
)
func (self Symmattype) String() string {
  switch self {
    case MSK_SYMMAT_TYPE_SPARSE: return "MSK_SYMMAT_TYPE_SPARSE"
    default: return "<unknown>"
  }
} // Symmattype.String()

type Dataformat int32
// Data format types
const (
    MSK_DATA_FORMAT_EXTENSION Dataformat = 0 // The file extension is used to determine the data file format.
    MSK_DATA_FORMAT_MPS Dataformat = 1 // The data file is MPS formatted.
    MSK_DATA_FORMAT_LP Dataformat = 2 // The data file is LP formatted.
    MSK_DATA_FORMAT_OP Dataformat = 3 // The data file is an optimization problem formatted file.
    MSK_DATA_FORMAT_FREE_MPS Dataformat = 4 // The data a free MPS formatted file.
    MSK_DATA_FORMAT_TASK Dataformat = 5 // Generic task dump file.
    MSK_DATA_FORMAT_PTF Dataformat = 6 // (P)retty (T)ext (F)format.
    MSK_DATA_FORMAT_CB Dataformat = 7 // Conic benchmark format,
    MSK_DATA_FORMAT_JSON_TASK Dataformat = 8 // JSON based task format.
    MSK_DATA_FORMAT_END Dataformat = 8
)
func (self Dataformat) String() string {
  switch self {
    case MSK_DATA_FORMAT_EXTENSION: return "MSK_DATA_FORMAT_EXTENSION"
    case MSK_DATA_FORMAT_MPS: return "MSK_DATA_FORMAT_MPS"
    case MSK_DATA_FORMAT_LP: return "MSK_DATA_FORMAT_LP"
    case MSK_DATA_FORMAT_OP: return "MSK_DATA_FORMAT_OP"
    case MSK_DATA_FORMAT_FREE_MPS: return "MSK_DATA_FORMAT_FREE_MPS"
    case MSK_DATA_FORMAT_TASK: return "MSK_DATA_FORMAT_TASK"
    case MSK_DATA_FORMAT_PTF: return "MSK_DATA_FORMAT_PTF"
    case MSK_DATA_FORMAT_CB: return "MSK_DATA_FORMAT_CB"
    case MSK_DATA_FORMAT_JSON_TASK: return "MSK_DATA_FORMAT_JSON_TASK"
    default: return "<unknown>"
  }
} // Dataformat.String()

type Solformat int32
// Data format types
const (
    MSK_SOL_FORMAT_EXTENSION Solformat = 0 // The file extension is used to determine the data file format.
    MSK_SOL_FORMAT_B Solformat = 1 // Simple binary format
    MSK_SOL_FORMAT_TASK Solformat = 2 // Tar based format.
    MSK_SOL_FORMAT_JSON_TASK Solformat = 3 // JSON based format.
    MSK_SOL_FORMAT_END Solformat = 3
)
func (self Solformat) String() string {
  switch self {
    case MSK_SOL_FORMAT_EXTENSION: return "MSK_SOL_FORMAT_EXTENSION"
    case MSK_SOL_FORMAT_B: return "MSK_SOL_FORMAT_B"
    case MSK_SOL_FORMAT_TASK: return "MSK_SOL_FORMAT_TASK"
    case MSK_SOL_FORMAT_JSON_TASK: return "MSK_SOL_FORMAT_JSON_TASK"
    default: return "<unknown>"
  }
} // Solformat.String()

type Dinfitem int32
// Double information items
const (
    MSK_DINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_DENSITY Dinfitem = 0 // Density percentage of the scalarized constraint matrix.
    MSK_DINF_BI_CLEAN_DUAL_TIME Dinfitem = 1 // Time  spent within the dual clean-up optimizer of the basis identification procedure since its invocation.
    MSK_DINF_BI_CLEAN_PRIMAL_TIME Dinfitem = 2 // Time spent within the primal clean-up optimizer of the basis identification procedure since its invocation.
    MSK_DINF_BI_CLEAN_TIME Dinfitem = 3 // Time spent within the clean-up phase of the basis identification procedure since its invocation.
    MSK_DINF_BI_DUAL_TIME Dinfitem = 4 // Time spent within the dual phase basis identification procedure since its invocation.
    MSK_DINF_BI_PRIMAL_TIME Dinfitem = 5 // Time  spent within the primal phase of the basis identification procedure since its invocation.
    MSK_DINF_BI_TIME Dinfitem = 6 // Time spent within the basis identification procedure since its invocation.
    MSK_DINF_INTPNT_DUAL_FEAS Dinfitem = 7 // Dual feasibility measure reported by the interior-point optimizer.
    MSK_DINF_INTPNT_DUAL_OBJ Dinfitem = 8 // Dual objective value reported by the interior-point optimizer.
    MSK_DINF_INTPNT_FACTOR_NUM_FLOPS Dinfitem = 9 // An estimate of the number of flops used in the factorization.
    MSK_DINF_INTPNT_OPT_STATUS Dinfitem = 10 // A measure of optimality of the solution.
    MSK_DINF_INTPNT_ORDER_TIME Dinfitem = 11 // Order time (in seconds).
    MSK_DINF_INTPNT_PRIMAL_FEAS Dinfitem = 12 // Primal feasibility measure reported by the interior-point optimizer.
    MSK_DINF_INTPNT_PRIMAL_OBJ Dinfitem = 13 // Primal objective value reported by the interior-point optimizer.
    MSK_DINF_INTPNT_TIME Dinfitem = 14 // Time spent within the interior-point optimizer since its invocation.
    MSK_DINF_MIO_CLIQUE_SELECTION_TIME Dinfitem = 15 // Selection time for clique cuts.
    MSK_DINF_MIO_CLIQUE_SEPARATION_TIME Dinfitem = 16 // Separation time for clique cuts.
    MSK_DINF_MIO_CMIR_SELECTION_TIME Dinfitem = 17 // Selection time for CMIR cuts.
    MSK_DINF_MIO_CMIR_SEPARATION_TIME Dinfitem = 18 // Separation time for CMIR cuts.
    MSK_DINF_MIO_CONSTRUCT_SOLUTION_OBJ Dinfitem = 19 // Optimal objective value corresponding to the feasible solution.
    MSK_DINF_MIO_DUAL_BOUND_AFTER_PRESOLVE Dinfitem = 20 // Value of the dual bound after presolve but before cut generation.
    MSK_DINF_MIO_GMI_SELECTION_TIME Dinfitem = 21 // Selection time for GMI cuts.
    MSK_DINF_MIO_GMI_SEPARATION_TIME Dinfitem = 22 // Separation time for GMI cuts.
    MSK_DINF_MIO_IMPLIED_BOUND_SELECTION_TIME Dinfitem = 23 // Selection time for implied bound cuts.
    MSK_DINF_MIO_IMPLIED_BOUND_SEPARATION_TIME Dinfitem = 24 // Separation time for implied bound cuts.
    MSK_DINF_MIO_INITIAL_FEASIBLE_SOLUTION_OBJ Dinfitem = 25 // Optimal objective value corresponding to the user provided initial solution.
    MSK_DINF_MIO_KNAPSACK_COVER_SELECTION_TIME Dinfitem = 26 // Selection time for knapsack cover.
    MSK_DINF_MIO_KNAPSACK_COVER_SEPARATION_TIME Dinfitem = 27 // Separation time for knapsack cover.
    MSK_DINF_MIO_LIPRO_SELECTION_TIME Dinfitem = 28 // Selection time for lift-and-project cuts.
    MSK_DINF_MIO_LIPRO_SEPARATION_TIME Dinfitem = 29 // Separation time for lift-and-project cuts.
    MSK_DINF_MIO_OBJ_ABS_GAP Dinfitem = 30 // If the mixed-integer optimizer has computed a feasible solution and a bound, this contains the absolute gap.
    MSK_DINF_MIO_OBJ_BOUND Dinfitem = 31 // The best bound on the objective value known.
    MSK_DINF_MIO_OBJ_INT Dinfitem = 32 // The primal objective value corresponding to the best integer feasible solution.
    MSK_DINF_MIO_OBJ_REL_GAP Dinfitem = 33 // If the mixed-integer optimizer has computed a feasible solution and a bound, this contains the relative gap.
    MSK_DINF_MIO_PROBING_TIME Dinfitem = 34 // Total time for probing.
    MSK_DINF_MIO_ROOT_CUT_SELECTION_TIME Dinfitem = 35 // Total time for cut selection.
    MSK_DINF_MIO_ROOT_CUT_SEPARATION_TIME Dinfitem = 36 // Total time for cut separation.
    MSK_DINF_MIO_ROOT_OPTIMIZER_TIME Dinfitem = 37 // Time spent in the contiuous optimizer while processing the root node relaxation.
    MSK_DINF_MIO_ROOT_PRESOLVE_TIME Dinfitem = 38 // Time spent presolving the problem at the root node.
    MSK_DINF_MIO_ROOT_TIME Dinfitem = 39 // Time spent processing the root node.
    MSK_DINF_MIO_SYMMETRY_DETECTION_TIME Dinfitem = 40 // Total time for symmetry detection.
    MSK_DINF_MIO_SYMMETRY_FACTOR Dinfitem = 41 // Degree to which the problem is affected by detected symmetry.
    MSK_DINF_MIO_TIME Dinfitem = 42 // Time spent in the mixed-integer optimizer.
    MSK_DINF_MIO_USER_OBJ_CUT Dinfitem = 43 // If the objective cut is used, then this information item has the value of the cut.
    MSK_DINF_OPTIMIZER_TICKS Dinfitem = 44 // Total number of ticks spent in the optimizer since it was invoked. It is strictly negative if it is not available.
    MSK_DINF_OPTIMIZER_TIME Dinfitem = 45 // Total time spent in the optimizer since it was invoked.
    MSK_DINF_PRESOLVE_ELI_TIME Dinfitem = 46 // Total time spent in the eliminator since the presolve was invoked.
    MSK_DINF_PRESOLVE_LINDEP_TIME Dinfitem = 47 // Total time spent  in the linear dependency checker since the presolve was invoked.
    MSK_DINF_PRESOLVE_TIME Dinfitem = 48 // Total time (in seconds) spent in the presolve since it was invoked.
    MSK_DINF_PRESOLVE_TOTAL_PRIMAL_PERTURBATION Dinfitem = 49 // Total perturbation of the bounds of the primal problem.
    MSK_DINF_PRIMAL_REPAIR_PENALTY_OBJ Dinfitem = 50 // The optimal objective value of the penalty function.
    MSK_DINF_QCQO_REFORMULATE_MAX_PERTURBATION Dinfitem = 51 // Maximum absolute diagonal perturbation occurring during the QCQO reformulation.
    MSK_DINF_QCQO_REFORMULATE_TIME Dinfitem = 52 // Time spent with conic quadratic reformulation.
    MSK_DINF_QCQO_REFORMULATE_WORST_CHOLESKY_COLUMN_SCALING Dinfitem = 53 // Worst Cholesky column scaling.
    MSK_DINF_QCQO_REFORMULATE_WORST_CHOLESKY_DIAG_SCALING Dinfitem = 54 // Worst Cholesky diagonal scaling.
    MSK_DINF_READ_DATA_TIME Dinfitem = 55 // Time spent reading the data file.
    MSK_DINF_REMOTE_TIME Dinfitem = 56 // The total real time in seconds spent when optimizing on a server by the process performing the optimization on the server
    MSK_DINF_SIM_DUAL_TIME Dinfitem = 57 // Time spent in the dual simplex optimizer since invoking it.
    MSK_DINF_SIM_FEAS Dinfitem = 58 // Feasibility measure reported by the simplex optimizer.
    MSK_DINF_SIM_OBJ Dinfitem = 59 // Objective value reported by the simplex optimizer.
    MSK_DINF_SIM_PRIMAL_TIME Dinfitem = 60 // Time spent in the primal simplex optimizer since invoking it.
    MSK_DINF_SIM_TIME Dinfitem = 61 // Time spent in the simplex optimizer since invoking it.
    MSK_DINF_SOL_BAS_DUAL_OBJ Dinfitem = 62 // Dual objective value of the basic solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_BAS_DVIOLCON Dinfitem = 63 // Maximal dual bound violation for xx in the basic solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_BAS_DVIOLVAR Dinfitem = 64 // Maximal dual bound violation for xx in the basic solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_BAS_NRM_BARX Dinfitem = 65 // Infinity norm of barx in the basic solution.
    MSK_DINF_SOL_BAS_NRM_SLC Dinfitem = 66 // Infinity norm of slc in the basic solution.
    MSK_DINF_SOL_BAS_NRM_SLX Dinfitem = 67 // Infinity norm of slx in the basic solution.
    MSK_DINF_SOL_BAS_NRM_SUC Dinfitem = 68 // Infinity norm of suc in the basic solution.
    MSK_DINF_SOL_BAS_NRM_SUX Dinfitem = 69 // Infinity norm of sux in the basic solution.
    MSK_DINF_SOL_BAS_NRM_XC Dinfitem = 70 // Infinity norm of xc in the basic solution.
    MSK_DINF_SOL_BAS_NRM_XX Dinfitem = 71 // Infinity norm of xx in the basic solution.
    MSK_DINF_SOL_BAS_NRM_Y Dinfitem = 72 // Infinity norm of Y in the basic solution.
    MSK_DINF_SOL_BAS_PRIMAL_OBJ Dinfitem = 73 // Primal objective value of the basic solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_BAS_PVIOLCON Dinfitem = 74 // Maximal primal bound violation for xc in the basic solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_BAS_PVIOLVAR Dinfitem = 75 // Maximal primal bound violation for xx in the basic solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITG_NRM_BARX Dinfitem = 76 // Infinity norm of barx in the integer solution.
    MSK_DINF_SOL_ITG_NRM_XC Dinfitem = 77 // Infinity norm of xc in the integer solution.
    MSK_DINF_SOL_ITG_NRM_XX Dinfitem = 78 // Infinity norm of xx in the integer solution.
    MSK_DINF_SOL_ITG_PRIMAL_OBJ Dinfitem = 79 // Primal objective value of the integer solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITG_PVIOLACC Dinfitem = 80 // Maximal primal violation for affine conic constraints in the integer solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITG_PVIOLBARVAR Dinfitem = 81 // Maximal primal bound violation for barx in the integer solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITG_PVIOLCON Dinfitem = 82 // Maximal primal bound violation for xc in the integer solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITG_PVIOLCONES Dinfitem = 83 // Maximal primal violation for primal conic constraints in the integer solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITG_PVIOLDJC Dinfitem = 84 // Maximal primal violation for disjunctive constraints in the integer solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITG_PVIOLITG Dinfitem = 85 // Maximal violation for the integer constraints in the integer solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITG_PVIOLVAR Dinfitem = 86 // Maximal primal bound violation for xx in the integer solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITR_DUAL_OBJ Dinfitem = 87 // Dual objective value of the interior-point solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITR_DVIOLACC Dinfitem = 88 // Maximal dual violation for affine conic constraints in the interior-point solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITR_DVIOLBARVAR Dinfitem = 89 // Maximal dual bound violation for barx in the interior-point solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITR_DVIOLCON Dinfitem = 90 // Maximal dual bound violation for xc in the interior-point solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITR_DVIOLCONES Dinfitem = 91 // Maximal dual violation for conic constraints in the interior-point solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITR_DVIOLVAR Dinfitem = 92 // Maximal dual bound violation for xx in the interior-point solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITR_NRM_BARS Dinfitem = 93 // Infinity norm of bars in the interior-point solution.
    MSK_DINF_SOL_ITR_NRM_BARX Dinfitem = 94 // Infinity norm of barx in the interior-point solution.
    MSK_DINF_SOL_ITR_NRM_SLC Dinfitem = 95 // Infinity norm of slc in the interior-point solution.
    MSK_DINF_SOL_ITR_NRM_SLX Dinfitem = 96 // Infinity norm of slx in the interior-point solution.
    MSK_DINF_SOL_ITR_NRM_SNX Dinfitem = 97 // Infinity norm of snx in the interior-point solution.
    MSK_DINF_SOL_ITR_NRM_SUC Dinfitem = 98 // Infinity norm of suc in the interior-point solution.
    MSK_DINF_SOL_ITR_NRM_SUX Dinfitem = 99 // Infinity norm of sux in the interior-point solution.
    MSK_DINF_SOL_ITR_NRM_XC Dinfitem = 100 // Infinity norm of xc in the interior-point solution.
    MSK_DINF_SOL_ITR_NRM_XX Dinfitem = 101 // Infinity norm of xx in the interior-point solution.
    MSK_DINF_SOL_ITR_NRM_Y Dinfitem = 102 // Infinity norm of Y in the interior-point solution.
    MSK_DINF_SOL_ITR_PRIMAL_OBJ Dinfitem = 103 // Primal objective value of the interior-point solution.
    MSK_DINF_SOL_ITR_PVIOLACC Dinfitem = 104 // Maximal primal violation for affine conic constraints in the interior-point solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITR_PVIOLBARVAR Dinfitem = 105 // Maximal primal bound violation for barx in the interior-point solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITR_PVIOLCON Dinfitem = 106 // Maximal primal bound violation for xc in the interior-point solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITR_PVIOLCONES Dinfitem = 107 // Maximal primal violation for conic constraints in the interior-point solution. Updated by the function updatesolutioninfo.
    MSK_DINF_SOL_ITR_PVIOLVAR Dinfitem = 108 // Maximal primal bound violation for xx in the interior-point solution. Updated by the function updatesolutioninfo.
    MSK_DINF_TO_CONIC_TIME Dinfitem = 109 // Time spent in the last to conic reformulation.
    MSK_DINF_WRITE_DATA_TIME Dinfitem = 110 // Time spent writing the data file.
    MSK_DINF_END Dinfitem = 110
)
func (self Dinfitem) String() string {
  switch self {
    case MSK_DINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_DENSITY: return "MSK_DINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_DENSITY"
    case MSK_DINF_BI_CLEAN_DUAL_TIME: return "MSK_DINF_BI_CLEAN_DUAL_TIME"
    case MSK_DINF_BI_CLEAN_PRIMAL_TIME: return "MSK_DINF_BI_CLEAN_PRIMAL_TIME"
    case MSK_DINF_BI_CLEAN_TIME: return "MSK_DINF_BI_CLEAN_TIME"
    case MSK_DINF_BI_DUAL_TIME: return "MSK_DINF_BI_DUAL_TIME"
    case MSK_DINF_BI_PRIMAL_TIME: return "MSK_DINF_BI_PRIMAL_TIME"
    case MSK_DINF_BI_TIME: return "MSK_DINF_BI_TIME"
    case MSK_DINF_INTPNT_DUAL_FEAS: return "MSK_DINF_INTPNT_DUAL_FEAS"
    case MSK_DINF_INTPNT_DUAL_OBJ: return "MSK_DINF_INTPNT_DUAL_OBJ"
    case MSK_DINF_INTPNT_FACTOR_NUM_FLOPS: return "MSK_DINF_INTPNT_FACTOR_NUM_FLOPS"
    case MSK_DINF_INTPNT_OPT_STATUS: return "MSK_DINF_INTPNT_OPT_STATUS"
    case MSK_DINF_INTPNT_ORDER_TIME: return "MSK_DINF_INTPNT_ORDER_TIME"
    case MSK_DINF_INTPNT_PRIMAL_FEAS: return "MSK_DINF_INTPNT_PRIMAL_FEAS"
    case MSK_DINF_INTPNT_PRIMAL_OBJ: return "MSK_DINF_INTPNT_PRIMAL_OBJ"
    case MSK_DINF_INTPNT_TIME: return "MSK_DINF_INTPNT_TIME"
    case MSK_DINF_MIO_CLIQUE_SELECTION_TIME: return "MSK_DINF_MIO_CLIQUE_SELECTION_TIME"
    case MSK_DINF_MIO_CLIQUE_SEPARATION_TIME: return "MSK_DINF_MIO_CLIQUE_SEPARATION_TIME"
    case MSK_DINF_MIO_CMIR_SELECTION_TIME: return "MSK_DINF_MIO_CMIR_SELECTION_TIME"
    case MSK_DINF_MIO_CMIR_SEPARATION_TIME: return "MSK_DINF_MIO_CMIR_SEPARATION_TIME"
    case MSK_DINF_MIO_CONSTRUCT_SOLUTION_OBJ: return "MSK_DINF_MIO_CONSTRUCT_SOLUTION_OBJ"
    case MSK_DINF_MIO_DUAL_BOUND_AFTER_PRESOLVE: return "MSK_DINF_MIO_DUAL_BOUND_AFTER_PRESOLVE"
    case MSK_DINF_MIO_GMI_SELECTION_TIME: return "MSK_DINF_MIO_GMI_SELECTION_TIME"
    case MSK_DINF_MIO_GMI_SEPARATION_TIME: return "MSK_DINF_MIO_GMI_SEPARATION_TIME"
    case MSK_DINF_MIO_IMPLIED_BOUND_SELECTION_TIME: return "MSK_DINF_MIO_IMPLIED_BOUND_SELECTION_TIME"
    case MSK_DINF_MIO_IMPLIED_BOUND_SEPARATION_TIME: return "MSK_DINF_MIO_IMPLIED_BOUND_SEPARATION_TIME"
    case MSK_DINF_MIO_INITIAL_FEASIBLE_SOLUTION_OBJ: return "MSK_DINF_MIO_INITIAL_FEASIBLE_SOLUTION_OBJ"
    case MSK_DINF_MIO_KNAPSACK_COVER_SELECTION_TIME: return "MSK_DINF_MIO_KNAPSACK_COVER_SELECTION_TIME"
    case MSK_DINF_MIO_KNAPSACK_COVER_SEPARATION_TIME: return "MSK_DINF_MIO_KNAPSACK_COVER_SEPARATION_TIME"
    case MSK_DINF_MIO_LIPRO_SELECTION_TIME: return "MSK_DINF_MIO_LIPRO_SELECTION_TIME"
    case MSK_DINF_MIO_LIPRO_SEPARATION_TIME: return "MSK_DINF_MIO_LIPRO_SEPARATION_TIME"
    case MSK_DINF_MIO_OBJ_ABS_GAP: return "MSK_DINF_MIO_OBJ_ABS_GAP"
    case MSK_DINF_MIO_OBJ_BOUND: return "MSK_DINF_MIO_OBJ_BOUND"
    case MSK_DINF_MIO_OBJ_INT: return "MSK_DINF_MIO_OBJ_INT"
    case MSK_DINF_MIO_OBJ_REL_GAP: return "MSK_DINF_MIO_OBJ_REL_GAP"
    case MSK_DINF_MIO_PROBING_TIME: return "MSK_DINF_MIO_PROBING_TIME"
    case MSK_DINF_MIO_ROOT_CUT_SELECTION_TIME: return "MSK_DINF_MIO_ROOT_CUT_SELECTION_TIME"
    case MSK_DINF_MIO_ROOT_CUT_SEPARATION_TIME: return "MSK_DINF_MIO_ROOT_CUT_SEPARATION_TIME"
    case MSK_DINF_MIO_ROOT_OPTIMIZER_TIME: return "MSK_DINF_MIO_ROOT_OPTIMIZER_TIME"
    case MSK_DINF_MIO_ROOT_PRESOLVE_TIME: return "MSK_DINF_MIO_ROOT_PRESOLVE_TIME"
    case MSK_DINF_MIO_ROOT_TIME: return "MSK_DINF_MIO_ROOT_TIME"
    case MSK_DINF_MIO_SYMMETRY_DETECTION_TIME: return "MSK_DINF_MIO_SYMMETRY_DETECTION_TIME"
    case MSK_DINF_MIO_SYMMETRY_FACTOR: return "MSK_DINF_MIO_SYMMETRY_FACTOR"
    case MSK_DINF_MIO_TIME: return "MSK_DINF_MIO_TIME"
    case MSK_DINF_MIO_USER_OBJ_CUT: return "MSK_DINF_MIO_USER_OBJ_CUT"
    case MSK_DINF_OPTIMIZER_TICKS: return "MSK_DINF_OPTIMIZER_TICKS"
    case MSK_DINF_OPTIMIZER_TIME: return "MSK_DINF_OPTIMIZER_TIME"
    case MSK_DINF_PRESOLVE_ELI_TIME: return "MSK_DINF_PRESOLVE_ELI_TIME"
    case MSK_DINF_PRESOLVE_LINDEP_TIME: return "MSK_DINF_PRESOLVE_LINDEP_TIME"
    case MSK_DINF_PRESOLVE_TIME: return "MSK_DINF_PRESOLVE_TIME"
    case MSK_DINF_PRESOLVE_TOTAL_PRIMAL_PERTURBATION: return "MSK_DINF_PRESOLVE_TOTAL_PRIMAL_PERTURBATION"
    case MSK_DINF_PRIMAL_REPAIR_PENALTY_OBJ: return "MSK_DINF_PRIMAL_REPAIR_PENALTY_OBJ"
    case MSK_DINF_QCQO_REFORMULATE_MAX_PERTURBATION: return "MSK_DINF_QCQO_REFORMULATE_MAX_PERTURBATION"
    case MSK_DINF_QCQO_REFORMULATE_TIME: return "MSK_DINF_QCQO_REFORMULATE_TIME"
    case MSK_DINF_QCQO_REFORMULATE_WORST_CHOLESKY_COLUMN_SCALING: return "MSK_DINF_QCQO_REFORMULATE_WORST_CHOLESKY_COLUMN_SCALING"
    case MSK_DINF_QCQO_REFORMULATE_WORST_CHOLESKY_DIAG_SCALING: return "MSK_DINF_QCQO_REFORMULATE_WORST_CHOLESKY_DIAG_SCALING"
    case MSK_DINF_READ_DATA_TIME: return "MSK_DINF_READ_DATA_TIME"
    case MSK_DINF_REMOTE_TIME: return "MSK_DINF_REMOTE_TIME"
    case MSK_DINF_SIM_DUAL_TIME: return "MSK_DINF_SIM_DUAL_TIME"
    case MSK_DINF_SIM_FEAS: return "MSK_DINF_SIM_FEAS"
    case MSK_DINF_SIM_OBJ: return "MSK_DINF_SIM_OBJ"
    case MSK_DINF_SIM_PRIMAL_TIME: return "MSK_DINF_SIM_PRIMAL_TIME"
    case MSK_DINF_SIM_TIME: return "MSK_DINF_SIM_TIME"
    case MSK_DINF_SOL_BAS_DUAL_OBJ: return "MSK_DINF_SOL_BAS_DUAL_OBJ"
    case MSK_DINF_SOL_BAS_DVIOLCON: return "MSK_DINF_SOL_BAS_DVIOLCON"
    case MSK_DINF_SOL_BAS_DVIOLVAR: return "MSK_DINF_SOL_BAS_DVIOLVAR"
    case MSK_DINF_SOL_BAS_NRM_BARX: return "MSK_DINF_SOL_BAS_NRM_BARX"
    case MSK_DINF_SOL_BAS_NRM_SLC: return "MSK_DINF_SOL_BAS_NRM_SLC"
    case MSK_DINF_SOL_BAS_NRM_SLX: return "MSK_DINF_SOL_BAS_NRM_SLX"
    case MSK_DINF_SOL_BAS_NRM_SUC: return "MSK_DINF_SOL_BAS_NRM_SUC"
    case MSK_DINF_SOL_BAS_NRM_SUX: return "MSK_DINF_SOL_BAS_NRM_SUX"
    case MSK_DINF_SOL_BAS_NRM_XC: return "MSK_DINF_SOL_BAS_NRM_XC"
    case MSK_DINF_SOL_BAS_NRM_XX: return "MSK_DINF_SOL_BAS_NRM_XX"
    case MSK_DINF_SOL_BAS_NRM_Y: return "MSK_DINF_SOL_BAS_NRM_Y"
    case MSK_DINF_SOL_BAS_PRIMAL_OBJ: return "MSK_DINF_SOL_BAS_PRIMAL_OBJ"
    case MSK_DINF_SOL_BAS_PVIOLCON: return "MSK_DINF_SOL_BAS_PVIOLCON"
    case MSK_DINF_SOL_BAS_PVIOLVAR: return "MSK_DINF_SOL_BAS_PVIOLVAR"
    case MSK_DINF_SOL_ITG_NRM_BARX: return "MSK_DINF_SOL_ITG_NRM_BARX"
    case MSK_DINF_SOL_ITG_NRM_XC: return "MSK_DINF_SOL_ITG_NRM_XC"
    case MSK_DINF_SOL_ITG_NRM_XX: return "MSK_DINF_SOL_ITG_NRM_XX"
    case MSK_DINF_SOL_ITG_PRIMAL_OBJ: return "MSK_DINF_SOL_ITG_PRIMAL_OBJ"
    case MSK_DINF_SOL_ITG_PVIOLACC: return "MSK_DINF_SOL_ITG_PVIOLACC"
    case MSK_DINF_SOL_ITG_PVIOLBARVAR: return "MSK_DINF_SOL_ITG_PVIOLBARVAR"
    case MSK_DINF_SOL_ITG_PVIOLCON: return "MSK_DINF_SOL_ITG_PVIOLCON"
    case MSK_DINF_SOL_ITG_PVIOLCONES: return "MSK_DINF_SOL_ITG_PVIOLCONES"
    case MSK_DINF_SOL_ITG_PVIOLDJC: return "MSK_DINF_SOL_ITG_PVIOLDJC"
    case MSK_DINF_SOL_ITG_PVIOLITG: return "MSK_DINF_SOL_ITG_PVIOLITG"
    case MSK_DINF_SOL_ITG_PVIOLVAR: return "MSK_DINF_SOL_ITG_PVIOLVAR"
    case MSK_DINF_SOL_ITR_DUAL_OBJ: return "MSK_DINF_SOL_ITR_DUAL_OBJ"
    case MSK_DINF_SOL_ITR_DVIOLACC: return "MSK_DINF_SOL_ITR_DVIOLACC"
    case MSK_DINF_SOL_ITR_DVIOLBARVAR: return "MSK_DINF_SOL_ITR_DVIOLBARVAR"
    case MSK_DINF_SOL_ITR_DVIOLCON: return "MSK_DINF_SOL_ITR_DVIOLCON"
    case MSK_DINF_SOL_ITR_DVIOLCONES: return "MSK_DINF_SOL_ITR_DVIOLCONES"
    case MSK_DINF_SOL_ITR_DVIOLVAR: return "MSK_DINF_SOL_ITR_DVIOLVAR"
    case MSK_DINF_SOL_ITR_NRM_BARS: return "MSK_DINF_SOL_ITR_NRM_BARS"
    case MSK_DINF_SOL_ITR_NRM_BARX: return "MSK_DINF_SOL_ITR_NRM_BARX"
    case MSK_DINF_SOL_ITR_NRM_SLC: return "MSK_DINF_SOL_ITR_NRM_SLC"
    case MSK_DINF_SOL_ITR_NRM_SLX: return "MSK_DINF_SOL_ITR_NRM_SLX"
    case MSK_DINF_SOL_ITR_NRM_SNX: return "MSK_DINF_SOL_ITR_NRM_SNX"
    case MSK_DINF_SOL_ITR_NRM_SUC: return "MSK_DINF_SOL_ITR_NRM_SUC"
    case MSK_DINF_SOL_ITR_NRM_SUX: return "MSK_DINF_SOL_ITR_NRM_SUX"
    case MSK_DINF_SOL_ITR_NRM_XC: return "MSK_DINF_SOL_ITR_NRM_XC"
    case MSK_DINF_SOL_ITR_NRM_XX: return "MSK_DINF_SOL_ITR_NRM_XX"
    case MSK_DINF_SOL_ITR_NRM_Y: return "MSK_DINF_SOL_ITR_NRM_Y"
    case MSK_DINF_SOL_ITR_PRIMAL_OBJ: return "MSK_DINF_SOL_ITR_PRIMAL_OBJ"
    case MSK_DINF_SOL_ITR_PVIOLACC: return "MSK_DINF_SOL_ITR_PVIOLACC"
    case MSK_DINF_SOL_ITR_PVIOLBARVAR: return "MSK_DINF_SOL_ITR_PVIOLBARVAR"
    case MSK_DINF_SOL_ITR_PVIOLCON: return "MSK_DINF_SOL_ITR_PVIOLCON"
    case MSK_DINF_SOL_ITR_PVIOLCONES: return "MSK_DINF_SOL_ITR_PVIOLCONES"
    case MSK_DINF_SOL_ITR_PVIOLVAR: return "MSK_DINF_SOL_ITR_PVIOLVAR"
    case MSK_DINF_TO_CONIC_TIME: return "MSK_DINF_TO_CONIC_TIME"
    case MSK_DINF_WRITE_DATA_TIME: return "MSK_DINF_WRITE_DATA_TIME"
    default: return "<unknown>"
  }
} // Dinfitem.String()

type Feature int32
// License feature
const (
    MSK_FEATURE_PTS Feature = 0 // Base system.
    MSK_FEATURE_PTON Feature = 1 // Conic extension.
    MSK_FEATURE_END Feature = 1
)
func (self Feature) String() string {
  switch self {
    case MSK_FEATURE_PTS: return "MSK_FEATURE_PTS"
    case MSK_FEATURE_PTON: return "MSK_FEATURE_PTON"
    default: return "<unknown>"
  }
} // Feature.String()

type Dparam int32
// Double parameters
const (
    MSK_DPAR_ANA_SOL_INFEAS_TOL Dparam = 0 // If a constraint violates its bound with an amount larger than this value, the constraint name, index and violation will be printed by the solution analyzer.
    MSK_DPAR_BASIS_REL_TOL_S Dparam = 1 // Maximum relative dual bound violation allowed in an optimal basic solution.
    MSK_DPAR_BASIS_TOL_S Dparam = 2 // Maximum absolute dual bound violation in an optimal basic solution.
    MSK_DPAR_BASIS_TOL_X Dparam = 3 // Maximum absolute primal bound violation allowed in an optimal basic solution.
    MSK_DPAR_CHECK_CONVEXITY_REL_TOL Dparam = 4 // Convexity check tolerance.
    MSK_DPAR_DATA_SYM_MAT_TOL Dparam = 5 // Zero tolerance threshold for symmetric matrices.
    MSK_DPAR_DATA_SYM_MAT_TOL_HUGE Dparam = 6 // Data tolerance threshold.
    MSK_DPAR_DATA_SYM_MAT_TOL_LARGE Dparam = 7 // Data tolerance threshold.
    MSK_DPAR_DATA_TOL_AIJ_HUGE Dparam = 8 // Data tolerance threshold.
    MSK_DPAR_DATA_TOL_AIJ_LARGE Dparam = 9 // Data tolerance threshold.
    MSK_DPAR_DATA_TOL_BOUND_INF Dparam = 10 // Data tolerance threshold.
    MSK_DPAR_DATA_TOL_BOUND_WRN Dparam = 11 // Data tolerance threshold.
    MSK_DPAR_DATA_TOL_C_HUGE Dparam = 12 // Data tolerance threshold.
    MSK_DPAR_DATA_TOL_CJ_LARGE Dparam = 13 // Data tolerance threshold.
    MSK_DPAR_DATA_TOL_QIJ Dparam = 14 // Data tolerance threshold.
    MSK_DPAR_DATA_TOL_X Dparam = 15 // Data tolerance threshold.
    MSK_DPAR_INTPNT_CO_TOL_DFEAS Dparam = 16 // Dual feasibility tolerance used by the interior-point optimizer for conic problems.
    MSK_DPAR_INTPNT_CO_TOL_INFEAS Dparam = 17 // Infeasibility tolerance used by the interior-point optimizer for conic problems.
    MSK_DPAR_INTPNT_CO_TOL_MU_RED Dparam = 18 // Relative complementarity gap tolerance used by the interior-point optimizer for conic problems.
    MSK_DPAR_INTPNT_CO_TOL_NEAR_REL Dparam = 19 // Optimality tolerance used by the interior-point optimizer for conic problems.
    MSK_DPAR_INTPNT_CO_TOL_PFEAS Dparam = 20 // Primal feasibility tolerance used by the interior-point optimizer for conic problems.
    MSK_DPAR_INTPNT_CO_TOL_REL_GAP Dparam = 21 // Relative gap termination tolerance used by the interior-point optimizer for conic problems.
    MSK_DPAR_INTPNT_QO_TOL_DFEAS Dparam = 22 // Dual feasibility tolerance used by the interior-point optimizer for quadratic problems.
    MSK_DPAR_INTPNT_QO_TOL_INFEAS Dparam = 23 // Infeasibility tolerance used by the interior-point optimizer for quadratic problems.
    MSK_DPAR_INTPNT_QO_TOL_MU_RED Dparam = 24 // Relative complementarity gap tolerance used by the interior-point optimizer for quadratic problems.
    MSK_DPAR_INTPNT_QO_TOL_NEAR_REL Dparam = 25 // Optimality tolerance used by the interior-point optimizer for quadratic problems.
    MSK_DPAR_INTPNT_QO_TOL_PFEAS Dparam = 26 // Primal feasibility tolerance used by the interior-point optimizer for quadratic problems.
    MSK_DPAR_INTPNT_QO_TOL_REL_GAP Dparam = 27 // Relative gap termination tolerance used by the interior-point optimizer for quadratic problems.
    MSK_DPAR_INTPNT_TOL_DFEAS Dparam = 28 // Dual feasibility tolerance used by the interior-point optimizer for linear problems.
    MSK_DPAR_INTPNT_TOL_DSAFE Dparam = 29 // Controls the interior-point dual starting point.
    MSK_DPAR_INTPNT_TOL_INFEAS Dparam = 30 // Infeasibility tolerance used by the interior-point optimizer for linear problems.
    MSK_DPAR_INTPNT_TOL_MU_RED Dparam = 31 // Relative complementarity gap tolerance used by the interior-point optimizer for linear problems.
    MSK_DPAR_INTPNT_TOL_PATH Dparam = 32 // Interior-point centering aggressiveness.
    MSK_DPAR_INTPNT_TOL_PFEAS Dparam = 33 // Primal feasibility tolerance used by the interior-point optimizer for linear problems.
    MSK_DPAR_INTPNT_TOL_PSAFE Dparam = 34 // Controls the interior-point primal starting point.
    MSK_DPAR_INTPNT_TOL_REL_GAP Dparam = 35 // Relative gap termination tolerance used by the interior-point optimizer for linear problems.
    MSK_DPAR_INTPNT_TOL_REL_STEP Dparam = 36 // Relative step size to the boundary for linear and quadratic optimization problems.
    MSK_DPAR_INTPNT_TOL_STEP_SIZE Dparam = 37 // Minimal step size tolerance for the interior-point optimizer.
    MSK_DPAR_LOWER_OBJ_CUT Dparam = 38 // Objective bound.
    MSK_DPAR_LOWER_OBJ_CUT_FINITE_TRH Dparam = 39 // Objective bound.
    MSK_DPAR_MIO_DJC_MAX_BIGM Dparam = 40 // Maximum allowed big-M value when reformulating disjunctive constraints to linear constraints.
    MSK_DPAR_MIO_MAX_TIME Dparam = 41 // Time limit for the mixed-integer optimizer.
    MSK_DPAR_MIO_REL_GAP_CONST Dparam = 42 // This value is used to compute the relative gap for the solution to an integer optimization problem.
    MSK_DPAR_MIO_TOL_ABS_GAP Dparam = 43 // Absolute optimality tolerance employed by the mixed-integer optimizer.
    MSK_DPAR_MIO_TOL_ABS_RELAX_INT Dparam = 44 // Integer feasibility tolerance.
    MSK_DPAR_MIO_TOL_FEAS Dparam = 45 // Feasibility tolerance for mixed integer solver.
    MSK_DPAR_MIO_TOL_REL_DUAL_BOUND_IMPROVEMENT Dparam = 46 // Controls cut generation for mixed-integer optimizer.
    MSK_DPAR_MIO_TOL_REL_GAP Dparam = 47 // Relative optimality tolerance employed by the mixed-integer optimizer.
    MSK_DPAR_OPTIMIZER_MAX_TICKS Dparam = 48 // Solver ticks limit.
    MSK_DPAR_OPTIMIZER_MAX_TIME Dparam = 49 // Solver time limit.
    MSK_DPAR_PRESOLVE_TOL_ABS_LINDEP Dparam = 50 // Absolute tolerance employed by the linear dependency checker.
    MSK_DPAR_PRESOLVE_TOL_AIJ Dparam = 51 // Absolute zero tolerance employed for constraint coefficients in the presolve.
    MSK_DPAR_PRESOLVE_TOL_PRIMAL_INFEAS_PERTURBATION Dparam = 52 // The presolve is allowed to perturb a bound on a constraint or variable by this amount if it removes an infeasibility.
    MSK_DPAR_PRESOLVE_TOL_REL_LINDEP Dparam = 53 // Relative tolerance employed by the linear dependency checker.
    MSK_DPAR_PRESOLVE_TOL_S Dparam = 54 // Absolute zero tolerance employed for slack variables in the presolve.
    MSK_DPAR_PRESOLVE_TOL_X Dparam = 55 // Absolute zero tolerance employed for variables in the presolve.
    MSK_DPAR_QCQO_REFORMULATE_REL_DROP_TOL Dparam = 56 // This parameter determines when columns are dropped in incomplete Cholesky factorization during reformulation of quadratic problems.
    MSK_DPAR_SEMIDEFINITE_TOL_APPROX Dparam = 57 // Tolerance to define a matrix to be positive semidefinite.
    MSK_DPAR_SIM_LU_TOL_REL_PIV Dparam = 58 // Relative pivot tolerance employed when computing the LU factorization of the basis matrix.
    MSK_DPAR_SIMPLEX_ABS_TOL_PIV Dparam = 59 // Absolute pivot tolerance employed by the simplex optimizers.
    MSK_DPAR_UPPER_OBJ_CUT Dparam = 60 // Objective bound.
    MSK_DPAR_UPPER_OBJ_CUT_FINITE_TRH Dparam = 61 // Objective bound.
    MSK_DPAR_END Dparam = 61
)
func (self Dparam) String() string {
  switch self {
    case MSK_DPAR_ANA_SOL_INFEAS_TOL: return "MSK_DPAR_ANA_SOL_INFEAS_TOL"
    case MSK_DPAR_BASIS_REL_TOL_S: return "MSK_DPAR_BASIS_REL_TOL_S"
    case MSK_DPAR_BASIS_TOL_S: return "MSK_DPAR_BASIS_TOL_S"
    case MSK_DPAR_BASIS_TOL_X: return "MSK_DPAR_BASIS_TOL_X"
    case MSK_DPAR_CHECK_CONVEXITY_REL_TOL: return "MSK_DPAR_CHECK_CONVEXITY_REL_TOL"
    case MSK_DPAR_DATA_SYM_MAT_TOL: return "MSK_DPAR_DATA_SYM_MAT_TOL"
    case MSK_DPAR_DATA_SYM_MAT_TOL_HUGE: return "MSK_DPAR_DATA_SYM_MAT_TOL_HUGE"
    case MSK_DPAR_DATA_SYM_MAT_TOL_LARGE: return "MSK_DPAR_DATA_SYM_MAT_TOL_LARGE"
    case MSK_DPAR_DATA_TOL_AIJ_HUGE: return "MSK_DPAR_DATA_TOL_AIJ_HUGE"
    case MSK_DPAR_DATA_TOL_AIJ_LARGE: return "MSK_DPAR_DATA_TOL_AIJ_LARGE"
    case MSK_DPAR_DATA_TOL_BOUND_INF: return "MSK_DPAR_DATA_TOL_BOUND_INF"
    case MSK_DPAR_DATA_TOL_BOUND_WRN: return "MSK_DPAR_DATA_TOL_BOUND_WRN"
    case MSK_DPAR_DATA_TOL_C_HUGE: return "MSK_DPAR_DATA_TOL_C_HUGE"
    case MSK_DPAR_DATA_TOL_CJ_LARGE: return "MSK_DPAR_DATA_TOL_CJ_LARGE"
    case MSK_DPAR_DATA_TOL_QIJ: return "MSK_DPAR_DATA_TOL_QIJ"
    case MSK_DPAR_DATA_TOL_X: return "MSK_DPAR_DATA_TOL_X"
    case MSK_DPAR_INTPNT_CO_TOL_DFEAS: return "MSK_DPAR_INTPNT_CO_TOL_DFEAS"
    case MSK_DPAR_INTPNT_CO_TOL_INFEAS: return "MSK_DPAR_INTPNT_CO_TOL_INFEAS"
    case MSK_DPAR_INTPNT_CO_TOL_MU_RED: return "MSK_DPAR_INTPNT_CO_TOL_MU_RED"
    case MSK_DPAR_INTPNT_CO_TOL_NEAR_REL: return "MSK_DPAR_INTPNT_CO_TOL_NEAR_REL"
    case MSK_DPAR_INTPNT_CO_TOL_PFEAS: return "MSK_DPAR_INTPNT_CO_TOL_PFEAS"
    case MSK_DPAR_INTPNT_CO_TOL_REL_GAP: return "MSK_DPAR_INTPNT_CO_TOL_REL_GAP"
    case MSK_DPAR_INTPNT_QO_TOL_DFEAS: return "MSK_DPAR_INTPNT_QO_TOL_DFEAS"
    case MSK_DPAR_INTPNT_QO_TOL_INFEAS: return "MSK_DPAR_INTPNT_QO_TOL_INFEAS"
    case MSK_DPAR_INTPNT_QO_TOL_MU_RED: return "MSK_DPAR_INTPNT_QO_TOL_MU_RED"
    case MSK_DPAR_INTPNT_QO_TOL_NEAR_REL: return "MSK_DPAR_INTPNT_QO_TOL_NEAR_REL"
    case MSK_DPAR_INTPNT_QO_TOL_PFEAS: return "MSK_DPAR_INTPNT_QO_TOL_PFEAS"
    case MSK_DPAR_INTPNT_QO_TOL_REL_GAP: return "MSK_DPAR_INTPNT_QO_TOL_REL_GAP"
    case MSK_DPAR_INTPNT_TOL_DFEAS: return "MSK_DPAR_INTPNT_TOL_DFEAS"
    case MSK_DPAR_INTPNT_TOL_DSAFE: return "MSK_DPAR_INTPNT_TOL_DSAFE"
    case MSK_DPAR_INTPNT_TOL_INFEAS: return "MSK_DPAR_INTPNT_TOL_INFEAS"
    case MSK_DPAR_INTPNT_TOL_MU_RED: return "MSK_DPAR_INTPNT_TOL_MU_RED"
    case MSK_DPAR_INTPNT_TOL_PATH: return "MSK_DPAR_INTPNT_TOL_PATH"
    case MSK_DPAR_INTPNT_TOL_PFEAS: return "MSK_DPAR_INTPNT_TOL_PFEAS"
    case MSK_DPAR_INTPNT_TOL_PSAFE: return "MSK_DPAR_INTPNT_TOL_PSAFE"
    case MSK_DPAR_INTPNT_TOL_REL_GAP: return "MSK_DPAR_INTPNT_TOL_REL_GAP"
    case MSK_DPAR_INTPNT_TOL_REL_STEP: return "MSK_DPAR_INTPNT_TOL_REL_STEP"
    case MSK_DPAR_INTPNT_TOL_STEP_SIZE: return "MSK_DPAR_INTPNT_TOL_STEP_SIZE"
    case MSK_DPAR_LOWER_OBJ_CUT: return "MSK_DPAR_LOWER_OBJ_CUT"
    case MSK_DPAR_LOWER_OBJ_CUT_FINITE_TRH: return "MSK_DPAR_LOWER_OBJ_CUT_FINITE_TRH"
    case MSK_DPAR_MIO_DJC_MAX_BIGM: return "MSK_DPAR_MIO_DJC_MAX_BIGM"
    case MSK_DPAR_MIO_MAX_TIME: return "MSK_DPAR_MIO_MAX_TIME"
    case MSK_DPAR_MIO_REL_GAP_CONST: return "MSK_DPAR_MIO_REL_GAP_CONST"
    case MSK_DPAR_MIO_TOL_ABS_GAP: return "MSK_DPAR_MIO_TOL_ABS_GAP"
    case MSK_DPAR_MIO_TOL_ABS_RELAX_INT: return "MSK_DPAR_MIO_TOL_ABS_RELAX_INT"
    case MSK_DPAR_MIO_TOL_FEAS: return "MSK_DPAR_MIO_TOL_FEAS"
    case MSK_DPAR_MIO_TOL_REL_DUAL_BOUND_IMPROVEMENT: return "MSK_DPAR_MIO_TOL_REL_DUAL_BOUND_IMPROVEMENT"
    case MSK_DPAR_MIO_TOL_REL_GAP: return "MSK_DPAR_MIO_TOL_REL_GAP"
    case MSK_DPAR_OPTIMIZER_MAX_TICKS: return "MSK_DPAR_OPTIMIZER_MAX_TICKS"
    case MSK_DPAR_OPTIMIZER_MAX_TIME: return "MSK_DPAR_OPTIMIZER_MAX_TIME"
    case MSK_DPAR_PRESOLVE_TOL_ABS_LINDEP: return "MSK_DPAR_PRESOLVE_TOL_ABS_LINDEP"
    case MSK_DPAR_PRESOLVE_TOL_AIJ: return "MSK_DPAR_PRESOLVE_TOL_AIJ"
    case MSK_DPAR_PRESOLVE_TOL_PRIMAL_INFEAS_PERTURBATION: return "MSK_DPAR_PRESOLVE_TOL_PRIMAL_INFEAS_PERTURBATION"
    case MSK_DPAR_PRESOLVE_TOL_REL_LINDEP: return "MSK_DPAR_PRESOLVE_TOL_REL_LINDEP"
    case MSK_DPAR_PRESOLVE_TOL_S: return "MSK_DPAR_PRESOLVE_TOL_S"
    case MSK_DPAR_PRESOLVE_TOL_X: return "MSK_DPAR_PRESOLVE_TOL_X"
    case MSK_DPAR_QCQO_REFORMULATE_REL_DROP_TOL: return "MSK_DPAR_QCQO_REFORMULATE_REL_DROP_TOL"
    case MSK_DPAR_SEMIDEFINITE_TOL_APPROX: return "MSK_DPAR_SEMIDEFINITE_TOL_APPROX"
    case MSK_DPAR_SIM_LU_TOL_REL_PIV: return "MSK_DPAR_SIM_LU_TOL_REL_PIV"
    case MSK_DPAR_SIMPLEX_ABS_TOL_PIV: return "MSK_DPAR_SIMPLEX_ABS_TOL_PIV"
    case MSK_DPAR_UPPER_OBJ_CUT: return "MSK_DPAR_UPPER_OBJ_CUT"
    case MSK_DPAR_UPPER_OBJ_CUT_FINITE_TRH: return "MSK_DPAR_UPPER_OBJ_CUT_FINITE_TRH"
    default: return "<unknown>"
  }
} // Dparam.String()

type Liinfitem int32
// Long integer information items.
const (
    MSK_LIINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_NUM_COLUMNS Liinfitem = 0 // Number of columns in the scalarized constraint matrix.
    MSK_LIINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_NUM_NZ Liinfitem = 1 // Number of non-zero entries in the scalarized constraint matrix.
    MSK_LIINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_NUM_ROWS Liinfitem = 2 // Number of rows in the scalarized constraint matrix.
    MSK_LIINF_BI_CLEAN_DUAL_DEG_ITER Liinfitem = 3 // Number of dual degenerate clean iterations performed in the basis identification.
    MSK_LIINF_BI_CLEAN_DUAL_ITER Liinfitem = 4 // Number of dual clean iterations performed in the basis identification.
    MSK_LIINF_BI_CLEAN_PRIMAL_DEG_ITER Liinfitem = 5 // Number of primal degenerate clean iterations performed in the basis identification.
    MSK_LIINF_BI_CLEAN_PRIMAL_ITER Liinfitem = 6 // Number of primal clean iterations performed in the basis identification.
    MSK_LIINF_BI_DUAL_ITER Liinfitem = 7 // Number of dual pivots performed in the basis identification.
    MSK_LIINF_BI_PRIMAL_ITER Liinfitem = 8 // Number of primal pivots performed in the basis identification.
    MSK_LIINF_INTPNT_FACTOR_NUM_NZ Liinfitem = 9 // Number of non-zeros in factorization.
    MSK_LIINF_MIO_ANZ Liinfitem = 10 // Number of non-zero entries in the constraint matrix of the problem to be solved by the mixed-integer optimizer.
    MSK_LIINF_MIO_INTPNT_ITER Liinfitem = 11 // Number of interior-point iterations performed by the mixed-integer optimizer.
    MSK_LIINF_MIO_NUM_DUAL_ILLPOSED_CER Liinfitem = 12 // Number of dual illposed certificates encountered by the mixed-integer optimizer.
    MSK_LIINF_MIO_NUM_PRIM_ILLPOSED_CER Liinfitem = 13 // Number of primal illposed certificates encountered by the mixed-integer optimizer.
    MSK_LIINF_MIO_PRESOLVED_ANZ Liinfitem = 14 // Number of non-zero entries in the constraint matrix of the problem after the mixed-integer optimizer's presolve.
    MSK_LIINF_MIO_SIMPLEX_ITER Liinfitem = 15 // Number of simplex iterations performed by the mixed-integer optimizer.
    MSK_LIINF_RD_NUMACC Liinfitem = 16 // Number of affince conic constraints.
    MSK_LIINF_RD_NUMANZ Liinfitem = 17 // Number of non-zeros in A that is read.
    MSK_LIINF_RD_NUMDJC Liinfitem = 18 // Number of disjuncive constraints.
    MSK_LIINF_RD_NUMQNZ Liinfitem = 19 // Number of Q non-zeros.
    MSK_LIINF_SIMPLEX_ITER Liinfitem = 20 // Number of iterations performed by the simplex optimizer.
    MSK_LIINF_END Liinfitem = 20
)
func (self Liinfitem) String() string {
  switch self {
    case MSK_LIINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_NUM_COLUMNS: return "MSK_LIINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_NUM_COLUMNS"
    case MSK_LIINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_NUM_NZ: return "MSK_LIINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_NUM_NZ"
    case MSK_LIINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_NUM_ROWS: return "MSK_LIINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_NUM_ROWS"
    case MSK_LIINF_BI_CLEAN_DUAL_DEG_ITER: return "MSK_LIINF_BI_CLEAN_DUAL_DEG_ITER"
    case MSK_LIINF_BI_CLEAN_DUAL_ITER: return "MSK_LIINF_BI_CLEAN_DUAL_ITER"
    case MSK_LIINF_BI_CLEAN_PRIMAL_DEG_ITER: return "MSK_LIINF_BI_CLEAN_PRIMAL_DEG_ITER"
    case MSK_LIINF_BI_CLEAN_PRIMAL_ITER: return "MSK_LIINF_BI_CLEAN_PRIMAL_ITER"
    case MSK_LIINF_BI_DUAL_ITER: return "MSK_LIINF_BI_DUAL_ITER"
    case MSK_LIINF_BI_PRIMAL_ITER: return "MSK_LIINF_BI_PRIMAL_ITER"
    case MSK_LIINF_INTPNT_FACTOR_NUM_NZ: return "MSK_LIINF_INTPNT_FACTOR_NUM_NZ"
    case MSK_LIINF_MIO_ANZ: return "MSK_LIINF_MIO_ANZ"
    case MSK_LIINF_MIO_INTPNT_ITER: return "MSK_LIINF_MIO_INTPNT_ITER"
    case MSK_LIINF_MIO_NUM_DUAL_ILLPOSED_CER: return "MSK_LIINF_MIO_NUM_DUAL_ILLPOSED_CER"
    case MSK_LIINF_MIO_NUM_PRIM_ILLPOSED_CER: return "MSK_LIINF_MIO_NUM_PRIM_ILLPOSED_CER"
    case MSK_LIINF_MIO_PRESOLVED_ANZ: return "MSK_LIINF_MIO_PRESOLVED_ANZ"
    case MSK_LIINF_MIO_SIMPLEX_ITER: return "MSK_LIINF_MIO_SIMPLEX_ITER"
    case MSK_LIINF_RD_NUMACC: return "MSK_LIINF_RD_NUMACC"
    case MSK_LIINF_RD_NUMANZ: return "MSK_LIINF_RD_NUMANZ"
    case MSK_LIINF_RD_NUMDJC: return "MSK_LIINF_RD_NUMDJC"
    case MSK_LIINF_RD_NUMQNZ: return "MSK_LIINF_RD_NUMQNZ"
    case MSK_LIINF_SIMPLEX_ITER: return "MSK_LIINF_SIMPLEX_ITER"
    default: return "<unknown>"
  }
} // Liinfitem.String()

type Iinfitem int32
// Integer information items.
const (
    MSK_IINF_ANA_PRO_NUM_CON Iinfitem = 0 // Number of constraints in the problem.
    MSK_IINF_ANA_PRO_NUM_CON_EQ Iinfitem = 1 // Number of equality constraints.
    MSK_IINF_ANA_PRO_NUM_CON_FR Iinfitem = 2 // Number of unbounded constraints.
    MSK_IINF_ANA_PRO_NUM_CON_LO Iinfitem = 3 // Number of constraints with a lower bound and an infinite upper bound.
    MSK_IINF_ANA_PRO_NUM_CON_RA Iinfitem = 4 // Number of constraints with finite lower and upper bounds.
    MSK_IINF_ANA_PRO_NUM_CON_UP Iinfitem = 5 // Number of constraints with an upper bound and an infinite lower bound.
    MSK_IINF_ANA_PRO_NUM_VAR Iinfitem = 6 // Number of variables in the problem.
    MSK_IINF_ANA_PRO_NUM_VAR_BIN Iinfitem = 7 // Number of binary variables.
    MSK_IINF_ANA_PRO_NUM_VAR_CONT Iinfitem = 8 // Number of continuous variables.
    MSK_IINF_ANA_PRO_NUM_VAR_EQ Iinfitem = 9 // Number of fixed variables.
    MSK_IINF_ANA_PRO_NUM_VAR_FR Iinfitem = 10 // Number of unbounded constraints.
    MSK_IINF_ANA_PRO_NUM_VAR_INT Iinfitem = 11 // Number of general integer variables.
    MSK_IINF_ANA_PRO_NUM_VAR_LO Iinfitem = 12 // Number of variables with a lower bound and an infinite upper bound.
    MSK_IINF_ANA_PRO_NUM_VAR_RA Iinfitem = 13 // Number of variables with finite lower and upper bounds.
    MSK_IINF_ANA_PRO_NUM_VAR_UP Iinfitem = 14 // Number of variables with an upper bound and an infinite lower bound.
    MSK_IINF_INTPNT_FACTOR_DIM_DENSE Iinfitem = 15 // Dimension of the dense sub system in factorization.
    MSK_IINF_INTPNT_ITER Iinfitem = 16 // Number of interior-point iterations since invoking the interior-point optimizer.
    MSK_IINF_INTPNT_NUM_THREADS Iinfitem = 17 // Number of threads that the interior-point optimizer is using.
    MSK_IINF_INTPNT_SOLVE_DUAL Iinfitem = 18 // Non-zero if the interior-point optimizer is solving the dual problem.
    MSK_IINF_MIO_ABSGAP_SATISFIED Iinfitem = 19 // Non-zero if absolute gap is within tolerances.
    MSK_IINF_MIO_CLIQUE_TABLE_SIZE Iinfitem = 20 // Size of the clique table.
    MSK_IINF_MIO_CONSTRUCT_SOLUTION Iinfitem = 21 // Informs if MOSEK successfully constructed an initial integer feasible solution.
    MSK_IINF_MIO_INITIAL_FEASIBLE_SOLUTION Iinfitem = 22 // Informs if MOSEK found the solution provided by the user to be feasible
    MSK_IINF_MIO_NODE_DEPTH Iinfitem = 23 // Depth of the last node solved.
    MSK_IINF_MIO_NUM_ACTIVE_NODES Iinfitem = 24 // Number of active branch and bound nodes.
    MSK_IINF_MIO_NUM_ACTIVE_ROOT_CUTS Iinfitem = 25 // Number of active cuts in the final relaxation after the mixed-integer optimizer's root cut generation.
    MSK_IINF_MIO_NUM_BRANCH Iinfitem = 26 // Number of branches performed during the optimization.
    MSK_IINF_MIO_NUM_INT_SOLUTIONS Iinfitem = 27 // Number of integer feasible solutions that have been found.
    MSK_IINF_MIO_NUM_RELAX Iinfitem = 28 // Number of relaxations solved during the optimization.
    MSK_IINF_MIO_NUM_REPEATED_PRESOLVE Iinfitem = 29 // Number of times presolve was repeated at root.
    MSK_IINF_MIO_NUM_RESTARTS Iinfitem = 30 // Number of restarts performed during the optimization.
    MSK_IINF_MIO_NUM_ROOT_CUT_ROUNDS Iinfitem = 31 // Number of cut separation rounds at the root node of the mixed-integer optimizer.
    MSK_IINF_MIO_NUM_SELECTED_CLIQUE_CUTS Iinfitem = 32 // Number of clique cuts selected to be included in the relaxation.
    MSK_IINF_MIO_NUM_SELECTED_CMIR_CUTS Iinfitem = 33 // Number of Complemented Mixed Integer Rounding (CMIR) cuts selected to be included in the relaxation.
    MSK_IINF_MIO_NUM_SELECTED_GOMORY_CUTS Iinfitem = 34 // Number of Gomory cuts selected to be included in the relaxation.
    MSK_IINF_MIO_NUM_SELECTED_IMPLIED_BOUND_CUTS Iinfitem = 35 // Number of implied bound cuts selected to be included in the relaxation.
    MSK_IINF_MIO_NUM_SELECTED_KNAPSACK_COVER_CUTS Iinfitem = 36 // Number of clique cuts selected to be included in the relaxation.
    MSK_IINF_MIO_NUM_SELECTED_LIPRO_CUTS Iinfitem = 37 // Number of lift-and-project cuts selected to be included in the relaxation.
    MSK_IINF_MIO_NUM_SEPARATED_CLIQUE_CUTS Iinfitem = 38 // Number of separated clique cuts.
    MSK_IINF_MIO_NUM_SEPARATED_CMIR_CUTS Iinfitem = 39 // Number of separated Complemented Mixed Integer Rounding (CMIR) cuts.
    MSK_IINF_MIO_NUM_SEPARATED_GOMORY_CUTS Iinfitem = 40 // Number of separated Gomory cuts.
    MSK_IINF_MIO_NUM_SEPARATED_IMPLIED_BOUND_CUTS Iinfitem = 41 // Number of separated implied bound cuts.
    MSK_IINF_MIO_NUM_SEPARATED_KNAPSACK_COVER_CUTS Iinfitem = 42 // Number of separated clique cuts.
    MSK_IINF_MIO_NUM_SEPARATED_LIPRO_CUTS Iinfitem = 43 // Number of separated lift-and-project cuts.
    MSK_IINF_MIO_NUM_SOLVED_NODES Iinfitem = 44 // Number of branch and bounds nodes solved in the main branch and bound tree.
    MSK_IINF_MIO_NUMBIN Iinfitem = 45 // Number of binary variables in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMBINCONEVAR Iinfitem = 46 // Number of binary cone variables in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMCON Iinfitem = 47 // Number of constraints in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMCONE Iinfitem = 48 // Number of cones in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMCONEVAR Iinfitem = 49 // Number of cone variables in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMCONT Iinfitem = 50 // Number of continuous variables in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMCONTCONEVAR Iinfitem = 51 // Number of continuous cone variables in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMDEXPCONES Iinfitem = 52 // Number of dual exponential cones in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMDJC Iinfitem = 53 // Number of disjunctive constraints in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMDPOWCONES Iinfitem = 54 // Number of dual power cones in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMINT Iinfitem = 55 // Number of integer variables in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMINTCONEVAR Iinfitem = 56 // Number of integer cone variables in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMPEXPCONES Iinfitem = 57 // Number of primal exponential cones in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMPPOWCONES Iinfitem = 58 // Number of primal power cones in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMQCONES Iinfitem = 59 // Number of quadratic cones in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMRQCONES Iinfitem = 60 // Number of rotated quadratic cones in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_NUMVAR Iinfitem = 61 // Number of variables in the problem to be solved by the mixed-integer optimizer.
    MSK_IINF_MIO_OBJ_BOUND_DEFINED Iinfitem = 62 // Non-zero if a valid objective bound has been found, otherwise zero.
    MSK_IINF_MIO_PRESOLVED_NUMBIN Iinfitem = 63 // Number of binary variables in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMBINCONEVAR Iinfitem = 64 // Number of binary cone variables in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMCON Iinfitem = 65 // Number of constraints in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMCONE Iinfitem = 66 // Number of cones in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMCONEVAR Iinfitem = 67 // Number of cone variables in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMCONT Iinfitem = 68 // Number of continuous variables in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMCONTCONEVAR Iinfitem = 69 // Number of continuous cone variables in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMDEXPCONES Iinfitem = 70 // Number of dual exponential cones in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMDJC Iinfitem = 71 // Number of disjunctive constraints in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMDPOWCONES Iinfitem = 72 // Number of dual power cones in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMINT Iinfitem = 73 // Number of integer variables in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMINTCONEVAR Iinfitem = 74 // Number of integer cone variables in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMPEXPCONES Iinfitem = 75 // Number of primal exponential cones in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMPPOWCONES Iinfitem = 76 // Number of primal power cones in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMQCONES Iinfitem = 77 // Number of quadratic cones in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMRQCONES Iinfitem = 78 // Number of rotated quadratic cones in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_PRESOLVED_NUMVAR Iinfitem = 79 // Number of variables in the problem after the mixed-integer optimizer's presolve.
    MSK_IINF_MIO_RELGAP_SATISFIED Iinfitem = 80 // Non-zero if relative gap is within tolerances.
    MSK_IINF_MIO_TOTAL_NUM_SELECTED_CUTS Iinfitem = 81 // Total number of cuts selected to be included in the relaxation by the mixed-integer optimizer.
    MSK_IINF_MIO_TOTAL_NUM_SEPARATED_CUTS Iinfitem = 82 // Total number of cuts separated by the mixed-integer optimizer.
    MSK_IINF_MIO_USER_OBJ_CUT Iinfitem = 83 // If it is non-zero, then the objective cut is used.
    MSK_IINF_OPT_NUMCON Iinfitem = 84 // Number of constraints in the problem solved when the optimizer is called.
    MSK_IINF_OPT_NUMVAR Iinfitem = 85 // Number of variables in the problem solved when the optimizer is called
    MSK_IINF_OPTIMIZE_RESPONSE Iinfitem = 86 // The response code returned by optimize.
    MSK_IINF_PRESOLVE_NUM_PRIMAL_PERTURBATIONS Iinfitem = 87 // Number perturbations to thhe bounds of the primal problem.
    MSK_IINF_PURIFY_DUAL_SUCCESS Iinfitem = 88 // Is nonzero if the dual solution is purified.
    MSK_IINF_PURIFY_PRIMAL_SUCCESS Iinfitem = 89 // Is nonzero if the primal solution is purified.
    MSK_IINF_RD_NUMBARVAR Iinfitem = 90 // Number of symmetric variables read.
    MSK_IINF_RD_NUMCON Iinfitem = 91 // Number of constraints read.
    MSK_IINF_RD_NUMCONE Iinfitem = 92 // Number of conic constraints read.
    MSK_IINF_RD_NUMINTVAR Iinfitem = 93 // Number of integer-constrained variables read.
    MSK_IINF_RD_NUMQ Iinfitem = 94 // Number of nonempty Q matrices read.
    MSK_IINF_RD_NUMVAR Iinfitem = 95 // Number of variables read.
    MSK_IINF_RD_PROTYPE Iinfitem = 96 // Problem type.
    MSK_IINF_SIM_DUAL_DEG_ITER Iinfitem = 97 // The number of dual degenerate iterations.
    MSK_IINF_SIM_DUAL_HOTSTART Iinfitem = 98 // If 1 then the dual simplex algorithm is solving from an advanced basis.
    MSK_IINF_SIM_DUAL_HOTSTART_LU Iinfitem = 99 // If 1 then a valid basis factorization of full rank was located and used by the dual simplex algorithm.
    MSK_IINF_SIM_DUAL_INF_ITER Iinfitem = 100 // The number of iterations taken with dual infeasibility.
    MSK_IINF_SIM_DUAL_ITER Iinfitem = 101 // Number of dual simplex iterations during the last optimization.
    MSK_IINF_SIM_NUMCON Iinfitem = 102 // Number of constraints in the problem solved by the simplex optimizer.
    MSK_IINF_SIM_NUMVAR Iinfitem = 103 // Number of variables in the problem solved by the simplex optimizer.
    MSK_IINF_SIM_PRIMAL_DEG_ITER Iinfitem = 104 // The number of primal degenerate iterations.
    MSK_IINF_SIM_PRIMAL_HOTSTART Iinfitem = 105 // If 1 then the primal simplex algorithm is solving from an advanced basis.
    MSK_IINF_SIM_PRIMAL_HOTSTART_LU Iinfitem = 106 // If 1 then a valid basis factorization of full rank was located and used by the primal simplex algorithm.
    MSK_IINF_SIM_PRIMAL_INF_ITER Iinfitem = 107 // The number of iterations taken with primal infeasibility.
    MSK_IINF_SIM_PRIMAL_ITER Iinfitem = 108 // Number of primal simplex iterations during the last optimization.
    MSK_IINF_SIM_SOLVE_DUAL Iinfitem = 109 // Is non-zero if dual problem is solved.
    MSK_IINF_SOL_BAS_PROSTA Iinfitem = 110 // Problem status of the basic solution. Updated after each optimization.
    MSK_IINF_SOL_BAS_SOLSTA Iinfitem = 111 // Solution status of the basic solution. Updated after each optimization.
    MSK_IINF_SOL_ITG_PROSTA Iinfitem = 112 // Problem status of the integer solution. Updated after each optimization.
    MSK_IINF_SOL_ITG_SOLSTA Iinfitem = 113 // Solution status of the integer solution. Updated after each optimization.
    MSK_IINF_SOL_ITR_PROSTA Iinfitem = 114 // Problem status of the interior-point solution. Updated after each optimization.
    MSK_IINF_SOL_ITR_SOLSTA Iinfitem = 115 // Solution status of the interior-point solution. Updated after each optimization.
    MSK_IINF_STO_NUM_A_REALLOC Iinfitem = 116 // Number of times the storage for storing the linear coefficient matrix has been changed.
    MSK_IINF_END Iinfitem = 116
)
func (self Iinfitem) String() string {
  switch self {
    case MSK_IINF_ANA_PRO_NUM_CON: return "MSK_IINF_ANA_PRO_NUM_CON"
    case MSK_IINF_ANA_PRO_NUM_CON_EQ: return "MSK_IINF_ANA_PRO_NUM_CON_EQ"
    case MSK_IINF_ANA_PRO_NUM_CON_FR: return "MSK_IINF_ANA_PRO_NUM_CON_FR"
    case MSK_IINF_ANA_PRO_NUM_CON_LO: return "MSK_IINF_ANA_PRO_NUM_CON_LO"
    case MSK_IINF_ANA_PRO_NUM_CON_RA: return "MSK_IINF_ANA_PRO_NUM_CON_RA"
    case MSK_IINF_ANA_PRO_NUM_CON_UP: return "MSK_IINF_ANA_PRO_NUM_CON_UP"
    case MSK_IINF_ANA_PRO_NUM_VAR: return "MSK_IINF_ANA_PRO_NUM_VAR"
    case MSK_IINF_ANA_PRO_NUM_VAR_BIN: return "MSK_IINF_ANA_PRO_NUM_VAR_BIN"
    case MSK_IINF_ANA_PRO_NUM_VAR_CONT: return "MSK_IINF_ANA_PRO_NUM_VAR_CONT"
    case MSK_IINF_ANA_PRO_NUM_VAR_EQ: return "MSK_IINF_ANA_PRO_NUM_VAR_EQ"
    case MSK_IINF_ANA_PRO_NUM_VAR_FR: return "MSK_IINF_ANA_PRO_NUM_VAR_FR"
    case MSK_IINF_ANA_PRO_NUM_VAR_INT: return "MSK_IINF_ANA_PRO_NUM_VAR_INT"
    case MSK_IINF_ANA_PRO_NUM_VAR_LO: return "MSK_IINF_ANA_PRO_NUM_VAR_LO"
    case MSK_IINF_ANA_PRO_NUM_VAR_RA: return "MSK_IINF_ANA_PRO_NUM_VAR_RA"
    case MSK_IINF_ANA_PRO_NUM_VAR_UP: return "MSK_IINF_ANA_PRO_NUM_VAR_UP"
    case MSK_IINF_INTPNT_FACTOR_DIM_DENSE: return "MSK_IINF_INTPNT_FACTOR_DIM_DENSE"
    case MSK_IINF_INTPNT_ITER: return "MSK_IINF_INTPNT_ITER"
    case MSK_IINF_INTPNT_NUM_THREADS: return "MSK_IINF_INTPNT_NUM_THREADS"
    case MSK_IINF_INTPNT_SOLVE_DUAL: return "MSK_IINF_INTPNT_SOLVE_DUAL"
    case MSK_IINF_MIO_ABSGAP_SATISFIED: return "MSK_IINF_MIO_ABSGAP_SATISFIED"
    case MSK_IINF_MIO_CLIQUE_TABLE_SIZE: return "MSK_IINF_MIO_CLIQUE_TABLE_SIZE"
    case MSK_IINF_MIO_CONSTRUCT_SOLUTION: return "MSK_IINF_MIO_CONSTRUCT_SOLUTION"
    case MSK_IINF_MIO_INITIAL_FEASIBLE_SOLUTION: return "MSK_IINF_MIO_INITIAL_FEASIBLE_SOLUTION"
    case MSK_IINF_MIO_NODE_DEPTH: return "MSK_IINF_MIO_NODE_DEPTH"
    case MSK_IINF_MIO_NUM_ACTIVE_NODES: return "MSK_IINF_MIO_NUM_ACTIVE_NODES"
    case MSK_IINF_MIO_NUM_ACTIVE_ROOT_CUTS: return "MSK_IINF_MIO_NUM_ACTIVE_ROOT_CUTS"
    case MSK_IINF_MIO_NUM_BRANCH: return "MSK_IINF_MIO_NUM_BRANCH"
    case MSK_IINF_MIO_NUM_INT_SOLUTIONS: return "MSK_IINF_MIO_NUM_INT_SOLUTIONS"
    case MSK_IINF_MIO_NUM_RELAX: return "MSK_IINF_MIO_NUM_RELAX"
    case MSK_IINF_MIO_NUM_REPEATED_PRESOLVE: return "MSK_IINF_MIO_NUM_REPEATED_PRESOLVE"
    case MSK_IINF_MIO_NUM_RESTARTS: return "MSK_IINF_MIO_NUM_RESTARTS"
    case MSK_IINF_MIO_NUM_ROOT_CUT_ROUNDS: return "MSK_IINF_MIO_NUM_ROOT_CUT_ROUNDS"
    case MSK_IINF_MIO_NUM_SELECTED_CLIQUE_CUTS: return "MSK_IINF_MIO_NUM_SELECTED_CLIQUE_CUTS"
    case MSK_IINF_MIO_NUM_SELECTED_CMIR_CUTS: return "MSK_IINF_MIO_NUM_SELECTED_CMIR_CUTS"
    case MSK_IINF_MIO_NUM_SELECTED_GOMORY_CUTS: return "MSK_IINF_MIO_NUM_SELECTED_GOMORY_CUTS"
    case MSK_IINF_MIO_NUM_SELECTED_IMPLIED_BOUND_CUTS: return "MSK_IINF_MIO_NUM_SELECTED_IMPLIED_BOUND_CUTS"
    case MSK_IINF_MIO_NUM_SELECTED_KNAPSACK_COVER_CUTS: return "MSK_IINF_MIO_NUM_SELECTED_KNAPSACK_COVER_CUTS"
    case MSK_IINF_MIO_NUM_SELECTED_LIPRO_CUTS: return "MSK_IINF_MIO_NUM_SELECTED_LIPRO_CUTS"
    case MSK_IINF_MIO_NUM_SEPARATED_CLIQUE_CUTS: return "MSK_IINF_MIO_NUM_SEPARATED_CLIQUE_CUTS"
    case MSK_IINF_MIO_NUM_SEPARATED_CMIR_CUTS: return "MSK_IINF_MIO_NUM_SEPARATED_CMIR_CUTS"
    case MSK_IINF_MIO_NUM_SEPARATED_GOMORY_CUTS: return "MSK_IINF_MIO_NUM_SEPARATED_GOMORY_CUTS"
    case MSK_IINF_MIO_NUM_SEPARATED_IMPLIED_BOUND_CUTS: return "MSK_IINF_MIO_NUM_SEPARATED_IMPLIED_BOUND_CUTS"
    case MSK_IINF_MIO_NUM_SEPARATED_KNAPSACK_COVER_CUTS: return "MSK_IINF_MIO_NUM_SEPARATED_KNAPSACK_COVER_CUTS"
    case MSK_IINF_MIO_NUM_SEPARATED_LIPRO_CUTS: return "MSK_IINF_MIO_NUM_SEPARATED_LIPRO_CUTS"
    case MSK_IINF_MIO_NUM_SOLVED_NODES: return "MSK_IINF_MIO_NUM_SOLVED_NODES"
    case MSK_IINF_MIO_NUMBIN: return "MSK_IINF_MIO_NUMBIN"
    case MSK_IINF_MIO_NUMBINCONEVAR: return "MSK_IINF_MIO_NUMBINCONEVAR"
    case MSK_IINF_MIO_NUMCON: return "MSK_IINF_MIO_NUMCON"
    case MSK_IINF_MIO_NUMCONE: return "MSK_IINF_MIO_NUMCONE"
    case MSK_IINF_MIO_NUMCONEVAR: return "MSK_IINF_MIO_NUMCONEVAR"
    case MSK_IINF_MIO_NUMCONT: return "MSK_IINF_MIO_NUMCONT"
    case MSK_IINF_MIO_NUMCONTCONEVAR: return "MSK_IINF_MIO_NUMCONTCONEVAR"
    case MSK_IINF_MIO_NUMDEXPCONES: return "MSK_IINF_MIO_NUMDEXPCONES"
    case MSK_IINF_MIO_NUMDJC: return "MSK_IINF_MIO_NUMDJC"
    case MSK_IINF_MIO_NUMDPOWCONES: return "MSK_IINF_MIO_NUMDPOWCONES"
    case MSK_IINF_MIO_NUMINT: return "MSK_IINF_MIO_NUMINT"
    case MSK_IINF_MIO_NUMINTCONEVAR: return "MSK_IINF_MIO_NUMINTCONEVAR"
    case MSK_IINF_MIO_NUMPEXPCONES: return "MSK_IINF_MIO_NUMPEXPCONES"
    case MSK_IINF_MIO_NUMPPOWCONES: return "MSK_IINF_MIO_NUMPPOWCONES"
    case MSK_IINF_MIO_NUMQCONES: return "MSK_IINF_MIO_NUMQCONES"
    case MSK_IINF_MIO_NUMRQCONES: return "MSK_IINF_MIO_NUMRQCONES"
    case MSK_IINF_MIO_NUMVAR: return "MSK_IINF_MIO_NUMVAR"
    case MSK_IINF_MIO_OBJ_BOUND_DEFINED: return "MSK_IINF_MIO_OBJ_BOUND_DEFINED"
    case MSK_IINF_MIO_PRESOLVED_NUMBIN: return "MSK_IINF_MIO_PRESOLVED_NUMBIN"
    case MSK_IINF_MIO_PRESOLVED_NUMBINCONEVAR: return "MSK_IINF_MIO_PRESOLVED_NUMBINCONEVAR"
    case MSK_IINF_MIO_PRESOLVED_NUMCON: return "MSK_IINF_MIO_PRESOLVED_NUMCON"
    case MSK_IINF_MIO_PRESOLVED_NUMCONE: return "MSK_IINF_MIO_PRESOLVED_NUMCONE"
    case MSK_IINF_MIO_PRESOLVED_NUMCONEVAR: return "MSK_IINF_MIO_PRESOLVED_NUMCONEVAR"
    case MSK_IINF_MIO_PRESOLVED_NUMCONT: return "MSK_IINF_MIO_PRESOLVED_NUMCONT"
    case MSK_IINF_MIO_PRESOLVED_NUMCONTCONEVAR: return "MSK_IINF_MIO_PRESOLVED_NUMCONTCONEVAR"
    case MSK_IINF_MIO_PRESOLVED_NUMDEXPCONES: return "MSK_IINF_MIO_PRESOLVED_NUMDEXPCONES"
    case MSK_IINF_MIO_PRESOLVED_NUMDJC: return "MSK_IINF_MIO_PRESOLVED_NUMDJC"
    case MSK_IINF_MIO_PRESOLVED_NUMDPOWCONES: return "MSK_IINF_MIO_PRESOLVED_NUMDPOWCONES"
    case MSK_IINF_MIO_PRESOLVED_NUMINT: return "MSK_IINF_MIO_PRESOLVED_NUMINT"
    case MSK_IINF_MIO_PRESOLVED_NUMINTCONEVAR: return "MSK_IINF_MIO_PRESOLVED_NUMINTCONEVAR"
    case MSK_IINF_MIO_PRESOLVED_NUMPEXPCONES: return "MSK_IINF_MIO_PRESOLVED_NUMPEXPCONES"
    case MSK_IINF_MIO_PRESOLVED_NUMPPOWCONES: return "MSK_IINF_MIO_PRESOLVED_NUMPPOWCONES"
    case MSK_IINF_MIO_PRESOLVED_NUMQCONES: return "MSK_IINF_MIO_PRESOLVED_NUMQCONES"
    case MSK_IINF_MIO_PRESOLVED_NUMRQCONES: return "MSK_IINF_MIO_PRESOLVED_NUMRQCONES"
    case MSK_IINF_MIO_PRESOLVED_NUMVAR: return "MSK_IINF_MIO_PRESOLVED_NUMVAR"
    case MSK_IINF_MIO_RELGAP_SATISFIED: return "MSK_IINF_MIO_RELGAP_SATISFIED"
    case MSK_IINF_MIO_TOTAL_NUM_SELECTED_CUTS: return "MSK_IINF_MIO_TOTAL_NUM_SELECTED_CUTS"
    case MSK_IINF_MIO_TOTAL_NUM_SEPARATED_CUTS: return "MSK_IINF_MIO_TOTAL_NUM_SEPARATED_CUTS"
    case MSK_IINF_MIO_USER_OBJ_CUT: return "MSK_IINF_MIO_USER_OBJ_CUT"
    case MSK_IINF_OPT_NUMCON: return "MSK_IINF_OPT_NUMCON"
    case MSK_IINF_OPT_NUMVAR: return "MSK_IINF_OPT_NUMVAR"
    case MSK_IINF_OPTIMIZE_RESPONSE: return "MSK_IINF_OPTIMIZE_RESPONSE"
    case MSK_IINF_PRESOLVE_NUM_PRIMAL_PERTURBATIONS: return "MSK_IINF_PRESOLVE_NUM_PRIMAL_PERTURBATIONS"
    case MSK_IINF_PURIFY_DUAL_SUCCESS: return "MSK_IINF_PURIFY_DUAL_SUCCESS"
    case MSK_IINF_PURIFY_PRIMAL_SUCCESS: return "MSK_IINF_PURIFY_PRIMAL_SUCCESS"
    case MSK_IINF_RD_NUMBARVAR: return "MSK_IINF_RD_NUMBARVAR"
    case MSK_IINF_RD_NUMCON: return "MSK_IINF_RD_NUMCON"
    case MSK_IINF_RD_NUMCONE: return "MSK_IINF_RD_NUMCONE"
    case MSK_IINF_RD_NUMINTVAR: return "MSK_IINF_RD_NUMINTVAR"
    case MSK_IINF_RD_NUMQ: return "MSK_IINF_RD_NUMQ"
    case MSK_IINF_RD_NUMVAR: return "MSK_IINF_RD_NUMVAR"
    case MSK_IINF_RD_PROTYPE: return "MSK_IINF_RD_PROTYPE"
    case MSK_IINF_SIM_DUAL_DEG_ITER: return "MSK_IINF_SIM_DUAL_DEG_ITER"
    case MSK_IINF_SIM_DUAL_HOTSTART: return "MSK_IINF_SIM_DUAL_HOTSTART"
    case MSK_IINF_SIM_DUAL_HOTSTART_LU: return "MSK_IINF_SIM_DUAL_HOTSTART_LU"
    case MSK_IINF_SIM_DUAL_INF_ITER: return "MSK_IINF_SIM_DUAL_INF_ITER"
    case MSK_IINF_SIM_DUAL_ITER: return "MSK_IINF_SIM_DUAL_ITER"
    case MSK_IINF_SIM_NUMCON: return "MSK_IINF_SIM_NUMCON"
    case MSK_IINF_SIM_NUMVAR: return "MSK_IINF_SIM_NUMVAR"
    case MSK_IINF_SIM_PRIMAL_DEG_ITER: return "MSK_IINF_SIM_PRIMAL_DEG_ITER"
    case MSK_IINF_SIM_PRIMAL_HOTSTART: return "MSK_IINF_SIM_PRIMAL_HOTSTART"
    case MSK_IINF_SIM_PRIMAL_HOTSTART_LU: return "MSK_IINF_SIM_PRIMAL_HOTSTART_LU"
    case MSK_IINF_SIM_PRIMAL_INF_ITER: return "MSK_IINF_SIM_PRIMAL_INF_ITER"
    case MSK_IINF_SIM_PRIMAL_ITER: return "MSK_IINF_SIM_PRIMAL_ITER"
    case MSK_IINF_SIM_SOLVE_DUAL: return "MSK_IINF_SIM_SOLVE_DUAL"
    case MSK_IINF_SOL_BAS_PROSTA: return "MSK_IINF_SOL_BAS_PROSTA"
    case MSK_IINF_SOL_BAS_SOLSTA: return "MSK_IINF_SOL_BAS_SOLSTA"
    case MSK_IINF_SOL_ITG_PROSTA: return "MSK_IINF_SOL_ITG_PROSTA"
    case MSK_IINF_SOL_ITG_SOLSTA: return "MSK_IINF_SOL_ITG_SOLSTA"
    case MSK_IINF_SOL_ITR_PROSTA: return "MSK_IINF_SOL_ITR_PROSTA"
    case MSK_IINF_SOL_ITR_SOLSTA: return "MSK_IINF_SOL_ITR_SOLSTA"
    case MSK_IINF_STO_NUM_A_REALLOC: return "MSK_IINF_STO_NUM_A_REALLOC"
    default: return "<unknown>"
  }
} // Iinfitem.String()

type Inftype int32
// Information item types
const (
    MSK_INF_DOU_TYPE Inftype = 0 // Is a double information type.
    MSK_INF_INT_TYPE Inftype = 1 // Is an integer.
    MSK_INF_LINT_TYPE Inftype = 2 // Is a long integer.
    MSK_INF_END Inftype = 2
)
func (self Inftype) String() string {
  switch self {
    case MSK_INF_DOU_TYPE: return "MSK_INF_DOU_TYPE"
    case MSK_INF_INT_TYPE: return "MSK_INF_INT_TYPE"
    case MSK_INF_LINT_TYPE: return "MSK_INF_LINT_TYPE"
    default: return "<unknown>"
  }
} // Inftype.String()

type Iomode int32
// Input/output modes
const (
    MSK_IOMODE_READ Iomode = 0 // The file is read-only.
    MSK_IOMODE_WRITE Iomode = 1 // The file is write-only. If the file exists then it is truncated when it is opened. Otherwise it is created when it is opened.
    MSK_IOMODE_READWRITE Iomode = 2 // The file is to read and write.
    MSK_IOMODE_END Iomode = 2
)
func (self Iomode) String() string {
  switch self {
    case MSK_IOMODE_READ: return "MSK_IOMODE_READ"
    case MSK_IOMODE_WRITE: return "MSK_IOMODE_WRITE"
    case MSK_IOMODE_READWRITE: return "MSK_IOMODE_READWRITE"
    default: return "<unknown>"
  }
} // Iomode.String()

type Iparam int32
// Integer parameters
const (
    MSK_IPAR_ANA_SOL_BASIS Iparam = 0 // Controls whether the basis matrix is analyzed in solution analyzer.
    MSK_IPAR_ANA_SOL_PRINT_VIOLATED Iparam = 1 // Controls whether a list of violated constraints is printed.
    MSK_IPAR_AUTO_SORT_A_BEFORE_OPT Iparam = 2 // Controls whether the elements in each column of A are sorted before an optimization is performed.
    MSK_IPAR_AUTO_UPDATE_SOL_INFO Iparam = 3 // Controls whether the solution information items are automatically updated after an optimization is performed.
    MSK_IPAR_BASIS_SOLVE_USE_PLUS_ONE Iparam = 4 // Controls the sign of the columns in the basis matrix corresponding to slack variables.
    MSK_IPAR_BI_CLEAN_OPTIMIZER Iparam = 5 // Controls which simplex optimizer is used in the clean-up phase.
    MSK_IPAR_BI_IGNORE_MAX_ITER Iparam = 6 // Turns on basis identification in case the interior-point optimizer is terminated due to maximum number of iterations.
    MSK_IPAR_BI_IGNORE_NUM_ERROR Iparam = 7 // Turns on basis identification in case the interior-point optimizer is terminated due to a numerical problem.
    MSK_IPAR_BI_MAX_ITERATIONS Iparam = 8 // Maximum number of iterations after basis identification.
    MSK_IPAR_CACHE_LICENSE Iparam = 9 // Control license caching.
    MSK_IPAR_COMPRESS_STATFILE Iparam = 10 // Control compression of stat files.
    MSK_IPAR_INFEAS_GENERIC_NAMES Iparam = 11 // Controls the contents of the infeasibility report.
    MSK_IPAR_INFEAS_PREFER_PRIMAL Iparam = 12 // Controls which certificate is used if both primal- and dual- certificate of infeasibility is available.
    MSK_IPAR_INFEAS_REPORT_AUTO Iparam = 13 // Turns the feasibility report on or off.
    MSK_IPAR_INFEAS_REPORT_LEVEL Iparam = 14 // Controls the contents of the infeasibility report.
    MSK_IPAR_INTPNT_BASIS Iparam = 15 // Controls whether basis identification is performed.
    MSK_IPAR_INTPNT_DIFF_STEP Iparam = 16 // Controls whether different step sizes are allowed in the primal and dual space.
    MSK_IPAR_INTPNT_HOTSTART Iparam = 17 // Currently not in use.
    MSK_IPAR_INTPNT_MAX_ITERATIONS Iparam = 18 // Controls the maximum number of iterations allowed in the interior-point optimizer.
    MSK_IPAR_INTPNT_MAX_NUM_COR Iparam = 19 // Maximum number of correction steps.
    MSK_IPAR_INTPNT_MAX_NUM_REFINEMENT_STEPS Iparam = 20 // Maximum number of steps to be used by the iterative search direction refinement.
    MSK_IPAR_INTPNT_OFF_COL_TRH Iparam = 21 // Controls the aggressiveness of the offending column detection.
    MSK_IPAR_INTPNT_ORDER_GP_NUM_SEEDS Iparam = 22 // This parameter controls the number of random seeds tried.
    MSK_IPAR_INTPNT_ORDER_METHOD Iparam = 23 // Controls the ordering strategy.
    MSK_IPAR_INTPNT_PURIFY Iparam = 24 // Currently not in use.
    MSK_IPAR_INTPNT_REGULARIZATION_USE Iparam = 25 // Controls whether regularization is allowed.
    MSK_IPAR_INTPNT_SCALING Iparam = 26 // Controls how the problem is scaled before the interior-point optimizer is used.
    MSK_IPAR_INTPNT_SOLVE_FORM Iparam = 27 // Controls whether the primal or the dual problem is solved.
    MSK_IPAR_INTPNT_STARTING_POINT Iparam = 28 // Starting point used by the interior-point optimizer.
    MSK_IPAR_LICENSE_DEBUG Iparam = 29 // Controls the license manager client debugging behavior.
    MSK_IPAR_LICENSE_PAUSE_TIME Iparam = 30 // Controls license manager client behavior.
    MSK_IPAR_LICENSE_SUPPRESS_EXPIRE_WRNS Iparam = 31 // Controls license manager client behavior.
    MSK_IPAR_LICENSE_TRH_EXPIRY_WRN Iparam = 32 // Controls when expiry warnings are issued.
    MSK_IPAR_LICENSE_WAIT Iparam = 33 // Controls if MOSEK should queue for a license if none is available.
    MSK_IPAR_LOG Iparam = 34 // Controls the amount of log information.
    MSK_IPAR_LOG_ANA_PRO Iparam = 35 // Controls amount of output from the problem analyzer.
    MSK_IPAR_LOG_BI Iparam = 36 // Controls the amount of output printed by the basis identification procedure. A higher level implies that more information is logged.
    MSK_IPAR_LOG_BI_FREQ Iparam = 37 // Controls the logging frequency.
    MSK_IPAR_LOG_CUT_SECOND_OPT Iparam = 38 // Controls the reduction in the log levels for the second and any subsequent optimizations.
    MSK_IPAR_LOG_EXPAND Iparam = 39 // Controls the amount of logging when a data item such as the maximum number constrains is expanded.
    MSK_IPAR_LOG_FEAS_REPAIR Iparam = 40 // Controls the amount of output printed when performing feasibility repair. A value higher than one means extensive logging.
    MSK_IPAR_LOG_FILE Iparam = 41 // If turned on, then some log info is printed when a file is written or read.
    MSK_IPAR_LOG_INCLUDE_SUMMARY Iparam = 42 // Controls whether solution summary should be printed by the optimizer.
    MSK_IPAR_LOG_INFEAS_ANA Iparam = 43 // Controls log level for the infeasibility analyzer.
    MSK_IPAR_LOG_INTPNT Iparam = 44 // Controls the amount of log information from the interior-point optimizers.
    MSK_IPAR_LOG_LOCAL_INFO Iparam = 45 // Control whether local identifying information is printed to the log.
    MSK_IPAR_LOG_MIO Iparam = 46 // Controls the amount of log information from the mixed-integer optimizers.
    MSK_IPAR_LOG_MIO_FREQ Iparam = 47 // The mixed-integer optimizer logging frequency.
    MSK_IPAR_LOG_ORDER Iparam = 48 // If turned on, then factor lines are added to the log.
    MSK_IPAR_LOG_PRESOLVE Iparam = 49 // Controls amount of output printed by the presolve procedure. A higher level implies that more information is logged.
    MSK_IPAR_LOG_RESPONSE Iparam = 50 // Controls amount of output printed when response codes are reported. A higher level implies that more information is logged.
    MSK_IPAR_LOG_SENSITIVITY Iparam = 51 // Control logging in sensitivity analyzer.
    MSK_IPAR_LOG_SENSITIVITY_OPT Iparam = 52 // Control logging in sensitivity analyzer.
    MSK_IPAR_LOG_SIM Iparam = 53 // Controls the amount of log information from the simplex optimizers.
    MSK_IPAR_LOG_SIM_FREQ Iparam = 54 // Controls simplex logging frequency.
    MSK_IPAR_LOG_SIM_MINOR Iparam = 55 // Currently not in use.
    MSK_IPAR_LOG_STORAGE Iparam = 56 // Controls the memory related log information.
    MSK_IPAR_MAX_NUM_WARNINGS Iparam = 57 // Each warning is shown a limited number of times controlled by this parameter. A negative value is identical to infinite number of times.
    MSK_IPAR_MIO_BRANCH_DIR Iparam = 58 // Controls whether the mixed-integer optimizer is branching up or down by default.
    MSK_IPAR_MIO_CONIC_OUTER_APPROXIMATION Iparam = 59 // Toggles outer approximation for conic problems.
    MSK_IPAR_MIO_CONSTRUCT_SOL Iparam = 60 // Controls if an initial mixed integer solution should be constructed from the values of the integer variables.
    MSK_IPAR_MIO_CUT_CLIQUE Iparam = 61 // Controls whether clique cuts should be generated.
    MSK_IPAR_MIO_CUT_CMIR Iparam = 62 // Controls whether mixed integer rounding cuts should be generated.
    MSK_IPAR_MIO_CUT_GMI Iparam = 63 // Controls whether GMI cuts should be generated.
    MSK_IPAR_MIO_CUT_IMPLIED_BOUND Iparam = 64 // Controls whether implied bound cuts should be generated.
    MSK_IPAR_MIO_CUT_KNAPSACK_COVER Iparam = 65 // Controls whether knapsack cover cuts should be generated.
    MSK_IPAR_MIO_CUT_LIPRO Iparam = 66 // Controls whether lift-and-project cuts should be generated.
    MSK_IPAR_MIO_CUT_SELECTION_LEVEL Iparam = 67 // Controls how aggressively generated cuts are selected to be included in the relaxation.
    MSK_IPAR_MIO_DATA_PERMUTATION_METHOD Iparam = 68 // Controls what problem data permutation method is appplied to mixed-integer problems.
    MSK_IPAR_MIO_DUAL_RAY_ANALYSIS_LEVEL Iparam = 69 // Controls the amount of dual ray analysis employed by the mixed-integer optimizer in presolve.
    MSK_IPAR_MIO_FEASPUMP_LEVEL Iparam = 70 // Controls the way the Feasibility Pump heuristic is employed by the mixed-integer optimizer.
    MSK_IPAR_MIO_HEURISTIC_LEVEL Iparam = 71 // Controls the heuristic employed by the mixed-integer optimizer to locate an initial integer feasible solution.
    MSK_IPAR_MIO_MAX_NUM_BRANCHES Iparam = 72 // Maximum number of branches allowed during the branch and bound search.
    MSK_IPAR_MIO_MAX_NUM_RELAXS Iparam = 73 // Maximum number of relaxations in branch and bound search.
    MSK_IPAR_MIO_MAX_NUM_RESTARTS Iparam = 74 // Maximum number of restarts allowed during the branch and bound search.
    MSK_IPAR_MIO_MAX_NUM_ROOT_CUT_ROUNDS Iparam = 75 // Maximum number of cut separation rounds at the root node.
    MSK_IPAR_MIO_MAX_NUM_SOLUTIONS Iparam = 76 // Controls how many feasible solutions the mixed-integer optimizer investigates.
    MSK_IPAR_MIO_MEMORY_EMPHASIS_LEVEL Iparam = 77 // Controls how much emphasis is put on reducing memory usage.
    MSK_IPAR_MIO_MIN_REL Iparam = 78 // Number of times a variable must have been branched on for its pseudocost to be cosidered reliable.
    MSK_IPAR_MIO_MODE Iparam = 79 // Turns on/off the mixed-integer mode.
    MSK_IPAR_MIO_NODE_OPTIMIZER Iparam = 80 // Controls which optimizer is employed at the non-root nodes in the mixed-integer optimizer.
    MSK_IPAR_MIO_NODE_SELECTION Iparam = 81 // Controls the node selection strategy employed by the mixed-integer optimizer.
    MSK_IPAR_MIO_NUMERICAL_EMPHASIS_LEVEL Iparam = 82 // Controls how much emphasis is put on reducing numerical problems
    MSK_IPAR_MIO_PERSPECTIVE_REFORMULATE Iparam = 83 // Enables or disables perspective reformulation in presolve.
    MSK_IPAR_MIO_PRESOLVE_AGGREGATOR_USE Iparam = 84 // Controls if the aggregator should be used.
    MSK_IPAR_MIO_PROBING_LEVEL Iparam = 85 // Controls the amount of probing employed by the mixed-integer optimizer in presolve.
    MSK_IPAR_MIO_PROPAGATE_OBJECTIVE_CONSTRAINT Iparam = 86 // Use objective domain propagation.
    MSK_IPAR_MIO_QCQO_REFORMULATION_METHOD Iparam = 87 // Controls what reformulation method is applied to mixed-integer quadratic problems.
    MSK_IPAR_MIO_RINS_MAX_NODES Iparam = 88 // Maximum number of nodes in each call to RINS.
    MSK_IPAR_MIO_ROOT_OPTIMIZER Iparam = 89 // Controls which optimizer is employed at the root node in the mixed-integer optimizer.
    MSK_IPAR_MIO_ROOT_REPEAT_PRESOLVE_LEVEL Iparam = 90 // Controls whether presolve can be repeated at root node.
    MSK_IPAR_MIO_SEED Iparam = 91 // Sets the random seed used for randomization in the mixed integer optimizer.
    MSK_IPAR_MIO_SYMMETRY_LEVEL Iparam = 92 // Controls the amount of symmetry detection and handling employed by the mixed-integer optimizer in presolve.
    MSK_IPAR_MIO_VAR_SELECTION Iparam = 93 // Controls the variable selection strategy employed by the mixed-integer optimizer.
    MSK_IPAR_MIO_VB_DETECTION_LEVEL Iparam = 94 // Controls how much effort is put into detecting variable bounds.
    MSK_IPAR_MT_SPINCOUNT Iparam = 95 // Set the number of iterations to spin before sleeping.
    MSK_IPAR_NG Iparam = 96 // Not in use
    MSK_IPAR_NUM_THREADS Iparam = 97 // The number of threads employed by the optimizer.
    MSK_IPAR_OPF_WRITE_HEADER Iparam = 98 // Write a text header with date and MOSEK version in an OPF file.
    MSK_IPAR_OPF_WRITE_HINTS Iparam = 99 // Write a hint section with problem dimensions in the beginning of an OPF file.
    MSK_IPAR_OPF_WRITE_LINE_LENGTH Iparam = 100 // Aim to keep lines in OPF files not much longer than this.
    MSK_IPAR_OPF_WRITE_PARAMETERS Iparam = 101 // Write a parameter section in an OPF file.
    MSK_IPAR_OPF_WRITE_PROBLEM Iparam = 102 // Write objective, constraints, bounds etc. to an OPF file.
    MSK_IPAR_OPF_WRITE_SOL_BAS Iparam = 103 // Controls what is written to the OPF files.
    MSK_IPAR_OPF_WRITE_SOL_ITG Iparam = 104 // Controls what is written to the OPF files.
    MSK_IPAR_OPF_WRITE_SOL_ITR Iparam = 105 // Controls what is written to the OPF files.
    MSK_IPAR_OPF_WRITE_SOLUTIONS Iparam = 106 // Enable inclusion of solutions in the OPF files.
    MSK_IPAR_OPTIMIZER Iparam = 107 // Controls which optimizer is used to optimize the task.
    MSK_IPAR_PARAM_READ_CASE_NAME Iparam = 108 // If turned on, then names in the parameter file are case sensitive.
    MSK_IPAR_PARAM_READ_IGN_ERROR Iparam = 109 // If turned on, then errors in parameter settings is ignored.
    MSK_IPAR_PRESOLVE_ELIMINATOR_MAX_FILL Iparam = 110 // Maximum amount of fill-in created in one pivot during the elimination phase.
    MSK_IPAR_PRESOLVE_ELIMINATOR_MAX_NUM_TRIES Iparam = 111 // Control the maximum number of times the eliminator is tried.
    MSK_IPAR_PRESOLVE_LEVEL Iparam = 112 // Currently not used.
    MSK_IPAR_PRESOLVE_LINDEP_ABS_WORK_TRH Iparam = 113 // Controls linear dependency check in presolve.
    MSK_IPAR_PRESOLVE_LINDEP_NEW Iparam = 114 // Controls whether whether a new experimental linear dependency checker is employed.
    MSK_IPAR_PRESOLVE_LINDEP_REL_WORK_TRH Iparam = 115 // Controls linear dependency check in presolve.
    MSK_IPAR_PRESOLVE_LINDEP_USE Iparam = 116 // Controls whether the linear constraints are checked for linear dependencies.
    MSK_IPAR_PRESOLVE_MAX_NUM_PASS Iparam = 117 // Control the maximum number of times presolve passes over the problem.
    MSK_IPAR_PRESOLVE_MAX_NUM_REDUCTIONS Iparam = 118 // Controls the maximum number of reductions performed by the presolve.
    MSK_IPAR_PRESOLVE_USE Iparam = 119 // Controls whether the presolve is applied to a problem before it is optimized.
    MSK_IPAR_PRIMAL_REPAIR_OPTIMIZER Iparam = 120 // Controls which optimizer that is used to find the optimal repair.
    MSK_IPAR_PTF_WRITE_PARAMETERS Iparam = 121 // Controls whether parameters section is written in PTF files.
    MSK_IPAR_PTF_WRITE_SOLUTIONS Iparam = 122 // Controls whether solution section is written in PTF files.
    MSK_IPAR_PTF_WRITE_TRANSFORM Iparam = 123 // Controls if simple transformation are done when writing PTF files.
    MSK_IPAR_READ_DEBUG Iparam = 124 // Turns on additional debugging information when reading files.
    MSK_IPAR_READ_KEEP_FREE_CON Iparam = 125 // Controls whether the free constraints are included in the problem.
    MSK_IPAR_READ_MPS_FORMAT Iparam = 126 // Controls how strictly the MPS file reader interprets the MPS format.
    MSK_IPAR_READ_MPS_WIDTH Iparam = 127 // Controls the maximal number of characters allowed in one line of the MPS file.
    MSK_IPAR_READ_TASK_IGNORE_PARAM Iparam = 128 // Controls what information is used from the task files.
    MSK_IPAR_REMOTE_USE_COMPRESSION Iparam = 129 // Use compression when sending data to an optimization server
    MSK_IPAR_REMOVE_UNUSED_SOLUTIONS Iparam = 130 // Removes unused solutions before the optimization is performed.
    MSK_IPAR_SENSITIVITY_ALL Iparam = 131 // Controls sensitivity report behavior.
    MSK_IPAR_SENSITIVITY_OPTIMIZER Iparam = 132 // Controls which optimizer is used for optimal partition sensitivity analysis.
    MSK_IPAR_SENSITIVITY_TYPE Iparam = 133 // Controls which type of sensitivity analysis is to be performed.
    MSK_IPAR_SIM_BASIS_FACTOR_USE Iparam = 134 // Controls whether an LU factorization of the basis is used in a hot-start.
    MSK_IPAR_SIM_DEGEN Iparam = 135 // Controls how aggressively degeneration is handled.
    MSK_IPAR_SIM_DETECT_PWL Iparam = 136 // Not in use.
    MSK_IPAR_SIM_DUAL_CRASH Iparam = 137 // Controls whether crashing is performed in the dual simplex optimizer.
    MSK_IPAR_SIM_DUAL_PHASEONE_METHOD Iparam = 138 // An experimental feature.
    MSK_IPAR_SIM_DUAL_RESTRICT_SELECTION Iparam = 139 // Controls how aggressively restricted selection is used.
    MSK_IPAR_SIM_DUAL_SELECTION Iparam = 140 // Controls the dual simplex strategy.
    MSK_IPAR_SIM_EXPLOIT_DUPVEC Iparam = 141 // Controls if the simplex optimizers are allowed to exploit duplicated columns.
    MSK_IPAR_SIM_HOTSTART Iparam = 142 // Controls the type of hot-start that the simplex optimizer perform.
    MSK_IPAR_SIM_HOTSTART_LU Iparam = 143 // Determines if the simplex optimizer should exploit the initial factorization.
    MSK_IPAR_SIM_MAX_ITERATIONS Iparam = 144 // Maximum number of iterations that can be used by a simplex optimizer.
    MSK_IPAR_SIM_MAX_NUM_SETBACKS Iparam = 145 // Controls how many set-backs that are allowed within a simplex optimizer.
    MSK_IPAR_SIM_NON_SINGULAR Iparam = 146 // Controls if the simplex optimizer ensures a non-singular basis, if possible.
    MSK_IPAR_SIM_PRIMAL_CRASH Iparam = 147 // Controls the simplex crash.
    MSK_IPAR_SIM_PRIMAL_PHASEONE_METHOD Iparam = 148 // An experimental feature.
    MSK_IPAR_SIM_PRIMAL_RESTRICT_SELECTION Iparam = 149 // Controls how aggressively restricted selection is used.
    MSK_IPAR_SIM_PRIMAL_SELECTION Iparam = 150 // Controls the primal simplex strategy.
    MSK_IPAR_SIM_REFACTOR_FREQ Iparam = 151 // Controls the basis refactoring frequency.
    MSK_IPAR_SIM_REFORMULATION Iparam = 152 // Controls if the simplex optimizers are allowed to reformulate the problem.
    MSK_IPAR_SIM_SAVE_LU Iparam = 153 // Controls if the LU factorization stored should be replaced with the LU factorization corresponding to the initial basis.
    MSK_IPAR_SIM_SCALING Iparam = 154 // Controls how much effort is used in scaling the problem before a simplex optimizer is used.
    MSK_IPAR_SIM_SCALING_METHOD Iparam = 155 // Controls how the problem is scaled before a simplex optimizer is used.
    MSK_IPAR_SIM_SEED Iparam = 156 // Sets the random seed used for randomization in the simplex optimizers.
    MSK_IPAR_SIM_SOLVE_FORM Iparam = 157 // Controls whether the primal or the dual problem is solved by the primal-/dual-simplex optimizer.
    MSK_IPAR_SIM_STABILITY_PRIORITY Iparam = 158 // Controls how high priority the numerical stability should be given.
    MSK_IPAR_SIM_SWITCH_OPTIMIZER Iparam = 159 // Controls the simplex behavior.
    MSK_IPAR_SOL_FILTER_KEEP_BASIC Iparam = 160 // Control the contents of the solution files.
    MSK_IPAR_SOL_FILTER_KEEP_RANGED Iparam = 161 // Control the contents of the solution files.
    MSK_IPAR_SOL_READ_NAME_WIDTH Iparam = 162 // Controls the input solution file format.
    MSK_IPAR_SOL_READ_WIDTH Iparam = 163 // Controls the input solution file format.
    MSK_IPAR_SOLUTION_CALLBACK Iparam = 164 // Indicates whether solution callbacks will be performed during the optimization.
    MSK_IPAR_TIMING_LEVEL Iparam = 165 // Controls the amount of timing performed inside MOSEK.
    MSK_IPAR_WRITE_BAS_CONSTRAINTS Iparam = 166 // Controls the basic solution file format.
    MSK_IPAR_WRITE_BAS_HEAD Iparam = 167 // Controls the basic solution file format.
    MSK_IPAR_WRITE_BAS_VARIABLES Iparam = 168 // Controls the basic solution file format.
    MSK_IPAR_WRITE_COMPRESSION Iparam = 169 // Controls output file compression.
    MSK_IPAR_WRITE_DATA_PARAM Iparam = 170 // Controls output file data.
    MSK_IPAR_WRITE_FREE_CON Iparam = 171 // Controls the output file data.
    MSK_IPAR_WRITE_GENERIC_NAMES Iparam = 172 // Controls the output file data.
    MSK_IPAR_WRITE_GENERIC_NAMES_IO Iparam = 173 // Index origin used in  generic names.
    MSK_IPAR_WRITE_IGNORE_INCOMPATIBLE_ITEMS Iparam = 174 // Controls if the writer ignores incompatible problem items when writing files.
    MSK_IPAR_WRITE_INT_CONSTRAINTS Iparam = 175 // Controls the integer solution file format.
    MSK_IPAR_WRITE_INT_HEAD Iparam = 176 // Controls the integer solution file format.
    MSK_IPAR_WRITE_INT_VARIABLES Iparam = 177 // Controls the integer solution file format.
    MSK_IPAR_WRITE_JSON_INDENTATION Iparam = 178 // When set, the JSON task and solution files are written with indentation for better readability.
    MSK_IPAR_WRITE_LP_FULL_OBJ Iparam = 179 // Write full linear objective
    MSK_IPAR_WRITE_LP_LINE_WIDTH Iparam = 180 // Controls the LP output file format.
    MSK_IPAR_WRITE_MPS_FORMAT Iparam = 181 // Controls in which format the MPS is written.
    MSK_IPAR_WRITE_MPS_INT Iparam = 182 // Controls the output file data.
    MSK_IPAR_WRITE_SOL_BARVARIABLES Iparam = 183 // Controls the solution file format.
    MSK_IPAR_WRITE_SOL_CONSTRAINTS Iparam = 184 // Controls the solution file format.
    MSK_IPAR_WRITE_SOL_HEAD Iparam = 185 // Controls solution file format.
    MSK_IPAR_WRITE_SOL_IGNORE_INVALID_NAMES Iparam = 186 // Controls whether the user specified names are employed even if they are invalid names.
    MSK_IPAR_WRITE_SOL_VARIABLES Iparam = 187 // Controls the solution file format.
    MSK_IPAR_WRITE_TASK_INC_SOL Iparam = 188 // Controls whether the solutions are stored in the task file too.
    MSK_IPAR_WRITE_XML_MODE Iparam = 189 // Controls if linear coefficients should be written by row or column when writing in the XML file format.
    MSK_IPAR_END Iparam = 189
)
func (self Iparam) String() string {
  switch self {
    case MSK_IPAR_ANA_SOL_BASIS: return "MSK_IPAR_ANA_SOL_BASIS"
    case MSK_IPAR_ANA_SOL_PRINT_VIOLATED: return "MSK_IPAR_ANA_SOL_PRINT_VIOLATED"
    case MSK_IPAR_AUTO_SORT_A_BEFORE_OPT: return "MSK_IPAR_AUTO_SORT_A_BEFORE_OPT"
    case MSK_IPAR_AUTO_UPDATE_SOL_INFO: return "MSK_IPAR_AUTO_UPDATE_SOL_INFO"
    case MSK_IPAR_BASIS_SOLVE_USE_PLUS_ONE: return "MSK_IPAR_BASIS_SOLVE_USE_PLUS_ONE"
    case MSK_IPAR_BI_CLEAN_OPTIMIZER: return "MSK_IPAR_BI_CLEAN_OPTIMIZER"
    case MSK_IPAR_BI_IGNORE_MAX_ITER: return "MSK_IPAR_BI_IGNORE_MAX_ITER"
    case MSK_IPAR_BI_IGNORE_NUM_ERROR: return "MSK_IPAR_BI_IGNORE_NUM_ERROR"
    case MSK_IPAR_BI_MAX_ITERATIONS: return "MSK_IPAR_BI_MAX_ITERATIONS"
    case MSK_IPAR_CACHE_LICENSE: return "MSK_IPAR_CACHE_LICENSE"
    case MSK_IPAR_COMPRESS_STATFILE: return "MSK_IPAR_COMPRESS_STATFILE"
    case MSK_IPAR_INFEAS_GENERIC_NAMES: return "MSK_IPAR_INFEAS_GENERIC_NAMES"
    case MSK_IPAR_INFEAS_PREFER_PRIMAL: return "MSK_IPAR_INFEAS_PREFER_PRIMAL"
    case MSK_IPAR_INFEAS_REPORT_AUTO: return "MSK_IPAR_INFEAS_REPORT_AUTO"
    case MSK_IPAR_INFEAS_REPORT_LEVEL: return "MSK_IPAR_INFEAS_REPORT_LEVEL"
    case MSK_IPAR_INTPNT_BASIS: return "MSK_IPAR_INTPNT_BASIS"
    case MSK_IPAR_INTPNT_DIFF_STEP: return "MSK_IPAR_INTPNT_DIFF_STEP"
    case MSK_IPAR_INTPNT_HOTSTART: return "MSK_IPAR_INTPNT_HOTSTART"
    case MSK_IPAR_INTPNT_MAX_ITERATIONS: return "MSK_IPAR_INTPNT_MAX_ITERATIONS"
    case MSK_IPAR_INTPNT_MAX_NUM_COR: return "MSK_IPAR_INTPNT_MAX_NUM_COR"
    case MSK_IPAR_INTPNT_MAX_NUM_REFINEMENT_STEPS: return "MSK_IPAR_INTPNT_MAX_NUM_REFINEMENT_STEPS"
    case MSK_IPAR_INTPNT_OFF_COL_TRH: return "MSK_IPAR_INTPNT_OFF_COL_TRH"
    case MSK_IPAR_INTPNT_ORDER_GP_NUM_SEEDS: return "MSK_IPAR_INTPNT_ORDER_GP_NUM_SEEDS"
    case MSK_IPAR_INTPNT_ORDER_METHOD: return "MSK_IPAR_INTPNT_ORDER_METHOD"
    case MSK_IPAR_INTPNT_PURIFY: return "MSK_IPAR_INTPNT_PURIFY"
    case MSK_IPAR_INTPNT_REGULARIZATION_USE: return "MSK_IPAR_INTPNT_REGULARIZATION_USE"
    case MSK_IPAR_INTPNT_SCALING: return "MSK_IPAR_INTPNT_SCALING"
    case MSK_IPAR_INTPNT_SOLVE_FORM: return "MSK_IPAR_INTPNT_SOLVE_FORM"
    case MSK_IPAR_INTPNT_STARTING_POINT: return "MSK_IPAR_INTPNT_STARTING_POINT"
    case MSK_IPAR_LICENSE_DEBUG: return "MSK_IPAR_LICENSE_DEBUG"
    case MSK_IPAR_LICENSE_PAUSE_TIME: return "MSK_IPAR_LICENSE_PAUSE_TIME"
    case MSK_IPAR_LICENSE_SUPPRESS_EXPIRE_WRNS: return "MSK_IPAR_LICENSE_SUPPRESS_EXPIRE_WRNS"
    case MSK_IPAR_LICENSE_TRH_EXPIRY_WRN: return "MSK_IPAR_LICENSE_TRH_EXPIRY_WRN"
    case MSK_IPAR_LICENSE_WAIT: return "MSK_IPAR_LICENSE_WAIT"
    case MSK_IPAR_LOG: return "MSK_IPAR_LOG"
    case MSK_IPAR_LOG_ANA_PRO: return "MSK_IPAR_LOG_ANA_PRO"
    case MSK_IPAR_LOG_BI: return "MSK_IPAR_LOG_BI"
    case MSK_IPAR_LOG_BI_FREQ: return "MSK_IPAR_LOG_BI_FREQ"
    case MSK_IPAR_LOG_CUT_SECOND_OPT: return "MSK_IPAR_LOG_CUT_SECOND_OPT"
    case MSK_IPAR_LOG_EXPAND: return "MSK_IPAR_LOG_EXPAND"
    case MSK_IPAR_LOG_FEAS_REPAIR: return "MSK_IPAR_LOG_FEAS_REPAIR"
    case MSK_IPAR_LOG_FILE: return "MSK_IPAR_LOG_FILE"
    case MSK_IPAR_LOG_INCLUDE_SUMMARY: return "MSK_IPAR_LOG_INCLUDE_SUMMARY"
    case MSK_IPAR_LOG_INFEAS_ANA: return "MSK_IPAR_LOG_INFEAS_ANA"
    case MSK_IPAR_LOG_INTPNT: return "MSK_IPAR_LOG_INTPNT"
    case MSK_IPAR_LOG_LOCAL_INFO: return "MSK_IPAR_LOG_LOCAL_INFO"
    case MSK_IPAR_LOG_MIO: return "MSK_IPAR_LOG_MIO"
    case MSK_IPAR_LOG_MIO_FREQ: return "MSK_IPAR_LOG_MIO_FREQ"
    case MSK_IPAR_LOG_ORDER: return "MSK_IPAR_LOG_ORDER"
    case MSK_IPAR_LOG_PRESOLVE: return "MSK_IPAR_LOG_PRESOLVE"
    case MSK_IPAR_LOG_RESPONSE: return "MSK_IPAR_LOG_RESPONSE"
    case MSK_IPAR_LOG_SENSITIVITY: return "MSK_IPAR_LOG_SENSITIVITY"
    case MSK_IPAR_LOG_SENSITIVITY_OPT: return "MSK_IPAR_LOG_SENSITIVITY_OPT"
    case MSK_IPAR_LOG_SIM: return "MSK_IPAR_LOG_SIM"
    case MSK_IPAR_LOG_SIM_FREQ: return "MSK_IPAR_LOG_SIM_FREQ"
    case MSK_IPAR_LOG_SIM_MINOR: return "MSK_IPAR_LOG_SIM_MINOR"
    case MSK_IPAR_LOG_STORAGE: return "MSK_IPAR_LOG_STORAGE"
    case MSK_IPAR_MAX_NUM_WARNINGS: return "MSK_IPAR_MAX_NUM_WARNINGS"
    case MSK_IPAR_MIO_BRANCH_DIR: return "MSK_IPAR_MIO_BRANCH_DIR"
    case MSK_IPAR_MIO_CONIC_OUTER_APPROXIMATION: return "MSK_IPAR_MIO_CONIC_OUTER_APPROXIMATION"
    case MSK_IPAR_MIO_CONSTRUCT_SOL: return "MSK_IPAR_MIO_CONSTRUCT_SOL"
    case MSK_IPAR_MIO_CUT_CLIQUE: return "MSK_IPAR_MIO_CUT_CLIQUE"
    case MSK_IPAR_MIO_CUT_CMIR: return "MSK_IPAR_MIO_CUT_CMIR"
    case MSK_IPAR_MIO_CUT_GMI: return "MSK_IPAR_MIO_CUT_GMI"
    case MSK_IPAR_MIO_CUT_IMPLIED_BOUND: return "MSK_IPAR_MIO_CUT_IMPLIED_BOUND"
    case MSK_IPAR_MIO_CUT_KNAPSACK_COVER: return "MSK_IPAR_MIO_CUT_KNAPSACK_COVER"
    case MSK_IPAR_MIO_CUT_LIPRO: return "MSK_IPAR_MIO_CUT_LIPRO"
    case MSK_IPAR_MIO_CUT_SELECTION_LEVEL: return "MSK_IPAR_MIO_CUT_SELECTION_LEVEL"
    case MSK_IPAR_MIO_DATA_PERMUTATION_METHOD: return "MSK_IPAR_MIO_DATA_PERMUTATION_METHOD"
    case MSK_IPAR_MIO_DUAL_RAY_ANALYSIS_LEVEL: return "MSK_IPAR_MIO_DUAL_RAY_ANALYSIS_LEVEL"
    case MSK_IPAR_MIO_FEASPUMP_LEVEL: return "MSK_IPAR_MIO_FEASPUMP_LEVEL"
    case MSK_IPAR_MIO_HEURISTIC_LEVEL: return "MSK_IPAR_MIO_HEURISTIC_LEVEL"
    case MSK_IPAR_MIO_MAX_NUM_BRANCHES: return "MSK_IPAR_MIO_MAX_NUM_BRANCHES"
    case MSK_IPAR_MIO_MAX_NUM_RELAXS: return "MSK_IPAR_MIO_MAX_NUM_RELAXS"
    case MSK_IPAR_MIO_MAX_NUM_RESTARTS: return "MSK_IPAR_MIO_MAX_NUM_RESTARTS"
    case MSK_IPAR_MIO_MAX_NUM_ROOT_CUT_ROUNDS: return "MSK_IPAR_MIO_MAX_NUM_ROOT_CUT_ROUNDS"
    case MSK_IPAR_MIO_MAX_NUM_SOLUTIONS: return "MSK_IPAR_MIO_MAX_NUM_SOLUTIONS"
    case MSK_IPAR_MIO_MEMORY_EMPHASIS_LEVEL: return "MSK_IPAR_MIO_MEMORY_EMPHASIS_LEVEL"
    case MSK_IPAR_MIO_MIN_REL: return "MSK_IPAR_MIO_MIN_REL"
    case MSK_IPAR_MIO_MODE: return "MSK_IPAR_MIO_MODE"
    case MSK_IPAR_MIO_NODE_OPTIMIZER: return "MSK_IPAR_MIO_NODE_OPTIMIZER"
    case MSK_IPAR_MIO_NODE_SELECTION: return "MSK_IPAR_MIO_NODE_SELECTION"
    case MSK_IPAR_MIO_NUMERICAL_EMPHASIS_LEVEL: return "MSK_IPAR_MIO_NUMERICAL_EMPHASIS_LEVEL"
    case MSK_IPAR_MIO_PERSPECTIVE_REFORMULATE: return "MSK_IPAR_MIO_PERSPECTIVE_REFORMULATE"
    case MSK_IPAR_MIO_PRESOLVE_AGGREGATOR_USE: return "MSK_IPAR_MIO_PRESOLVE_AGGREGATOR_USE"
    case MSK_IPAR_MIO_PROBING_LEVEL: return "MSK_IPAR_MIO_PROBING_LEVEL"
    case MSK_IPAR_MIO_PROPAGATE_OBJECTIVE_CONSTRAINT: return "MSK_IPAR_MIO_PROPAGATE_OBJECTIVE_CONSTRAINT"
    case MSK_IPAR_MIO_QCQO_REFORMULATION_METHOD: return "MSK_IPAR_MIO_QCQO_REFORMULATION_METHOD"
    case MSK_IPAR_MIO_RINS_MAX_NODES: return "MSK_IPAR_MIO_RINS_MAX_NODES"
    case MSK_IPAR_MIO_ROOT_OPTIMIZER: return "MSK_IPAR_MIO_ROOT_OPTIMIZER"
    case MSK_IPAR_MIO_ROOT_REPEAT_PRESOLVE_LEVEL: return "MSK_IPAR_MIO_ROOT_REPEAT_PRESOLVE_LEVEL"
    case MSK_IPAR_MIO_SEED: return "MSK_IPAR_MIO_SEED"
    case MSK_IPAR_MIO_SYMMETRY_LEVEL: return "MSK_IPAR_MIO_SYMMETRY_LEVEL"
    case MSK_IPAR_MIO_VAR_SELECTION: return "MSK_IPAR_MIO_VAR_SELECTION"
    case MSK_IPAR_MIO_VB_DETECTION_LEVEL: return "MSK_IPAR_MIO_VB_DETECTION_LEVEL"
    case MSK_IPAR_MT_SPINCOUNT: return "MSK_IPAR_MT_SPINCOUNT"
    case MSK_IPAR_NG: return "MSK_IPAR_NG"
    case MSK_IPAR_NUM_THREADS: return "MSK_IPAR_NUM_THREADS"
    case MSK_IPAR_OPF_WRITE_HEADER: return "MSK_IPAR_OPF_WRITE_HEADER"
    case MSK_IPAR_OPF_WRITE_HINTS: return "MSK_IPAR_OPF_WRITE_HINTS"
    case MSK_IPAR_OPF_WRITE_LINE_LENGTH: return "MSK_IPAR_OPF_WRITE_LINE_LENGTH"
    case MSK_IPAR_OPF_WRITE_PARAMETERS: return "MSK_IPAR_OPF_WRITE_PARAMETERS"
    case MSK_IPAR_OPF_WRITE_PROBLEM: return "MSK_IPAR_OPF_WRITE_PROBLEM"
    case MSK_IPAR_OPF_WRITE_SOL_BAS: return "MSK_IPAR_OPF_WRITE_SOL_BAS"
    case MSK_IPAR_OPF_WRITE_SOL_ITG: return "MSK_IPAR_OPF_WRITE_SOL_ITG"
    case MSK_IPAR_OPF_WRITE_SOL_ITR: return "MSK_IPAR_OPF_WRITE_SOL_ITR"
    case MSK_IPAR_OPF_WRITE_SOLUTIONS: return "MSK_IPAR_OPF_WRITE_SOLUTIONS"
    case MSK_IPAR_OPTIMIZER: return "MSK_IPAR_OPTIMIZER"
    case MSK_IPAR_PARAM_READ_CASE_NAME: return "MSK_IPAR_PARAM_READ_CASE_NAME"
    case MSK_IPAR_PARAM_READ_IGN_ERROR: return "MSK_IPAR_PARAM_READ_IGN_ERROR"
    case MSK_IPAR_PRESOLVE_ELIMINATOR_MAX_FILL: return "MSK_IPAR_PRESOLVE_ELIMINATOR_MAX_FILL"
    case MSK_IPAR_PRESOLVE_ELIMINATOR_MAX_NUM_TRIES: return "MSK_IPAR_PRESOLVE_ELIMINATOR_MAX_NUM_TRIES"
    case MSK_IPAR_PRESOLVE_LEVEL: return "MSK_IPAR_PRESOLVE_LEVEL"
    case MSK_IPAR_PRESOLVE_LINDEP_ABS_WORK_TRH: return "MSK_IPAR_PRESOLVE_LINDEP_ABS_WORK_TRH"
    case MSK_IPAR_PRESOLVE_LINDEP_NEW: return "MSK_IPAR_PRESOLVE_LINDEP_NEW"
    case MSK_IPAR_PRESOLVE_LINDEP_REL_WORK_TRH: return "MSK_IPAR_PRESOLVE_LINDEP_REL_WORK_TRH"
    case MSK_IPAR_PRESOLVE_LINDEP_USE: return "MSK_IPAR_PRESOLVE_LINDEP_USE"
    case MSK_IPAR_PRESOLVE_MAX_NUM_PASS: return "MSK_IPAR_PRESOLVE_MAX_NUM_PASS"
    case MSK_IPAR_PRESOLVE_MAX_NUM_REDUCTIONS: return "MSK_IPAR_PRESOLVE_MAX_NUM_REDUCTIONS"
    case MSK_IPAR_PRESOLVE_USE: return "MSK_IPAR_PRESOLVE_USE"
    case MSK_IPAR_PRIMAL_REPAIR_OPTIMIZER: return "MSK_IPAR_PRIMAL_REPAIR_OPTIMIZER"
    case MSK_IPAR_PTF_WRITE_PARAMETERS: return "MSK_IPAR_PTF_WRITE_PARAMETERS"
    case MSK_IPAR_PTF_WRITE_SOLUTIONS: return "MSK_IPAR_PTF_WRITE_SOLUTIONS"
    case MSK_IPAR_PTF_WRITE_TRANSFORM: return "MSK_IPAR_PTF_WRITE_TRANSFORM"
    case MSK_IPAR_READ_DEBUG: return "MSK_IPAR_READ_DEBUG"
    case MSK_IPAR_READ_KEEP_FREE_CON: return "MSK_IPAR_READ_KEEP_FREE_CON"
    case MSK_IPAR_READ_MPS_FORMAT: return "MSK_IPAR_READ_MPS_FORMAT"
    case MSK_IPAR_READ_MPS_WIDTH: return "MSK_IPAR_READ_MPS_WIDTH"
    case MSK_IPAR_READ_TASK_IGNORE_PARAM: return "MSK_IPAR_READ_TASK_IGNORE_PARAM"
    case MSK_IPAR_REMOTE_USE_COMPRESSION: return "MSK_IPAR_REMOTE_USE_COMPRESSION"
    case MSK_IPAR_REMOVE_UNUSED_SOLUTIONS: return "MSK_IPAR_REMOVE_UNUSED_SOLUTIONS"
    case MSK_IPAR_SENSITIVITY_ALL: return "MSK_IPAR_SENSITIVITY_ALL"
    case MSK_IPAR_SENSITIVITY_OPTIMIZER: return "MSK_IPAR_SENSITIVITY_OPTIMIZER"
    case MSK_IPAR_SENSITIVITY_TYPE: return "MSK_IPAR_SENSITIVITY_TYPE"
    case MSK_IPAR_SIM_BASIS_FACTOR_USE: return "MSK_IPAR_SIM_BASIS_FACTOR_USE"
    case MSK_IPAR_SIM_DEGEN: return "MSK_IPAR_SIM_DEGEN"
    case MSK_IPAR_SIM_DETECT_PWL: return "MSK_IPAR_SIM_DETECT_PWL"
    case MSK_IPAR_SIM_DUAL_CRASH: return "MSK_IPAR_SIM_DUAL_CRASH"
    case MSK_IPAR_SIM_DUAL_PHASEONE_METHOD: return "MSK_IPAR_SIM_DUAL_PHASEONE_METHOD"
    case MSK_IPAR_SIM_DUAL_RESTRICT_SELECTION: return "MSK_IPAR_SIM_DUAL_RESTRICT_SELECTION"
    case MSK_IPAR_SIM_DUAL_SELECTION: return "MSK_IPAR_SIM_DUAL_SELECTION"
    case MSK_IPAR_SIM_EXPLOIT_DUPVEC: return "MSK_IPAR_SIM_EXPLOIT_DUPVEC"
    case MSK_IPAR_SIM_HOTSTART: return "MSK_IPAR_SIM_HOTSTART"
    case MSK_IPAR_SIM_HOTSTART_LU: return "MSK_IPAR_SIM_HOTSTART_LU"
    case MSK_IPAR_SIM_MAX_ITERATIONS: return "MSK_IPAR_SIM_MAX_ITERATIONS"
    case MSK_IPAR_SIM_MAX_NUM_SETBACKS: return "MSK_IPAR_SIM_MAX_NUM_SETBACKS"
    case MSK_IPAR_SIM_NON_SINGULAR: return "MSK_IPAR_SIM_NON_SINGULAR"
    case MSK_IPAR_SIM_PRIMAL_CRASH: return "MSK_IPAR_SIM_PRIMAL_CRASH"
    case MSK_IPAR_SIM_PRIMAL_PHASEONE_METHOD: return "MSK_IPAR_SIM_PRIMAL_PHASEONE_METHOD"
    case MSK_IPAR_SIM_PRIMAL_RESTRICT_SELECTION: return "MSK_IPAR_SIM_PRIMAL_RESTRICT_SELECTION"
    case MSK_IPAR_SIM_PRIMAL_SELECTION: return "MSK_IPAR_SIM_PRIMAL_SELECTION"
    case MSK_IPAR_SIM_REFACTOR_FREQ: return "MSK_IPAR_SIM_REFACTOR_FREQ"
    case MSK_IPAR_SIM_REFORMULATION: return "MSK_IPAR_SIM_REFORMULATION"
    case MSK_IPAR_SIM_SAVE_LU: return "MSK_IPAR_SIM_SAVE_LU"
    case MSK_IPAR_SIM_SCALING: return "MSK_IPAR_SIM_SCALING"
    case MSK_IPAR_SIM_SCALING_METHOD: return "MSK_IPAR_SIM_SCALING_METHOD"
    case MSK_IPAR_SIM_SEED: return "MSK_IPAR_SIM_SEED"
    case MSK_IPAR_SIM_SOLVE_FORM: return "MSK_IPAR_SIM_SOLVE_FORM"
    case MSK_IPAR_SIM_STABILITY_PRIORITY: return "MSK_IPAR_SIM_STABILITY_PRIORITY"
    case MSK_IPAR_SIM_SWITCH_OPTIMIZER: return "MSK_IPAR_SIM_SWITCH_OPTIMIZER"
    case MSK_IPAR_SOL_FILTER_KEEP_BASIC: return "MSK_IPAR_SOL_FILTER_KEEP_BASIC"
    case MSK_IPAR_SOL_FILTER_KEEP_RANGED: return "MSK_IPAR_SOL_FILTER_KEEP_RANGED"
    case MSK_IPAR_SOL_READ_NAME_WIDTH: return "MSK_IPAR_SOL_READ_NAME_WIDTH"
    case MSK_IPAR_SOL_READ_WIDTH: return "MSK_IPAR_SOL_READ_WIDTH"
    case MSK_IPAR_SOLUTION_CALLBACK: return "MSK_IPAR_SOLUTION_CALLBACK"
    case MSK_IPAR_TIMING_LEVEL: return "MSK_IPAR_TIMING_LEVEL"
    case MSK_IPAR_WRITE_BAS_CONSTRAINTS: return "MSK_IPAR_WRITE_BAS_CONSTRAINTS"
    case MSK_IPAR_WRITE_BAS_HEAD: return "MSK_IPAR_WRITE_BAS_HEAD"
    case MSK_IPAR_WRITE_BAS_VARIABLES: return "MSK_IPAR_WRITE_BAS_VARIABLES"
    case MSK_IPAR_WRITE_COMPRESSION: return "MSK_IPAR_WRITE_COMPRESSION"
    case MSK_IPAR_WRITE_DATA_PARAM: return "MSK_IPAR_WRITE_DATA_PARAM"
    case MSK_IPAR_WRITE_FREE_CON: return "MSK_IPAR_WRITE_FREE_CON"
    case MSK_IPAR_WRITE_GENERIC_NAMES: return "MSK_IPAR_WRITE_GENERIC_NAMES"
    case MSK_IPAR_WRITE_GENERIC_NAMES_IO: return "MSK_IPAR_WRITE_GENERIC_NAMES_IO"
    case MSK_IPAR_WRITE_IGNORE_INCOMPATIBLE_ITEMS: return "MSK_IPAR_WRITE_IGNORE_INCOMPATIBLE_ITEMS"
    case MSK_IPAR_WRITE_INT_CONSTRAINTS: return "MSK_IPAR_WRITE_INT_CONSTRAINTS"
    case MSK_IPAR_WRITE_INT_HEAD: return "MSK_IPAR_WRITE_INT_HEAD"
    case MSK_IPAR_WRITE_INT_VARIABLES: return "MSK_IPAR_WRITE_INT_VARIABLES"
    case MSK_IPAR_WRITE_JSON_INDENTATION: return "MSK_IPAR_WRITE_JSON_INDENTATION"
    case MSK_IPAR_WRITE_LP_FULL_OBJ: return "MSK_IPAR_WRITE_LP_FULL_OBJ"
    case MSK_IPAR_WRITE_LP_LINE_WIDTH: return "MSK_IPAR_WRITE_LP_LINE_WIDTH"
    case MSK_IPAR_WRITE_MPS_FORMAT: return "MSK_IPAR_WRITE_MPS_FORMAT"
    case MSK_IPAR_WRITE_MPS_INT: return "MSK_IPAR_WRITE_MPS_INT"
    case MSK_IPAR_WRITE_SOL_BARVARIABLES: return "MSK_IPAR_WRITE_SOL_BARVARIABLES"
    case MSK_IPAR_WRITE_SOL_CONSTRAINTS: return "MSK_IPAR_WRITE_SOL_CONSTRAINTS"
    case MSK_IPAR_WRITE_SOL_HEAD: return "MSK_IPAR_WRITE_SOL_HEAD"
    case MSK_IPAR_WRITE_SOL_IGNORE_INVALID_NAMES: return "MSK_IPAR_WRITE_SOL_IGNORE_INVALID_NAMES"
    case MSK_IPAR_WRITE_SOL_VARIABLES: return "MSK_IPAR_WRITE_SOL_VARIABLES"
    case MSK_IPAR_WRITE_TASK_INC_SOL: return "MSK_IPAR_WRITE_TASK_INC_SOL"
    case MSK_IPAR_WRITE_XML_MODE: return "MSK_IPAR_WRITE_XML_MODE"
    default: return "<unknown>"
  }
} // Iparam.String()

type Branchdir int32
// Specifies the branching direction.
const (
    MSK_BRANCH_DIR_FREE Branchdir = 0 // The mixed-integer optimizer decides which branch to choose.
    MSK_BRANCH_DIR_UP Branchdir = 1 // The mixed-integer optimizer always chooses the up branch first.
    MSK_BRANCH_DIR_DOWN Branchdir = 2 // The mixed-integer optimizer always chooses the down branch first.
    MSK_BRANCH_DIR_NEAR Branchdir = 3 // Branch in direction nearest to selected fractional variable.
    MSK_BRANCH_DIR_FAR Branchdir = 4 // Branch in direction farthest from selected fractional variable.
    MSK_BRANCH_DIR_ROOT_LP Branchdir = 5 // Chose direction based on root lp value of selected variable.
    MSK_BRANCH_DIR_GUIDED Branchdir = 6 // Branch in direction of current incumbent.
    MSK_BRANCH_DIR_PSEUDOCOST Branchdir = 7 // Branch based on the pseudocost of the variable.
    MSK_BRANCH_DIR_END Branchdir = 7
)
func (self Branchdir) String() string {
  switch self {
    case MSK_BRANCH_DIR_FREE: return "MSK_BRANCH_DIR_FREE"
    case MSK_BRANCH_DIR_UP: return "MSK_BRANCH_DIR_UP"
    case MSK_BRANCH_DIR_DOWN: return "MSK_BRANCH_DIR_DOWN"
    case MSK_BRANCH_DIR_NEAR: return "MSK_BRANCH_DIR_NEAR"
    case MSK_BRANCH_DIR_FAR: return "MSK_BRANCH_DIR_FAR"
    case MSK_BRANCH_DIR_ROOT_LP: return "MSK_BRANCH_DIR_ROOT_LP"
    case MSK_BRANCH_DIR_GUIDED: return "MSK_BRANCH_DIR_GUIDED"
    case MSK_BRANCH_DIR_PSEUDOCOST: return "MSK_BRANCH_DIR_PSEUDOCOST"
    default: return "<unknown>"
  }
} // Branchdir.String()

type Miqcqoreformmethod int32
// Specifies the reformulation method for mixed-integer quadratic problems.
const (
    MSK_MIO_QCQO_REFORMULATION_METHOD_FREE Miqcqoreformmethod = 0 // The mixed-integer optimizer decides which reformulation method to apply.
    MSK_MIO_QCQO_REFORMULATION_METHOD_NONE Miqcqoreformmethod = 1 // No reformulation method is applied.
    MSK_MIO_QCQO_REFORMULATION_METHOD_LINEARIZATION Miqcqoreformmethod = 2 // A reformulation via linearization is applied.
    MSK_MIO_QCQO_REFORMULATION_METHOD_EIGEN_VAL_METHOD Miqcqoreformmethod = 3 // The eigenvalue method is applied.
    MSK_MIO_QCQO_REFORMULATION_METHOD_DIAG_SDP Miqcqoreformmethod = 4 // A perturbation of matrix diagonals via the solution of SDPs is applied.
    MSK_MIO_QCQO_REFORMULATION_METHOD_RELAX_SDP Miqcqoreformmethod = 5 // A Reformulation based on the solution of an SDP-relaxation of the problem is applied.
    MSK_MIO_QCQO_REFORMULATION_METHOD_END Miqcqoreformmethod = 5
)
func (self Miqcqoreformmethod) String() string {
  switch self {
    case MSK_MIO_QCQO_REFORMULATION_METHOD_FREE: return "MSK_MIO_QCQO_REFORMULATION_METHOD_FREE"
    case MSK_MIO_QCQO_REFORMULATION_METHOD_NONE: return "MSK_MIO_QCQO_REFORMULATION_METHOD_NONE"
    case MSK_MIO_QCQO_REFORMULATION_METHOD_LINEARIZATION: return "MSK_MIO_QCQO_REFORMULATION_METHOD_LINEARIZATION"
    case MSK_MIO_QCQO_REFORMULATION_METHOD_EIGEN_VAL_METHOD: return "MSK_MIO_QCQO_REFORMULATION_METHOD_EIGEN_VAL_METHOD"
    case MSK_MIO_QCQO_REFORMULATION_METHOD_DIAG_SDP: return "MSK_MIO_QCQO_REFORMULATION_METHOD_DIAG_SDP"
    case MSK_MIO_QCQO_REFORMULATION_METHOD_RELAX_SDP: return "MSK_MIO_QCQO_REFORMULATION_METHOD_RELAX_SDP"
    default: return "<unknown>"
  }
} // Miqcqoreformmethod.String()

type Miodatapermmethod int32
// Specifies the problem data permutation method for mixed-integer problems.
const (
    MSK_MIO_DATA_PERMUTATION_METHOD_NONE Miodatapermmethod = 0 // No problem data permutation is applied.
    MSK_MIO_DATA_PERMUTATION_METHOD_CYCLIC_SHIFT Miodatapermmethod = 1 // A random cyclic shift is applied to permute the problem data.
    MSK_MIO_DATA_PERMUTATION_METHOD_RANDOM Miodatapermmethod = 2 // A random permutation is applied to the problem data.
    MSK_MIO_DATA_PERMUTATION_METHOD_END Miodatapermmethod = 2
)
func (self Miodatapermmethod) String() string {
  switch self {
    case MSK_MIO_DATA_PERMUTATION_METHOD_NONE: return "MSK_MIO_DATA_PERMUTATION_METHOD_NONE"
    case MSK_MIO_DATA_PERMUTATION_METHOD_CYCLIC_SHIFT: return "MSK_MIO_DATA_PERMUTATION_METHOD_CYCLIC_SHIFT"
    case MSK_MIO_DATA_PERMUTATION_METHOD_RANDOM: return "MSK_MIO_DATA_PERMUTATION_METHOD_RANDOM"
    default: return "<unknown>"
  }
} // Miodatapermmethod.String()

type Miocontsoltype int32
// Continuous mixed-integer solution type
const (
    MSK_MIO_CONT_SOL_NONE Miocontsoltype = 0 // No interior-point or basic solution.
    MSK_MIO_CONT_SOL_ROOT Miocontsoltype = 1 // Solutions to the root node problem.
    MSK_MIO_CONT_SOL_ITG Miocontsoltype = 2 // A feasible primal solution.
    MSK_MIO_CONT_SOL_ITG_REL Miocontsoltype = 3 // A feasible primal solution or a root node solution if the problem is infeasible.
    MSK_MIO_CONT_SOL_END Miocontsoltype = 3
)
func (self Miocontsoltype) String() string {
  switch self {
    case MSK_MIO_CONT_SOL_NONE: return "MSK_MIO_CONT_SOL_NONE"
    case MSK_MIO_CONT_SOL_ROOT: return "MSK_MIO_CONT_SOL_ROOT"
    case MSK_MIO_CONT_SOL_ITG: return "MSK_MIO_CONT_SOL_ITG"
    case MSK_MIO_CONT_SOL_ITG_REL: return "MSK_MIO_CONT_SOL_ITG_REL"
    default: return "<unknown>"
  }
} // Miocontsoltype.String()

type Miomode int32
// Integer restrictions
const (
    MSK_MIO_MODE_IGNORED Miomode = 0 // The integer constraints are ignored and the problem is solved as a continuous problem.
    MSK_MIO_MODE_SATISFIED Miomode = 1 // Integer restrictions should be satisfied.
    MSK_MIO_MODE_END Miomode = 1
)
func (self Miomode) String() string {
  switch self {
    case MSK_MIO_MODE_IGNORED: return "MSK_MIO_MODE_IGNORED"
    case MSK_MIO_MODE_SATISFIED: return "MSK_MIO_MODE_SATISFIED"
    default: return "<unknown>"
  }
} // Miomode.String()

type Mionodeseltype int32
// Mixed-integer node selection types
const (
    MSK_MIO_NODE_SELECTION_FREE Mionodeseltype = 0 // The optimizer decides the node selection strategy.
    MSK_MIO_NODE_SELECTION_FIRST Mionodeseltype = 1 // The optimizer employs a depth first node selection strategy.
    MSK_MIO_NODE_SELECTION_BEST Mionodeseltype = 2 // The optimizer employs a best bound node selection strategy.
    MSK_MIO_NODE_SELECTION_PSEUDO Mionodeseltype = 3 // The optimizer employs selects the node based on a pseudo cost estimate.
    MSK_MIO_NODE_SELECTION_END Mionodeseltype = 3
)
func (self Mionodeseltype) String() string {
  switch self {
    case MSK_MIO_NODE_SELECTION_FREE: return "MSK_MIO_NODE_SELECTION_FREE"
    case MSK_MIO_NODE_SELECTION_FIRST: return "MSK_MIO_NODE_SELECTION_FIRST"
    case MSK_MIO_NODE_SELECTION_BEST: return "MSK_MIO_NODE_SELECTION_BEST"
    case MSK_MIO_NODE_SELECTION_PSEUDO: return "MSK_MIO_NODE_SELECTION_PSEUDO"
    default: return "<unknown>"
  }
} // Mionodeseltype.String()

type Miovarseltype int32
// Mixed-integer variable selection types
const (
    MSK_MIO_VAR_SELECTION_FREE Miovarseltype = 0 // The optimizer decides the variable selection strategy.
    MSK_MIO_VAR_SELECTION_PSEUDOCOST Miovarseltype = 1 // The optimizer employs pseudocost variable selection.
    MSK_MIO_VAR_SELECTION_STRONG Miovarseltype = 2 // The optimizer employs strong branching variable selection.
    MSK_MIO_VAR_SELECTION_END Miovarseltype = 2
)
func (self Miovarseltype) String() string {
  switch self {
    case MSK_MIO_VAR_SELECTION_FREE: return "MSK_MIO_VAR_SELECTION_FREE"
    case MSK_MIO_VAR_SELECTION_PSEUDOCOST: return "MSK_MIO_VAR_SELECTION_PSEUDOCOST"
    case MSK_MIO_VAR_SELECTION_STRONG: return "MSK_MIO_VAR_SELECTION_STRONG"
    default: return "<unknown>"
  }
} // Miovarseltype.String()

type Mpsformat int32
// MPS file format type
const (
    MSK_MPS_FORMAT_STRICT Mpsformat = 0 // It is assumed that the input file satisfies the MPS format strictly.
    MSK_MPS_FORMAT_RELAXED Mpsformat = 1 // It is assumed that the input file satisfies a slightly relaxed version of the MPS format.
    MSK_MPS_FORMAT_FREE Mpsformat = 2 // It is assumed that the input file satisfies the free MPS format. This implies that spaces are not allowed in names. Otherwise the format is free.
    MSK_MPS_FORMAT_CPLEX Mpsformat = 3 // The CPLEX compatible version of the MPS format is employed.
    MSK_MPS_FORMAT_END Mpsformat = 3
)
func (self Mpsformat) String() string {
  switch self {
    case MSK_MPS_FORMAT_STRICT: return "MSK_MPS_FORMAT_STRICT"
    case MSK_MPS_FORMAT_RELAXED: return "MSK_MPS_FORMAT_RELAXED"
    case MSK_MPS_FORMAT_FREE: return "MSK_MPS_FORMAT_FREE"
    case MSK_MPS_FORMAT_CPLEX: return "MSK_MPS_FORMAT_CPLEX"
    default: return "<unknown>"
  }
} // Mpsformat.String()

type Objsense int32
// Objective sense types
const (
    MSK_OBJECTIVE_SENSE_MINIMIZE Objsense = 0 // The problem should be minimized.
    MSK_OBJECTIVE_SENSE_MAXIMIZE Objsense = 1 // The problem should be maximized.
    MSK_OBJECTIVE_SENSE_END Objsense = 1
)
func (self Objsense) String() string {
  switch self {
    case MSK_OBJECTIVE_SENSE_MINIMIZE: return "MSK_OBJECTIVE_SENSE_MINIMIZE"
    case MSK_OBJECTIVE_SENSE_MAXIMIZE: return "MSK_OBJECTIVE_SENSE_MAXIMIZE"
    default: return "<unknown>"
  }
} // Objsense.String()

type Onoffkey int32
// On/off
const (
    MSK_OFF Onoffkey = 0 // Switch the option off.
    MSK_ON Onoffkey = 1 // Switch the option on.
)
func (self Onoffkey) String() string {
  switch self {
    case MSK_OFF: return "MSK_OFF"
    case MSK_ON: return "MSK_ON"
    default: return "<unknown>"
  }
} // Onoffkey.String()

type Optimizertype int32
// Optimizer types
const (
    MSK_OPTIMIZER_CONIC Optimizertype = 0 // The optimizer for problems having conic constraints.
    MSK_OPTIMIZER_DUAL_SIMPLEX Optimizertype = 1 // The dual simplex optimizer is used.
    MSK_OPTIMIZER_FREE Optimizertype = 2 // The optimizer is chosen automatically.
    MSK_OPTIMIZER_FREE_SIMPLEX Optimizertype = 3 // One of the simplex optimizers is used.
    MSK_OPTIMIZER_INTPNT Optimizertype = 4 // The interior-point optimizer is used.
    MSK_OPTIMIZER_MIXED_INT Optimizertype = 5 // The mixed-integer optimizer.
    MSK_OPTIMIZER_PRIMAL_SIMPLEX Optimizertype = 6 // The primal simplex optimizer is used.
    MSK_OPTIMIZER_END Optimizertype = 6
)
func (self Optimizertype) String() string {
  switch self {
    case MSK_OPTIMIZER_CONIC: return "MSK_OPTIMIZER_CONIC"
    case MSK_OPTIMIZER_DUAL_SIMPLEX: return "MSK_OPTIMIZER_DUAL_SIMPLEX"
    case MSK_OPTIMIZER_FREE: return "MSK_OPTIMIZER_FREE"
    case MSK_OPTIMIZER_FREE_SIMPLEX: return "MSK_OPTIMIZER_FREE_SIMPLEX"
    case MSK_OPTIMIZER_INTPNT: return "MSK_OPTIMIZER_INTPNT"
    case MSK_OPTIMIZER_MIXED_INT: return "MSK_OPTIMIZER_MIXED_INT"
    case MSK_OPTIMIZER_PRIMAL_SIMPLEX: return "MSK_OPTIMIZER_PRIMAL_SIMPLEX"
    default: return "<unknown>"
  }
} // Optimizertype.String()

type Orderingtype int32
// Ordering strategies
const (
    MSK_ORDER_METHOD_FREE Orderingtype = 0 // The ordering method is chosen automatically.
    MSK_ORDER_METHOD_APPMINLOC Orderingtype = 1 // Approximate minimum local fill-in ordering is employed.
    MSK_ORDER_METHOD_EXPERIMENTAL Orderingtype = 2 // This option should not be used.
    MSK_ORDER_METHOD_TRY_GRAPHPAR Orderingtype = 3 // Always try the graph partitioning based ordering.
    MSK_ORDER_METHOD_FORCE_GRAPHPAR Orderingtype = 4 // Always use the graph partitioning based ordering even if it is worse than the approximate minimum local fill ordering.
    MSK_ORDER_METHOD_NONE Orderingtype = 5 // No ordering is used. Note using this value almost always leads to a significantly slow down.
    MSK_ORDER_METHOD_END Orderingtype = 5
)
func (self Orderingtype) String() string {
  switch self {
    case MSK_ORDER_METHOD_FREE: return "MSK_ORDER_METHOD_FREE"
    case MSK_ORDER_METHOD_APPMINLOC: return "MSK_ORDER_METHOD_APPMINLOC"
    case MSK_ORDER_METHOD_EXPERIMENTAL: return "MSK_ORDER_METHOD_EXPERIMENTAL"
    case MSK_ORDER_METHOD_TRY_GRAPHPAR: return "MSK_ORDER_METHOD_TRY_GRAPHPAR"
    case MSK_ORDER_METHOD_FORCE_GRAPHPAR: return "MSK_ORDER_METHOD_FORCE_GRAPHPAR"
    case MSK_ORDER_METHOD_NONE: return "MSK_ORDER_METHOD_NONE"
    default: return "<unknown>"
  }
} // Orderingtype.String()

type Presolvemode int32
// Presolve method.
const (
    MSK_PRESOLVE_MODE_OFF Presolvemode = 0 // The problem is not presolved before it is optimized.
    MSK_PRESOLVE_MODE_ON Presolvemode = 1 // The problem is presolved before it is optimized.
    MSK_PRESOLVE_MODE_FREE Presolvemode = 2 // It is decided automatically whether to presolve before the problem is optimized.
    MSK_PRESOLVE_MODE_END Presolvemode = 2
)
func (self Presolvemode) String() string {
  switch self {
    case MSK_PRESOLVE_MODE_OFF: return "MSK_PRESOLVE_MODE_OFF"
    case MSK_PRESOLVE_MODE_ON: return "MSK_PRESOLVE_MODE_ON"
    case MSK_PRESOLVE_MODE_FREE: return "MSK_PRESOLVE_MODE_FREE"
    default: return "<unknown>"
  }
} // Presolvemode.String()

type Parametertype int32
// Parameter type
const (
    MSK_PAR_INVALID_TYPE Parametertype = 0 // Not a valid parameter.
    MSK_PAR_DOU_TYPE Parametertype = 1 // Is a double parameter.
    MSK_PAR_INT_TYPE Parametertype = 2 // Is an integer parameter.
    MSK_PAR_STR_TYPE Parametertype = 3 // Is a string parameter.
    MSK_PAR_END Parametertype = 3
)
func (self Parametertype) String() string {
  switch self {
    case MSK_PAR_INVALID_TYPE: return "MSK_PAR_INVALID_TYPE"
    case MSK_PAR_DOU_TYPE: return "MSK_PAR_DOU_TYPE"
    case MSK_PAR_INT_TYPE: return "MSK_PAR_INT_TYPE"
    case MSK_PAR_STR_TYPE: return "MSK_PAR_STR_TYPE"
    default: return "<unknown>"
  }
} // Parametertype.String()

type Problemitem int32
// Problem data items
const (
    MSK_PI_VAR Problemitem = 0 // Item is a variable.
    MSK_PI_CON Problemitem = 1 // Item is a constraint.
    MSK_PI_CONE Problemitem = 2 // Item is a cone.
    MSK_PI_END Problemitem = 2
)
func (self Problemitem) String() string {
  switch self {
    case MSK_PI_VAR: return "MSK_PI_VAR"
    case MSK_PI_CON: return "MSK_PI_CON"
    case MSK_PI_CONE: return "MSK_PI_CONE"
    default: return "<unknown>"
  }
} // Problemitem.String()

type Problemtype int32
// Problem types
const (
    MSK_PROBTYPE_LO Problemtype = 0 // The problem is a linear optimization problem.
    MSK_PROBTYPE_QO Problemtype = 1 // The problem is a quadratic optimization problem.
    MSK_PROBTYPE_QCQO Problemtype = 2 // The problem is a quadratically constrained optimization problem.
    MSK_PROBTYPE_CONIC Problemtype = 3 // A conic optimization.
    MSK_PROBTYPE_MIXED Problemtype = 4 // General nonlinear constraints and conic constraints. This combination can not be solved by MOSEK.
    MSK_PROBTYPE_END Problemtype = 4
)
func (self Problemtype) String() string {
  switch self {
    case MSK_PROBTYPE_LO: return "MSK_PROBTYPE_LO"
    case MSK_PROBTYPE_QO: return "MSK_PROBTYPE_QO"
    case MSK_PROBTYPE_QCQO: return "MSK_PROBTYPE_QCQO"
    case MSK_PROBTYPE_CONIC: return "MSK_PROBTYPE_CONIC"
    case MSK_PROBTYPE_MIXED: return "MSK_PROBTYPE_MIXED"
    default: return "<unknown>"
  }
} // Problemtype.String()

type Prosta int32
// Problem status keys
const (
    MSK_PRO_STA_UNKNOWN Prosta = 0 // Unknown problem status.
    MSK_PRO_STA_PRIM_AND_DUAL_FEAS Prosta = 1 // The problem is primal and dual feasible.
    MSK_PRO_STA_PRIM_FEAS Prosta = 2 // The problem is primal feasible.
    MSK_PRO_STA_DUAL_FEAS Prosta = 3 // The problem is dual feasible.
    MSK_PRO_STA_PRIM_INFEAS Prosta = 4 // The problem is primal infeasible.
    MSK_PRO_STA_DUAL_INFEAS Prosta = 5 // The problem is dual infeasible.
    MSK_PRO_STA_PRIM_AND_DUAL_INFEAS Prosta = 6 // The problem is primal and dual infeasible.
    MSK_PRO_STA_ILL_POSED Prosta = 7 // The problem is ill-posed. For example, it may be primal and dual feasible but have a positive duality gap.
    MSK_PRO_STA_PRIM_INFEAS_OR_UNBOUNDED Prosta = 8 // The problem is either primal infeasible or unbounded. This may occur for mixed-integer problems.
    MSK_PRO_STA_END Prosta = 8
)
func (self Prosta) String() string {
  switch self {
    case MSK_PRO_STA_UNKNOWN: return "MSK_PRO_STA_UNKNOWN"
    case MSK_PRO_STA_PRIM_AND_DUAL_FEAS: return "MSK_PRO_STA_PRIM_AND_DUAL_FEAS"
    case MSK_PRO_STA_PRIM_FEAS: return "MSK_PRO_STA_PRIM_FEAS"
    case MSK_PRO_STA_DUAL_FEAS: return "MSK_PRO_STA_DUAL_FEAS"
    case MSK_PRO_STA_PRIM_INFEAS: return "MSK_PRO_STA_PRIM_INFEAS"
    case MSK_PRO_STA_DUAL_INFEAS: return "MSK_PRO_STA_DUAL_INFEAS"
    case MSK_PRO_STA_PRIM_AND_DUAL_INFEAS: return "MSK_PRO_STA_PRIM_AND_DUAL_INFEAS"
    case MSK_PRO_STA_ILL_POSED: return "MSK_PRO_STA_ILL_POSED"
    case MSK_PRO_STA_PRIM_INFEAS_OR_UNBOUNDED: return "MSK_PRO_STA_PRIM_INFEAS_OR_UNBOUNDED"
    default: return "<unknown>"
  }
} // Prosta.String()

type Xmlwriteroutputtype int32
// XML writer output mode
const (
    MSK_WRITE_XML_MODE_ROW Xmlwriteroutputtype = 0 // Write in row order.
    MSK_WRITE_XML_MODE_COL Xmlwriteroutputtype = 1 // Write in column order.
    MSK_WRITE_XML_MODE_END Xmlwriteroutputtype = 1
)
func (self Xmlwriteroutputtype) String() string {
  switch self {
    case MSK_WRITE_XML_MODE_ROW: return "MSK_WRITE_XML_MODE_ROW"
    case MSK_WRITE_XML_MODE_COL: return "MSK_WRITE_XML_MODE_COL"
    default: return "<unknown>"
  }
} // Xmlwriteroutputtype.String()

type Rescode int32
// Response codes
const (
    MSK_RES_OK Rescode = 0 // No error occurred.
    MSK_RES_WRN_OPEN_PARAM_FILE Rescode = 50 // The parameter file could not be opened.
    MSK_RES_WRN_LARGE_BOUND Rescode = 51 // A numerically large bound value is specified.
    MSK_RES_WRN_LARGE_LO_BOUND Rescode = 52 // A numerically large lower bound value is specified.
    MSK_RES_WRN_LARGE_UP_BOUND Rescode = 53 // A numerically large upper bound value is specified.
    MSK_RES_WRN_LARGE_CON_FX Rescode = 54 // A equality constraint is fixed to numerically large value.
    MSK_RES_WRN_LARGE_CJ Rescode = 57 // A numerically large value is specified for one element in c.
    MSK_RES_WRN_LARGE_AIJ Rescode = 62 // A numerically large value is specified for an element in A.
    MSK_RES_WRN_ZERO_AIJ Rescode = 63 // One or more zero elements are specified in A.
    MSK_RES_WRN_NAME_MAX_LEN Rescode = 65 // A name is longer than the buffer that is supposed to hold it.
    MSK_RES_WRN_SPAR_MAX_LEN Rescode = 66 // A value for a string parameter is longer than the buffer that is supposed to hold it.
    MSK_RES_WRN_MPS_SPLIT_RHS_VECTOR Rescode = 70 // An RHS vector is split into several nonadjacent parts.
    MSK_RES_WRN_MPS_SPLIT_RAN_VECTOR Rescode = 71 // A RANGE vector is split into several nonadjacent parts in an MPS file.
    MSK_RES_WRN_MPS_SPLIT_BOU_VECTOR Rescode = 72 // A BOUNDS vector is split into several nonadjacent parts in an MPS file.
    MSK_RES_WRN_LP_OLD_QUAD_FORMAT Rescode = 80 // Missing '/2' after quadratic expressions in bound or objective.
    MSK_RES_WRN_LP_DROP_VARIABLE Rescode = 85 // Ignore a variable because the variable was not previously defined.
    MSK_RES_WRN_NZ_IN_UPR_TRI Rescode = 200 // Non-zero elements specified in the upper triangle of a matrix were ignored.
    MSK_RES_WRN_DROPPED_NZ_QOBJ Rescode = 201 // One or more non-zero elements were dropped in the Q matrix in the objective.
    MSK_RES_WRN_IGNORE_INTEGER Rescode = 250 // Ignored integer constraints.
    MSK_RES_WRN_NO_GLOBAL_OPTIMIZER Rescode = 251 // No global optimizer is available.
    MSK_RES_WRN_MIO_INFEASIBLE_FINAL Rescode = 270 // The final mixed-integer problem with all the integer variables fixed at their optimal values is infeasible.
    MSK_RES_WRN_SOL_FILTER Rescode = 300 // Invalid solution filter is specified.
    MSK_RES_WRN_UNDEF_SOL_FILE_NAME Rescode = 350 // Undefined name occurred in a solution.
    MSK_RES_WRN_SOL_FILE_IGNORED_CON Rescode = 351 // One or more lines in the constraint section were ignored when reading a solution file.
    MSK_RES_WRN_SOL_FILE_IGNORED_VAR Rescode = 352 // One or more lines in the variable section were ignored when reading a solution file.
    MSK_RES_WRN_TOO_FEW_BASIS_VARS Rescode = 400 // An incomplete basis is specified.
    MSK_RES_WRN_TOO_MANY_BASIS_VARS Rescode = 405 // A basis with too many variables is specified.
    MSK_RES_WRN_LICENSE_EXPIRE Rescode = 500 // The license expires.
    MSK_RES_WRN_LICENSE_SERVER Rescode = 501 // The license server is not responding.
    MSK_RES_WRN_EMPTY_NAME Rescode = 502 // A variable or constraint name is empty. The output file may be invalid.
    MSK_RES_WRN_USING_GENERIC_NAMES Rescode = 503 // Generic names are used because a name is invalid for requested format.
    MSK_RES_WRN_INVALID_MPS_NAME Rescode = 504 // A name e.g. a row name is not a valid MPS name.
    MSK_RES_WRN_INVALID_MPS_OBJ_NAME Rescode = 505 // The objective name is not a valid MPS name.
    MSK_RES_WRN_LICENSE_FEATURE_EXPIRE Rescode = 509 // The license expires.
    MSK_RES_WRN_PARAM_NAME_DOU Rescode = 510 // Parameter name not recognized.
    MSK_RES_WRN_PARAM_NAME_INT Rescode = 511 // Parameter name not recognized.
    MSK_RES_WRN_PARAM_NAME_STR Rescode = 512 // Parameter name not recognized.
    MSK_RES_WRN_PARAM_STR_VALUE Rescode = 515 // A parameter value is not correct.
    MSK_RES_WRN_PARAM_IGNORED_CMIO Rescode = 516 // A parameter was ignored by the conic mixed integer optimizer.
    MSK_RES_WRN_ZEROS_IN_SPARSE_ROW Rescode = 705 // One or more (near) zero elements are specified in a sparse row of a matrix.
    MSK_RES_WRN_ZEROS_IN_SPARSE_COL Rescode = 710 // One or more (near) zero elements are specified in a sparse column of a matrix.
    MSK_RES_WRN_INCOMPLETE_LINEAR_DEPENDENCY_CHECK Rescode = 800 // The linear dependency check(s) is incomplete.
    MSK_RES_WRN_ELIMINATOR_SPACE Rescode = 801 // The eliminator is skipped at least once due to lack of space.
    MSK_RES_WRN_PRESOLVE_OUTOFSPACE Rescode = 802 // The presolve is incomplete due to lack of space.
    MSK_RES_WRN_PRESOLVE_PRIMAL_PERTUBATIONS Rescode = 803 // The presolve perturbed the bounds of the primal problem. This is an indication that the problem is nearly infeasible.
    MSK_RES_WRN_WRITE_CHANGED_NAMES Rescode = 830 // Some names were changed because they were invalid for the output file format.
    MSK_RES_WRN_WRITE_DISCARDED_CFIX Rescode = 831 // The fixed objective term was discarded in the output file.
    MSK_RES_WRN_DUPLICATE_CONSTRAINT_NAMES Rescode = 850 // Two constraint names are identical.
    MSK_RES_WRN_DUPLICATE_VARIABLE_NAMES Rescode = 851 // Two variable names are identical.
    MSK_RES_WRN_DUPLICATE_BARVARIABLE_NAMES Rescode = 852 // Two barvariable names are identical.
    MSK_RES_WRN_DUPLICATE_CONE_NAMES Rescode = 853 // Two cone names are identical.
    MSK_RES_WRN_WRITE_LP_INVALID_VAR_NAMES Rescode = 854 // LP file will be written with generic variable names.
    MSK_RES_WRN_WRITE_LP_DUPLICATE_VAR_NAMES Rescode = 855 // LP file will be written with generic variable names.
    MSK_RES_WRN_WRITE_LP_INVALID_CON_NAMES Rescode = 856 // LP file will be written with generic constraint names.
    MSK_RES_WRN_WRITE_LP_DUPLICATE_CON_NAMES Rescode = 857 // LP file will be written with generic constraint names.
    MSK_RES_WRN_ANA_LARGE_BOUNDS Rescode = 900 // Warn against very large bounds.
    MSK_RES_WRN_ANA_C_ZERO Rescode = 901 // Warn against all objective coefficients being zero.
    MSK_RES_WRN_ANA_EMPTY_COLS Rescode = 902 // Warn against empty columns.
    MSK_RES_WRN_ANA_CLOSE_BOUNDS Rescode = 903 // Warn against close bounds.
    MSK_RES_WRN_ANA_ALMOST_INT_BOUNDS Rescode = 904 // Warn against almost integral bounds.
    MSK_RES_WRN_NO_INFEASIBILITY_REPORT_WHEN_MATRIX_VARIABLES Rescode = 930 // An infeasibility report is not available when the problem contains matrix variables.
    MSK_RES_WRN_NO_DUALIZER Rescode = 950 // No automatic dualizer is available for the specified problem.
    MSK_RES_WRN_SYM_MAT_LARGE Rescode = 960 // A numerically large value is specified for an element in E.
    MSK_RES_WRN_MODIFIED_DOUBLE_PARAMETER Rescode = 970 // A double parameter related to solver tolerances has a non-default value.
    MSK_RES_WRN_LARGE_FIJ Rescode = 980 // A numerically large value is specified for an element in F.
    MSK_RES_ERR_LICENSE Rescode = 1000 // Invalid license.
    MSK_RES_ERR_LICENSE_EXPIRED Rescode = 1001 // The license has expired.
    MSK_RES_ERR_LICENSE_VERSION Rescode = 1002 // Invalid license version.
    MSK_RES_ERR_LICENSE_OLD_SERVER_VERSION Rescode = 1003 // The license server version is too old.
    MSK_RES_ERR_SIZE_LICENSE Rescode = 1005 // The problem is bigger than the license.
    MSK_RES_ERR_PROB_LICENSE Rescode = 1006 // The software is not licensed to solve the problem.
    MSK_RES_ERR_FILE_LICENSE Rescode = 1007 // Invalid license file.
    MSK_RES_ERR_MISSING_LICENSE_FILE Rescode = 1008 // A license cannot be located.
    MSK_RES_ERR_SIZE_LICENSE_CON Rescode = 1010 // The problem has too many constraints.
    MSK_RES_ERR_SIZE_LICENSE_VAR Rescode = 1011 // The problem has too many variables.
    MSK_RES_ERR_SIZE_LICENSE_INTVAR Rescode = 1012 // The problem contains too many integer variables.
    MSK_RES_ERR_OPTIMIZER_LICENSE Rescode = 1013 // The optimizer required is not licensed.
    MSK_RES_ERR_FLEXLM Rescode = 1014 // The license manager reported an error.
    MSK_RES_ERR_LICENSE_SERVER Rescode = 1015 // The license server is not responding.
    MSK_RES_ERR_LICENSE_MAX Rescode = 1016 // Maximum number of licenses is reached.
    MSK_RES_ERR_LICENSE_MOSEKLM_DAEMON Rescode = 1017 // The MOSEKLM license manager daemon is not up and running.
    MSK_RES_ERR_LICENSE_FEATURE Rescode = 1018 // A requested feature is not available in the license file(s).
    MSK_RES_ERR_PLATFORM_NOT_LICENSED Rescode = 1019 // A requested license feature is not available for the required platform.
    MSK_RES_ERR_LICENSE_CANNOT_ALLOCATE Rescode = 1020 // The license system cannot allocate the memory required.
    MSK_RES_ERR_LICENSE_CANNOT_CONNECT Rescode = 1021 // MOSEK cannot connect to the license server.
    MSK_RES_ERR_LICENSE_INVALID_HOSTID Rescode = 1025 // The host ID specified in the license file does not match the host ID of the computer.
    MSK_RES_ERR_LICENSE_SERVER_VERSION Rescode = 1026 // The version specified in the checkout request is greater than the highest version number the daemon supports.
    MSK_RES_ERR_LICENSE_NO_SERVER_SUPPORT Rescode = 1027 // The license server does not support the requested feature.
    MSK_RES_ERR_LICENSE_NO_SERVER_LINE Rescode = 1028 // No SERVER lines in license file.
    MSK_RES_ERR_OLDER_DLL Rescode = 1035 // The dynamic link library is older than the specified version.
    MSK_RES_ERR_NEWER_DLL Rescode = 1036 // The dynamic link library is newer than the specified version.
    MSK_RES_ERR_LINK_FILE_DLL Rescode = 1040 // A file cannot be linked to a stream in the DLL version.
    MSK_RES_ERR_THREAD_MUTEX_INIT Rescode = 1045 // Could not initialize a mutex.
    MSK_RES_ERR_THREAD_MUTEX_LOCK Rescode = 1046 // Could not lock a mutex.
    MSK_RES_ERR_THREAD_MUTEX_UNLOCK Rescode = 1047 // Could not unlock a mutex.
    MSK_RES_ERR_THREAD_CREATE Rescode = 1048 // Could not create a thread.
    MSK_RES_ERR_THREAD_COND_INIT Rescode = 1049 // Could not initialize a condition.
    MSK_RES_ERR_UNKNOWN Rescode = 1050 // Unknown error.
    MSK_RES_ERR_SPACE Rescode = 1051 // Out of space.
    MSK_RES_ERR_FILE_OPEN Rescode = 1052 // An error occurred while opening a file.
    MSK_RES_ERR_FILE_READ Rescode = 1053 // An error occurred while reading file.
    MSK_RES_ERR_FILE_WRITE Rescode = 1054 // An error occurred while writing to a file.
    MSK_RES_ERR_DATA_FILE_EXT Rescode = 1055 // The data file format cannot be determined from the file name.
    MSK_RES_ERR_INVALID_FILE_NAME Rescode = 1056 // An invalid file name has been specified.
    MSK_RES_ERR_INVALID_SOL_FILE_NAME Rescode = 1057 // An invalid file name has been specified.
    MSK_RES_ERR_END_OF_FILE Rescode = 1059 // End of file reached.
    MSK_RES_ERR_NULL_ENV Rescode = 1060 // env is a null pointer.
    MSK_RES_ERR_NULL_TASK Rescode = 1061 // task is a null pointer.
    MSK_RES_ERR_INVALID_STREAM Rescode = 1062 // An invalid stream is referenced.
    MSK_RES_ERR_NO_INIT_ENV Rescode = 1063 // Environment is not initialized.
    MSK_RES_ERR_INVALID_TASK Rescode = 1064 // The task is invalid.
    MSK_RES_ERR_NULL_POINTER Rescode = 1065 // An argument to a function is unexpectedly a null pointer.
    MSK_RES_ERR_LIVING_TASKS Rescode = 1066 // Not all tasks associated with the environment have been deleted.
    MSK_RES_ERR_READ_GZIP Rescode = 1067 // Error encountered in GZIP stream.
    MSK_RES_ERR_READ_ZSTD Rescode = 1068 // Error encountered in ZSTD stream.
    MSK_RES_ERR_BLANK_NAME Rescode = 1070 // An all blank name has been specified.
    MSK_RES_ERR_DUP_NAME Rescode = 1071 // Duplicate names specified.
    MSK_RES_ERR_FORMAT_STRING Rescode = 1072 // The name format string is invalid.
    MSK_RES_ERR_SPARSITY_SPECIFICATION Rescode = 1073 // The sparsity included an index that was out of bounds of the shape.
    MSK_RES_ERR_MISMATCHING_DIMENSION Rescode = 1074 // Mismatching dimensions specified in arguments
    MSK_RES_ERR_INVALID_OBJ_NAME Rescode = 1075 // An invalid objective name is specified.
    MSK_RES_ERR_INVALID_CON_NAME Rescode = 1076 // An invalid constraint name is used.
    MSK_RES_ERR_INVALID_VAR_NAME Rescode = 1077 // An invalid variable name is used.
    MSK_RES_ERR_INVALID_CONE_NAME Rescode = 1078 // An invalid cone name is used.
    MSK_RES_ERR_INVALID_BARVAR_NAME Rescode = 1079 // An invalid symmetric matrix variable name is used.
    MSK_RES_ERR_SPACE_LEAKING Rescode = 1080 // MOSEK is leaking memory.
    MSK_RES_ERR_SPACE_NO_INFO Rescode = 1081 // No available information about the space usage.
    MSK_RES_ERR_DIMENSION_SPECIFICATION Rescode = 1082 // Invalid dimension specification
    MSK_RES_ERR_AXIS_NAME_SPECIFICATION Rescode = 1083 // Invalid axis names specification
    MSK_RES_ERR_READ_FORMAT Rescode = 1090 // The specified format cannot be read.
    MSK_RES_ERR_MPS_FILE Rescode = 1100 // An error occurred while reading an MPS file.
    MSK_RES_ERR_MPS_INV_FIELD Rescode = 1101 // Invalid field occurred while reading an MPS file.
    MSK_RES_ERR_MPS_INV_MARKER Rescode = 1102 // An invalid marker has been specified in the MPS file.
    MSK_RES_ERR_MPS_NULL_CON_NAME Rescode = 1103 // An empty constraint name is used in an MPS file.
    MSK_RES_ERR_MPS_NULL_VAR_NAME Rescode = 1104 // An empty variable name is used in an MPS file.
    MSK_RES_ERR_MPS_UNDEF_CON_NAME Rescode = 1105 // An undefined constraint name occurred in an MPS file.
    MSK_RES_ERR_MPS_UNDEF_VAR_NAME Rescode = 1106 // An undefined variable name occurred in an MPS file.
    MSK_RES_ERR_MPS_INVALID_CON_KEY Rescode = 1107 // An invalid constraint key occurred in an MPS file.
    MSK_RES_ERR_MPS_INVALID_BOUND_KEY Rescode = 1108 // An invalid bound key occurred in an MPS file.
    MSK_RES_ERR_MPS_INVALID_SEC_NAME Rescode = 1109 // An invalid section name occurred in an MPS file.
    MSK_RES_ERR_MPS_NO_OBJECTIVE Rescode = 1110 // No objective is defined in an MPS file.
    MSK_RES_ERR_MPS_SPLITTED_VAR Rescode = 1111 // The non-zero elements in A corresponding to a variable in an MPS file must be specified consecutively.
    MSK_RES_ERR_MPS_MUL_CON_NAME Rescode = 1112 // A constraint name is specified multiple times in the ROWS section in an MPS file.
    MSK_RES_ERR_MPS_MUL_QSEC Rescode = 1113 // Multiple QSECTIONs are specified for a constraint.
    MSK_RES_ERR_MPS_MUL_QOBJ Rescode = 1114 // The Q term in the objective is specified multiple times.
    MSK_RES_ERR_MPS_INV_SEC_ORDER Rescode = 1115 // The sections in an MPS file is not in the correct order.
    MSK_RES_ERR_MPS_MUL_CSEC Rescode = 1116 // Multiple CSECTIONs are given the same name.
    MSK_RES_ERR_MPS_CONE_TYPE Rescode = 1117 // Invalid cone type specified in a  CSECTION.
    MSK_RES_ERR_MPS_CONE_OVERLAP Rescode = 1118 // A variable is specified to be a member of several cones.
    MSK_RES_ERR_MPS_CONE_REPEAT Rescode = 1119 // A variable is repeated within the CSECTION.
    MSK_RES_ERR_MPS_NON_SYMMETRIC_Q Rescode = 1120 // A non symmetric matrix has been speciefied.
    MSK_RES_ERR_MPS_DUPLICATE_Q_ELEMENT Rescode = 1121 // Duplicate elements is specified in a Q matrix.
    MSK_RES_ERR_MPS_INVALID_OBJSENSE Rescode = 1122 // An invalid objective sense is specified.
    MSK_RES_ERR_MPS_TAB_IN_FIELD2 Rescode = 1125 // A tab char occurred in field 2.
    MSK_RES_ERR_MPS_TAB_IN_FIELD3 Rescode = 1126 // A tab char occurred in field 3.
    MSK_RES_ERR_MPS_TAB_IN_FIELD5 Rescode = 1127 // A tab char occurred in field 5.
    MSK_RES_ERR_MPS_INVALID_OBJ_NAME Rescode = 1128 // An invalid objective name is specified.
    MSK_RES_ERR_MPS_INVALID_KEY Rescode = 1129 // An invalid indicator key occurred in an MPS file.
    MSK_RES_ERR_MPS_INVALID_INDICATOR_CONSTRAINT Rescode = 1130 // An invalid indicator constraint is used. It must not be a ranged constraint.
    MSK_RES_ERR_MPS_INVALID_INDICATOR_VARIABLE Rescode = 1131 // An invalid indicator variable is specfied. It must be a binary variable.
    MSK_RES_ERR_MPS_INVALID_INDICATOR_VALUE Rescode = 1132 // An invalid indicator value is specfied. It must be either 0 or 1.
    MSK_RES_ERR_MPS_INVALID_INDICATOR_QUADRATIC_CONSTRAINT Rescode = 1133 // A quadratic constraint can be be an indicator constraint.
    MSK_RES_ERR_OPF_SYNTAX Rescode = 1134 // Syntax error in an OPF file
    MSK_RES_ERR_OPF_PREMATURE_EOF Rescode = 1136 // Premature end of file in an OPF file.
    MSK_RES_ERR_OPF_MISMATCHED_TAG Rescode = 1137 // Mismatched end-tag in OPF file
    MSK_RES_ERR_OPF_DUPLICATE_BOUND Rescode = 1138 // Either upper or lower bound was specified twice in OPF file
    MSK_RES_ERR_OPF_DUPLICATE_CONSTRAINT_NAME Rescode = 1139 // Duplicate constraint name in OPF File
    MSK_RES_ERR_OPF_INVALID_CONE_TYPE Rescode = 1140 // Invalid cone type in OPF File
    MSK_RES_ERR_OPF_INCORRECT_TAG_PARAM Rescode = 1141 // Invalid number of parameters in start-tag in OPF File
    MSK_RES_ERR_OPF_INVALID_TAG Rescode = 1142 // Invalid start-tag in OPF File
    MSK_RES_ERR_OPF_DUPLICATE_CONE_ENTRY Rescode = 1143 // Same variable appears in multiple cones in OPF File
    MSK_RES_ERR_OPF_TOO_LARGE Rescode = 1144 // The problem is too large to be correctly loaded
    MSK_RES_ERR_OPF_DUAL_INTEGER_SOLUTION Rescode = 1146 // Dual solution values are not allowed in OPF File
    MSK_RES_ERR_LP_EMPTY Rescode = 1151 // The problem cannot be written to an LP formatted file.
    MSK_RES_ERR_WRITE_MPS_INVALID_NAME Rescode = 1153 // An invalid name is created while writing an MPS file.
    MSK_RES_ERR_LP_INVALID_VAR_NAME Rescode = 1154 // A variable name is invalid when used in an LP formatted file.
    MSK_RES_ERR_WRITE_OPF_INVALID_VAR_NAME Rescode = 1156 // Empty variable names cannot be written to OPF files.
    MSK_RES_ERR_LP_FILE_FORMAT Rescode = 1157 // Syntax error in an LP file.
    MSK_RES_ERR_LP_EXPECTED_NUMBER Rescode = 1158 // Expected a number in LP file
    MSK_RES_ERR_READ_LP_MISSING_END_TAG Rescode = 1159 // Syntax error in LP fil. Possibly missing End tag.
    MSK_RES_ERR_LP_INDICATOR_VAR Rescode = 1160 // An indicator variable was not declared binary
    MSK_RES_ERR_LP_EXPECTED_OBJECTIVE Rescode = 1161 // Expected an objective section in LP file
    MSK_RES_ERR_LP_EXPECTED_CONSTRAINT_RELATION Rescode = 1162 // Expected constraint relation
    MSK_RES_ERR_LP_AMBIGUOUS_CONSTRAINT_BOUND Rescode = 1163 // Constraint has ambiguous or invalid bound
    MSK_RES_ERR_LP_DUPLICATE_SECTION Rescode = 1164 // Duplicate section
    MSK_RES_ERR_READ_LP_DELAYED_ROWS_NOT_SUPPORTED Rescode = 1165 // Duplicate section
    MSK_RES_ERR_WRITING_FILE Rescode = 1166 // An error occurred while writing file
    MSK_RES_ERR_INVALID_NAME_IN_SOL_FILE Rescode = 1170 // An invalid name occurred in a solution file.
    MSK_RES_ERR_JSON_SYNTAX Rescode = 1175 // Syntax error in an JSON data
    MSK_RES_ERR_JSON_STRING Rescode = 1176 // Error in JSON string.
    MSK_RES_ERR_JSON_NUMBER_OVERFLOW Rescode = 1177 // Invalid number entry - wrong type or value overflow.
    MSK_RES_ERR_JSON_FORMAT Rescode = 1178 // Error in an JSON Task file
    MSK_RES_ERR_JSON_DATA Rescode = 1179 // Inconsistent data in JSON Task file
    MSK_RES_ERR_JSON_MISSING_DATA Rescode = 1180 // Missing data section in JSON task file.
    MSK_RES_ERR_PTF_INCOMPATIBILITY Rescode = 1181 // Incompatible item
    MSK_RES_ERR_PTF_UNDEFINED_ITEM Rescode = 1182 // Undefined symbol referenced
    MSK_RES_ERR_PTF_INCONSISTENCY Rescode = 1183 // Inconsistent size of item
    MSK_RES_ERR_PTF_FORMAT Rescode = 1184 // Syntax error in an PTF file
    MSK_RES_ERR_ARGUMENT_LENNEQ Rescode = 1197 // Incorrect length of arguments.
    MSK_RES_ERR_ARGUMENT_TYPE Rescode = 1198 // Incorrect argument type.
    MSK_RES_ERR_NUM_ARGUMENTS Rescode = 1199 // Incorrect number of function arguments.
    MSK_RES_ERR_IN_ARGUMENT Rescode = 1200 // A function argument is incorrect.
    MSK_RES_ERR_ARGUMENT_DIMENSION Rescode = 1201 // A function argument is of incorrect dimension.
    MSK_RES_ERR_SHAPE_IS_TOO_LARGE Rescode = 1202 // The size of the n-dimensional shape is too large.
    MSK_RES_ERR_INDEX_IS_TOO_SMALL Rescode = 1203 // An index in an argument is too small.
    MSK_RES_ERR_INDEX_IS_TOO_LARGE Rescode = 1204 // An index in an argument is too large.
    MSK_RES_ERR_INDEX_IS_NOT_UNIQUE Rescode = 1205 // An index in an argument is is unique.
    MSK_RES_ERR_PARAM_NAME Rescode = 1206 // A parameter name is not correct.
    MSK_RES_ERR_PARAM_NAME_DOU Rescode = 1207 // A parameter name is not correct.
    MSK_RES_ERR_PARAM_NAME_INT Rescode = 1208 // A parameter name is not correct.
    MSK_RES_ERR_PARAM_NAME_STR Rescode = 1209 // A parameter name is not correct.
    MSK_RES_ERR_PARAM_INDEX Rescode = 1210 // Parameter index is out of range.
    MSK_RES_ERR_PARAM_IS_TOO_LARGE Rescode = 1215 // A parameter value is too large.
    MSK_RES_ERR_PARAM_IS_TOO_SMALL Rescode = 1216 // A parameter value is too small.
    MSK_RES_ERR_PARAM_VALUE_STR Rescode = 1217 // A parameter value string is incorrect.
    MSK_RES_ERR_PARAM_TYPE Rescode = 1218 // A parameter type is invalid.
    MSK_RES_ERR_INF_DOU_INDEX Rescode = 1219 // A double information index is out of range for the specified type.
    MSK_RES_ERR_INF_INT_INDEX Rescode = 1220 // An integer information index is out of range for the specified type.
    MSK_RES_ERR_INDEX_ARR_IS_TOO_SMALL Rescode = 1221 // An index in an array argument is too small.
    MSK_RES_ERR_INDEX_ARR_IS_TOO_LARGE Rescode = 1222 // An index in an array argument is too large.
    MSK_RES_ERR_INF_LINT_INDEX Rescode = 1225 // A long integer information index is out of range for the specified type.
    MSK_RES_ERR_ARG_IS_TOO_SMALL Rescode = 1226 // The value of a argument is too small.
    MSK_RES_ERR_ARG_IS_TOO_LARGE Rescode = 1227 // The value of a argument is too large.
    MSK_RES_ERR_INVALID_WHICHSOL Rescode = 1228 // whichsol is invalid.
    MSK_RES_ERR_INF_DOU_NAME Rescode = 1230 // A double information name is invalid.
    MSK_RES_ERR_INF_INT_NAME Rescode = 1231 // An integer information name is invalid.
    MSK_RES_ERR_INF_TYPE Rescode = 1232 // The information type is invalid.
    MSK_RES_ERR_INF_LINT_NAME Rescode = 1234 // A long integer information name is invalid.
    MSK_RES_ERR_INDEX Rescode = 1235 // An index is out of range.
    MSK_RES_ERR_WHICHSOL Rescode = 1236 // The solution defined by whichsol does not exists.
    MSK_RES_ERR_SOLITEM Rescode = 1237 // The solution number  solemn does not exists.
    MSK_RES_ERR_WHICHITEM_NOT_ALLOWED Rescode = 1238 // whichitem is unacceptable.
    MSK_RES_ERR_MAXNUMCON Rescode = 1240 // Invalid maximum number of constraints specified.
    MSK_RES_ERR_MAXNUMVAR Rescode = 1241 // The maximum number of variables limit is too small.
    MSK_RES_ERR_MAXNUMBARVAR Rescode = 1242 // The maximum number of semidefinite variables limit is too small.
    MSK_RES_ERR_MAXNUMQNZ Rescode = 1243 // Too small maximum number of non-zeros for the Q matrices is specified.
    MSK_RES_ERR_TOO_SMALL_MAX_NUM_NZ Rescode = 1245 // The maximum number of non-zeros specified is too small.
    MSK_RES_ERR_INVALID_IDX Rescode = 1246 // A specified index is invalid.
    MSK_RES_ERR_INVALID_MAX_NUM Rescode = 1247 // A specified index is invalid.
    MSK_RES_ERR_UNALLOWED_WHICHSOL Rescode = 1248 // The value of whichsol is not allowed.
    MSK_RES_ERR_NUMCONLIM Rescode = 1250 // Maximum number of constraints limit is exceeded.
    MSK_RES_ERR_NUMVARLIM Rescode = 1251 // Maximum number of variables limit is exceeded.
    MSK_RES_ERR_TOO_SMALL_MAXNUMANZ Rescode = 1252 // Too small maximum number of non-zeros in A specified.
    MSK_RES_ERR_INV_APTRE Rescode = 1253 // aptre[j] is strictly smaller than aptrb[j] for some j.
    MSK_RES_ERR_MUL_A_ELEMENT Rescode = 1254 // An element in A is defined multiple times.
    MSK_RES_ERR_INV_BK Rescode = 1255 // Invalid bound key.
    MSK_RES_ERR_INV_BKC Rescode = 1256 // Invalid bound key is specified for a constraint.
    MSK_RES_ERR_INV_BKX Rescode = 1257 // An invalid bound key is specified for a variable.
    MSK_RES_ERR_INV_VAR_TYPE Rescode = 1258 // An invalid variable type is specified for a variable.
    MSK_RES_ERR_SOLVER_PROBTYPE Rescode = 1259 // Problem type does not match the chosen optimizer.
    MSK_RES_ERR_OBJECTIVE_RANGE Rescode = 1260 // Empty objective range.
    MSK_RES_ERR_INV_RESCODE Rescode = 1261 // Invalid response code.
    MSK_RES_ERR_INV_IINF Rescode = 1262 // Invalid integer information item.
    MSK_RES_ERR_INV_LIINF Rescode = 1263 // Invalid long integer information item.
    MSK_RES_ERR_INV_DINF Rescode = 1264 // Invalid double information item.
    MSK_RES_ERR_BASIS Rescode = 1266 // Invalid basis is specified.
    MSK_RES_ERR_INV_SKC Rescode = 1267 // Invalid value in skc encountered.
    MSK_RES_ERR_INV_SKX Rescode = 1268 // Invalid value in skx encountered.
    MSK_RES_ERR_INV_SK_STR Rescode = 1269 // Invalid status key string encountered.
    MSK_RES_ERR_INV_SK Rescode = 1270 // Invalid status key code encountered.
    MSK_RES_ERR_INV_CONE_TYPE_STR Rescode = 1271 // Invalid cone type string encountered.
    MSK_RES_ERR_INV_CONE_TYPE Rescode = 1272 // Invalid cone type code encountered.
    MSK_RES_ERR_INV_SKN Rescode = 1274 // Invalid value in skn encountered.
    MSK_RES_ERR_INVALID_SURPLUS Rescode = 1275 // Invalid surplus.
    MSK_RES_ERR_INV_NAME_ITEM Rescode = 1280 // An invalid name item code is used.
    MSK_RES_ERR_PRO_ITEM Rescode = 1281 // An invalid problem item is used.
    MSK_RES_ERR_INVALID_FORMAT_TYPE Rescode = 1283 // Invalid format type.
    MSK_RES_ERR_FIRSTI Rescode = 1285 // Invalid firsti.
    MSK_RES_ERR_LASTI Rescode = 1286 // Invalid lasti.
    MSK_RES_ERR_FIRSTJ Rescode = 1287 // Invalid firstj.
    MSK_RES_ERR_LASTJ Rescode = 1288 // Invalid lastj.
    MSK_RES_ERR_MAX_LEN_IS_TOO_SMALL Rescode = 1289 // A maximum length that is too small has been specified.
    MSK_RES_ERR_NONLINEAR_EQUALITY Rescode = 1290 // The model contains a nonlinear equality.
    MSK_RES_ERR_NONCONVEX Rescode = 1291 // The optimization problem is nonconvex.
    MSK_RES_ERR_NONLINEAR_RANGED Rescode = 1292 // The problem contains a nonlinear constraint with inite lower and upper bound.
    MSK_RES_ERR_CON_Q_NOT_PSD Rescode = 1293 // The quadratic constraint matrix is not PSD.
    MSK_RES_ERR_CON_Q_NOT_NSD Rescode = 1294 // The quadratic constraint matrix is not NSD.
    MSK_RES_ERR_OBJ_Q_NOT_PSD Rescode = 1295 // The quadratic coefficient matrix in the objective is not PSD.
    MSK_RES_ERR_OBJ_Q_NOT_NSD Rescode = 1296 // The quadratic coefficient matrix in the objective is not NSD.
    MSK_RES_ERR_ARGUMENT_PERM_ARRAY Rescode = 1299 // An invalid permutation array is specified.
    MSK_RES_ERR_CONE_INDEX Rescode = 1300 // An index of a non-existing cone has been specified.
    MSK_RES_ERR_CONE_SIZE Rescode = 1301 // A cone with incorrect number of members is specified.
    MSK_RES_ERR_CONE_OVERLAP Rescode = 1302 // One or more of variables in the cone to be added is already member of another cone.
    MSK_RES_ERR_CONE_REP_VAR Rescode = 1303 // A variable is included multiple times in the cone.
    MSK_RES_ERR_MAXNUMCONE Rescode = 1304 // The value specified for maxnumcone is too small.
    MSK_RES_ERR_CONE_TYPE Rescode = 1305 // Invalid cone type specified.
    MSK_RES_ERR_CONE_TYPE_STR Rescode = 1306 // Invalid cone type specified.
    MSK_RES_ERR_CONE_OVERLAP_APPEND Rescode = 1307 // The cone to be appended has one variable which is already member of another cone.
    MSK_RES_ERR_REMOVE_CONE_VARIABLE Rescode = 1310 // A variable cannot be removed because it will make a cone invalid.
    MSK_RES_ERR_APPENDING_TOO_BIG_CONE Rescode = 1311 // Trying to append a too big cone.
    MSK_RES_ERR_CONE_PARAMETER Rescode = 1320 // An invalid cone parameter.
    MSK_RES_ERR_SOL_FILE_INVALID_NUMBER Rescode = 1350 // An invalid number is specified in a solution file.
    MSK_RES_ERR_HUGE_C Rescode = 1375 // A huge value in absolute size is specified for an objective coefficient.
    MSK_RES_ERR_HUGE_AIJ Rescode = 1380 // A numerically huge value is specified for an element in A.
    MSK_RES_ERR_DUPLICATE_AIJ Rescode = 1385 // An element in the A matrix is specified twice.
    MSK_RES_ERR_LOWER_BOUND_IS_A_NAN Rescode = 1390 // The lower bound specified is not a number (nan).
    MSK_RES_ERR_UPPER_BOUND_IS_A_NAN Rescode = 1391 // The upper bound specified is not a number (nan).
    MSK_RES_ERR_INFINITE_BOUND Rescode = 1400 // A numerically huge bound value is specified.
    MSK_RES_ERR_INV_QOBJ_SUBI Rescode = 1401 // Invalid value %d at qosubi.
    MSK_RES_ERR_INV_QOBJ_SUBJ Rescode = 1402 // Invalid value in qosubj.
    MSK_RES_ERR_INV_QOBJ_VAL Rescode = 1403 // Invalid value in qoval.
    MSK_RES_ERR_INV_QCON_SUBK Rescode = 1404 // Invalid value in qcsubk.
    MSK_RES_ERR_INV_QCON_SUBI Rescode = 1405 // Invalid value in qcsubi.
    MSK_RES_ERR_INV_QCON_SUBJ Rescode = 1406 // Invalid value in qcsubj.
    MSK_RES_ERR_INV_QCON_VAL Rescode = 1407 // Invalid value in qcval.
    MSK_RES_ERR_QCON_SUBI_TOO_SMALL Rescode = 1408 // Invalid value in qcsubi.
    MSK_RES_ERR_QCON_SUBI_TOO_LARGE Rescode = 1409 // Invalid value in qcsubi.
    MSK_RES_ERR_QOBJ_UPPER_TRIANGLE Rescode = 1415 // An element in the upper triangle of the quadratic term in the objective is specified.
    MSK_RES_ERR_QCON_UPPER_TRIANGLE Rescode = 1417 // An element in the upper triangle of the quadratic term in a constraint.
    MSK_RES_ERR_FIXED_BOUND_VALUES Rescode = 1420 // A fixed constraint/variable has been specified using the bound keys but the numerical bounds are different.
    MSK_RES_ERR_TOO_SMALL_A_TRUNCATION_VALUE Rescode = 1421 // A too small value for the A trucation value is specified.
    MSK_RES_ERR_INVALID_OBJECTIVE_SENSE Rescode = 1445 // An invalid objective sense is specified.
    MSK_RES_ERR_UNDEFINED_OBJECTIVE_SENSE Rescode = 1446 // The objective sense has not been specified before the optimization.
    MSK_RES_ERR_Y_IS_UNDEFINED Rescode = 1449 // The solution item y is undefined.
    MSK_RES_ERR_NAN_IN_DOUBLE_DATA Rescode = 1450 // An invalid floating value was used in some double data.
    MSK_RES_ERR_INF_IN_DOUBLE_DATA Rescode = 1451 // An infinite floating value was used in some double data.
    MSK_RES_ERR_NAN_IN_BLC Rescode = 1461 // blc contains an invalid floating point value, i.e. a NaN.
    MSK_RES_ERR_NAN_IN_BUC Rescode = 1462 // buc contains an invalid floating point value, i.e. a NaN.
    MSK_RES_ERR_INVALID_CFIX Rescode = 1469 // An invalid fixed term in the objective is speficied.
    MSK_RES_ERR_NAN_IN_C Rescode = 1470 // c contains an invalid floating point value, i.e. a NaN.
    MSK_RES_ERR_NAN_IN_BLX Rescode = 1471 // blx contains an invalid floating point value, i.e. a NaN.
    MSK_RES_ERR_NAN_IN_BUX Rescode = 1472 // bux contains an invalid floating point value, i.e. a NaN.
    MSK_RES_ERR_INVALID_AIJ Rescode = 1473 // a[i,j] contains an invalid floating point value, i.e. a NaN or an infinite value.
    MSK_RES_ERR_INVALID_CJ Rescode = 1474 // c[j] contains an invalid floating point value, i.e. a NaN or an infinite value.
    MSK_RES_ERR_SYM_MAT_INVALID Rescode = 1480 // A symmetric matrix contains an invalid floating point value, i.e. a NaN or an infinite value.
    MSK_RES_ERR_SYM_MAT_HUGE Rescode = 1482 // A numerically huge value is specified for an element in E.
    MSK_RES_ERR_INV_PROBLEM Rescode = 1500 // Invalid problem type.
    MSK_RES_ERR_MIXED_CONIC_AND_NL Rescode = 1501 // The problem contains both conic and nonlinear constraints.
    MSK_RES_ERR_GLOBAL_INV_CONIC_PROBLEM Rescode = 1503 // The global optimizer can only be applied to problems without semidefinite variables.
    MSK_RES_ERR_INV_OPTIMIZER Rescode = 1550 // An invalid optimizer has been chosen for the problem.
    MSK_RES_ERR_MIO_NO_OPTIMIZER Rescode = 1551 // No optimizer is available for the current class of integer optimization problems.
    MSK_RES_ERR_NO_OPTIMIZER_VAR_TYPE Rescode = 1552 // No optimizer is available for this class of optimization problems.
    MSK_RES_ERR_FINAL_SOLUTION Rescode = 1560 // An error occurred during the solution finalization.
    MSK_RES_ERR_FIRST Rescode = 1570 // Invalid first.
    MSK_RES_ERR_LAST Rescode = 1571 // Invalid last.
    MSK_RES_ERR_SLICE_SIZE Rescode = 1572 // Invalid slice size specified.
    MSK_RES_ERR_NEGATIVE_SURPLUS Rescode = 1573 // Negative surplus.
    MSK_RES_ERR_NEGATIVE_APPEND Rescode = 1578 // Cannot append a negative number.
    MSK_RES_ERR_POSTSOLVE Rescode = 1580 // An error occurred during the postsolve.
    MSK_RES_ERR_OVERFLOW Rescode = 1590 // A computation produced an overflow.
    MSK_RES_ERR_NO_BASIS_SOL Rescode = 1600 // No basic solution is defined.
    MSK_RES_ERR_BASIS_FACTOR Rescode = 1610 // The factorization of the basis is invalid.
    MSK_RES_ERR_BASIS_SINGULAR Rescode = 1615 // The basis is singular.
    MSK_RES_ERR_FACTOR Rescode = 1650 // An error occurred while factorizing a matrix.
    MSK_RES_ERR_FEASREPAIR_CANNOT_RELAX Rescode = 1700 // An optimization problem cannot be relaxed.
    MSK_RES_ERR_FEASREPAIR_SOLVING_RELAXED Rescode = 1701 // The relaxed problem could not be solved to optimality.
    MSK_RES_ERR_FEASREPAIR_INCONSISTENT_BOUND Rescode = 1702 // The upper bound is less than the lower bound for a variable or a constraint.
    MSK_RES_ERR_REPAIR_INVALID_PROBLEM Rescode = 1710 // The feasibility repair does not support the specified problem type.
    MSK_RES_ERR_REPAIR_OPTIMIZATION_FAILED Rescode = 1711 // Computation the optimal relaxation failed.
    MSK_RES_ERR_NAME_MAX_LEN Rescode = 1750 // A name is longer than the buffer that is supposed to hold it.
    MSK_RES_ERR_NAME_IS_NULL Rescode = 1760 // The name buffer is a null pointer.
    MSK_RES_ERR_INVALID_COMPRESSION Rescode = 1800 // Invalid compression type.
    MSK_RES_ERR_INVALID_IOMODE Rescode = 1801 // Invalid io mode.
    MSK_RES_ERR_NO_PRIMAL_INFEAS_CER Rescode = 2000 // A certificate of primal infeasibility is not available.
    MSK_RES_ERR_NO_DUAL_INFEAS_CER Rescode = 2001 // A certificate of dual infeasibility is not available.
    MSK_RES_ERR_NO_SOLUTION_IN_CALLBACK Rescode = 2500 // The required solution is not available.
    MSK_RES_ERR_INV_MARKI Rescode = 2501 // Invalid value in marki.
    MSK_RES_ERR_INV_MARKJ Rescode = 2502 // Invalid value in markj.
    MSK_RES_ERR_INV_NUMI Rescode = 2503 // Invalid numi.
    MSK_RES_ERR_INV_NUMJ Rescode = 2504 // Invalid numj.
    MSK_RES_ERR_TASK_INCOMPATIBLE Rescode = 2560 // The Task file is incompatible with this platform.
    MSK_RES_ERR_TASK_INVALID Rescode = 2561 // The Task file is invalid.
    MSK_RES_ERR_TASK_WRITE Rescode = 2562 // Failed to write the task file.
    MSK_RES_ERR_LU_MAX_NUM_TRIES Rescode = 2800 // Could not compute the LU factors of the matrix within the maximum number of allowed tries.
    MSK_RES_ERR_INVALID_UTF8 Rescode = 2900 // An invalid UTF8 string is encountered.
    MSK_RES_ERR_INVALID_WCHAR Rescode = 2901 // An invalid wchar string is encountered.
    MSK_RES_ERR_NO_DUAL_FOR_ITG_SOL Rescode = 2950 // No dual information is available for the integer solution.
    MSK_RES_ERR_NO_SNX_FOR_BAS_SOL Rescode = 2953 // snx is not available for the basis solution.
    MSK_RES_ERR_INTERNAL Rescode = 3000 // An internal error occurred.
    MSK_RES_ERR_API_ARRAY_TOO_SMALL Rescode = 3001 // An input array was too short.
    MSK_RES_ERR_API_CB_CONNECT Rescode = 3002 // Failed to connect a callback object.
    MSK_RES_ERR_API_FATAL_ERROR Rescode = 3005 // An internal error occurred in the API. Please report this problem.
    MSK_RES_ERR_SEN_FORMAT Rescode = 3050 // Syntax error in sensitivity analysis file.
    MSK_RES_ERR_SEN_UNDEF_NAME Rescode = 3051 // An undefined name was encountered in the sensitivity analysis file.
    MSK_RES_ERR_SEN_INDEX_RANGE Rescode = 3052 // Index out of range in the sensitivity analysis file.
    MSK_RES_ERR_SEN_BOUND_INVALID_UP Rescode = 3053 // Analysis of upper bound requested for an index, where no upper bound exists.
    MSK_RES_ERR_SEN_BOUND_INVALID_LO Rescode = 3054 // Analysis of lower bound requested for an index, where no lower bound exists.
    MSK_RES_ERR_SEN_INDEX_INVALID Rescode = 3055 // Invalid range given in the sensitivity file.
    MSK_RES_ERR_SEN_INVALID_REGEXP Rescode = 3056 // Syntax error in regexp or regexp longer than 1024.
    MSK_RES_ERR_SEN_SOLUTION_STATUS Rescode = 3057 // No optimal solution found to the original problem given for sensitivity analysis.
    MSK_RES_ERR_SEN_NUMERICAL Rescode = 3058 // Numerical difficulties encountered performing the sensitivity analysis.
    MSK_RES_ERR_SEN_UNHANDLED_PROBLEM_TYPE Rescode = 3080 // Sensitivity analysis cannot be performed for the specified problem.
    MSK_RES_ERR_UNB_STEP_SIZE Rescode = 3100 // A step-size in an optimizer was unexpectedly unbounded.
    MSK_RES_ERR_IDENTICAL_TASKS Rescode = 3101 // Some tasks related to this function call were identical. Unique tasks were expected.
    MSK_RES_ERR_AD_INVALID_CODELIST Rescode = 3102 // The code list data was invalid.
    MSK_RES_ERR_INTERNAL_TEST_FAILED Rescode = 3500 // An internal unit test function failed.
    MSK_RES_ERR_XML_INVALID_PROBLEM_TYPE Rescode = 3600 // The problem type is not supported by the XML format.
    MSK_RES_ERR_INVALID_AMPL_STUB Rescode = 3700 // Invalid AMPL stub.
    MSK_RES_ERR_INT64_TO_INT32_CAST Rescode = 3800 // A 64 bit integer could not be cast to a 32 bit integer.
    MSK_RES_ERR_SIZE_LICENSE_NUMCORES Rescode = 3900 // The computer contains more cpu cores than the license allows for.
    MSK_RES_ERR_INFEAS_UNDEFINED Rescode = 3910 // The requested value is not defined for this solution type.
    MSK_RES_ERR_NO_BARX_FOR_SOLUTION Rescode = 3915 // There is no barx available for the solution specified.
    MSK_RES_ERR_NO_BARS_FOR_SOLUTION Rescode = 3916 // There is no bars available for the solution specified.
    MSK_RES_ERR_BAR_VAR_DIM Rescode = 3920 // The dimension of a symmetric matrix variable has to be greater than 0.
    MSK_RES_ERR_SYM_MAT_INVALID_ROW_INDEX Rescode = 3940 // A row index specified for sparse symmetric matrix is invalid.
    MSK_RES_ERR_SYM_MAT_INVALID_COL_INDEX Rescode = 3941 // A column index specified for sparse symmetric matrix is invalid.
    MSK_RES_ERR_SYM_MAT_NOT_LOWER_TRINGULAR Rescode = 3942 // Only the lower triangular part of sparse symmetric matrix should be specified.
    MSK_RES_ERR_SYM_MAT_INVALID_VALUE Rescode = 3943 // The numerical value specified in a sparse symmetric matrix is not a floating point value.
    MSK_RES_ERR_SYM_MAT_DUPLICATE Rescode = 3944 // A value in a symmetric matric as been specified more than once.
    MSK_RES_ERR_INVALID_SYM_MAT_DIM Rescode = 3950 // A sparse symmetric matrix of invalid dimension is specified.
    MSK_RES_ERR_API_INTERNAL Rescode = 3999 // An internal fatal error occurred in an interface function.
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_SYM_MAT Rescode = 4000 // The file format does not support a problem with symmetric matrix variables.
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_CFIX Rescode = 4001 // The file format does not support a problem with nonzero fixed term in c.
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_RANGED_CONSTRAINTS Rescode = 4002 // The file format does not support a problem with ranged constraints.
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_FREE_CONSTRAINTS Rescode = 4003 // The file format does not support a problem with free constraints.
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_CONES Rescode = 4005 // The file format does not support a problem with the simple cones (deprecated).
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_QUADRATIC_TERMS Rescode = 4006 // The file format does not support a problem with quadratic terms.
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_NONLINEAR Rescode = 4010 // The file format does not support a problem with nonlinear terms.
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_DISJUNCTIVE_CONSTRAINTS Rescode = 4011 // The file format does not support a problem with disjunctive constraints.
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_AFFINE_CONIC_CONSTRAINTS Rescode = 4012 // The file format does not support a problem with affine conic constraints.
    MSK_RES_ERR_DUPLICATE_CONSTRAINT_NAMES Rescode = 4500 // Two constraint names are identical.
    MSK_RES_ERR_DUPLICATE_VARIABLE_NAMES Rescode = 4501 // Two variable names are identical.
    MSK_RES_ERR_DUPLICATE_BARVARIABLE_NAMES Rescode = 4502 // Two barvariable names are identical.
    MSK_RES_ERR_DUPLICATE_CONE_NAMES Rescode = 4503 // Two cone names are identical.
    MSK_RES_ERR_DUPLICATE_DOMAIN_NAMES Rescode = 4504 // Two domain names are identical.
    MSK_RES_ERR_DUPLICATE_DJC_NAMES Rescode = 4505 // Two disjunctive constraint names are identical.
    MSK_RES_ERR_NON_UNIQUE_ARRAY Rescode = 5000 // An array does not contain unique elements.
    MSK_RES_ERR_ARGUMENT_IS_TOO_SMALL Rescode = 5004 // The value of a function argument is too small.
    MSK_RES_ERR_ARGUMENT_IS_TOO_LARGE Rescode = 5005 // The value of a function argument is too large.
    MSK_RES_ERR_MIO_INTERNAL Rescode = 5010 // A fatal error occurred in the mixed integer optimizer.  Please contact MOSEK support.
    MSK_RES_ERR_INVALID_PROBLEM_TYPE Rescode = 6000 // An invalid problem type.
    MSK_RES_ERR_UNHANDLED_SOLUTION_STATUS Rescode = 6010 // Unhandled solution status.
    MSK_RES_ERR_UPPER_TRIANGLE Rescode = 6020 // An element in the upper triangle of a lower triangular matrix is specified.
    MSK_RES_ERR_LAU_SINGULAR_MATRIX Rescode = 7000 // A matrix is singular.
    MSK_RES_ERR_LAU_NOT_POSITIVE_DEFINITE Rescode = 7001 // A matrix is not positive definite.
    MSK_RES_ERR_LAU_INVALID_LOWER_TRIANGULAR_MATRIX Rescode = 7002 // An invalid lower triangular matrix.
    MSK_RES_ERR_LAU_UNKNOWN Rescode = 7005 // An unknown error.
    MSK_RES_ERR_LAU_ARG_M Rescode = 7010 // Invalid argument m.
    MSK_RES_ERR_LAU_ARG_N Rescode = 7011 // Invalid argument n.
    MSK_RES_ERR_LAU_ARG_K Rescode = 7012 // Invalid argument k.
    MSK_RES_ERR_LAU_ARG_TRANSA Rescode = 7015 // Invalid argument transa.
    MSK_RES_ERR_LAU_ARG_TRANSB Rescode = 7016 // Invalid argument transb.
    MSK_RES_ERR_LAU_ARG_UPLO Rescode = 7017 // Invalid argument uplo.
    MSK_RES_ERR_LAU_ARG_TRANS Rescode = 7018 // Invalid argument trans.
    MSK_RES_ERR_LAU_INVALID_SPARSE_SYMMETRIC_MATRIX Rescode = 7019 // An invalid sparse symmetric matrix is specfified.
    MSK_RES_ERR_CBF_PARSE Rescode = 7100 // An error occurred while parsing an CBF file.
    MSK_RES_ERR_CBF_OBJ_SENSE Rescode = 7101 // An invalid objective sense is specified.
    MSK_RES_ERR_CBF_NO_VARIABLES Rescode = 7102 // An invalid objective sense is specified.
    MSK_RES_ERR_CBF_TOO_MANY_CONSTRAINTS Rescode = 7103 // Too many constraints specified.
    MSK_RES_ERR_CBF_TOO_MANY_VARIABLES Rescode = 7104 // Too many variables specified.
    MSK_RES_ERR_CBF_NO_VERSION_SPECIFIED Rescode = 7105 // No version specified.
    MSK_RES_ERR_CBF_SYNTAX Rescode = 7106 // Invalid syntax.
    MSK_RES_ERR_CBF_DUPLICATE_OBJ Rescode = 7107 // Duplicate OBJ keyword.
    MSK_RES_ERR_CBF_DUPLICATE_CON Rescode = 7108 // Duplicate CON keyword.
    MSK_RES_ERR_CBF_DUPLICATE_VAR Rescode = 7110 // Duplicate VAR keyword.
    MSK_RES_ERR_CBF_DUPLICATE_INT Rescode = 7111 // Duplicate INT keyword.
    MSK_RES_ERR_CBF_INVALID_VAR_TYPE Rescode = 7112 // Invalid variable type.
    MSK_RES_ERR_CBF_INVALID_CON_TYPE Rescode = 7113 // Invalid constraint type.
    MSK_RES_ERR_CBF_INVALID_DOMAIN_DIMENSION Rescode = 7114 // Invalid domain dimension.
    MSK_RES_ERR_CBF_DUPLICATE_OBJACOORD Rescode = 7115 // Duplicate index in OBJCOORD.
    MSK_RES_ERR_CBF_DUPLICATE_BCOORD Rescode = 7116 // Duplicate index in BCOORD.
    MSK_RES_ERR_CBF_DUPLICATE_ACOORD Rescode = 7117 // Duplicate index in ACOORD.
    MSK_RES_ERR_CBF_TOO_FEW_VARIABLES Rescode = 7118 // Too few variables defined.
    MSK_RES_ERR_CBF_TOO_FEW_CONSTRAINTS Rescode = 7119 // Too few constraints defined.
    MSK_RES_ERR_CBF_TOO_FEW_INTS Rescode = 7120 // Too ints specified.
    MSK_RES_ERR_CBF_TOO_MANY_INTS Rescode = 7121 // Too ints specified.
    MSK_RES_ERR_CBF_INVALID_INT_INDEX Rescode = 7122 // Invalid INT index.
    MSK_RES_ERR_CBF_UNSUPPORTED Rescode = 7123 // Unsupported feature is present.
    MSK_RES_ERR_CBF_DUPLICATE_PSDVAR Rescode = 7124 // Duplicate PSDVAR keyword.
    MSK_RES_ERR_CBF_INVALID_PSDVAR_DIMENSION Rescode = 7125 // Invalid PSDVAR dimension.
    MSK_RES_ERR_CBF_TOO_FEW_PSDVAR Rescode = 7126 // Too few variables defined.
    MSK_RES_ERR_CBF_INVALID_EXP_DIMENSION Rescode = 7127 // Invalid dimension of a exponential cone.
    MSK_RES_ERR_CBF_DUPLICATE_POW_CONES Rescode = 7130 // Multiple POWCONES specified.
    MSK_RES_ERR_CBF_DUPLICATE_POW_STAR_CONES Rescode = 7131 // Multiple POW*CONES specified.
    MSK_RES_ERR_CBF_INVALID_POWER Rescode = 7132 // Invalid power specified.
    MSK_RES_ERR_CBF_POWER_CONE_IS_TOO_LONG Rescode = 7133 // Power cone is too long.
    MSK_RES_ERR_CBF_INVALID_POWER_CONE_INDEX Rescode = 7134 // Invalid power cone index.
    MSK_RES_ERR_CBF_INVALID_POWER_STAR_CONE_INDEX Rescode = 7135 // Invalid power star cone index.
    MSK_RES_ERR_CBF_UNHANDLED_POWER_CONE_TYPE Rescode = 7136 // An unhandled power cone type.
    MSK_RES_ERR_CBF_UNHANDLED_POWER_STAR_CONE_TYPE Rescode = 7137 // An unhandled power star cone type.
    MSK_RES_ERR_CBF_POWER_CONE_MISMATCH Rescode = 7138 // The power cone does not match with it definition.
    MSK_RES_ERR_CBF_POWER_STAR_CONE_MISMATCH Rescode = 7139 // The power star cone does not match with it definition.
    MSK_RES_ERR_CBF_INVALID_NUMBER_OF_CONES Rescode = 7140 // Invalid number of cones.
    MSK_RES_ERR_CBF_INVALID_DIMENSION_OF_CONES Rescode = 7141 // Invalid number of cones.
    MSK_RES_ERR_CBF_INVALID_NUM_OBJACOORD Rescode = 7150 // Invalid number of OBJACOORD.
    MSK_RES_ERR_CBF_INVALID_NUM_OBJFCOORD Rescode = 7151 // Invalid number of OBJFCOORD.
    MSK_RES_ERR_CBF_INVALID_NUM_ACOORD Rescode = 7152 // Invalid number of ACOORD.
    MSK_RES_ERR_CBF_INVALID_NUM_BCOORD Rescode = 7153 // Invalid number of BCOORD.
    MSK_RES_ERR_CBF_INVALID_NUM_FCOORD Rescode = 7155 // Invalid number of FCOORD.
    MSK_RES_ERR_CBF_INVALID_NUM_HCOORD Rescode = 7156 // Invalid number of HCOORD.
    MSK_RES_ERR_CBF_INVALID_NUM_DCOORD Rescode = 7157 // Invalid number of DCOORD.
    MSK_RES_ERR_CBF_EXPECTED_A_KEYWORD Rescode = 7158 // Expected a key word.
    MSK_RES_ERR_CBF_INVALID_NUM_PSDCON Rescode = 7200 // Invalid number of PSDCON.
    MSK_RES_ERR_CBF_DUPLICATE_PSDCON Rescode = 7201 // Duplicate CON keyword.
    MSK_RES_ERR_CBF_INVALID_DIMENSION_OF_PSDCON Rescode = 7202 // Invalid PSDCON dimension.
    MSK_RES_ERR_CBF_INVALID_PSDCON_INDEX Rescode = 7203 // Invalid PSDCON index.
    MSK_RES_ERR_CBF_INVALID_PSDCON_VARIABLE_INDEX Rescode = 7204 // Invalid PSDCON index.
    MSK_RES_ERR_CBF_INVALID_PSDCON_BLOCK_INDEX Rescode = 7205 // Invalid PSDCON index.
    MSK_RES_ERR_CBF_UNSUPPORTED_CHANGE Rescode = 7210 // The CHANGE section is not supported.
    MSK_RES_ERR_MIO_INVALID_ROOT_OPTIMIZER Rescode = 7700 // An invalid root optimizer was selected for the problem type.
    MSK_RES_ERR_MIO_INVALID_NODE_OPTIMIZER Rescode = 7701 // An invalid node optimizer was selected for the problem type.
    MSK_RES_ERR_MPS_WRITE_CPLEX_INVALID_CONE_TYPE Rescode = 7750 // An invalid cone type occurs when writing a CPLEX formatted MPS file.
    MSK_RES_ERR_TOCONIC_CONSTR_Q_NOT_PSD Rescode = 7800 // The matrix defining the quadratric part of constraint is not positive semidefinite.
    MSK_RES_ERR_TOCONIC_CONSTRAINT_FX Rescode = 7801 // The quadratic constraint is an equality, thus not convex.
    MSK_RES_ERR_TOCONIC_CONSTRAINT_RA Rescode = 7802 // The quadratic constraint has finite lower and upper bound, and therefore it is not convex.
    MSK_RES_ERR_TOCONIC_CONSTR_NOT_CONIC Rescode = 7803 // The constraint is not conic representable.
    MSK_RES_ERR_TOCONIC_OBJECTIVE_NOT_PSD Rescode = 7804 // The matrix defining the quadratric part of the objective function is not positive semidefinite.
    MSK_RES_ERR_SERVER_CONNECT Rescode = 8000 // Failed to connect to remote solver server.
    MSK_RES_ERR_SERVER_PROTOCOL Rescode = 8001 // Unexpected message or data from solver server.
    MSK_RES_ERR_SERVER_STATUS Rescode = 8002 // Server returned non-ok status code
    MSK_RES_ERR_SERVER_TOKEN Rescode = 8003 // Invalid job ID
    MSK_RES_ERR_SERVER_ADDRESS Rescode = 8004 // Invalid address
    MSK_RES_ERR_SERVER_CERTIFICATE Rescode = 8005 // Invalid TLS certificate format or path
    MSK_RES_ERR_SERVER_TLS_CLIENT Rescode = 8006 // Failed to create TLS client
    MSK_RES_ERR_SERVER_ACCESS_TOKEN Rescode = 8007 // Invalid access token
    MSK_RES_ERR_SERVER_PROBLEM_SIZE Rescode = 8008 // The problem is too large.
    MSK_RES_ERR_DUPLICATE_INDEX_IN_A_SPARSE_MATRIX Rescode = 20050 // An element in a sparse matrix is specified twice.
    MSK_RES_ERR_DUPLICATE_INDEX_IN_AFEIDX_LIST Rescode = 20060 // An index is specified twice in an affine expression list.
    MSK_RES_ERR_DUPLICATE_FIJ Rescode = 20100 // An element in the F matrix is specified twice.
    MSK_RES_ERR_INVALID_FIJ Rescode = 20101 // f[i,j] contains an invalid floating point value, i.e. a NaN or an infinite value.
    MSK_RES_ERR_HUGE_FIJ Rescode = 20102 // A numerically huge value is specified for an element in F.
    MSK_RES_ERR_INVALID_G Rescode = 20103 // g contains an invalid floating point value, i.e. a NaN or an infinite value.
    MSK_RES_ERR_INVALID_B Rescode = 20150 // b contains an invalid floating point value, i.e. a NaN or an infinite value.
    MSK_RES_ERR_DOMAIN_INVALID_INDEX Rescode = 20400 // A domain index is invalid.
    MSK_RES_ERR_DOMAIN_DIMENSION Rescode = 20401 // A domain dimension is invalid.
    MSK_RES_ERR_DOMAIN_DIMENSION_PSD Rescode = 20402 // A PSD domain dimension is invalid.
    MSK_RES_ERR_NOT_POWER_DOMAIN Rescode = 20403 // The function is only applicable to primal and dual power cone domains.
    MSK_RES_ERR_DOMAIN_POWER_INVALID_ALPHA Rescode = 20404 // Alpha contains an invalid floating point value, i.e. a NaN or an infinite value.
    MSK_RES_ERR_DOMAIN_POWER_NEGATIVE_ALPHA Rescode = 20405 // Alpha contains a negative value or zero.
    MSK_RES_ERR_DOMAIN_POWER_NLEFT Rescode = 20406 // The value of nleft is too small or too large.
    MSK_RES_ERR_AFE_INVALID_INDEX Rescode = 20500 // An affine expression index is invalid.
    MSK_RES_ERR_ACC_INVALID_INDEX Rescode = 20600 // A affine conic constraint index is invalid.
    MSK_RES_ERR_ACC_INVALID_ENTRY_INDEX Rescode = 20601 // The index of an element in an affine conic constraint is invalid.
    MSK_RES_ERR_ACC_AFE_DOMAIN_MISMATCH Rescode = 20602 // There is a mismatch between between the number of affine expressions and total dimension of the domain(s).
    MSK_RES_ERR_DJC_INVALID_INDEX Rescode = 20700 // A disjunctive constraint index is invalid.
    MSK_RES_ERR_DJC_UNSUPPORTED_DOMAIN_TYPE Rescode = 20701 // An unsupported domain type has been used in a disjunctive constraint.
    MSK_RES_ERR_DJC_AFE_DOMAIN_MISMATCH Rescode = 20702 // There is a mismatch between the number of affine expressions and total dimension of the domain(s).
    MSK_RES_ERR_DJC_INVALID_TERM_SIZE Rescode = 20703 // A termize is invalid.
    MSK_RES_ERR_DJC_DOMAIN_TERMSIZE_MISMATCH Rescode = 20704 // There is a mismatch between the number of domains and the term sizes.
    MSK_RES_ERR_DJC_TOTAL_NUM_TERMS_MISMATCH Rescode = 20705 // There total number of terms in all domains does not match.
    MSK_RES_ERR_UNDEF_SOLUTION Rescode = 22000 // The required solution is not defined.
    MSK_RES_ERR_NO_DOTY Rescode = 22010 // No doty is available.
    MSK_RES_TRM_MAX_ITERATIONS Rescode = 100000 // The optimizer terminated at the maximum number of iterations.
    MSK_RES_TRM_MAX_TIME Rescode = 100001 // The optimizer terminated at the maximum amount of time.
    MSK_RES_TRM_OBJECTIVE_RANGE Rescode = 100002 // The optimizer terminated with an objective value outside the objective range.
    MSK_RES_TRM_STALL Rescode = 100006 // The optimizer is terminated due to slow progress.
    MSK_RES_TRM_USER_CALLBACK Rescode = 100007 // The user-defined progress callback function terminated the optimization.
    MSK_RES_TRM_MIO_NUM_RELAXS Rescode = 100008 // The mixed-integer optimizer terminated as the maximum number of relaxations was reached.
    MSK_RES_TRM_MIO_NUM_BRANCHES Rescode = 100009 // The mixed-integer optimizer terminated as the maximum number of branches was reached.
    MSK_RES_TRM_NUM_MAX_NUM_INT_SOLUTIONS Rescode = 100015 // The mixed-integer optimizer terminated as the maximum number of feasible solutions was reached.
    MSK_RES_TRM_MAX_NUM_SETBACKS Rescode = 100020 // The optimizer terminated as the maximum number of set-backs was reached.
    MSK_RES_TRM_NUMERICAL_PROBLEM Rescode = 100025 // The optimizer terminated due to a numerical problem.
    MSK_RES_TRM_LOST_RACE Rescode = 100027 // Lost a race.
    MSK_RES_TRM_INTERNAL Rescode = 100030 // The optimizer terminated due to some internal reason.
    MSK_RES_TRM_INTERNAL_STOP Rescode = 100031 // The optimizer terminated for internal reasons.
    MSK_RES_END Rescode = 100031
)
func (self Rescode) String() string {
  switch self {
    case MSK_RES_OK: return "MSK_RES_OK"
    case MSK_RES_WRN_OPEN_PARAM_FILE: return "MSK_RES_WRN_OPEN_PARAM_FILE"
    case MSK_RES_WRN_LARGE_BOUND: return "MSK_RES_WRN_LARGE_BOUND"
    case MSK_RES_WRN_LARGE_LO_BOUND: return "MSK_RES_WRN_LARGE_LO_BOUND"
    case MSK_RES_WRN_LARGE_UP_BOUND: return "MSK_RES_WRN_LARGE_UP_BOUND"
    case MSK_RES_WRN_LARGE_CON_FX: return "MSK_RES_WRN_LARGE_CON_FX"
    case MSK_RES_WRN_LARGE_CJ: return "MSK_RES_WRN_LARGE_CJ"
    case MSK_RES_WRN_LARGE_AIJ: return "MSK_RES_WRN_LARGE_AIJ"
    case MSK_RES_WRN_ZERO_AIJ: return "MSK_RES_WRN_ZERO_AIJ"
    case MSK_RES_WRN_NAME_MAX_LEN: return "MSK_RES_WRN_NAME_MAX_LEN"
    case MSK_RES_WRN_SPAR_MAX_LEN: return "MSK_RES_WRN_SPAR_MAX_LEN"
    case MSK_RES_WRN_MPS_SPLIT_RHS_VECTOR: return "MSK_RES_WRN_MPS_SPLIT_RHS_VECTOR"
    case MSK_RES_WRN_MPS_SPLIT_RAN_VECTOR: return "MSK_RES_WRN_MPS_SPLIT_RAN_VECTOR"
    case MSK_RES_WRN_MPS_SPLIT_BOU_VECTOR: return "MSK_RES_WRN_MPS_SPLIT_BOU_VECTOR"
    case MSK_RES_WRN_LP_OLD_QUAD_FORMAT: return "MSK_RES_WRN_LP_OLD_QUAD_FORMAT"
    case MSK_RES_WRN_LP_DROP_VARIABLE: return "MSK_RES_WRN_LP_DROP_VARIABLE"
    case MSK_RES_WRN_NZ_IN_UPR_TRI: return "MSK_RES_WRN_NZ_IN_UPR_TRI"
    case MSK_RES_WRN_DROPPED_NZ_QOBJ: return "MSK_RES_WRN_DROPPED_NZ_QOBJ"
    case MSK_RES_WRN_IGNORE_INTEGER: return "MSK_RES_WRN_IGNORE_INTEGER"
    case MSK_RES_WRN_NO_GLOBAL_OPTIMIZER: return "MSK_RES_WRN_NO_GLOBAL_OPTIMIZER"
    case MSK_RES_WRN_MIO_INFEASIBLE_FINAL: return "MSK_RES_WRN_MIO_INFEASIBLE_FINAL"
    case MSK_RES_WRN_SOL_FILTER: return "MSK_RES_WRN_SOL_FILTER"
    case MSK_RES_WRN_UNDEF_SOL_FILE_NAME: return "MSK_RES_WRN_UNDEF_SOL_FILE_NAME"
    case MSK_RES_WRN_SOL_FILE_IGNORED_CON: return "MSK_RES_WRN_SOL_FILE_IGNORED_CON"
    case MSK_RES_WRN_SOL_FILE_IGNORED_VAR: return "MSK_RES_WRN_SOL_FILE_IGNORED_VAR"
    case MSK_RES_WRN_TOO_FEW_BASIS_VARS: return "MSK_RES_WRN_TOO_FEW_BASIS_VARS"
    case MSK_RES_WRN_TOO_MANY_BASIS_VARS: return "MSK_RES_WRN_TOO_MANY_BASIS_VARS"
    case MSK_RES_WRN_LICENSE_EXPIRE: return "MSK_RES_WRN_LICENSE_EXPIRE"
    case MSK_RES_WRN_LICENSE_SERVER: return "MSK_RES_WRN_LICENSE_SERVER"
    case MSK_RES_WRN_EMPTY_NAME: return "MSK_RES_WRN_EMPTY_NAME"
    case MSK_RES_WRN_USING_GENERIC_NAMES: return "MSK_RES_WRN_USING_GENERIC_NAMES"
    case MSK_RES_WRN_INVALID_MPS_NAME: return "MSK_RES_WRN_INVALID_MPS_NAME"
    case MSK_RES_WRN_INVALID_MPS_OBJ_NAME: return "MSK_RES_WRN_INVALID_MPS_OBJ_NAME"
    case MSK_RES_WRN_LICENSE_FEATURE_EXPIRE: return "MSK_RES_WRN_LICENSE_FEATURE_EXPIRE"
    case MSK_RES_WRN_PARAM_NAME_DOU: return "MSK_RES_WRN_PARAM_NAME_DOU"
    case MSK_RES_WRN_PARAM_NAME_INT: return "MSK_RES_WRN_PARAM_NAME_INT"
    case MSK_RES_WRN_PARAM_NAME_STR: return "MSK_RES_WRN_PARAM_NAME_STR"
    case MSK_RES_WRN_PARAM_STR_VALUE: return "MSK_RES_WRN_PARAM_STR_VALUE"
    case MSK_RES_WRN_PARAM_IGNORED_CMIO: return "MSK_RES_WRN_PARAM_IGNORED_CMIO"
    case MSK_RES_WRN_ZEROS_IN_SPARSE_ROW: return "MSK_RES_WRN_ZEROS_IN_SPARSE_ROW"
    case MSK_RES_WRN_ZEROS_IN_SPARSE_COL: return "MSK_RES_WRN_ZEROS_IN_SPARSE_COL"
    case MSK_RES_WRN_INCOMPLETE_LINEAR_DEPENDENCY_CHECK: return "MSK_RES_WRN_INCOMPLETE_LINEAR_DEPENDENCY_CHECK"
    case MSK_RES_WRN_ELIMINATOR_SPACE: return "MSK_RES_WRN_ELIMINATOR_SPACE"
    case MSK_RES_WRN_PRESOLVE_OUTOFSPACE: return "MSK_RES_WRN_PRESOLVE_OUTOFSPACE"
    case MSK_RES_WRN_PRESOLVE_PRIMAL_PERTUBATIONS: return "MSK_RES_WRN_PRESOLVE_PRIMAL_PERTUBATIONS"
    case MSK_RES_WRN_WRITE_CHANGED_NAMES: return "MSK_RES_WRN_WRITE_CHANGED_NAMES"
    case MSK_RES_WRN_WRITE_DISCARDED_CFIX: return "MSK_RES_WRN_WRITE_DISCARDED_CFIX"
    case MSK_RES_WRN_DUPLICATE_CONSTRAINT_NAMES: return "MSK_RES_WRN_DUPLICATE_CONSTRAINT_NAMES"
    case MSK_RES_WRN_DUPLICATE_VARIABLE_NAMES: return "MSK_RES_WRN_DUPLICATE_VARIABLE_NAMES"
    case MSK_RES_WRN_DUPLICATE_BARVARIABLE_NAMES: return "MSK_RES_WRN_DUPLICATE_BARVARIABLE_NAMES"
    case MSK_RES_WRN_DUPLICATE_CONE_NAMES: return "MSK_RES_WRN_DUPLICATE_CONE_NAMES"
    case MSK_RES_WRN_WRITE_LP_INVALID_VAR_NAMES: return "MSK_RES_WRN_WRITE_LP_INVALID_VAR_NAMES"
    case MSK_RES_WRN_WRITE_LP_DUPLICATE_VAR_NAMES: return "MSK_RES_WRN_WRITE_LP_DUPLICATE_VAR_NAMES"
    case MSK_RES_WRN_WRITE_LP_INVALID_CON_NAMES: return "MSK_RES_WRN_WRITE_LP_INVALID_CON_NAMES"
    case MSK_RES_WRN_WRITE_LP_DUPLICATE_CON_NAMES: return "MSK_RES_WRN_WRITE_LP_DUPLICATE_CON_NAMES"
    case MSK_RES_WRN_ANA_LARGE_BOUNDS: return "MSK_RES_WRN_ANA_LARGE_BOUNDS"
    case MSK_RES_WRN_ANA_C_ZERO: return "MSK_RES_WRN_ANA_C_ZERO"
    case MSK_RES_WRN_ANA_EMPTY_COLS: return "MSK_RES_WRN_ANA_EMPTY_COLS"
    case MSK_RES_WRN_ANA_CLOSE_BOUNDS: return "MSK_RES_WRN_ANA_CLOSE_BOUNDS"
    case MSK_RES_WRN_ANA_ALMOST_INT_BOUNDS: return "MSK_RES_WRN_ANA_ALMOST_INT_BOUNDS"
    case MSK_RES_WRN_NO_INFEASIBILITY_REPORT_WHEN_MATRIX_VARIABLES: return "MSK_RES_WRN_NO_INFEASIBILITY_REPORT_WHEN_MATRIX_VARIABLES"
    case MSK_RES_WRN_NO_DUALIZER: return "MSK_RES_WRN_NO_DUALIZER"
    case MSK_RES_WRN_SYM_MAT_LARGE: return "MSK_RES_WRN_SYM_MAT_LARGE"
    case MSK_RES_WRN_MODIFIED_DOUBLE_PARAMETER: return "MSK_RES_WRN_MODIFIED_DOUBLE_PARAMETER"
    case MSK_RES_WRN_LARGE_FIJ: return "MSK_RES_WRN_LARGE_FIJ"
    case MSK_RES_ERR_LICENSE: return "MSK_RES_ERR_LICENSE"
    case MSK_RES_ERR_LICENSE_EXPIRED: return "MSK_RES_ERR_LICENSE_EXPIRED"
    case MSK_RES_ERR_LICENSE_VERSION: return "MSK_RES_ERR_LICENSE_VERSION"
    case MSK_RES_ERR_LICENSE_OLD_SERVER_VERSION: return "MSK_RES_ERR_LICENSE_OLD_SERVER_VERSION"
    case MSK_RES_ERR_SIZE_LICENSE: return "MSK_RES_ERR_SIZE_LICENSE"
    case MSK_RES_ERR_PROB_LICENSE: return "MSK_RES_ERR_PROB_LICENSE"
    case MSK_RES_ERR_FILE_LICENSE: return "MSK_RES_ERR_FILE_LICENSE"
    case MSK_RES_ERR_MISSING_LICENSE_FILE: return "MSK_RES_ERR_MISSING_LICENSE_FILE"
    case MSK_RES_ERR_SIZE_LICENSE_CON: return "MSK_RES_ERR_SIZE_LICENSE_CON"
    case MSK_RES_ERR_SIZE_LICENSE_VAR: return "MSK_RES_ERR_SIZE_LICENSE_VAR"
    case MSK_RES_ERR_SIZE_LICENSE_INTVAR: return "MSK_RES_ERR_SIZE_LICENSE_INTVAR"
    case MSK_RES_ERR_OPTIMIZER_LICENSE: return "MSK_RES_ERR_OPTIMIZER_LICENSE"
    case MSK_RES_ERR_FLEXLM: return "MSK_RES_ERR_FLEXLM"
    case MSK_RES_ERR_LICENSE_SERVER: return "MSK_RES_ERR_LICENSE_SERVER"
    case MSK_RES_ERR_LICENSE_MAX: return "MSK_RES_ERR_LICENSE_MAX"
    case MSK_RES_ERR_LICENSE_MOSEKLM_DAEMON: return "MSK_RES_ERR_LICENSE_MOSEKLM_DAEMON"
    case MSK_RES_ERR_LICENSE_FEATURE: return "MSK_RES_ERR_LICENSE_FEATURE"
    case MSK_RES_ERR_PLATFORM_NOT_LICENSED: return "MSK_RES_ERR_PLATFORM_NOT_LICENSED"
    case MSK_RES_ERR_LICENSE_CANNOT_ALLOCATE: return "MSK_RES_ERR_LICENSE_CANNOT_ALLOCATE"
    case MSK_RES_ERR_LICENSE_CANNOT_CONNECT: return "MSK_RES_ERR_LICENSE_CANNOT_CONNECT"
    case MSK_RES_ERR_LICENSE_INVALID_HOSTID: return "MSK_RES_ERR_LICENSE_INVALID_HOSTID"
    case MSK_RES_ERR_LICENSE_SERVER_VERSION: return "MSK_RES_ERR_LICENSE_SERVER_VERSION"
    case MSK_RES_ERR_LICENSE_NO_SERVER_SUPPORT: return "MSK_RES_ERR_LICENSE_NO_SERVER_SUPPORT"
    case MSK_RES_ERR_LICENSE_NO_SERVER_LINE: return "MSK_RES_ERR_LICENSE_NO_SERVER_LINE"
    case MSK_RES_ERR_OLDER_DLL: return "MSK_RES_ERR_OLDER_DLL"
    case MSK_RES_ERR_NEWER_DLL: return "MSK_RES_ERR_NEWER_DLL"
    case MSK_RES_ERR_LINK_FILE_DLL: return "MSK_RES_ERR_LINK_FILE_DLL"
    case MSK_RES_ERR_THREAD_MUTEX_INIT: return "MSK_RES_ERR_THREAD_MUTEX_INIT"
    case MSK_RES_ERR_THREAD_MUTEX_LOCK: return "MSK_RES_ERR_THREAD_MUTEX_LOCK"
    case MSK_RES_ERR_THREAD_MUTEX_UNLOCK: return "MSK_RES_ERR_THREAD_MUTEX_UNLOCK"
    case MSK_RES_ERR_THREAD_CREATE: return "MSK_RES_ERR_THREAD_CREATE"
    case MSK_RES_ERR_THREAD_COND_INIT: return "MSK_RES_ERR_THREAD_COND_INIT"
    case MSK_RES_ERR_UNKNOWN: return "MSK_RES_ERR_UNKNOWN"
    case MSK_RES_ERR_SPACE: return "MSK_RES_ERR_SPACE"
    case MSK_RES_ERR_FILE_OPEN: return "MSK_RES_ERR_FILE_OPEN"
    case MSK_RES_ERR_FILE_READ: return "MSK_RES_ERR_FILE_READ"
    case MSK_RES_ERR_FILE_WRITE: return "MSK_RES_ERR_FILE_WRITE"
    case MSK_RES_ERR_DATA_FILE_EXT: return "MSK_RES_ERR_DATA_FILE_EXT"
    case MSK_RES_ERR_INVALID_FILE_NAME: return "MSK_RES_ERR_INVALID_FILE_NAME"
    case MSK_RES_ERR_INVALID_SOL_FILE_NAME: return "MSK_RES_ERR_INVALID_SOL_FILE_NAME"
    case MSK_RES_ERR_END_OF_FILE: return "MSK_RES_ERR_END_OF_FILE"
    case MSK_RES_ERR_NULL_ENV: return "MSK_RES_ERR_NULL_ENV"
    case MSK_RES_ERR_NULL_TASK: return "MSK_RES_ERR_NULL_TASK"
    case MSK_RES_ERR_INVALID_STREAM: return "MSK_RES_ERR_INVALID_STREAM"
    case MSK_RES_ERR_NO_INIT_ENV: return "MSK_RES_ERR_NO_INIT_ENV"
    case MSK_RES_ERR_INVALID_TASK: return "MSK_RES_ERR_INVALID_TASK"
    case MSK_RES_ERR_NULL_POINTER: return "MSK_RES_ERR_NULL_POINTER"
    case MSK_RES_ERR_LIVING_TASKS: return "MSK_RES_ERR_LIVING_TASKS"
    case MSK_RES_ERR_READ_GZIP: return "MSK_RES_ERR_READ_GZIP"
    case MSK_RES_ERR_READ_ZSTD: return "MSK_RES_ERR_READ_ZSTD"
    case MSK_RES_ERR_BLANK_NAME: return "MSK_RES_ERR_BLANK_NAME"
    case MSK_RES_ERR_DUP_NAME: return "MSK_RES_ERR_DUP_NAME"
    case MSK_RES_ERR_FORMAT_STRING: return "MSK_RES_ERR_FORMAT_STRING"
    case MSK_RES_ERR_SPARSITY_SPECIFICATION: return "MSK_RES_ERR_SPARSITY_SPECIFICATION"
    case MSK_RES_ERR_MISMATCHING_DIMENSION: return "MSK_RES_ERR_MISMATCHING_DIMENSION"
    case MSK_RES_ERR_INVALID_OBJ_NAME: return "MSK_RES_ERR_INVALID_OBJ_NAME"
    case MSK_RES_ERR_INVALID_CON_NAME: return "MSK_RES_ERR_INVALID_CON_NAME"
    case MSK_RES_ERR_INVALID_VAR_NAME: return "MSK_RES_ERR_INVALID_VAR_NAME"
    case MSK_RES_ERR_INVALID_CONE_NAME: return "MSK_RES_ERR_INVALID_CONE_NAME"
    case MSK_RES_ERR_INVALID_BARVAR_NAME: return "MSK_RES_ERR_INVALID_BARVAR_NAME"
    case MSK_RES_ERR_SPACE_LEAKING: return "MSK_RES_ERR_SPACE_LEAKING"
    case MSK_RES_ERR_SPACE_NO_INFO: return "MSK_RES_ERR_SPACE_NO_INFO"
    case MSK_RES_ERR_DIMENSION_SPECIFICATION: return "MSK_RES_ERR_DIMENSION_SPECIFICATION"
    case MSK_RES_ERR_AXIS_NAME_SPECIFICATION: return "MSK_RES_ERR_AXIS_NAME_SPECIFICATION"
    case MSK_RES_ERR_READ_FORMAT: return "MSK_RES_ERR_READ_FORMAT"
    case MSK_RES_ERR_MPS_FILE: return "MSK_RES_ERR_MPS_FILE"
    case MSK_RES_ERR_MPS_INV_FIELD: return "MSK_RES_ERR_MPS_INV_FIELD"
    case MSK_RES_ERR_MPS_INV_MARKER: return "MSK_RES_ERR_MPS_INV_MARKER"
    case MSK_RES_ERR_MPS_NULL_CON_NAME: return "MSK_RES_ERR_MPS_NULL_CON_NAME"
    case MSK_RES_ERR_MPS_NULL_VAR_NAME: return "MSK_RES_ERR_MPS_NULL_VAR_NAME"
    case MSK_RES_ERR_MPS_UNDEF_CON_NAME: return "MSK_RES_ERR_MPS_UNDEF_CON_NAME"
    case MSK_RES_ERR_MPS_UNDEF_VAR_NAME: return "MSK_RES_ERR_MPS_UNDEF_VAR_NAME"
    case MSK_RES_ERR_MPS_INVALID_CON_KEY: return "MSK_RES_ERR_MPS_INVALID_CON_KEY"
    case MSK_RES_ERR_MPS_INVALID_BOUND_KEY: return "MSK_RES_ERR_MPS_INVALID_BOUND_KEY"
    case MSK_RES_ERR_MPS_INVALID_SEC_NAME: return "MSK_RES_ERR_MPS_INVALID_SEC_NAME"
    case MSK_RES_ERR_MPS_NO_OBJECTIVE: return "MSK_RES_ERR_MPS_NO_OBJECTIVE"
    case MSK_RES_ERR_MPS_SPLITTED_VAR: return "MSK_RES_ERR_MPS_SPLITTED_VAR"
    case MSK_RES_ERR_MPS_MUL_CON_NAME: return "MSK_RES_ERR_MPS_MUL_CON_NAME"
    case MSK_RES_ERR_MPS_MUL_QSEC: return "MSK_RES_ERR_MPS_MUL_QSEC"
    case MSK_RES_ERR_MPS_MUL_QOBJ: return "MSK_RES_ERR_MPS_MUL_QOBJ"
    case MSK_RES_ERR_MPS_INV_SEC_ORDER: return "MSK_RES_ERR_MPS_INV_SEC_ORDER"
    case MSK_RES_ERR_MPS_MUL_CSEC: return "MSK_RES_ERR_MPS_MUL_CSEC"
    case MSK_RES_ERR_MPS_CONE_TYPE: return "MSK_RES_ERR_MPS_CONE_TYPE"
    case MSK_RES_ERR_MPS_CONE_OVERLAP: return "MSK_RES_ERR_MPS_CONE_OVERLAP"
    case MSK_RES_ERR_MPS_CONE_REPEAT: return "MSK_RES_ERR_MPS_CONE_REPEAT"
    case MSK_RES_ERR_MPS_NON_SYMMETRIC_Q: return "MSK_RES_ERR_MPS_NON_SYMMETRIC_Q"
    case MSK_RES_ERR_MPS_DUPLICATE_Q_ELEMENT: return "MSK_RES_ERR_MPS_DUPLICATE_Q_ELEMENT"
    case MSK_RES_ERR_MPS_INVALID_OBJSENSE: return "MSK_RES_ERR_MPS_INVALID_OBJSENSE"
    case MSK_RES_ERR_MPS_TAB_IN_FIELD2: return "MSK_RES_ERR_MPS_TAB_IN_FIELD2"
    case MSK_RES_ERR_MPS_TAB_IN_FIELD3: return "MSK_RES_ERR_MPS_TAB_IN_FIELD3"
    case MSK_RES_ERR_MPS_TAB_IN_FIELD5: return "MSK_RES_ERR_MPS_TAB_IN_FIELD5"
    case MSK_RES_ERR_MPS_INVALID_OBJ_NAME: return "MSK_RES_ERR_MPS_INVALID_OBJ_NAME"
    case MSK_RES_ERR_MPS_INVALID_KEY: return "MSK_RES_ERR_MPS_INVALID_KEY"
    case MSK_RES_ERR_MPS_INVALID_INDICATOR_CONSTRAINT: return "MSK_RES_ERR_MPS_INVALID_INDICATOR_CONSTRAINT"
    case MSK_RES_ERR_MPS_INVALID_INDICATOR_VARIABLE: return "MSK_RES_ERR_MPS_INVALID_INDICATOR_VARIABLE"
    case MSK_RES_ERR_MPS_INVALID_INDICATOR_VALUE: return "MSK_RES_ERR_MPS_INVALID_INDICATOR_VALUE"
    case MSK_RES_ERR_MPS_INVALID_INDICATOR_QUADRATIC_CONSTRAINT: return "MSK_RES_ERR_MPS_INVALID_INDICATOR_QUADRATIC_CONSTRAINT"
    case MSK_RES_ERR_OPF_SYNTAX: return "MSK_RES_ERR_OPF_SYNTAX"
    case MSK_RES_ERR_OPF_PREMATURE_EOF: return "MSK_RES_ERR_OPF_PREMATURE_EOF"
    case MSK_RES_ERR_OPF_MISMATCHED_TAG: return "MSK_RES_ERR_OPF_MISMATCHED_TAG"
    case MSK_RES_ERR_OPF_DUPLICATE_BOUND: return "MSK_RES_ERR_OPF_DUPLICATE_BOUND"
    case MSK_RES_ERR_OPF_DUPLICATE_CONSTRAINT_NAME: return "MSK_RES_ERR_OPF_DUPLICATE_CONSTRAINT_NAME"
    case MSK_RES_ERR_OPF_INVALID_CONE_TYPE: return "MSK_RES_ERR_OPF_INVALID_CONE_TYPE"
    case MSK_RES_ERR_OPF_INCORRECT_TAG_PARAM: return "MSK_RES_ERR_OPF_INCORRECT_TAG_PARAM"
    case MSK_RES_ERR_OPF_INVALID_TAG: return "MSK_RES_ERR_OPF_INVALID_TAG"
    case MSK_RES_ERR_OPF_DUPLICATE_CONE_ENTRY: return "MSK_RES_ERR_OPF_DUPLICATE_CONE_ENTRY"
    case MSK_RES_ERR_OPF_TOO_LARGE: return "MSK_RES_ERR_OPF_TOO_LARGE"
    case MSK_RES_ERR_OPF_DUAL_INTEGER_SOLUTION: return "MSK_RES_ERR_OPF_DUAL_INTEGER_SOLUTION"
    case MSK_RES_ERR_LP_EMPTY: return "MSK_RES_ERR_LP_EMPTY"
    case MSK_RES_ERR_WRITE_MPS_INVALID_NAME: return "MSK_RES_ERR_WRITE_MPS_INVALID_NAME"
    case MSK_RES_ERR_LP_INVALID_VAR_NAME: return "MSK_RES_ERR_LP_INVALID_VAR_NAME"
    case MSK_RES_ERR_WRITE_OPF_INVALID_VAR_NAME: return "MSK_RES_ERR_WRITE_OPF_INVALID_VAR_NAME"
    case MSK_RES_ERR_LP_FILE_FORMAT: return "MSK_RES_ERR_LP_FILE_FORMAT"
    case MSK_RES_ERR_LP_EXPECTED_NUMBER: return "MSK_RES_ERR_LP_EXPECTED_NUMBER"
    case MSK_RES_ERR_READ_LP_MISSING_END_TAG: return "MSK_RES_ERR_READ_LP_MISSING_END_TAG"
    case MSK_RES_ERR_LP_INDICATOR_VAR: return "MSK_RES_ERR_LP_INDICATOR_VAR"
    case MSK_RES_ERR_LP_EXPECTED_OBJECTIVE: return "MSK_RES_ERR_LP_EXPECTED_OBJECTIVE"
    case MSK_RES_ERR_LP_EXPECTED_CONSTRAINT_RELATION: return "MSK_RES_ERR_LP_EXPECTED_CONSTRAINT_RELATION"
    case MSK_RES_ERR_LP_AMBIGUOUS_CONSTRAINT_BOUND: return "MSK_RES_ERR_LP_AMBIGUOUS_CONSTRAINT_BOUND"
    case MSK_RES_ERR_LP_DUPLICATE_SECTION: return "MSK_RES_ERR_LP_DUPLICATE_SECTION"
    case MSK_RES_ERR_READ_LP_DELAYED_ROWS_NOT_SUPPORTED: return "MSK_RES_ERR_READ_LP_DELAYED_ROWS_NOT_SUPPORTED"
    case MSK_RES_ERR_WRITING_FILE: return "MSK_RES_ERR_WRITING_FILE"
    case MSK_RES_ERR_INVALID_NAME_IN_SOL_FILE: return "MSK_RES_ERR_INVALID_NAME_IN_SOL_FILE"
    case MSK_RES_ERR_JSON_SYNTAX: return "MSK_RES_ERR_JSON_SYNTAX"
    case MSK_RES_ERR_JSON_STRING: return "MSK_RES_ERR_JSON_STRING"
    case MSK_RES_ERR_JSON_NUMBER_OVERFLOW: return "MSK_RES_ERR_JSON_NUMBER_OVERFLOW"
    case MSK_RES_ERR_JSON_FORMAT: return "MSK_RES_ERR_JSON_FORMAT"
    case MSK_RES_ERR_JSON_DATA: return "MSK_RES_ERR_JSON_DATA"
    case MSK_RES_ERR_JSON_MISSING_DATA: return "MSK_RES_ERR_JSON_MISSING_DATA"
    case MSK_RES_ERR_PTF_INCOMPATIBILITY: return "MSK_RES_ERR_PTF_INCOMPATIBILITY"
    case MSK_RES_ERR_PTF_UNDEFINED_ITEM: return "MSK_RES_ERR_PTF_UNDEFINED_ITEM"
    case MSK_RES_ERR_PTF_INCONSISTENCY: return "MSK_RES_ERR_PTF_INCONSISTENCY"
    case MSK_RES_ERR_PTF_FORMAT: return "MSK_RES_ERR_PTF_FORMAT"
    case MSK_RES_ERR_ARGUMENT_LENNEQ: return "MSK_RES_ERR_ARGUMENT_LENNEQ"
    case MSK_RES_ERR_ARGUMENT_TYPE: return "MSK_RES_ERR_ARGUMENT_TYPE"
    case MSK_RES_ERR_NUM_ARGUMENTS: return "MSK_RES_ERR_NUM_ARGUMENTS"
    case MSK_RES_ERR_IN_ARGUMENT: return "MSK_RES_ERR_IN_ARGUMENT"
    case MSK_RES_ERR_ARGUMENT_DIMENSION: return "MSK_RES_ERR_ARGUMENT_DIMENSION"
    case MSK_RES_ERR_SHAPE_IS_TOO_LARGE: return "MSK_RES_ERR_SHAPE_IS_TOO_LARGE"
    case MSK_RES_ERR_INDEX_IS_TOO_SMALL: return "MSK_RES_ERR_INDEX_IS_TOO_SMALL"
    case MSK_RES_ERR_INDEX_IS_TOO_LARGE: return "MSK_RES_ERR_INDEX_IS_TOO_LARGE"
    case MSK_RES_ERR_INDEX_IS_NOT_UNIQUE: return "MSK_RES_ERR_INDEX_IS_NOT_UNIQUE"
    case MSK_RES_ERR_PARAM_NAME: return "MSK_RES_ERR_PARAM_NAME"
    case MSK_RES_ERR_PARAM_NAME_DOU: return "MSK_RES_ERR_PARAM_NAME_DOU"
    case MSK_RES_ERR_PARAM_NAME_INT: return "MSK_RES_ERR_PARAM_NAME_INT"
    case MSK_RES_ERR_PARAM_NAME_STR: return "MSK_RES_ERR_PARAM_NAME_STR"
    case MSK_RES_ERR_PARAM_INDEX: return "MSK_RES_ERR_PARAM_INDEX"
    case MSK_RES_ERR_PARAM_IS_TOO_LARGE: return "MSK_RES_ERR_PARAM_IS_TOO_LARGE"
    case MSK_RES_ERR_PARAM_IS_TOO_SMALL: return "MSK_RES_ERR_PARAM_IS_TOO_SMALL"
    case MSK_RES_ERR_PARAM_VALUE_STR: return "MSK_RES_ERR_PARAM_VALUE_STR"
    case MSK_RES_ERR_PARAM_TYPE: return "MSK_RES_ERR_PARAM_TYPE"
    case MSK_RES_ERR_INF_DOU_INDEX: return "MSK_RES_ERR_INF_DOU_INDEX"
    case MSK_RES_ERR_INF_INT_INDEX: return "MSK_RES_ERR_INF_INT_INDEX"
    case MSK_RES_ERR_INDEX_ARR_IS_TOO_SMALL: return "MSK_RES_ERR_INDEX_ARR_IS_TOO_SMALL"
    case MSK_RES_ERR_INDEX_ARR_IS_TOO_LARGE: return "MSK_RES_ERR_INDEX_ARR_IS_TOO_LARGE"
    case MSK_RES_ERR_INF_LINT_INDEX: return "MSK_RES_ERR_INF_LINT_INDEX"
    case MSK_RES_ERR_ARG_IS_TOO_SMALL: return "MSK_RES_ERR_ARG_IS_TOO_SMALL"
    case MSK_RES_ERR_ARG_IS_TOO_LARGE: return "MSK_RES_ERR_ARG_IS_TOO_LARGE"
    case MSK_RES_ERR_INVALID_WHICHSOL: return "MSK_RES_ERR_INVALID_WHICHSOL"
    case MSK_RES_ERR_INF_DOU_NAME: return "MSK_RES_ERR_INF_DOU_NAME"
    case MSK_RES_ERR_INF_INT_NAME: return "MSK_RES_ERR_INF_INT_NAME"
    case MSK_RES_ERR_INF_TYPE: return "MSK_RES_ERR_INF_TYPE"
    case MSK_RES_ERR_INF_LINT_NAME: return "MSK_RES_ERR_INF_LINT_NAME"
    case MSK_RES_ERR_INDEX: return "MSK_RES_ERR_INDEX"
    case MSK_RES_ERR_WHICHSOL: return "MSK_RES_ERR_WHICHSOL"
    case MSK_RES_ERR_SOLITEM: return "MSK_RES_ERR_SOLITEM"
    case MSK_RES_ERR_WHICHITEM_NOT_ALLOWED: return "MSK_RES_ERR_WHICHITEM_NOT_ALLOWED"
    case MSK_RES_ERR_MAXNUMCON: return "MSK_RES_ERR_MAXNUMCON"
    case MSK_RES_ERR_MAXNUMVAR: return "MSK_RES_ERR_MAXNUMVAR"
    case MSK_RES_ERR_MAXNUMBARVAR: return "MSK_RES_ERR_MAXNUMBARVAR"
    case MSK_RES_ERR_MAXNUMQNZ: return "MSK_RES_ERR_MAXNUMQNZ"
    case MSK_RES_ERR_TOO_SMALL_MAX_NUM_NZ: return "MSK_RES_ERR_TOO_SMALL_MAX_NUM_NZ"
    case MSK_RES_ERR_INVALID_IDX: return "MSK_RES_ERR_INVALID_IDX"
    case MSK_RES_ERR_INVALID_MAX_NUM: return "MSK_RES_ERR_INVALID_MAX_NUM"
    case MSK_RES_ERR_UNALLOWED_WHICHSOL: return "MSK_RES_ERR_UNALLOWED_WHICHSOL"
    case MSK_RES_ERR_NUMCONLIM: return "MSK_RES_ERR_NUMCONLIM"
    case MSK_RES_ERR_NUMVARLIM: return "MSK_RES_ERR_NUMVARLIM"
    case MSK_RES_ERR_TOO_SMALL_MAXNUMANZ: return "MSK_RES_ERR_TOO_SMALL_MAXNUMANZ"
    case MSK_RES_ERR_INV_APTRE: return "MSK_RES_ERR_INV_APTRE"
    case MSK_RES_ERR_MUL_A_ELEMENT: return "MSK_RES_ERR_MUL_A_ELEMENT"
    case MSK_RES_ERR_INV_BK: return "MSK_RES_ERR_INV_BK"
    case MSK_RES_ERR_INV_BKC: return "MSK_RES_ERR_INV_BKC"
    case MSK_RES_ERR_INV_BKX: return "MSK_RES_ERR_INV_BKX"
    case MSK_RES_ERR_INV_VAR_TYPE: return "MSK_RES_ERR_INV_VAR_TYPE"
    case MSK_RES_ERR_SOLVER_PROBTYPE: return "MSK_RES_ERR_SOLVER_PROBTYPE"
    case MSK_RES_ERR_OBJECTIVE_RANGE: return "MSK_RES_ERR_OBJECTIVE_RANGE"
    case MSK_RES_ERR_INV_RESCODE: return "MSK_RES_ERR_INV_RESCODE"
    case MSK_RES_ERR_INV_IINF: return "MSK_RES_ERR_INV_IINF"
    case MSK_RES_ERR_INV_LIINF: return "MSK_RES_ERR_INV_LIINF"
    case MSK_RES_ERR_INV_DINF: return "MSK_RES_ERR_INV_DINF"
    case MSK_RES_ERR_BASIS: return "MSK_RES_ERR_BASIS"
    case MSK_RES_ERR_INV_SKC: return "MSK_RES_ERR_INV_SKC"
    case MSK_RES_ERR_INV_SKX: return "MSK_RES_ERR_INV_SKX"
    case MSK_RES_ERR_INV_SK_STR: return "MSK_RES_ERR_INV_SK_STR"
    case MSK_RES_ERR_INV_SK: return "MSK_RES_ERR_INV_SK"
    case MSK_RES_ERR_INV_CONE_TYPE_STR: return "MSK_RES_ERR_INV_CONE_TYPE_STR"
    case MSK_RES_ERR_INV_CONE_TYPE: return "MSK_RES_ERR_INV_CONE_TYPE"
    case MSK_RES_ERR_INV_SKN: return "MSK_RES_ERR_INV_SKN"
    case MSK_RES_ERR_INVALID_SURPLUS: return "MSK_RES_ERR_INVALID_SURPLUS"
    case MSK_RES_ERR_INV_NAME_ITEM: return "MSK_RES_ERR_INV_NAME_ITEM"
    case MSK_RES_ERR_PRO_ITEM: return "MSK_RES_ERR_PRO_ITEM"
    case MSK_RES_ERR_INVALID_FORMAT_TYPE: return "MSK_RES_ERR_INVALID_FORMAT_TYPE"
    case MSK_RES_ERR_FIRSTI: return "MSK_RES_ERR_FIRSTI"
    case MSK_RES_ERR_LASTI: return "MSK_RES_ERR_LASTI"
    case MSK_RES_ERR_FIRSTJ: return "MSK_RES_ERR_FIRSTJ"
    case MSK_RES_ERR_LASTJ: return "MSK_RES_ERR_LASTJ"
    case MSK_RES_ERR_MAX_LEN_IS_TOO_SMALL: return "MSK_RES_ERR_MAX_LEN_IS_TOO_SMALL"
    case MSK_RES_ERR_NONLINEAR_EQUALITY: return "MSK_RES_ERR_NONLINEAR_EQUALITY"
    case MSK_RES_ERR_NONCONVEX: return "MSK_RES_ERR_NONCONVEX"
    case MSK_RES_ERR_NONLINEAR_RANGED: return "MSK_RES_ERR_NONLINEAR_RANGED"
    case MSK_RES_ERR_CON_Q_NOT_PSD: return "MSK_RES_ERR_CON_Q_NOT_PSD"
    case MSK_RES_ERR_CON_Q_NOT_NSD: return "MSK_RES_ERR_CON_Q_NOT_NSD"
    case MSK_RES_ERR_OBJ_Q_NOT_PSD: return "MSK_RES_ERR_OBJ_Q_NOT_PSD"
    case MSK_RES_ERR_OBJ_Q_NOT_NSD: return "MSK_RES_ERR_OBJ_Q_NOT_NSD"
    case MSK_RES_ERR_ARGUMENT_PERM_ARRAY: return "MSK_RES_ERR_ARGUMENT_PERM_ARRAY"
    case MSK_RES_ERR_CONE_INDEX: return "MSK_RES_ERR_CONE_INDEX"
    case MSK_RES_ERR_CONE_SIZE: return "MSK_RES_ERR_CONE_SIZE"
    case MSK_RES_ERR_CONE_OVERLAP: return "MSK_RES_ERR_CONE_OVERLAP"
    case MSK_RES_ERR_CONE_REP_VAR: return "MSK_RES_ERR_CONE_REP_VAR"
    case MSK_RES_ERR_MAXNUMCONE: return "MSK_RES_ERR_MAXNUMCONE"
    case MSK_RES_ERR_CONE_TYPE: return "MSK_RES_ERR_CONE_TYPE"
    case MSK_RES_ERR_CONE_TYPE_STR: return "MSK_RES_ERR_CONE_TYPE_STR"
    case MSK_RES_ERR_CONE_OVERLAP_APPEND: return "MSK_RES_ERR_CONE_OVERLAP_APPEND"
    case MSK_RES_ERR_REMOVE_CONE_VARIABLE: return "MSK_RES_ERR_REMOVE_CONE_VARIABLE"
    case MSK_RES_ERR_APPENDING_TOO_BIG_CONE: return "MSK_RES_ERR_APPENDING_TOO_BIG_CONE"
    case MSK_RES_ERR_CONE_PARAMETER: return "MSK_RES_ERR_CONE_PARAMETER"
    case MSK_RES_ERR_SOL_FILE_INVALID_NUMBER: return "MSK_RES_ERR_SOL_FILE_INVALID_NUMBER"
    case MSK_RES_ERR_HUGE_C: return "MSK_RES_ERR_HUGE_C"
    case MSK_RES_ERR_HUGE_AIJ: return "MSK_RES_ERR_HUGE_AIJ"
    case MSK_RES_ERR_DUPLICATE_AIJ: return "MSK_RES_ERR_DUPLICATE_AIJ"
    case MSK_RES_ERR_LOWER_BOUND_IS_A_NAN: return "MSK_RES_ERR_LOWER_BOUND_IS_A_NAN"
    case MSK_RES_ERR_UPPER_BOUND_IS_A_NAN: return "MSK_RES_ERR_UPPER_BOUND_IS_A_NAN"
    case MSK_RES_ERR_INFINITE_BOUND: return "MSK_RES_ERR_INFINITE_BOUND"
    case MSK_RES_ERR_INV_QOBJ_SUBI: return "MSK_RES_ERR_INV_QOBJ_SUBI"
    case MSK_RES_ERR_INV_QOBJ_SUBJ: return "MSK_RES_ERR_INV_QOBJ_SUBJ"
    case MSK_RES_ERR_INV_QOBJ_VAL: return "MSK_RES_ERR_INV_QOBJ_VAL"
    case MSK_RES_ERR_INV_QCON_SUBK: return "MSK_RES_ERR_INV_QCON_SUBK"
    case MSK_RES_ERR_INV_QCON_SUBI: return "MSK_RES_ERR_INV_QCON_SUBI"
    case MSK_RES_ERR_INV_QCON_SUBJ: return "MSK_RES_ERR_INV_QCON_SUBJ"
    case MSK_RES_ERR_INV_QCON_VAL: return "MSK_RES_ERR_INV_QCON_VAL"
    case MSK_RES_ERR_QCON_SUBI_TOO_SMALL: return "MSK_RES_ERR_QCON_SUBI_TOO_SMALL"
    case MSK_RES_ERR_QCON_SUBI_TOO_LARGE: return "MSK_RES_ERR_QCON_SUBI_TOO_LARGE"
    case MSK_RES_ERR_QOBJ_UPPER_TRIANGLE: return "MSK_RES_ERR_QOBJ_UPPER_TRIANGLE"
    case MSK_RES_ERR_QCON_UPPER_TRIANGLE: return "MSK_RES_ERR_QCON_UPPER_TRIANGLE"
    case MSK_RES_ERR_FIXED_BOUND_VALUES: return "MSK_RES_ERR_FIXED_BOUND_VALUES"
    case MSK_RES_ERR_TOO_SMALL_A_TRUNCATION_VALUE: return "MSK_RES_ERR_TOO_SMALL_A_TRUNCATION_VALUE"
    case MSK_RES_ERR_INVALID_OBJECTIVE_SENSE: return "MSK_RES_ERR_INVALID_OBJECTIVE_SENSE"
    case MSK_RES_ERR_UNDEFINED_OBJECTIVE_SENSE: return "MSK_RES_ERR_UNDEFINED_OBJECTIVE_SENSE"
    case MSK_RES_ERR_Y_IS_UNDEFINED: return "MSK_RES_ERR_Y_IS_UNDEFINED"
    case MSK_RES_ERR_NAN_IN_DOUBLE_DATA: return "MSK_RES_ERR_NAN_IN_DOUBLE_DATA"
    case MSK_RES_ERR_INF_IN_DOUBLE_DATA: return "MSK_RES_ERR_INF_IN_DOUBLE_DATA"
    case MSK_RES_ERR_NAN_IN_BLC: return "MSK_RES_ERR_NAN_IN_BLC"
    case MSK_RES_ERR_NAN_IN_BUC: return "MSK_RES_ERR_NAN_IN_BUC"
    case MSK_RES_ERR_INVALID_CFIX: return "MSK_RES_ERR_INVALID_CFIX"
    case MSK_RES_ERR_NAN_IN_C: return "MSK_RES_ERR_NAN_IN_C"
    case MSK_RES_ERR_NAN_IN_BLX: return "MSK_RES_ERR_NAN_IN_BLX"
    case MSK_RES_ERR_NAN_IN_BUX: return "MSK_RES_ERR_NAN_IN_BUX"
    case MSK_RES_ERR_INVALID_AIJ: return "MSK_RES_ERR_INVALID_AIJ"
    case MSK_RES_ERR_INVALID_CJ: return "MSK_RES_ERR_INVALID_CJ"
    case MSK_RES_ERR_SYM_MAT_INVALID: return "MSK_RES_ERR_SYM_MAT_INVALID"
    case MSK_RES_ERR_SYM_MAT_HUGE: return "MSK_RES_ERR_SYM_MAT_HUGE"
    case MSK_RES_ERR_INV_PROBLEM: return "MSK_RES_ERR_INV_PROBLEM"
    case MSK_RES_ERR_MIXED_CONIC_AND_NL: return "MSK_RES_ERR_MIXED_CONIC_AND_NL"
    case MSK_RES_ERR_GLOBAL_INV_CONIC_PROBLEM: return "MSK_RES_ERR_GLOBAL_INV_CONIC_PROBLEM"
    case MSK_RES_ERR_INV_OPTIMIZER: return "MSK_RES_ERR_INV_OPTIMIZER"
    case MSK_RES_ERR_MIO_NO_OPTIMIZER: return "MSK_RES_ERR_MIO_NO_OPTIMIZER"
    case MSK_RES_ERR_NO_OPTIMIZER_VAR_TYPE: return "MSK_RES_ERR_NO_OPTIMIZER_VAR_TYPE"
    case MSK_RES_ERR_FINAL_SOLUTION: return "MSK_RES_ERR_FINAL_SOLUTION"
    case MSK_RES_ERR_FIRST: return "MSK_RES_ERR_FIRST"
    case MSK_RES_ERR_LAST: return "MSK_RES_ERR_LAST"
    case MSK_RES_ERR_SLICE_SIZE: return "MSK_RES_ERR_SLICE_SIZE"
    case MSK_RES_ERR_NEGATIVE_SURPLUS: return "MSK_RES_ERR_NEGATIVE_SURPLUS"
    case MSK_RES_ERR_NEGATIVE_APPEND: return "MSK_RES_ERR_NEGATIVE_APPEND"
    case MSK_RES_ERR_POSTSOLVE: return "MSK_RES_ERR_POSTSOLVE"
    case MSK_RES_ERR_OVERFLOW: return "MSK_RES_ERR_OVERFLOW"
    case MSK_RES_ERR_NO_BASIS_SOL: return "MSK_RES_ERR_NO_BASIS_SOL"
    case MSK_RES_ERR_BASIS_FACTOR: return "MSK_RES_ERR_BASIS_FACTOR"
    case MSK_RES_ERR_BASIS_SINGULAR: return "MSK_RES_ERR_BASIS_SINGULAR"
    case MSK_RES_ERR_FACTOR: return "MSK_RES_ERR_FACTOR"
    case MSK_RES_ERR_FEASREPAIR_CANNOT_RELAX: return "MSK_RES_ERR_FEASREPAIR_CANNOT_RELAX"
    case MSK_RES_ERR_FEASREPAIR_SOLVING_RELAXED: return "MSK_RES_ERR_FEASREPAIR_SOLVING_RELAXED"
    case MSK_RES_ERR_FEASREPAIR_INCONSISTENT_BOUND: return "MSK_RES_ERR_FEASREPAIR_INCONSISTENT_BOUND"
    case MSK_RES_ERR_REPAIR_INVALID_PROBLEM: return "MSK_RES_ERR_REPAIR_INVALID_PROBLEM"
    case MSK_RES_ERR_REPAIR_OPTIMIZATION_FAILED: return "MSK_RES_ERR_REPAIR_OPTIMIZATION_FAILED"
    case MSK_RES_ERR_NAME_MAX_LEN: return "MSK_RES_ERR_NAME_MAX_LEN"
    case MSK_RES_ERR_NAME_IS_NULL: return "MSK_RES_ERR_NAME_IS_NULL"
    case MSK_RES_ERR_INVALID_COMPRESSION: return "MSK_RES_ERR_INVALID_COMPRESSION"
    case MSK_RES_ERR_INVALID_IOMODE: return "MSK_RES_ERR_INVALID_IOMODE"
    case MSK_RES_ERR_NO_PRIMAL_INFEAS_CER: return "MSK_RES_ERR_NO_PRIMAL_INFEAS_CER"
    case MSK_RES_ERR_NO_DUAL_INFEAS_CER: return "MSK_RES_ERR_NO_DUAL_INFEAS_CER"
    case MSK_RES_ERR_NO_SOLUTION_IN_CALLBACK: return "MSK_RES_ERR_NO_SOLUTION_IN_CALLBACK"
    case MSK_RES_ERR_INV_MARKI: return "MSK_RES_ERR_INV_MARKI"
    case MSK_RES_ERR_INV_MARKJ: return "MSK_RES_ERR_INV_MARKJ"
    case MSK_RES_ERR_INV_NUMI: return "MSK_RES_ERR_INV_NUMI"
    case MSK_RES_ERR_INV_NUMJ: return "MSK_RES_ERR_INV_NUMJ"
    case MSK_RES_ERR_TASK_INCOMPATIBLE: return "MSK_RES_ERR_TASK_INCOMPATIBLE"
    case MSK_RES_ERR_TASK_INVALID: return "MSK_RES_ERR_TASK_INVALID"
    case MSK_RES_ERR_TASK_WRITE: return "MSK_RES_ERR_TASK_WRITE"
    case MSK_RES_ERR_LU_MAX_NUM_TRIES: return "MSK_RES_ERR_LU_MAX_NUM_TRIES"
    case MSK_RES_ERR_INVALID_UTF8: return "MSK_RES_ERR_INVALID_UTF8"
    case MSK_RES_ERR_INVALID_WCHAR: return "MSK_RES_ERR_INVALID_WCHAR"
    case MSK_RES_ERR_NO_DUAL_FOR_ITG_SOL: return "MSK_RES_ERR_NO_DUAL_FOR_ITG_SOL"
    case MSK_RES_ERR_NO_SNX_FOR_BAS_SOL: return "MSK_RES_ERR_NO_SNX_FOR_BAS_SOL"
    case MSK_RES_ERR_INTERNAL: return "MSK_RES_ERR_INTERNAL"
    case MSK_RES_ERR_API_ARRAY_TOO_SMALL: return "MSK_RES_ERR_API_ARRAY_TOO_SMALL"
    case MSK_RES_ERR_API_CB_CONNECT: return "MSK_RES_ERR_API_CB_CONNECT"
    case MSK_RES_ERR_API_FATAL_ERROR: return "MSK_RES_ERR_API_FATAL_ERROR"
    case MSK_RES_ERR_SEN_FORMAT: return "MSK_RES_ERR_SEN_FORMAT"
    case MSK_RES_ERR_SEN_UNDEF_NAME: return "MSK_RES_ERR_SEN_UNDEF_NAME"
    case MSK_RES_ERR_SEN_INDEX_RANGE: return "MSK_RES_ERR_SEN_INDEX_RANGE"
    case MSK_RES_ERR_SEN_BOUND_INVALID_UP: return "MSK_RES_ERR_SEN_BOUND_INVALID_UP"
    case MSK_RES_ERR_SEN_BOUND_INVALID_LO: return "MSK_RES_ERR_SEN_BOUND_INVALID_LO"
    case MSK_RES_ERR_SEN_INDEX_INVALID: return "MSK_RES_ERR_SEN_INDEX_INVALID"
    case MSK_RES_ERR_SEN_INVALID_REGEXP: return "MSK_RES_ERR_SEN_INVALID_REGEXP"
    case MSK_RES_ERR_SEN_SOLUTION_STATUS: return "MSK_RES_ERR_SEN_SOLUTION_STATUS"
    case MSK_RES_ERR_SEN_NUMERICAL: return "MSK_RES_ERR_SEN_NUMERICAL"
    case MSK_RES_ERR_SEN_UNHANDLED_PROBLEM_TYPE: return "MSK_RES_ERR_SEN_UNHANDLED_PROBLEM_TYPE"
    case MSK_RES_ERR_UNB_STEP_SIZE: return "MSK_RES_ERR_UNB_STEP_SIZE"
    case MSK_RES_ERR_IDENTICAL_TASKS: return "MSK_RES_ERR_IDENTICAL_TASKS"
    case MSK_RES_ERR_AD_INVALID_CODELIST: return "MSK_RES_ERR_AD_INVALID_CODELIST"
    case MSK_RES_ERR_INTERNAL_TEST_FAILED: return "MSK_RES_ERR_INTERNAL_TEST_FAILED"
    case MSK_RES_ERR_XML_INVALID_PROBLEM_TYPE: return "MSK_RES_ERR_XML_INVALID_PROBLEM_TYPE"
    case MSK_RES_ERR_INVALID_AMPL_STUB: return "MSK_RES_ERR_INVALID_AMPL_STUB"
    case MSK_RES_ERR_INT64_TO_INT32_CAST: return "MSK_RES_ERR_INT64_TO_INT32_CAST"
    case MSK_RES_ERR_SIZE_LICENSE_NUMCORES: return "MSK_RES_ERR_SIZE_LICENSE_NUMCORES"
    case MSK_RES_ERR_INFEAS_UNDEFINED: return "MSK_RES_ERR_INFEAS_UNDEFINED"
    case MSK_RES_ERR_NO_BARX_FOR_SOLUTION: return "MSK_RES_ERR_NO_BARX_FOR_SOLUTION"
    case MSK_RES_ERR_NO_BARS_FOR_SOLUTION: return "MSK_RES_ERR_NO_BARS_FOR_SOLUTION"
    case MSK_RES_ERR_BAR_VAR_DIM: return "MSK_RES_ERR_BAR_VAR_DIM"
    case MSK_RES_ERR_SYM_MAT_INVALID_ROW_INDEX: return "MSK_RES_ERR_SYM_MAT_INVALID_ROW_INDEX"
    case MSK_RES_ERR_SYM_MAT_INVALID_COL_INDEX: return "MSK_RES_ERR_SYM_MAT_INVALID_COL_INDEX"
    case MSK_RES_ERR_SYM_MAT_NOT_LOWER_TRINGULAR: return "MSK_RES_ERR_SYM_MAT_NOT_LOWER_TRINGULAR"
    case MSK_RES_ERR_SYM_MAT_INVALID_VALUE: return "MSK_RES_ERR_SYM_MAT_INVALID_VALUE"
    case MSK_RES_ERR_SYM_MAT_DUPLICATE: return "MSK_RES_ERR_SYM_MAT_DUPLICATE"
    case MSK_RES_ERR_INVALID_SYM_MAT_DIM: return "MSK_RES_ERR_INVALID_SYM_MAT_DIM"
    case MSK_RES_ERR_API_INTERNAL: return "MSK_RES_ERR_API_INTERNAL"
    case MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_SYM_MAT: return "MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_SYM_MAT"
    case MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_CFIX: return "MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_CFIX"
    case MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_RANGED_CONSTRAINTS: return "MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_RANGED_CONSTRAINTS"
    case MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_FREE_CONSTRAINTS: return "MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_FREE_CONSTRAINTS"
    case MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_CONES: return "MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_CONES"
    case MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_QUADRATIC_TERMS: return "MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_QUADRATIC_TERMS"
    case MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_NONLINEAR: return "MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_NONLINEAR"
    case MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_DISJUNCTIVE_CONSTRAINTS: return "MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_DISJUNCTIVE_CONSTRAINTS"
    case MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_AFFINE_CONIC_CONSTRAINTS: return "MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_AFFINE_CONIC_CONSTRAINTS"
    case MSK_RES_ERR_DUPLICATE_CONSTRAINT_NAMES: return "MSK_RES_ERR_DUPLICATE_CONSTRAINT_NAMES"
    case MSK_RES_ERR_DUPLICATE_VARIABLE_NAMES: return "MSK_RES_ERR_DUPLICATE_VARIABLE_NAMES"
    case MSK_RES_ERR_DUPLICATE_BARVARIABLE_NAMES: return "MSK_RES_ERR_DUPLICATE_BARVARIABLE_NAMES"
    case MSK_RES_ERR_DUPLICATE_CONE_NAMES: return "MSK_RES_ERR_DUPLICATE_CONE_NAMES"
    case MSK_RES_ERR_DUPLICATE_DOMAIN_NAMES: return "MSK_RES_ERR_DUPLICATE_DOMAIN_NAMES"
    case MSK_RES_ERR_DUPLICATE_DJC_NAMES: return "MSK_RES_ERR_DUPLICATE_DJC_NAMES"
    case MSK_RES_ERR_NON_UNIQUE_ARRAY: return "MSK_RES_ERR_NON_UNIQUE_ARRAY"
    case MSK_RES_ERR_ARGUMENT_IS_TOO_SMALL: return "MSK_RES_ERR_ARGUMENT_IS_TOO_SMALL"
    case MSK_RES_ERR_ARGUMENT_IS_TOO_LARGE: return "MSK_RES_ERR_ARGUMENT_IS_TOO_LARGE"
    case MSK_RES_ERR_MIO_INTERNAL: return "MSK_RES_ERR_MIO_INTERNAL"
    case MSK_RES_ERR_INVALID_PROBLEM_TYPE: return "MSK_RES_ERR_INVALID_PROBLEM_TYPE"
    case MSK_RES_ERR_UNHANDLED_SOLUTION_STATUS: return "MSK_RES_ERR_UNHANDLED_SOLUTION_STATUS"
    case MSK_RES_ERR_UPPER_TRIANGLE: return "MSK_RES_ERR_UPPER_TRIANGLE"
    case MSK_RES_ERR_LAU_SINGULAR_MATRIX: return "MSK_RES_ERR_LAU_SINGULAR_MATRIX"
    case MSK_RES_ERR_LAU_NOT_POSITIVE_DEFINITE: return "MSK_RES_ERR_LAU_NOT_POSITIVE_DEFINITE"
    case MSK_RES_ERR_LAU_INVALID_LOWER_TRIANGULAR_MATRIX: return "MSK_RES_ERR_LAU_INVALID_LOWER_TRIANGULAR_MATRIX"
    case MSK_RES_ERR_LAU_UNKNOWN: return "MSK_RES_ERR_LAU_UNKNOWN"
    case MSK_RES_ERR_LAU_ARG_M: return "MSK_RES_ERR_LAU_ARG_M"
    case MSK_RES_ERR_LAU_ARG_N: return "MSK_RES_ERR_LAU_ARG_N"
    case MSK_RES_ERR_LAU_ARG_K: return "MSK_RES_ERR_LAU_ARG_K"
    case MSK_RES_ERR_LAU_ARG_TRANSA: return "MSK_RES_ERR_LAU_ARG_TRANSA"
    case MSK_RES_ERR_LAU_ARG_TRANSB: return "MSK_RES_ERR_LAU_ARG_TRANSB"
    case MSK_RES_ERR_LAU_ARG_UPLO: return "MSK_RES_ERR_LAU_ARG_UPLO"
    case MSK_RES_ERR_LAU_ARG_TRANS: return "MSK_RES_ERR_LAU_ARG_TRANS"
    case MSK_RES_ERR_LAU_INVALID_SPARSE_SYMMETRIC_MATRIX: return "MSK_RES_ERR_LAU_INVALID_SPARSE_SYMMETRIC_MATRIX"
    case MSK_RES_ERR_CBF_PARSE: return "MSK_RES_ERR_CBF_PARSE"
    case MSK_RES_ERR_CBF_OBJ_SENSE: return "MSK_RES_ERR_CBF_OBJ_SENSE"
    case MSK_RES_ERR_CBF_NO_VARIABLES: return "MSK_RES_ERR_CBF_NO_VARIABLES"
    case MSK_RES_ERR_CBF_TOO_MANY_CONSTRAINTS: return "MSK_RES_ERR_CBF_TOO_MANY_CONSTRAINTS"
    case MSK_RES_ERR_CBF_TOO_MANY_VARIABLES: return "MSK_RES_ERR_CBF_TOO_MANY_VARIABLES"
    case MSK_RES_ERR_CBF_NO_VERSION_SPECIFIED: return "MSK_RES_ERR_CBF_NO_VERSION_SPECIFIED"
    case MSK_RES_ERR_CBF_SYNTAX: return "MSK_RES_ERR_CBF_SYNTAX"
    case MSK_RES_ERR_CBF_DUPLICATE_OBJ: return "MSK_RES_ERR_CBF_DUPLICATE_OBJ"
    case MSK_RES_ERR_CBF_DUPLICATE_CON: return "MSK_RES_ERR_CBF_DUPLICATE_CON"
    case MSK_RES_ERR_CBF_DUPLICATE_VAR: return "MSK_RES_ERR_CBF_DUPLICATE_VAR"
    case MSK_RES_ERR_CBF_DUPLICATE_INT: return "MSK_RES_ERR_CBF_DUPLICATE_INT"
    case MSK_RES_ERR_CBF_INVALID_VAR_TYPE: return "MSK_RES_ERR_CBF_INVALID_VAR_TYPE"
    case MSK_RES_ERR_CBF_INVALID_CON_TYPE: return "MSK_RES_ERR_CBF_INVALID_CON_TYPE"
    case MSK_RES_ERR_CBF_INVALID_DOMAIN_DIMENSION: return "MSK_RES_ERR_CBF_INVALID_DOMAIN_DIMENSION"
    case MSK_RES_ERR_CBF_DUPLICATE_OBJACOORD: return "MSK_RES_ERR_CBF_DUPLICATE_OBJACOORD"
    case MSK_RES_ERR_CBF_DUPLICATE_BCOORD: return "MSK_RES_ERR_CBF_DUPLICATE_BCOORD"
    case MSK_RES_ERR_CBF_DUPLICATE_ACOORD: return "MSK_RES_ERR_CBF_DUPLICATE_ACOORD"
    case MSK_RES_ERR_CBF_TOO_FEW_VARIABLES: return "MSK_RES_ERR_CBF_TOO_FEW_VARIABLES"
    case MSK_RES_ERR_CBF_TOO_FEW_CONSTRAINTS: return "MSK_RES_ERR_CBF_TOO_FEW_CONSTRAINTS"
    case MSK_RES_ERR_CBF_TOO_FEW_INTS: return "MSK_RES_ERR_CBF_TOO_FEW_INTS"
    case MSK_RES_ERR_CBF_TOO_MANY_INTS: return "MSK_RES_ERR_CBF_TOO_MANY_INTS"
    case MSK_RES_ERR_CBF_INVALID_INT_INDEX: return "MSK_RES_ERR_CBF_INVALID_INT_INDEX"
    case MSK_RES_ERR_CBF_UNSUPPORTED: return "MSK_RES_ERR_CBF_UNSUPPORTED"
    case MSK_RES_ERR_CBF_DUPLICATE_PSDVAR: return "MSK_RES_ERR_CBF_DUPLICATE_PSDVAR"
    case MSK_RES_ERR_CBF_INVALID_PSDVAR_DIMENSION: return "MSK_RES_ERR_CBF_INVALID_PSDVAR_DIMENSION"
    case MSK_RES_ERR_CBF_TOO_FEW_PSDVAR: return "MSK_RES_ERR_CBF_TOO_FEW_PSDVAR"
    case MSK_RES_ERR_CBF_INVALID_EXP_DIMENSION: return "MSK_RES_ERR_CBF_INVALID_EXP_DIMENSION"
    case MSK_RES_ERR_CBF_DUPLICATE_POW_CONES: return "MSK_RES_ERR_CBF_DUPLICATE_POW_CONES"
    case MSK_RES_ERR_CBF_DUPLICATE_POW_STAR_CONES: return "MSK_RES_ERR_CBF_DUPLICATE_POW_STAR_CONES"
    case MSK_RES_ERR_CBF_INVALID_POWER: return "MSK_RES_ERR_CBF_INVALID_POWER"
    case MSK_RES_ERR_CBF_POWER_CONE_IS_TOO_LONG: return "MSK_RES_ERR_CBF_POWER_CONE_IS_TOO_LONG"
    case MSK_RES_ERR_CBF_INVALID_POWER_CONE_INDEX: return "MSK_RES_ERR_CBF_INVALID_POWER_CONE_INDEX"
    case MSK_RES_ERR_CBF_INVALID_POWER_STAR_CONE_INDEX: return "MSK_RES_ERR_CBF_INVALID_POWER_STAR_CONE_INDEX"
    case MSK_RES_ERR_CBF_UNHANDLED_POWER_CONE_TYPE: return "MSK_RES_ERR_CBF_UNHANDLED_POWER_CONE_TYPE"
    case MSK_RES_ERR_CBF_UNHANDLED_POWER_STAR_CONE_TYPE: return "MSK_RES_ERR_CBF_UNHANDLED_POWER_STAR_CONE_TYPE"
    case MSK_RES_ERR_CBF_POWER_CONE_MISMATCH: return "MSK_RES_ERR_CBF_POWER_CONE_MISMATCH"
    case MSK_RES_ERR_CBF_POWER_STAR_CONE_MISMATCH: return "MSK_RES_ERR_CBF_POWER_STAR_CONE_MISMATCH"
    case MSK_RES_ERR_CBF_INVALID_NUMBER_OF_CONES: return "MSK_RES_ERR_CBF_INVALID_NUMBER_OF_CONES"
    case MSK_RES_ERR_CBF_INVALID_DIMENSION_OF_CONES: return "MSK_RES_ERR_CBF_INVALID_DIMENSION_OF_CONES"
    case MSK_RES_ERR_CBF_INVALID_NUM_OBJACOORD: return "MSK_RES_ERR_CBF_INVALID_NUM_OBJACOORD"
    case MSK_RES_ERR_CBF_INVALID_NUM_OBJFCOORD: return "MSK_RES_ERR_CBF_INVALID_NUM_OBJFCOORD"
    case MSK_RES_ERR_CBF_INVALID_NUM_ACOORD: return "MSK_RES_ERR_CBF_INVALID_NUM_ACOORD"
    case MSK_RES_ERR_CBF_INVALID_NUM_BCOORD: return "MSK_RES_ERR_CBF_INVALID_NUM_BCOORD"
    case MSK_RES_ERR_CBF_INVALID_NUM_FCOORD: return "MSK_RES_ERR_CBF_INVALID_NUM_FCOORD"
    case MSK_RES_ERR_CBF_INVALID_NUM_HCOORD: return "MSK_RES_ERR_CBF_INVALID_NUM_HCOORD"
    case MSK_RES_ERR_CBF_INVALID_NUM_DCOORD: return "MSK_RES_ERR_CBF_INVALID_NUM_DCOORD"
    case MSK_RES_ERR_CBF_EXPECTED_A_KEYWORD: return "MSK_RES_ERR_CBF_EXPECTED_A_KEYWORD"
    case MSK_RES_ERR_CBF_INVALID_NUM_PSDCON: return "MSK_RES_ERR_CBF_INVALID_NUM_PSDCON"
    case MSK_RES_ERR_CBF_DUPLICATE_PSDCON: return "MSK_RES_ERR_CBF_DUPLICATE_PSDCON"
    case MSK_RES_ERR_CBF_INVALID_DIMENSION_OF_PSDCON: return "MSK_RES_ERR_CBF_INVALID_DIMENSION_OF_PSDCON"
    case MSK_RES_ERR_CBF_INVALID_PSDCON_INDEX: return "MSK_RES_ERR_CBF_INVALID_PSDCON_INDEX"
    case MSK_RES_ERR_CBF_INVALID_PSDCON_VARIABLE_INDEX: return "MSK_RES_ERR_CBF_INVALID_PSDCON_VARIABLE_INDEX"
    case MSK_RES_ERR_CBF_INVALID_PSDCON_BLOCK_INDEX: return "MSK_RES_ERR_CBF_INVALID_PSDCON_BLOCK_INDEX"
    case MSK_RES_ERR_CBF_UNSUPPORTED_CHANGE: return "MSK_RES_ERR_CBF_UNSUPPORTED_CHANGE"
    case MSK_RES_ERR_MIO_INVALID_ROOT_OPTIMIZER: return "MSK_RES_ERR_MIO_INVALID_ROOT_OPTIMIZER"
    case MSK_RES_ERR_MIO_INVALID_NODE_OPTIMIZER: return "MSK_RES_ERR_MIO_INVALID_NODE_OPTIMIZER"
    case MSK_RES_ERR_MPS_WRITE_CPLEX_INVALID_CONE_TYPE: return "MSK_RES_ERR_MPS_WRITE_CPLEX_INVALID_CONE_TYPE"
    case MSK_RES_ERR_TOCONIC_CONSTR_Q_NOT_PSD: return "MSK_RES_ERR_TOCONIC_CONSTR_Q_NOT_PSD"
    case MSK_RES_ERR_TOCONIC_CONSTRAINT_FX: return "MSK_RES_ERR_TOCONIC_CONSTRAINT_FX"
    case MSK_RES_ERR_TOCONIC_CONSTRAINT_RA: return "MSK_RES_ERR_TOCONIC_CONSTRAINT_RA"
    case MSK_RES_ERR_TOCONIC_CONSTR_NOT_CONIC: return "MSK_RES_ERR_TOCONIC_CONSTR_NOT_CONIC"
    case MSK_RES_ERR_TOCONIC_OBJECTIVE_NOT_PSD: return "MSK_RES_ERR_TOCONIC_OBJECTIVE_NOT_PSD"
    case MSK_RES_ERR_SERVER_CONNECT: return "MSK_RES_ERR_SERVER_CONNECT"
    case MSK_RES_ERR_SERVER_PROTOCOL: return "MSK_RES_ERR_SERVER_PROTOCOL"
    case MSK_RES_ERR_SERVER_STATUS: return "MSK_RES_ERR_SERVER_STATUS"
    case MSK_RES_ERR_SERVER_TOKEN: return "MSK_RES_ERR_SERVER_TOKEN"
    case MSK_RES_ERR_SERVER_ADDRESS: return "MSK_RES_ERR_SERVER_ADDRESS"
    case MSK_RES_ERR_SERVER_CERTIFICATE: return "MSK_RES_ERR_SERVER_CERTIFICATE"
    case MSK_RES_ERR_SERVER_TLS_CLIENT: return "MSK_RES_ERR_SERVER_TLS_CLIENT"
    case MSK_RES_ERR_SERVER_ACCESS_TOKEN: return "MSK_RES_ERR_SERVER_ACCESS_TOKEN"
    case MSK_RES_ERR_SERVER_PROBLEM_SIZE: return "MSK_RES_ERR_SERVER_PROBLEM_SIZE"
    case MSK_RES_ERR_DUPLICATE_INDEX_IN_A_SPARSE_MATRIX: return "MSK_RES_ERR_DUPLICATE_INDEX_IN_A_SPARSE_MATRIX"
    case MSK_RES_ERR_DUPLICATE_INDEX_IN_AFEIDX_LIST: return "MSK_RES_ERR_DUPLICATE_INDEX_IN_AFEIDX_LIST"
    case MSK_RES_ERR_DUPLICATE_FIJ: return "MSK_RES_ERR_DUPLICATE_FIJ"
    case MSK_RES_ERR_INVALID_FIJ: return "MSK_RES_ERR_INVALID_FIJ"
    case MSK_RES_ERR_HUGE_FIJ: return "MSK_RES_ERR_HUGE_FIJ"
    case MSK_RES_ERR_INVALID_G: return "MSK_RES_ERR_INVALID_G"
    case MSK_RES_ERR_INVALID_B: return "MSK_RES_ERR_INVALID_B"
    case MSK_RES_ERR_DOMAIN_INVALID_INDEX: return "MSK_RES_ERR_DOMAIN_INVALID_INDEX"
    case MSK_RES_ERR_DOMAIN_DIMENSION: return "MSK_RES_ERR_DOMAIN_DIMENSION"
    case MSK_RES_ERR_DOMAIN_DIMENSION_PSD: return "MSK_RES_ERR_DOMAIN_DIMENSION_PSD"
    case MSK_RES_ERR_NOT_POWER_DOMAIN: return "MSK_RES_ERR_NOT_POWER_DOMAIN"
    case MSK_RES_ERR_DOMAIN_POWER_INVALID_ALPHA: return "MSK_RES_ERR_DOMAIN_POWER_INVALID_ALPHA"
    case MSK_RES_ERR_DOMAIN_POWER_NEGATIVE_ALPHA: return "MSK_RES_ERR_DOMAIN_POWER_NEGATIVE_ALPHA"
    case MSK_RES_ERR_DOMAIN_POWER_NLEFT: return "MSK_RES_ERR_DOMAIN_POWER_NLEFT"
    case MSK_RES_ERR_AFE_INVALID_INDEX: return "MSK_RES_ERR_AFE_INVALID_INDEX"
    case MSK_RES_ERR_ACC_INVALID_INDEX: return "MSK_RES_ERR_ACC_INVALID_INDEX"
    case MSK_RES_ERR_ACC_INVALID_ENTRY_INDEX: return "MSK_RES_ERR_ACC_INVALID_ENTRY_INDEX"
    case MSK_RES_ERR_ACC_AFE_DOMAIN_MISMATCH: return "MSK_RES_ERR_ACC_AFE_DOMAIN_MISMATCH"
    case MSK_RES_ERR_DJC_INVALID_INDEX: return "MSK_RES_ERR_DJC_INVALID_INDEX"
    case MSK_RES_ERR_DJC_UNSUPPORTED_DOMAIN_TYPE: return "MSK_RES_ERR_DJC_UNSUPPORTED_DOMAIN_TYPE"
    case MSK_RES_ERR_DJC_AFE_DOMAIN_MISMATCH: return "MSK_RES_ERR_DJC_AFE_DOMAIN_MISMATCH"
    case MSK_RES_ERR_DJC_INVALID_TERM_SIZE: return "MSK_RES_ERR_DJC_INVALID_TERM_SIZE"
    case MSK_RES_ERR_DJC_DOMAIN_TERMSIZE_MISMATCH: return "MSK_RES_ERR_DJC_DOMAIN_TERMSIZE_MISMATCH"
    case MSK_RES_ERR_DJC_TOTAL_NUM_TERMS_MISMATCH: return "MSK_RES_ERR_DJC_TOTAL_NUM_TERMS_MISMATCH"
    case MSK_RES_ERR_UNDEF_SOLUTION: return "MSK_RES_ERR_UNDEF_SOLUTION"
    case MSK_RES_ERR_NO_DOTY: return "MSK_RES_ERR_NO_DOTY"
    case MSK_RES_TRM_MAX_ITERATIONS: return "MSK_RES_TRM_MAX_ITERATIONS"
    case MSK_RES_TRM_MAX_TIME: return "MSK_RES_TRM_MAX_TIME"
    case MSK_RES_TRM_OBJECTIVE_RANGE: return "MSK_RES_TRM_OBJECTIVE_RANGE"
    case MSK_RES_TRM_STALL: return "MSK_RES_TRM_STALL"
    case MSK_RES_TRM_USER_CALLBACK: return "MSK_RES_TRM_USER_CALLBACK"
    case MSK_RES_TRM_MIO_NUM_RELAXS: return "MSK_RES_TRM_MIO_NUM_RELAXS"
    case MSK_RES_TRM_MIO_NUM_BRANCHES: return "MSK_RES_TRM_MIO_NUM_BRANCHES"
    case MSK_RES_TRM_NUM_MAX_NUM_INT_SOLUTIONS: return "MSK_RES_TRM_NUM_MAX_NUM_INT_SOLUTIONS"
    case MSK_RES_TRM_MAX_NUM_SETBACKS: return "MSK_RES_TRM_MAX_NUM_SETBACKS"
    case MSK_RES_TRM_NUMERICAL_PROBLEM: return "MSK_RES_TRM_NUMERICAL_PROBLEM"
    case MSK_RES_TRM_LOST_RACE: return "MSK_RES_TRM_LOST_RACE"
    case MSK_RES_TRM_INTERNAL: return "MSK_RES_TRM_INTERNAL"
    case MSK_RES_TRM_INTERNAL_STOP: return "MSK_RES_TRM_INTERNAL_STOP"
    default: return "<unknown>"
  }
} // Rescode.String()

type Rescodetype int32
// Response code type
const (
    MSK_RESPONSE_OK Rescodetype = 0 // The response code is OK.
    MSK_RESPONSE_WRN Rescodetype = 1 // The response code is a warning.
    MSK_RESPONSE_TRM Rescodetype = 2 // The response code is an optimizer termination status.
    MSK_RESPONSE_ERR Rescodetype = 3 // The response code is an error.
    MSK_RESPONSE_UNK Rescodetype = 4 // The response code does not belong to any class.
    MSK_RESPONSE_END Rescodetype = 4
)
func (self Rescodetype) String() string {
  switch self {
    case MSK_RESPONSE_OK: return "MSK_RESPONSE_OK"
    case MSK_RESPONSE_WRN: return "MSK_RESPONSE_WRN"
    case MSK_RESPONSE_TRM: return "MSK_RESPONSE_TRM"
    case MSK_RESPONSE_ERR: return "MSK_RESPONSE_ERR"
    case MSK_RESPONSE_UNK: return "MSK_RESPONSE_UNK"
    default: return "<unknown>"
  }
} // Rescodetype.String()

type Scalingtype int32
// Scaling type
const (
    MSK_SCALING_FREE Scalingtype = 0 // The optimizer chooses the scaling heuristic.
    MSK_SCALING_NONE Scalingtype = 1 // No scaling is performed.
    MSK_SCALING_END Scalingtype = 1
)
func (self Scalingtype) String() string {
  switch self {
    case MSK_SCALING_FREE: return "MSK_SCALING_FREE"
    case MSK_SCALING_NONE: return "MSK_SCALING_NONE"
    default: return "<unknown>"
  }
} // Scalingtype.String()

type Scalingmethod int32
// Scaling method
const (
    MSK_SCALING_METHOD_POW2 Scalingmethod = 0 // Scales only with power of 2 leaving the mantissa untouched.
    MSK_SCALING_METHOD_FREE Scalingmethod = 1 // The optimizer chooses the scaling heuristic.
    MSK_SCALING_METHOD_END Scalingmethod = 1
)
func (self Scalingmethod) String() string {
  switch self {
    case MSK_SCALING_METHOD_POW2: return "MSK_SCALING_METHOD_POW2"
    case MSK_SCALING_METHOD_FREE: return "MSK_SCALING_METHOD_FREE"
    default: return "<unknown>"
  }
} // Scalingmethod.String()

type Sensitivitytype int32
// Sensitivity types
const (
    MSK_SENSITIVITY_TYPE_BASIS Sensitivitytype = 0 // Basis sensitivity analysis is performed.
    MSK_SENSITIVITY_TYPE_END Sensitivitytype = 0
)
func (self Sensitivitytype) String() string {
  switch self {
    case MSK_SENSITIVITY_TYPE_BASIS: return "MSK_SENSITIVITY_TYPE_BASIS"
    default: return "<unknown>"
  }
} // Sensitivitytype.String()

type Simseltype int32
// Simplex selection strategy
const (
    MSK_SIM_SELECTION_FREE Simseltype = 0 // The optimizer chooses the pricing strategy.
    MSK_SIM_SELECTION_FULL Simseltype = 1 // The optimizer uses full pricing.
    MSK_SIM_SELECTION_ASE Simseltype = 2 // The optimizer uses approximate steepest-edge pricing.
    MSK_SIM_SELECTION_DEVEX Simseltype = 3 // The optimizer uses devex steepest-edge pricing.
    MSK_SIM_SELECTION_SE Simseltype = 4 // The optimizer uses steepest-edge selection.
    MSK_SIM_SELECTION_PARTIAL Simseltype = 5 // The optimizer uses a partial selection approach.
    MSK_SIM_SELECTION_END Simseltype = 5
)
func (self Simseltype) String() string {
  switch self {
    case MSK_SIM_SELECTION_FREE: return "MSK_SIM_SELECTION_FREE"
    case MSK_SIM_SELECTION_FULL: return "MSK_SIM_SELECTION_FULL"
    case MSK_SIM_SELECTION_ASE: return "MSK_SIM_SELECTION_ASE"
    case MSK_SIM_SELECTION_DEVEX: return "MSK_SIM_SELECTION_DEVEX"
    case MSK_SIM_SELECTION_SE: return "MSK_SIM_SELECTION_SE"
    case MSK_SIM_SELECTION_PARTIAL: return "MSK_SIM_SELECTION_PARTIAL"
    default: return "<unknown>"
  }
} // Simseltype.String()

type Solitem int32
// Solution items
const (
    MSK_SOL_ITEM_XC Solitem = 0 // Solution for the constraints.
    MSK_SOL_ITEM_XX Solitem = 1 // Variable solution.
    MSK_SOL_ITEM_Y Solitem = 2 // Lagrange multipliers for equations.
    MSK_SOL_ITEM_SLC Solitem = 3 // Lagrange multipliers for lower bounds on the constraints.
    MSK_SOL_ITEM_SUC Solitem = 4 // Lagrange multipliers for upper bounds on the constraints.
    MSK_SOL_ITEM_SLX Solitem = 5 // Lagrange multipliers for lower bounds on the variables.
    MSK_SOL_ITEM_SUX Solitem = 6 // Lagrange multipliers for upper bounds on the variables.
    MSK_SOL_ITEM_SNX Solitem = 7 // Lagrange multipliers corresponding to the conic constraints on the variables.
    MSK_SOL_ITEM_END Solitem = 7
)
func (self Solitem) String() string {
  switch self {
    case MSK_SOL_ITEM_XC: return "MSK_SOL_ITEM_XC"
    case MSK_SOL_ITEM_XX: return "MSK_SOL_ITEM_XX"
    case MSK_SOL_ITEM_Y: return "MSK_SOL_ITEM_Y"
    case MSK_SOL_ITEM_SLC: return "MSK_SOL_ITEM_SLC"
    case MSK_SOL_ITEM_SUC: return "MSK_SOL_ITEM_SUC"
    case MSK_SOL_ITEM_SLX: return "MSK_SOL_ITEM_SLX"
    case MSK_SOL_ITEM_SUX: return "MSK_SOL_ITEM_SUX"
    case MSK_SOL_ITEM_SNX: return "MSK_SOL_ITEM_SNX"
    default: return "<unknown>"
  }
} // Solitem.String()

type Solsta int32
// Solution status keys
const (
    MSK_SOL_STA_UNKNOWN Solsta = 0 // Status of the solution is unknown.
    MSK_SOL_STA_OPTIMAL Solsta = 1 // The solution is optimal.
    MSK_SOL_STA_PRIM_FEAS Solsta = 2 // The solution is primal feasible.
    MSK_SOL_STA_DUAL_FEAS Solsta = 3 // The solution is dual feasible.
    MSK_SOL_STA_PRIM_AND_DUAL_FEAS Solsta = 4 // The solution is both primal and dual feasible.
    MSK_SOL_STA_PRIM_INFEAS_CER Solsta = 5 // The solution is a certificate of primal infeasibility.
    MSK_SOL_STA_DUAL_INFEAS_CER Solsta = 6 // The solution is a certificate of dual infeasibility.
    MSK_SOL_STA_PRIM_ILLPOSED_CER Solsta = 7 // The solution is a certificate that the primal problem is illposed.
    MSK_SOL_STA_DUAL_ILLPOSED_CER Solsta = 8 // The solution is a certificate that the dual problem is illposed.
    MSK_SOL_STA_INTEGER_OPTIMAL Solsta = 9 // The primal solution is integer optimal.
    MSK_SOL_STA_END Solsta = 9
)
func (self Solsta) String() string {
  switch self {
    case MSK_SOL_STA_UNKNOWN: return "MSK_SOL_STA_UNKNOWN"
    case MSK_SOL_STA_OPTIMAL: return "MSK_SOL_STA_OPTIMAL"
    case MSK_SOL_STA_PRIM_FEAS: return "MSK_SOL_STA_PRIM_FEAS"
    case MSK_SOL_STA_DUAL_FEAS: return "MSK_SOL_STA_DUAL_FEAS"
    case MSK_SOL_STA_PRIM_AND_DUAL_FEAS: return "MSK_SOL_STA_PRIM_AND_DUAL_FEAS"
    case MSK_SOL_STA_PRIM_INFEAS_CER: return "MSK_SOL_STA_PRIM_INFEAS_CER"
    case MSK_SOL_STA_DUAL_INFEAS_CER: return "MSK_SOL_STA_DUAL_INFEAS_CER"
    case MSK_SOL_STA_PRIM_ILLPOSED_CER: return "MSK_SOL_STA_PRIM_ILLPOSED_CER"
    case MSK_SOL_STA_DUAL_ILLPOSED_CER: return "MSK_SOL_STA_DUAL_ILLPOSED_CER"
    case MSK_SOL_STA_INTEGER_OPTIMAL: return "MSK_SOL_STA_INTEGER_OPTIMAL"
    default: return "<unknown>"
  }
} // Solsta.String()

type Soltype int32
// Solution types
const (
    MSK_SOL_ITR Soltype = 0 // The interior solution.
    MSK_SOL_BAS Soltype = 1 // The basic solution.
    MSK_SOL_ITG Soltype = 2 // The integer solution.
    MSK_SOL_END Soltype = 2
)
func (self Soltype) String() string {
  switch self {
    case MSK_SOL_ITR: return "MSK_SOL_ITR"
    case MSK_SOL_BAS: return "MSK_SOL_BAS"
    case MSK_SOL_ITG: return "MSK_SOL_ITG"
    default: return "<unknown>"
  }
} // Soltype.String()

type Solveform int32
// Solve primal or dual form
const (
    MSK_SOLVE_FREE Solveform = 0 // The optimizer is free to solve either the primal or the dual problem.
    MSK_SOLVE_PRIMAL Solveform = 1 // The optimizer should solve the primal problem.
    MSK_SOLVE_DUAL Solveform = 2 // The optimizer should solve the dual problem.
    MSK_SOLVE_END Solveform = 2
)
func (self Solveform) String() string {
  switch self {
    case MSK_SOLVE_FREE: return "MSK_SOLVE_FREE"
    case MSK_SOLVE_PRIMAL: return "MSK_SOLVE_PRIMAL"
    case MSK_SOLVE_DUAL: return "MSK_SOLVE_DUAL"
    default: return "<unknown>"
  }
} // Solveform.String()

type Sparam int32
// String parameters
const (
    MSK_SPAR_BAS_SOL_FILE_NAME Sparam = 0 // Name of the bas solution file.
    MSK_SPAR_DATA_FILE_NAME Sparam = 1 // Data are read and written to this file.
    MSK_SPAR_DEBUG_FILE_NAME Sparam = 2 // MOSEK debug file.
    MSK_SPAR_INT_SOL_FILE_NAME Sparam = 3 // Name of the int solution file.
    MSK_SPAR_ITR_SOL_FILE_NAME Sparam = 4 // Name of the itr solution file.
    MSK_SPAR_MIO_DEBUG_STRING Sparam = 5 // For internal debugging purposes.
    MSK_SPAR_PARAM_COMMENT_SIGN Sparam = 6 // Solution file comment character.
    MSK_SPAR_PARAM_READ_FILE_NAME Sparam = 7 // Modifications to the parameter database is read from this file.
    MSK_SPAR_PARAM_WRITE_FILE_NAME Sparam = 8 // The parameter database is written to this file.
    MSK_SPAR_READ_MPS_BOU_NAME Sparam = 9 // Name of the BOUNDS vector used. An empty name means that the first BOUNDS vector is used.
    MSK_SPAR_READ_MPS_OBJ_NAME Sparam = 10 // Objective name in the MPS file.
    MSK_SPAR_READ_MPS_RAN_NAME Sparam = 11 // Name of the RANGE vector  used. An empty name means that the first RANGE vector is used.
    MSK_SPAR_READ_MPS_RHS_NAME Sparam = 12 // Name of the RHS used. An empty name means that the first RHS vector is used.
    MSK_SPAR_REMOTE_OPTSERVER_HOST Sparam = 13 // URL of the remote optimization server.
    MSK_SPAR_REMOTE_TLS_CERT Sparam = 14 // Known server certificates in PEM format
    MSK_SPAR_REMOTE_TLS_CERT_PATH Sparam = 15 // Path to known server certificates in PEM format
    MSK_SPAR_SENSITIVITY_FILE_NAME Sparam = 16 // Sensitivity report file name.
    MSK_SPAR_SENSITIVITY_RES_FILE_NAME Sparam = 17 // Name of the sensitivity report output file.
    MSK_SPAR_SOL_FILTER_XC_LOW Sparam = 18 // Solution file filter.
    MSK_SPAR_SOL_FILTER_XC_UPR Sparam = 19 // Solution file filter.
    MSK_SPAR_SOL_FILTER_XX_LOW Sparam = 20 // Solution file filter.
    MSK_SPAR_SOL_FILTER_XX_UPR Sparam = 21 // Solution file filter.
    MSK_SPAR_STAT_KEY Sparam = 22 // Key used when writing the summary file.
    MSK_SPAR_STAT_NAME Sparam = 23 // Name used when writing the statistics file.
    MSK_SPAR_WRITE_LP_GEN_VAR_NAME Sparam = 24 // Added variable names in the LP files.
    MSK_SPAR_END Sparam = 24
)
func (self Sparam) String() string {
  switch self {
    case MSK_SPAR_BAS_SOL_FILE_NAME: return "MSK_SPAR_BAS_SOL_FILE_NAME"
    case MSK_SPAR_DATA_FILE_NAME: return "MSK_SPAR_DATA_FILE_NAME"
    case MSK_SPAR_DEBUG_FILE_NAME: return "MSK_SPAR_DEBUG_FILE_NAME"
    case MSK_SPAR_INT_SOL_FILE_NAME: return "MSK_SPAR_INT_SOL_FILE_NAME"
    case MSK_SPAR_ITR_SOL_FILE_NAME: return "MSK_SPAR_ITR_SOL_FILE_NAME"
    case MSK_SPAR_MIO_DEBUG_STRING: return "MSK_SPAR_MIO_DEBUG_STRING"
    case MSK_SPAR_PARAM_COMMENT_SIGN: return "MSK_SPAR_PARAM_COMMENT_SIGN"
    case MSK_SPAR_PARAM_READ_FILE_NAME: return "MSK_SPAR_PARAM_READ_FILE_NAME"
    case MSK_SPAR_PARAM_WRITE_FILE_NAME: return "MSK_SPAR_PARAM_WRITE_FILE_NAME"
    case MSK_SPAR_READ_MPS_BOU_NAME: return "MSK_SPAR_READ_MPS_BOU_NAME"
    case MSK_SPAR_READ_MPS_OBJ_NAME: return "MSK_SPAR_READ_MPS_OBJ_NAME"
    case MSK_SPAR_READ_MPS_RAN_NAME: return "MSK_SPAR_READ_MPS_RAN_NAME"
    case MSK_SPAR_READ_MPS_RHS_NAME: return "MSK_SPAR_READ_MPS_RHS_NAME"
    case MSK_SPAR_REMOTE_OPTSERVER_HOST: return "MSK_SPAR_REMOTE_OPTSERVER_HOST"
    case MSK_SPAR_REMOTE_TLS_CERT: return "MSK_SPAR_REMOTE_TLS_CERT"
    case MSK_SPAR_REMOTE_TLS_CERT_PATH: return "MSK_SPAR_REMOTE_TLS_CERT_PATH"
    case MSK_SPAR_SENSITIVITY_FILE_NAME: return "MSK_SPAR_SENSITIVITY_FILE_NAME"
    case MSK_SPAR_SENSITIVITY_RES_FILE_NAME: return "MSK_SPAR_SENSITIVITY_RES_FILE_NAME"
    case MSK_SPAR_SOL_FILTER_XC_LOW: return "MSK_SPAR_SOL_FILTER_XC_LOW"
    case MSK_SPAR_SOL_FILTER_XC_UPR: return "MSK_SPAR_SOL_FILTER_XC_UPR"
    case MSK_SPAR_SOL_FILTER_XX_LOW: return "MSK_SPAR_SOL_FILTER_XX_LOW"
    case MSK_SPAR_SOL_FILTER_XX_UPR: return "MSK_SPAR_SOL_FILTER_XX_UPR"
    case MSK_SPAR_STAT_KEY: return "MSK_SPAR_STAT_KEY"
    case MSK_SPAR_STAT_NAME: return "MSK_SPAR_STAT_NAME"
    case MSK_SPAR_WRITE_LP_GEN_VAR_NAME: return "MSK_SPAR_WRITE_LP_GEN_VAR_NAME"
    default: return "<unknown>"
  }
} // Sparam.String()

type Stakey int32
// Status keys
const (
    MSK_SK_UNK Stakey = 0 // The status for the constraint or variable is unknown.
    MSK_SK_BAS Stakey = 1 // The constraint or variable is in the basis.
    MSK_SK_SUPBAS Stakey = 2 // The constraint or variable is super basic.
    MSK_SK_LOW Stakey = 3 // The constraint or variable is at its lower bound.
    MSK_SK_UPR Stakey = 4 // The constraint or variable is at its upper bound.
    MSK_SK_FIX Stakey = 5 // The constraint or variable is fixed.
    MSK_SK_INF Stakey = 6 // The constraint or variable is infeasible in the bounds.
    MSK_SK_END Stakey = 6
)
func (self Stakey) String() string {
  switch self {
    case MSK_SK_UNK: return "MSK_SK_UNK"
    case MSK_SK_BAS: return "MSK_SK_BAS"
    case MSK_SK_SUPBAS: return "MSK_SK_SUPBAS"
    case MSK_SK_LOW: return "MSK_SK_LOW"
    case MSK_SK_UPR: return "MSK_SK_UPR"
    case MSK_SK_FIX: return "MSK_SK_FIX"
    case MSK_SK_INF: return "MSK_SK_INF"
    default: return "<unknown>"
  }
} // Stakey.String()

type Startpointtype int32
// Starting point types
const (
    MSK_STARTING_POINT_FREE Startpointtype = 0 // The starting point is chosen automatically.
    MSK_STARTING_POINT_GUESS Startpointtype = 1 // The optimizer guesses a starting point.
    MSK_STARTING_POINT_CONSTANT Startpointtype = 2 // The optimizer constructs a starting point by assigning a constant value to all primal and dual variables. This starting point is normally robust.
    MSK_STARTING_POINT_END Startpointtype = 2
)
func (self Startpointtype) String() string {
  switch self {
    case MSK_STARTING_POINT_FREE: return "MSK_STARTING_POINT_FREE"
    case MSK_STARTING_POINT_GUESS: return "MSK_STARTING_POINT_GUESS"
    case MSK_STARTING_POINT_CONSTANT: return "MSK_STARTING_POINT_CONSTANT"
    default: return "<unknown>"
  }
} // Startpointtype.String()

type Streamtype int32
// Stream types
const (
    MSK_STREAM_LOG Streamtype = 0 // Log stream. Contains the aggregated contents of all other streams. This means that a message written to any other stream will also be written to this stream.
    MSK_STREAM_MSG Streamtype = 1 // Message stream. Log information relating to performance and progress of the optimization is written to this stream.
    MSK_STREAM_ERR Streamtype = 2 // Error stream. Error messages are written to this stream.
    MSK_STREAM_WRN Streamtype = 3 // Warning stream. Warning messages are written to this stream.
    MSK_STREAM_END Streamtype = 3
)
func (self Streamtype) String() string {
  switch self {
    case MSK_STREAM_LOG: return "MSK_STREAM_LOG"
    case MSK_STREAM_MSG: return "MSK_STREAM_MSG"
    case MSK_STREAM_ERR: return "MSK_STREAM_ERR"
    case MSK_STREAM_WRN: return "MSK_STREAM_WRN"
    default: return "<unknown>"
  }
} // Streamtype.String()

type Value int32
// Integer values
const (
    MSK_LICENSE_BUFFER_LENGTH Value = 21 // The length of a license key buffer.
    MSK_MAX_STR_LEN Value = 1024 // Maximum string length allowed in MOSEK.
)
func (self Value) String() string {
  switch self {
    case MSK_LICENSE_BUFFER_LENGTH: return "MSK_LICENSE_BUFFER_LENGTH"
    case MSK_MAX_STR_LEN: return "MSK_MAX_STR_LEN"
    default: return "<unknown>"
  }
} // Value.String()

type Variabletype int32
// Variable types
const (
    MSK_VAR_TYPE_CONT Variabletype = 0 // Is a continuous variable.
    MSK_VAR_TYPE_INT Variabletype = 1 // Is an integer variable.
    MSK_VAR_END Variabletype = 1
)
func (self Variabletype) String() string {
  switch self {
    case MSK_VAR_TYPE_CONT: return "MSK_VAR_TYPE_CONT"
    case MSK_VAR_TYPE_INT: return "MSK_VAR_TYPE_INT"
    default: return "<unknown>"
  }
} // Variabletype.String()

// Methods

// Analyze the names and issue an error for the first invalid name.
//
// - whichstream Streamtype
//   Index of the stream.
// - nametype Nametype
//   The type of names e.g. valid in MPS or LP files.
func (self *Task) AnalyzeNames(whichstream Streamtype,nametype Nametype) (err error) {
  if _tmp0 := C.MSK_analyzenames(self.ptr(),C.int32_t(whichstream),C.int32_t(nametype)); _tmp0 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp0)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Analyze the data of a task.
//
// - whichstream Streamtype
//   Index of the stream.
func (self *Task) AnalyzeProblem(whichstream Streamtype) (err error) {
  if _tmp1 := C.MSK_analyzeproblem(self.ptr(),C.int32_t(whichstream)); _tmp1 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Print information related to the quality of the solution.
//
// - whichstream Streamtype
//   Index of the stream.
// - whichsol Soltype
//   Selects a solution.
func (self *Task) AnalyzeSolution(whichstream Streamtype,whichsol Soltype) (err error) {
  if _tmp2 := C.MSK_analyzesolution(self.ptr(),C.int32_t(whichstream),C.int32_t(whichsol)); _tmp2 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp2)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends an affine conic constraint to the task.
//
// - domidx int64
//   Domain index.
// - afeidxlist []int64
//   List of affine expression indexes.
// - b []float64
//   The vector of constant terms added to affine expressions. Optional, can be NULL.
func (self *Task) AppendAcc(domidx int64,afeidxlist []int64,b []float64) (err error) {
  _tmp3 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp3)
  var _tmp4 *int64
  if len(afeidxlist) > 0 { _tmp4 = (*int64)(&afeidxlist[0]) }
  if b != nil && int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"AppendAcc",arg:"b"}
    return
  }
  var _tmp5 *float64
  if len(b) > 0 { _tmp5 = (*float64)(&b[0]) }
  if _tmp6 := C.MSK_appendacc(self.ptr(),C.int64_t(domidx),C.int64_t(numafeidx),(*C.int64_t)(_tmp4),(*C.double)(_tmp5)); _tmp6 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp6)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends a number of affine conic constraint to the task.
//
// - domidxs []int64
//   Domain indices.
// - afeidxlist []int64
//   List of affine expression indexes.
// - b []float64
//   The vector of constant terms added to affine expressions. Optional, can be NULL.
func (self *Task) AppendAccs(domidxs []int64,afeidxlist []int64,b []float64) (err error) {
  _tmp7 := len(domidxs)
  var numaccs int64 = int64(_tmp7)
  var _tmp8 *int64
  if len(domidxs) > 0 { _tmp8 = (*int64)(&domidxs[0]) }
  _tmp9 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp9)
  var _tmp10 *int64
  if len(afeidxlist) > 0 { _tmp10 = (*int64)(&afeidxlist[0]) }
  if b != nil && int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"AppendAccs",arg:"b"}
    return
  }
  var _tmp11 *float64
  if len(b) > 0 { _tmp11 = (*float64)(&b[0]) }
  if _tmp12 := C.MSK_appendaccs(self.ptr(),C.int64_t(numaccs),(*C.int64_t)(_tmp8),C.int64_t(numafeidx),(*C.int64_t)(_tmp10),(*C.double)(_tmp11)); _tmp12 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp12)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends an affine conic constraint to the task.
//
// - domidx int64
//   Domain index.
// - afeidxfirst int64
//   Index of the first affine expression.
// - b []float64
//   The vector of constant terms added to affine expressions. Optional, can be NULL.
func (self *Task) AppendAccSeq(domidx int64,afeidxfirst int64,b []float64) (err error) {
  var _tmp13 C.int64_t
  if _tmp14 := C.MSK_getdomainn(self.ptr(),(C.int64_t)(domidx),(&_tmp13)); _tmp14 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp14)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var numafeidx int64 = int64(_tmp13)
  if b != nil && int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"AppendAccSeq",arg:"b"}
    return
  }
  var _tmp15 *float64
  if len(b) > 0 { _tmp15 = (*float64)(&b[0]) }
  if _tmp16 := C.MSK_appendaccseq(self.ptr(),C.int64_t(domidx),C.int64_t(numafeidx),C.int64_t(afeidxfirst),(*C.double)(_tmp15)); _tmp16 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp16)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends a number of affine conic constraint to the task.
//
// - domidxs []int64
//   Domain indices.
// - numafeidx int64
//   Number of affine expressions in the affine expression list (must equal the sum of dimensions of the domains).
// - afeidxfirst int64
//   Index of the first affine expression.
// - b []float64
//   The vector of constant terms added to affine expressions. Optional, can be NULL.
func (self *Task) AppendAccsSeq(domidxs []int64,numafeidx int64,afeidxfirst int64,b []float64) (err error) {
  _tmp17 := len(domidxs)
  var numaccs int64 = int64(_tmp17)
  var _tmp18 *int64
  if len(domidxs) > 0 { _tmp18 = (*int64)(&domidxs[0]) }
  if b != nil && int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"AppendAccsSeq",arg:"b"}
    return
  }
  var _tmp19 *float64
  if len(b) > 0 { _tmp19 = (*float64)(&b[0]) }
  if _tmp20 := C.MSK_appendaccsseq(self.ptr(),C.int64_t(numaccs),(*C.int64_t)(_tmp18),C.int64_t(numafeidx),C.int64_t(afeidxfirst),(*C.double)(_tmp19)); _tmp20 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp20)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends a number of empty affine expressions to the optimization task.
//
// - num int64
//   Number of empty affine expressions which should be appended.
func (self *Task) AppendAfes(num int64) (err error) {
  if _tmp21 := C.MSK_appendafes(self.ptr(),C.int64_t(num)); _tmp21 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp21)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends semidefinite variables to the problem.
//
// - dim []int32
//   Dimensions of symmetric matrix variables to be added.
func (self *Task) AppendBarvars(dim []int32) (err error) {
  _tmp22 := len(dim)
  var num int32 = int32(_tmp22)
  var _tmp23 *int32
  if len(dim) > 0 { _tmp23 = (*int32)(&dim[0]) }
  if _tmp24 := C.MSK_appendbarvars(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp23)); _tmp24 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp24)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends a new conic constraint to the problem.
//
// - ct Conetype
//   Specifies the type of the cone.
// - conepar float64
//   For the power cone it denotes the exponent alpha. For other cone types it is unused and can be set to 0.
// - submem []int32
//   Variable subscripts of the members in the cone.
func (self *Task) AppendCone(ct Conetype,conepar float64,submem []int32) (err error) {
  _tmp25 := len(submem)
  var nummem int32 = int32(_tmp25)
  var _tmp26 *int32
  if len(submem) > 0 { _tmp26 = (*int32)(&submem[0]) }
  if _tmp27 := C.MSK_appendcone(self.ptr(),C.int32_t(ct),C.double(conepar),C.int32_t(nummem),(*C.int32_t)(_tmp26)); _tmp27 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp27)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends a new conic constraint to the problem.
//
// - ct Conetype
//   Specifies the type of the cone.
// - conepar float64
//   For the power cone it denotes the exponent alpha. For other cone types it is unused and can be set to 0.
// - nummem int32
//   Number of member variables in the cone.
// - j int32
//   Index of the first variable in the conic constraint.
func (self *Task) AppendConeSeq(ct Conetype,conepar float64,nummem int32,j int32) (err error) {
  if _tmp28 := C.MSK_appendconeseq(self.ptr(),C.int32_t(ct),C.double(conepar),C.int32_t(nummem),C.int32_t(j)); _tmp28 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp28)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends multiple conic constraints to the problem.
//
// - ct []Conetype
//   Specifies the type of the cone.
// - conepar []float64
//   For the power cone it denotes the exponent alpha. For other cone types it is unused and can be set to 0.
// - nummem []int32
//   Numbers of member variables in the cones.
// - j int32
//   Index of the first variable in the first cone to be appended.
func (self *Task) AppendConesSeq(ct []Conetype,conepar []float64,nummem []int32,j int32) (err error) {
  _tmp29 := len(conepar)
  if _tmp29 < len(ct) { _tmp29 = len(ct) }
  if _tmp29 < len(nummem) { _tmp29 = len(nummem) }
  var num int32 = int32(_tmp29)
  var _tmp30 *Conetype
  if len(ct) > 0 { _tmp30 = (*Conetype)(&ct[0]) }
  var _tmp31 *float64
  if len(conepar) > 0 { _tmp31 = (*float64)(&conepar[0]) }
  var _tmp32 *int32
  if len(nummem) > 0 { _tmp32 = (*int32)(&nummem[0]) }
  if _tmp33 := C.MSK_appendconesseq(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp30),(*C.double)(_tmp31),(*C.int32_t)(_tmp32),C.int32_t(j)); _tmp33 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp33)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends a number of constraints to the optimization task.
//
// - num int32
//   Number of constraints which should be appended.
func (self *Task) AppendCons(num int32) (err error) {
  if _tmp34 := C.MSK_appendcons(self.ptr(),C.int32_t(num)); _tmp34 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp34)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends a number of empty disjunctive constraints to the task.
//
// - num int64
//   Number of empty disjunctive constraints which should be appended.
func (self *Task) AppendDjcs(num int64) (err error) {
  if _tmp35 := C.MSK_appenddjcs(self.ptr(),C.int64_t(num)); _tmp35 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp35)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the dual exponential cone domain.
//
// - domidx int64
//   Index of the domain.
func (self *Task) AppendDualExpConeDomain() (domidx int64,err error) {
  if _tmp36 := C.MSK_appenddualexpconedomain(self.ptr(),(*C.int64_t)(&domidx)); _tmp36 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp36)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the dual geometric mean cone domain.
//
// - n int64
//   Dimmension of the domain.
// - domidx int64
//   Index of the domain.
func (self *Task) AppendDualGeoMeanConeDomain(n int64) (domidx int64,err error) {
  if _tmp37 := C.MSK_appenddualgeomeanconedomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp37 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp37)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the dual power cone domain.
//
// - n int64
//   Dimension of the domain.
// - alpha []float64
//   The sequence proportional to exponents. Must be positive.
// - domidx int64
//   Index of the domain.
func (self *Task) AppendDualPowerConeDomain(n int64,alpha []float64) (domidx int64,err error) {
  _tmp38 := len(alpha)
  var nleft int64 = int64(_tmp38)
  var _tmp39 *float64
  if len(alpha) > 0 { _tmp39 = (*float64)(&alpha[0]) }
  if _tmp40 := C.MSK_appenddualpowerconedomain(self.ptr(),C.int64_t(n),C.int64_t(nleft),(*C.double)(_tmp39),(*C.int64_t)(&domidx)); _tmp40 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp40)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the primal exponential cone domain.
//
// - domidx int64
//   Index of the domain.
func (self *Task) AppendPrimalExpConeDomain() (domidx int64,err error) {
  if _tmp41 := C.MSK_appendprimalexpconedomain(self.ptr(),(*C.int64_t)(&domidx)); _tmp41 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp41)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the primal geometric mean cone domain.
//
// - n int64
//   Dimmension of the domain.
// - domidx int64
//   Index of the domain.
func (self *Task) AppendPrimalGeoMeanConeDomain(n int64) (domidx int64,err error) {
  if _tmp42 := C.MSK_appendprimalgeomeanconedomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp42 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp42)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the primal power cone domain.
//
// - n int64
//   Dimension of the domain.
// - alpha []float64
//   The sequence proportional to exponents. Must be positive.
// - domidx int64
//   Index of the domain.
func (self *Task) AppendPrimalPowerConeDomain(n int64,alpha []float64) (domidx int64,err error) {
  _tmp43 := len(alpha)
  var nleft int64 = int64(_tmp43)
  var _tmp44 *float64
  if len(alpha) > 0 { _tmp44 = (*float64)(&alpha[0]) }
  if _tmp45 := C.MSK_appendprimalpowerconedomain(self.ptr(),C.int64_t(n),C.int64_t(nleft),(*C.double)(_tmp44),(*C.int64_t)(&domidx)); _tmp45 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp45)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the n dimensional quadratic cone domain.
//
// - n int64
//   Dimmension of the domain.
// - domidx int64
//   Index of the domain.
func (self *Task) AppendQuadraticConeDomain(n int64) (domidx int64,err error) {
  if _tmp46 := C.MSK_appendquadraticconedomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp46 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp46)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the n dimensional real number domain.
//
// - n int64
//   Dimmension of the domain.
// - domidx int64
//   Index of the domain.
func (self *Task) AppendRDomain(n int64) (domidx int64,err error) {
  if _tmp47 := C.MSK_appendrdomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp47 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp47)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the n dimensional negative orthant to the list of domains.
//
// - n int64
//   Dimmension of the domain.
// - domidx int64
//   Index of the domain.
func (self *Task) AppendRminusDomain(n int64) (domidx int64,err error) {
  if _tmp48 := C.MSK_appendrminusdomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp48 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp48)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the n dimensional positive orthant to the list of domains.
//
// - n int64
//   Dimmension of the domain.
// - domidx int64
//   Index of the domain.
func (self *Task) AppendRplusDomain(n int64) (domidx int64,err error) {
  if _tmp49 := C.MSK_appendrplusdomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp49 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp49)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the n dimensional rotated quadratic cone domain.
//
// - n int64
//   Dimmension of the domain.
// - domidx int64
//   Index of the domain.
func (self *Task) AppendRQuadraticConeDomain(n int64) (domidx int64,err error) {
  if _tmp50 := C.MSK_appendrquadraticconedomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp50 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp50)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the n dimensional 0 domain.
//
// - n int64
//   Dimmension of the domain.
// - domidx int64
//   Index of the domain.
func (self *Task) AppendRzeroDomain(n int64) (domidx int64,err error) {
  if _tmp51 := C.MSK_appendrzerodomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp51 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp51)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends a general sparse symmetric matrix to the storage of symmetric matrices.
//
// - dim int32
//   Dimension of the symmetric matrix that is appended.
// - subi []int32
//   Row subscript in the triplets.
// - subj []int32
//   Column subscripts in the triplets.
// - valij []float64
//   Values of each triplet.
// - idx int64
//   Unique index assigned to the inputted matrix.
func (self *Task) AppendSparseSymMat(dim int32,subi []int32,subj []int32,valij []float64) (idx int64,err error) {
  _tmp52 := len(subi)
  if _tmp52 < len(valij) { _tmp52 = len(valij) }
  if _tmp52 < len(subj) { _tmp52 = len(subj) }
  var nz int64 = int64(_tmp52)
  var _tmp53 *int32
  if len(subi) > 0 { _tmp53 = (*int32)(&subi[0]) }
  var _tmp54 *int32
  if len(subj) > 0 { _tmp54 = (*int32)(&subj[0]) }
  var _tmp55 *float64
  if len(valij) > 0 { _tmp55 = (*float64)(&valij[0]) }
  if _tmp56 := C.MSK_appendsparsesymmat(self.ptr(),C.int32_t(dim),C.int64_t(nz),(*C.int32_t)(_tmp53),(*C.int32_t)(_tmp54),(*C.double)(_tmp55),(*C.int64_t)(&idx)); _tmp56 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp56)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends a general sparse symmetric matrix to the storage of symmetric matrices.
//
// - dims []int32
//   Dimensions of the symmetric matrixes.
// - nz []int64
//   Number of nonzeros for each matrix.
// - subi []int32
//   Row subscript in the triplets.
// - subj []int32
//   Column subscripts in the triplets.
// - valij []float64
//   Values of each triplet.
// - idx []int64
//   Unique index assigned to the inputted matrix.
func (self *Task) AppendSparseSymMatList(dims []int32,nz []int64,subi []int32,subj []int32,valij []float64) (idx []int64,err error) {
  _tmp57 := len(dims)
  if _tmp57 < len(nz) { _tmp57 = len(nz) }
  var num int32 = int32(_tmp57)
  var _tmp58 *int32
  if len(dims) > 0 { _tmp58 = (*int32)(&dims[0]) }
  var _tmp59 *int64
  if len(nz) > 0 { _tmp59 = (*int64)(&nz[0]) }
  if int64(len(subi)) < int64(sum(nz)) {
    err = &ArrayLengthError{fun:"AppendSparseSymMatList",arg:"subi"}
    return
  }
  var _tmp60 *int32
  if len(subi) > 0 { _tmp60 = (*int32)(&subi[0]) }
  if int64(len(subj)) < int64(sum(nz)) {
    err = &ArrayLengthError{fun:"AppendSparseSymMatList",arg:"subj"}
    return
  }
  var _tmp61 *int32
  if len(subj) > 0 { _tmp61 = (*int32)(&subj[0]) }
  if int64(len(valij)) < int64(sum(nz)) {
    err = &ArrayLengthError{fun:"AppendSparseSymMatList",arg:"valij"}
    return
  }
  var _tmp62 *float64
  if len(valij) > 0 { _tmp62 = (*float64)(&valij[0]) }
  var _tmp63 *int64
  idx = make([]int64,num)
  if len(idx) > 0 { _tmp63 = (*int64)(&idx[0]) }
  if _tmp64 := C.MSK_appendsparsesymmatlist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp58),(*C.int64_t)(_tmp59),(*C.int32_t)(_tmp60),(*C.int32_t)(_tmp61),(*C.double)(_tmp62),(*C.int64_t)(_tmp63)); _tmp64 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp64)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends the vectorized SVEC PSD cone domain.
//
// - n int64
//   Dimension of the domain.
// - domidx int64
//   Index of the domain.
func (self *Task) AppendSvecPsdConeDomain(n int64) (domidx int64,err error) {
  if _tmp65 := C.MSK_appendsvecpsdconedomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp65 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp65)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Appends a number of variables to the optimization task.
//
// - num int32
//   Number of variables which should be appended.
func (self *Task) AppendVars(num int32) (err error) {
  if _tmp66 := C.MSK_appendvars(self.ptr(),C.int32_t(num)); _tmp66 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp66)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes conditioning information for the basis matrix.
//
// - nrmbasis float64
//   An estimate for the 1-norm of the basis.
// - nrminvbasis float64
//   An estimate for the 1-norm of the inverse of the basis.
func (self *Task) BasisCond() (nrmbasis float64,nrminvbasis float64,err error) {
  if _tmp67 := C.MSK_basiscond(self.ptr(),(*C.double)(&nrmbasis),(*C.double)(&nrminvbasis)); _tmp67 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp67)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Checks the memory allocated by the task.
//
// - file string
//   File from which the function is called.
// - line int32
//   Line in the file from which the function is called.
func (self *Task) CheckMem(file string,line int32) (err error) {
  _tmp68 := C.CString(file)
  if _tmp69 := C.MSK_checkmemtask(self.ptr(),_tmp68,C.int32_t(line)); _tmp69 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp69)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes the bounds for one constraint.
//
// - i int32
//   Index of the constraint for which the bounds should be changed.
// - lower int32
//   If non-zero, then the lower bound is changed, otherwise the upper bound is changed.
// - finite int32
//   If non-zero, then the given value is assumed to be finite.
// - value float64
//   New value for the bound.
func (self *Task) ChgConBound(i int32,lower int32,finite int32,value float64) (err error) {
  if _tmp70 := C.MSK_chgconbound(self.ptr(),C.int32_t(i),C.int32_t(lower),C.int32_t(finite),C.double(value)); _tmp70 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp70)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes the bounds for one variable.
//
// - j int32
//   Index of the variable for which the bounds should be changed.
// - lower int32
//   If non-zero, then the lower bound is changed, otherwise the upper bound is changed.
// - finite int32
//   If non-zero, then the given value is assumed to be finite.
// - value float64
//   New value for the bound.
func (self *Task) ChgVarBound(j int32,lower int32,finite int32,value float64) (err error) {
  if _tmp71 := C.MSK_chgvarbound(self.ptr(),C.int32_t(j),C.int32_t(lower),C.int32_t(finite),C.double(value)); _tmp71 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp71)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Commits all cached problem changes.
func (self *Task) CommitChanges() (err error) {
  if _tmp72 := C.MSK_commitchanges(self.ptr()); _tmp72 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp72)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Undefine a solution and free the memory it uses.
//
// - whichsol Soltype
//   Selects a solution.
func (self *Task) DeleteSolution(whichsol Soltype) (err error) {
  if _tmp73 := C.MSK_deletesolution(self.ptr(),C.int32_t(whichsol)); _tmp73 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp73)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Performs sensitivity analysis on objective coefficients.
//
// - subj []int32
//   Indexes of objective coefficients to analyze.
// - leftpricej []float64
//   Left shadow prices for requested coefficients.
// - rightpricej []float64
//   Right shadow prices for requested coefficients.
// - leftrangej []float64
//   Left range for requested coefficients.
// - rightrangej []float64
//   Right range for requested coefficients.
func (self *Task) DualSensitivity(subj []int32) (leftpricej []float64,rightpricej []float64,leftrangej []float64,rightrangej []float64,err error) {
  _tmp74 := len(subj)
  var numj int32 = int32(_tmp74)
  var _tmp75 *int32
  if len(subj) > 0 { _tmp75 = (*int32)(&subj[0]) }
  var _tmp76 *float64
  leftpricej = make([]float64,numj)
  if len(leftpricej) > 0 { _tmp76 = (*float64)(&leftpricej[0]) }
  var _tmp77 *float64
  rightpricej = make([]float64,numj)
  if len(rightpricej) > 0 { _tmp77 = (*float64)(&rightpricej[0]) }
  var _tmp78 *float64
  leftrangej = make([]float64,numj)
  if len(leftrangej) > 0 { _tmp78 = (*float64)(&leftrangej[0]) }
  var _tmp79 *float64
  rightrangej = make([]float64,numj)
  if len(rightrangej) > 0 { _tmp79 = (*float64)(&rightrangej[0]) }
  if _tmp80 := C.MSK_dualsensitivity(self.ptr(),C.int32_t(numj),(*C.int32_t)(_tmp75),(*C.double)(_tmp76),(*C.double)(_tmp77),(*C.double)(_tmp78),(*C.double)(_tmp79)); _tmp80 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp80)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Clears a row in barF
//
// - afeidx int64
//   Row index of barF.
func (self *Task) EmptyAfeBarfRow(afeidx int64) (err error) {
  if _tmp81 := C.MSK_emptyafebarfrow(self.ptr(),C.int64_t(afeidx)); _tmp81 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp81)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Clears rows in barF.
//
// - afeidxlist []int64
//   Indices of rows in barF to clear.
func (self *Task) EmptyAfeBarfRowList(afeidxlist []int64) (err error) {
  _tmp82 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp82)
  var _tmp83 *int64
  if len(afeidxlist) > 0 { _tmp83 = (*int64)(&afeidxlist[0]) }
  if _tmp84 := C.MSK_emptyafebarfrowlist(self.ptr(),C.int64_t(numafeidx),(*C.int64_t)(_tmp83)); _tmp84 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp84)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Clears a column in F.
//
// - varidx int32
//   Variable index.
func (self *Task) EmptyAfeFCol(varidx int32) (err error) {
  if _tmp85 := C.MSK_emptyafefcol(self.ptr(),C.int32_t(varidx)); _tmp85 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp85)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Clears columns in F.
//
// - varidx []int32
//   Indices of variables in F to clear.
func (self *Task) EmptyAfeFColList(varidx []int32) (err error) {
  _tmp86 := len(varidx)
  var numvaridx int64 = int64(_tmp86)
  var _tmp87 *int32
  if len(varidx) > 0 { _tmp87 = (*int32)(&varidx[0]) }
  if _tmp88 := C.MSK_emptyafefcollist(self.ptr(),C.int64_t(numvaridx),(*C.int32_t)(_tmp87)); _tmp88 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp88)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Clears a row in F.
//
// - afeidx int64
//   Row index.
func (self *Task) EmptyAfeFRow(afeidx int64) (err error) {
  if _tmp89 := C.MSK_emptyafefrow(self.ptr(),C.int64_t(afeidx)); _tmp89 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp89)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Clears rows in F.
//
// - afeidx []int64
//   Indices of rows in F to clear.
func (self *Task) EmptyAfeFRowList(afeidx []int64) (err error) {
  _tmp90 := len(afeidx)
  var numafeidx int64 = int64(_tmp90)
  var _tmp91 *int64
  if len(afeidx) > 0 { _tmp91 = (*int64)(&afeidx[0]) }
  if _tmp92 := C.MSK_emptyafefrowlist(self.ptr(),C.int64_t(numafeidx),(*C.int64_t)(_tmp91)); _tmp92 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp92)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Evaluates the activity of an affine conic constraint.
//
// - whichsol Soltype
//   Selects a solution.
// - accidx int64
//   The index of the affine conic constraint.
// - activity []float64
//   The activity of the affine conic constraint. The array should have length equal to the dimension of the constraint.
func (self *Task) EvaluateAcc(whichsol Soltype,accidx int64) (activity []float64,err error) {
  var _tmp93 C.int64_t
  if _tmp94 := C.MSK_getaccn(self.ptr(),(C.int64_t)(accidx),(&_tmp93)); _tmp94 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp94)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp95 *float64
  activity = make([]float64,_tmp93)
  if len(activity) > 0 { _tmp95 = (*float64)(&activity[0]) }
  if _tmp96 := C.MSK_evaluateacc(self.ptr(),C.int32_t(whichsol),C.int64_t(accidx),(*C.double)(_tmp95)); _tmp96 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp96)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Evaluates the activities of all affine conic constraints.
//
// - whichsol Soltype
//   Selects a solution.
// - activity []float64
//   The activity of affine conic constraints. The array should have length equal to the sum of dimensions of all affine conic constraints.
func (self *Task) EvaluateAccs(whichsol Soltype) (activity []float64,err error) {
  var _tmp97 C.int64_t
  if _tmp98 := C.MSK_getaccntot(self.ptr(),(&_tmp97)); _tmp98 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp98)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp99 *float64
  activity = make([]float64,_tmp97)
  if len(activity) > 0 { _tmp99 = (*float64)(&activity[0]) }
  if _tmp100 := C.MSK_evaluateaccs(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp99)); _tmp100 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp100)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Generates systematic names for affine conic constraints.
//
// - sub []int64
//   Indexes of the affine conic constraints.
// - fmt string
//   The variable name formatting string.
// - dims []int32
//   Dimensions in the shape.
// - sp []int64
//   Items that should be named.
// - namedaxisidxs []int32
//   List if named index axes
// - names []string
//   All axis names.
func (self *Task) GenerateAccNames(sub []int64,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp101 := len(sub)
  var num int64 = int64(_tmp101)
  var _tmp102 *int64
  if len(sub) > 0 { _tmp102 = (*int64)(&sub[0]) }
  _tmp103 := C.CString(fmt)
  _tmp104 := len(dims)
  var ndims int32 = int32(_tmp104)
  var _tmp105 *int32
  if len(dims) > 0 { _tmp105 = (*int32)(&dims[0]) }
  if sp != nil && int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateAccNames",arg:"sp"}
    return
  }
  var _tmp106 *int64
  if len(sp) > 0 { _tmp106 = (*int64)(&sp[0]) }
  _tmp107 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp107)
  var _tmp108 *int32
  if len(namedaxisidxs) > 0 { _tmp108 = (*int32)(&namedaxisidxs[0]) }
  _tmp109 := len(names)
  var numnames int64 = int64(_tmp109)
  _tmp111 := make([]*C.char,len(names))
  for _tmp112,_tmp113 := range names {
    _tmp111[_tmp112] = C.CString(_tmp113)
  }
  var _tmp110 **C.char;
  if len(_tmp111) > 0 { _tmp110 = &_tmp111[0] }
  if _tmp114 := C.MSK_generateaccnames(self.ptr(),C.int64_t(num),(*C.int64_t)(_tmp102),_tmp103,C.int32_t(ndims),(*C.int32_t)(_tmp105),(*C.int64_t)(_tmp106),C.int32_t(numnamedaxis),(*C.int32_t)(_tmp108),C.int64_t(numnames),((**C.char))(_tmp110)); _tmp114 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp114)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Generates systematic names for variables.
//
// - subj []int32
//   Indexes of the variables.
// - fmt string
//   The variable name formatting string.
// - dims []int32
//   Dimensions in the shape.
// - sp []int64
//   Items that should be named.
// - namedaxisidxs []int32
//   List if named index axes
// - names []string
//   All axis names.
func (self *Task) GenerateBarvarNames(subj []int32,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp115 := len(subj)
  var num int32 = int32(_tmp115)
  var _tmp116 *int32
  if len(subj) > 0 { _tmp116 = (*int32)(&subj[0]) }
  _tmp117 := C.CString(fmt)
  _tmp118 := len(dims)
  var ndims int32 = int32(_tmp118)
  var _tmp119 *int32
  if len(dims) > 0 { _tmp119 = (*int32)(&dims[0]) }
  if sp != nil && int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateBarvarNames",arg:"sp"}
    return
  }
  var _tmp120 *int64
  if len(sp) > 0 { _tmp120 = (*int64)(&sp[0]) }
  _tmp121 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp121)
  var _tmp122 *int32
  if len(namedaxisidxs) > 0 { _tmp122 = (*int32)(&namedaxisidxs[0]) }
  _tmp123 := len(names)
  var numnames int64 = int64(_tmp123)
  _tmp125 := make([]*C.char,len(names))
  for _tmp126,_tmp127 := range names {
    _tmp125[_tmp126] = C.CString(_tmp127)
  }
  var _tmp124 **C.char;
  if len(_tmp125) > 0 { _tmp124 = &_tmp125[0] }
  if _tmp128 := C.MSK_generatebarvarnames(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp116),_tmp117,C.int32_t(ndims),(*C.int32_t)(_tmp119),(*C.int64_t)(_tmp120),C.int32_t(numnamedaxis),(*C.int32_t)(_tmp122),C.int64_t(numnames),((**C.char))(_tmp124)); _tmp128 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp128)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Generates systematic names for cone.
//
// - subk []int32
//   Indexes of the cone.
// - fmt string
//   The cone name formatting string.
// - dims []int32
//   Dimensions in the shape.
// - sp []int64
//   Items that should be named.
// - namedaxisidxs []int32
//   List if named index axes
// - names []string
//   All axis names.
func (self *Task) GenerateConeNames(subk []int32,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp129 := len(subk)
  var num int32 = int32(_tmp129)
  var _tmp130 *int32
  if len(subk) > 0 { _tmp130 = (*int32)(&subk[0]) }
  _tmp131 := C.CString(fmt)
  _tmp132 := len(dims)
  var ndims int32 = int32(_tmp132)
  var _tmp133 *int32
  if len(dims) > 0 { _tmp133 = (*int32)(&dims[0]) }
  if sp != nil && int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateConeNames",arg:"sp"}
    return
  }
  var _tmp134 *int64
  if len(sp) > 0 { _tmp134 = (*int64)(&sp[0]) }
  _tmp135 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp135)
  var _tmp136 *int32
  if len(namedaxisidxs) > 0 { _tmp136 = (*int32)(&namedaxisidxs[0]) }
  _tmp137 := len(names)
  var numnames int64 = int64(_tmp137)
  _tmp139 := make([]*C.char,len(names))
  for _tmp140,_tmp141 := range names {
    _tmp139[_tmp140] = C.CString(_tmp141)
  }
  var _tmp138 **C.char;
  if len(_tmp139) > 0 { _tmp138 = &_tmp139[0] }
  if _tmp142 := C.MSK_generateconenames(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp130),_tmp131,C.int32_t(ndims),(*C.int32_t)(_tmp133),(*C.int64_t)(_tmp134),C.int32_t(numnamedaxis),(*C.int32_t)(_tmp136),C.int64_t(numnames),((**C.char))(_tmp138)); _tmp142 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp142)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Generates systematic names for constraints.
//
// - subi []int32
//   Indexes of the constraints.
// - fmt string
//   The constraint name formatting string.
// - dims []int32
//   Dimensions in the shape.
// - sp []int64
//   Items that should be named.
// - namedaxisidxs []int32
//   List if named index axes
// - names []string
//   All axis names.
func (self *Task) GenerateConNames(subi []int32,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp143 := len(subi)
  var num int32 = int32(_tmp143)
  var _tmp144 *int32
  if len(subi) > 0 { _tmp144 = (*int32)(&subi[0]) }
  _tmp145 := C.CString(fmt)
  _tmp146 := len(dims)
  var ndims int32 = int32(_tmp146)
  var _tmp147 *int32
  if len(dims) > 0 { _tmp147 = (*int32)(&dims[0]) }
  if sp != nil && int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateConNames",arg:"sp"}
    return
  }
  var _tmp148 *int64
  if len(sp) > 0 { _tmp148 = (*int64)(&sp[0]) }
  _tmp149 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp149)
  var _tmp150 *int32
  if len(namedaxisidxs) > 0 { _tmp150 = (*int32)(&namedaxisidxs[0]) }
  _tmp151 := len(names)
  var numnames int64 = int64(_tmp151)
  _tmp153 := make([]*C.char,len(names))
  for _tmp154,_tmp155 := range names {
    _tmp153[_tmp154] = C.CString(_tmp155)
  }
  var _tmp152 **C.char;
  if len(_tmp153) > 0 { _tmp152 = &_tmp153[0] }
  if _tmp156 := C.MSK_generateconnames(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp144),_tmp145,C.int32_t(ndims),(*C.int32_t)(_tmp147),(*C.int64_t)(_tmp148),C.int32_t(numnamedaxis),(*C.int32_t)(_tmp150),C.int64_t(numnames),((**C.char))(_tmp152)); _tmp156 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp156)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Generates systematic names for affine conic constraints.
//
// - sub []int64
//   Indexes of the disjunctive constraints.
// - fmt string
//   The variable name formatting string.
// - dims []int32
//   Dimensions in the shape.
// - sp []int64
//   Items that should be named.
// - namedaxisidxs []int32
//   List if named index axes
// - names []string
//   All axis names.
func (self *Task) GenerateDjcNames(sub []int64,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp157 := len(sub)
  var num int64 = int64(_tmp157)
  var _tmp158 *int64
  if len(sub) > 0 { _tmp158 = (*int64)(&sub[0]) }
  _tmp159 := C.CString(fmt)
  _tmp160 := len(dims)
  var ndims int32 = int32(_tmp160)
  var _tmp161 *int32
  if len(dims) > 0 { _tmp161 = (*int32)(&dims[0]) }
  if sp != nil && int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateDjcNames",arg:"sp"}
    return
  }
  var _tmp162 *int64
  if len(sp) > 0 { _tmp162 = (*int64)(&sp[0]) }
  _tmp163 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp163)
  var _tmp164 *int32
  if len(namedaxisidxs) > 0 { _tmp164 = (*int32)(&namedaxisidxs[0]) }
  _tmp165 := len(names)
  var numnames int64 = int64(_tmp165)
  _tmp167 := make([]*C.char,len(names))
  for _tmp168,_tmp169 := range names {
    _tmp167[_tmp168] = C.CString(_tmp169)
  }
  var _tmp166 **C.char;
  if len(_tmp167) > 0 { _tmp166 = &_tmp167[0] }
  if _tmp170 := C.MSK_generatedjcnames(self.ptr(),C.int64_t(num),(*C.int64_t)(_tmp158),_tmp159,C.int32_t(ndims),(*C.int32_t)(_tmp161),(*C.int64_t)(_tmp162),C.int32_t(numnamedaxis),(*C.int32_t)(_tmp164),C.int64_t(numnames),((**C.char))(_tmp166)); _tmp170 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp170)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Generates systematic names for variables.
//
// - subj []int32
//   Indexes of the variables.
// - fmt string
//   The variable name formatting string.
// - dims []int32
//   Dimensions in the shape.
// - sp []int64
//   Items that should be named.
// - namedaxisidxs []int32
//   List if named index axes
// - names []string
//   All axis names.
func (self *Task) GenerateVarNames(subj []int32,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp171 := len(subj)
  var num int32 = int32(_tmp171)
  var _tmp172 *int32
  if len(subj) > 0 { _tmp172 = (*int32)(&subj[0]) }
  _tmp173 := C.CString(fmt)
  _tmp174 := len(dims)
  var ndims int32 = int32(_tmp174)
  var _tmp175 *int32
  if len(dims) > 0 { _tmp175 = (*int32)(&dims[0]) }
  if sp != nil && int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateVarNames",arg:"sp"}
    return
  }
  var _tmp176 *int64
  if len(sp) > 0 { _tmp176 = (*int64)(&sp[0]) }
  _tmp177 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp177)
  var _tmp178 *int32
  if len(namedaxisidxs) > 0 { _tmp178 = (*int32)(&namedaxisidxs[0]) }
  _tmp179 := len(names)
  var numnames int64 = int64(_tmp179)
  _tmp181 := make([]*C.char,len(names))
  for _tmp182,_tmp183 := range names {
    _tmp181[_tmp182] = C.CString(_tmp183)
  }
  var _tmp180 **C.char;
  if len(_tmp181) > 0 { _tmp180 = &_tmp181[0] }
  if _tmp184 := C.MSK_generatevarnames(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp172),_tmp173,C.int32_t(ndims),(*C.int32_t)(_tmp175),(*C.int64_t)(_tmp176),C.int32_t(numnamedaxis),(*C.int32_t)(_tmp178),C.int64_t(numnames),((**C.char))(_tmp180)); _tmp184 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp184)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the list of affine expressions appearing in the affine conic constraint.
//
// - accidx int64
//   Index of the affine conic constraint.
// - afeidxlist []int64
//   List of indexes of affine expressions appearing in the constraint.
func (self *Task) GetAccAfeIdxList(accidx int64) (afeidxlist []int64,err error) {
  var _tmp185 C.int64_t
  if _tmp186 := C.MSK_getaccn(self.ptr(),(C.int64_t)(accidx),(&_tmp185)); _tmp186 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp186)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp187 *int64
  afeidxlist = make([]int64,_tmp185)
  if len(afeidxlist) > 0 { _tmp187 = (*int64)(&afeidxlist[0]) }
  if _tmp188 := C.MSK_getaccafeidxlist(self.ptr(),C.int64_t(accidx),(*C.int64_t)(_tmp187)); _tmp188 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp188)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the additional constant term vector appearing in the affine conic constraint.
//
// - accidx int64
//   Index of the affine conic constraint.
// - b []float64
//   The vector b appearing in the constraint.
func (self *Task) GetAccB(accidx int64) (b []float64,err error) {
  var _tmp189 C.int64_t
  if _tmp190 := C.MSK_getaccn(self.ptr(),(C.int64_t)(accidx),(&_tmp189)); _tmp190 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp190)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp191 *float64
  b = make([]float64,_tmp189)
  if len(b) > 0 { _tmp191 = (*float64)(&b[0]) }
  if _tmp192 := C.MSK_getaccb(self.ptr(),C.int64_t(accidx),(*C.double)(_tmp191)); _tmp192 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp192)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains barF, implied by the ACCs, in block triplet form.
//
// - numtrip int64
//   Number of elements in the block triplet form.
// - acc_afe []int64
//   Index of the AFE within the concatenated list of AFEs in ACCs.
// - bar_var []int32
//   Symmetric matrix variable index.
// - blk_row []int32
//   Block row index.
// - blk_col []int32
//   Block column index.
// - blk_val []float64
//   The numerical value associated with each block triplet.
func (self *Task) GetAccBarfBlockTriplet() (numtrip int64,acc_afe []int64,bar_var []int32,blk_row []int32,blk_col []int32,blk_val []float64,err error) {
  var _tmp193 C.int64_t
  if _tmp194 := C.MSK_getaccbarfnumblocktriplets(self.ptr(),(&_tmp193)); _tmp194 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp194)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumtrip int64 = int64(_tmp193)
  var _tmp195 *int64
  acc_afe = make([]int64,maxnumtrip)
  if len(acc_afe) > 0 { _tmp195 = (*int64)(&acc_afe[0]) }
  var _tmp196 *int32
  bar_var = make([]int32,maxnumtrip)
  if len(bar_var) > 0 { _tmp196 = (*int32)(&bar_var[0]) }
  var _tmp197 *int32
  blk_row = make([]int32,maxnumtrip)
  if len(blk_row) > 0 { _tmp197 = (*int32)(&blk_row[0]) }
  var _tmp198 *int32
  blk_col = make([]int32,maxnumtrip)
  if len(blk_col) > 0 { _tmp198 = (*int32)(&blk_col[0]) }
  var _tmp199 *float64
  blk_val = make([]float64,maxnumtrip)
  if len(blk_val) > 0 { _tmp199 = (*float64)(&blk_val[0]) }
  if _tmp200 := C.MSK_getaccbarfblocktriplet(self.ptr(),C.int64_t(maxnumtrip),(*C.int64_t)(&numtrip),(*C.int64_t)(_tmp195),(*C.int32_t)(_tmp196),(*C.int32_t)(_tmp197),(*C.int32_t)(_tmp198),(*C.double)(_tmp199)); _tmp200 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp200)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains an upper bound on the number of elements in the block triplet form of barf, as used within the ACCs.
//
// - numtrip int64
//   An upper bound on the number of elements in the block triplet form of barf, as used within the ACCs.
func (self *Task) GetAccBarfNumBlockTriplets() (numtrip int64,err error) {
  if _tmp201 := C.MSK_getaccbarfnumblocktriplets(self.ptr(),(*C.int64_t)(&numtrip)); _tmp201 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp201)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the domain appearing in the affine conic constraint.
//
// - accidx int64
//   The index of the affine conic constraint.
// - domidx int64
//   The index of domain in the affine conic constraint.
func (self *Task) GetAccDomain(accidx int64) (domidx int64,err error) {
  if _tmp202 := C.MSK_getaccdomain(self.ptr(),C.int64_t(accidx),(*C.int64_t)(&domidx)); _tmp202 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp202)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the doty vector for an affine conic constraint.
//
// - whichsol Soltype
//   Selects a solution.
// - accidx int64
//   The index of the affine conic constraint.
// - doty []float64
//   The dual values for this affine conic constraint. The array should have length equal to the dimension of the constraint.
func (self *Task) GetAccDotY(whichsol Soltype,accidx int64) (doty []float64,err error) {
  var _tmp203 C.int64_t
  if _tmp204 := C.MSK_getaccn(self.ptr(),(C.int64_t)(accidx),(&_tmp203)); _tmp204 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp204)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp205 *float64
  doty = make([]float64,_tmp203)
  if len(doty) > 0 { _tmp205 = (*float64)(&doty[0]) }
  if _tmp206 := C.MSK_getaccdoty(self.ptr(),C.int32_t(whichsol),C.int64_t(accidx),(*C.double)(_tmp205)); _tmp206 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp206)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the doty vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - doty []float64
//   The dual values of affine conic constraints. The array should have length equal to the sum of dimensions of all affine conic constraints.
func (self *Task) GetAccDotYS(whichsol Soltype) (doty []float64,err error) {
  var _tmp207 C.int64_t
  if _tmp208 := C.MSK_getaccntot(self.ptr(),(&_tmp207)); _tmp208 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp208)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp209 *float64
  doty = make([]float64,_tmp207)
  if len(doty) > 0 { _tmp209 = (*float64)(&doty[0]) }
  if _tmp210 := C.MSK_getaccdotys(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp209)); _tmp210 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp210)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the total number of nonzeros in the ACC implied F matrix.
//
// - accfnnz int64
//   Number of nonzeros in the F matrix implied by ACCs.
func (self *Task) GetAccFNumnz() (accfnnz int64,err error) {
  if _tmp211 := C.MSK_getaccfnumnz(self.ptr(),(*C.int64_t)(&accfnnz)); _tmp211 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp211)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the F matrix (implied by the AFE ordering within the ACCs) in triplet format.
//
// - frow []int64
//   Row indices of nonzeros in the implied F matrix.
// - fcol []int32
//   Column indices of nonzeros in the implied F matrix.
// - fval []float64
//   Values of nonzero entries in the implied F matrix.
func (self *Task) GetAccFTrip() (frow []int64,fcol []int32,fval []float64,err error) {
  var _tmp212 C.int64_t
  if _tmp213 := C.MSK_getaccfnumnz(self.ptr(),(&_tmp212)); _tmp213 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp213)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp214 *int64
  frow = make([]int64,_tmp212)
  if len(frow) > 0 { _tmp214 = (*int64)(&frow[0]) }
  var _tmp215 C.int64_t
  if _tmp216 := C.MSK_getaccfnumnz(self.ptr(),(&_tmp215)); _tmp216 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp216)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp217 *int32
  fcol = make([]int32,_tmp215)
  if len(fcol) > 0 { _tmp217 = (*int32)(&fcol[0]) }
  var _tmp218 C.int64_t
  if _tmp219 := C.MSK_getaccfnumnz(self.ptr(),(&_tmp218)); _tmp219 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp219)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp220 *float64
  fval = make([]float64,_tmp218)
  if len(fval) > 0 { _tmp220 = (*float64)(&fval[0]) }
  if _tmp221 := C.MSK_getaccftrip(self.ptr(),(*C.int64_t)(_tmp214),(*C.int32_t)(_tmp217),(*C.double)(_tmp220)); _tmp221 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp221)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// The g vector as used within the ACCs.
//
// - g []float64
//   The g vector as used within the ACCs.
func (self *Task) GetAccGVector() (g []float64,err error) {
  var _tmp222 C.int64_t
  if _tmp223 := C.MSK_getaccntot(self.ptr(),(&_tmp222)); _tmp223 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp223)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp224 *float64
  g = make([]float64,_tmp222)
  if len(g) > 0 { _tmp224 = (*float64)(&g[0]) }
  if _tmp225 := C.MSK_getaccgvector(self.ptr(),(*C.double)(_tmp224)); _tmp225 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp225)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the dimension of the affine conic constraint.
//
// - accidx int64
//   The index of the affine conic constraint.
// - n int64
//   The dimension of the affine conic constraint (equal to the dimension of its domain).
func (self *Task) GetAccN(accidx int64) (n int64,err error) {
  if _tmp226 := C.MSK_getaccn(self.ptr(),C.int64_t(accidx),(*C.int64_t)(&n)); _tmp226 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp226)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the name of an affine conic constraint.
//
// - accidx int64
//   Index of an affine conic constraint.
// - name string
//   Returns the required name.
func (self *Task) GetAccName(accidx int64) (name string,err error) {
  var _tmp227 C.int32_t
  if _tmp228 := C.MSK_getaccnamelen(self.ptr(),(C.int64_t)(accidx),(&_tmp227)); _tmp228 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp228)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp227))
  _tmp229 := make([]byte,sizename)
  if _tmp230 := C.MSK_getaccname(self.ptr(),C.int64_t(accidx),C.int32_t(sizename),(*C.uchar)(&_tmp229[0])); _tmp230 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp230)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp229,byte(0)); p < 0 {
    name = string(_tmp229)
  } else {
    name = string(_tmp229[:p])
  }
  return
}

// Obtains the length of the name of an affine conic constraint.
//
// - accidx int64
//   Index of an affine conic constraint.
// - len int32
//   Returns the length of the indicated name.
func (self *Task) GetAccNameLen(accidx int64) (len int32,err error) {
  if _tmp231 := C.MSK_getaccnamelen(self.ptr(),C.int64_t(accidx),(*C.int32_t)(&len)); _tmp231 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp231)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the total dimension of all affine conic constraints.
//
// - n int64
//   The total dimension of all affine conic constraints.
func (self *Task) GetAccNTot() (n int64,err error) {
  if _tmp232 := C.MSK_getaccntot(self.ptr(),(*C.int64_t)(&n)); _tmp232 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp232)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains full data of all affine conic constraints.
//
// - domidxlist []int64
//   The list of domains appearing in all affine conic constraints.
// - afeidxlist []int64
//   The concatenation of index lists of affine expressions appearing in all affine conic constraints.
// - b []float64
//   The concatenation of vectors b appearing in all affine conic constraints.
func (self *Task) GetAccs() (domidxlist []int64,afeidxlist []int64,b []float64,err error) {
  var _tmp233 C.int64_t
  if _tmp234 := C.MSK_getnumacc(self.ptr(),(&_tmp233)); _tmp234 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp234)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp235 *int64
  domidxlist = make([]int64,_tmp233)
  if len(domidxlist) > 0 { _tmp235 = (*int64)(&domidxlist[0]) }
  var _tmp236 C.int64_t
  if _tmp237 := C.MSK_getaccntot(self.ptr(),(&_tmp236)); _tmp237 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp237)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp238 *int64
  afeidxlist = make([]int64,_tmp236)
  if len(afeidxlist) > 0 { _tmp238 = (*int64)(&afeidxlist[0]) }
  var _tmp239 C.int64_t
  if _tmp240 := C.MSK_getaccntot(self.ptr(),(&_tmp239)); _tmp240 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp240)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp241 *float64
  b = make([]float64,_tmp239)
  if len(b) > 0 { _tmp241 = (*float64)(&b[0]) }
  if _tmp242 := C.MSK_getaccs(self.ptr(),(*C.int64_t)(_tmp235),(*C.int64_t)(_tmp238),(*C.double)(_tmp241)); _tmp242 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp242)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains one column of the linear constraint matrix.
//
// - j int32
//   Index of the column.
// - nzj int32
//   Number of non-zeros in the column obtained.
// - subj []int32
//   Row indices of the non-zeros in the column obtained.
// - valj []float64
//   Numerical values in the column obtained.
func (self *Task) GetACol(j int32) (nzj int32,subj []int32,valj []float64,err error) {
  var _tmp243 C.int32_t
  if _tmp244 := C.MSK_getacolnumnz(self.ptr(),(C.int32_t)(j),(&_tmp243)); _tmp244 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp244)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp245 *int32
  subj = make([]int32,_tmp243)
  if len(subj) > 0 { _tmp245 = (*int32)(&subj[0]) }
  var _tmp246 C.int32_t
  if _tmp247 := C.MSK_getacolnumnz(self.ptr(),(C.int32_t)(j),(&_tmp246)); _tmp247 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp247)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp248 *float64
  valj = make([]float64,_tmp246)
  if len(valj) > 0 { _tmp248 = (*float64)(&valj[0]) }
  if _tmp249 := C.MSK_getacol(self.ptr(),C.int32_t(j),(*C.int32_t)(&nzj),(*C.int32_t)(_tmp245),(*C.double)(_tmp248)); _tmp249 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp249)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of non-zero elements in one column of the linear constraint matrix
//
// - i int32
//   Index of the column.
// - nzj int32
//   Number of non-zeros in the j'th column of (A).
func (self *Task) GetAColNumNz(i int32) (nzj int32,err error) {
  if _tmp250 := C.MSK_getacolnumnz(self.ptr(),C.int32_t(i),(*C.int32_t)(&nzj)); _tmp250 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp250)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a sequence of columns from the coefficient matrix.
//
// - first int32
//   Index of the first column in the sequence.
// - last int32
//   Index of the last column in the sequence plus one.
// - ptrb []int64
//   Column start pointers.
// - ptre []int64
//   Column end pointers.
// - sub []int32
//   Contains the row subscripts.
// - val []float64
//   Contains the coefficient values.
func (self *Task) GetAColSlice(first int32,last int32) (ptrb []int64,ptre []int64,sub []int32,val []float64,err error) {
  var _tmp251 C.int64_t
  if _tmp252 := C.MSK_getacolslicenumnz64(self.ptr(),(C.int32_t)(first),(C.int32_t)(last),(&_tmp251)); _tmp252 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp252)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp251)
  var _tmp253 *int64
  ptrb = make([]int64,(last - first))
  if len(ptrb) > 0 { _tmp253 = (*int64)(&ptrb[0]) }
  var _tmp254 *int64
  ptre = make([]int64,(last - first))
  if len(ptre) > 0 { _tmp254 = (*int64)(&ptre[0]) }
  var _tmp255 *int32
  sub = make([]int32,maxnumnz)
  if len(sub) > 0 { _tmp255 = (*int32)(&sub[0]) }
  var _tmp256 *float64
  val = make([]float64,maxnumnz)
  if len(val) > 0 { _tmp256 = (*float64)(&val[0]) }
  if _tmp257 := C.MSK_getacolslice64(self.ptr(),C.int32_t(first),C.int32_t(last),C.int64_t(maxnumnz),(*C.int64_t)(_tmp253),(*C.int64_t)(_tmp254),(*C.int32_t)(_tmp255),(*C.double)(_tmp256)); _tmp257 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp257)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of non-zeros in a slice of columns of the coefficient matrix.
//
// - first int32
//   Index of the first column in the sequence.
// - last int32
//   Index of the last column plus one in the sequence.
// - numnz int64
//   Number of non-zeros in the slice.
func (self *Task) GetAColSliceNumNz(first int32,last int32) (numnz int64,err error) {
  if _tmp258 := C.MSK_getacolslicenumnz64(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int64_t)(&numnz)); _tmp258 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp258)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a sequence of columns from the coefficient matrix in triplet format.
//
// - first int32
//   Index of the first column in the sequence.
// - last int32
//   Index of the last column in the sequence plus one.
// - subi []int32
//   Constraint subscripts.
// - subj []int32
//   Column subscripts.
// - val []float64
//   Values.
func (self *Task) GetAColSliceTrip(first int32,last int32) (subi []int32,subj []int32,val []float64,err error) {
  var _tmp259 C.int64_t
  if _tmp260 := C.MSK_getacolslicenumnz64(self.ptr(),(C.int32_t)(first),(C.int32_t)(last),(&_tmp259)); _tmp260 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp260)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp259)
  var _tmp261 *int32
  subi = make([]int32,maxnumnz)
  if len(subi) > 0 { _tmp261 = (*int32)(&subi[0]) }
  var _tmp262 *int32
  subj = make([]int32,maxnumnz)
  if len(subj) > 0 { _tmp262 = (*int32)(&subj[0]) }
  var _tmp263 *float64
  val = make([]float64,maxnumnz)
  if len(val) > 0 { _tmp263 = (*float64)(&val[0]) }
  if _tmp264 := C.MSK_getacolslicetrip(self.ptr(),C.int32_t(first),C.int32_t(last),C.int64_t(maxnumnz),(*C.int32_t)(_tmp261),(*C.int32_t)(_tmp262),(*C.double)(_tmp263)); _tmp264 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp264)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains barF in block triplet form.
//
// - numtrip int64
//   Number of elements in the block triplet form.
// - afeidx []int64
//   Constraint index.
// - barvaridx []int32
//   Symmetric matrix variable index.
// - subk []int32
//   Block row index.
// - subl []int32
//   Block column index.
// - valkl []float64
//   The numerical value associated with each block triplet.
func (self *Task) GetAfeBarfBlockTriplet() (numtrip int64,afeidx []int64,barvaridx []int32,subk []int32,subl []int32,valkl []float64,err error) {
  var _tmp265 C.int64_t
  if _tmp266 := C.MSK_getafebarfnumblocktriplets(self.ptr(),(&_tmp265)); _tmp266 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp266)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumtrip int64 = int64(_tmp265)
  var _tmp267 *int64
  afeidx = make([]int64,maxnumtrip)
  if len(afeidx) > 0 { _tmp267 = (*int64)(&afeidx[0]) }
  var _tmp268 *int32
  barvaridx = make([]int32,maxnumtrip)
  if len(barvaridx) > 0 { _tmp268 = (*int32)(&barvaridx[0]) }
  var _tmp269 *int32
  subk = make([]int32,maxnumtrip)
  if len(subk) > 0 { _tmp269 = (*int32)(&subk[0]) }
  var _tmp270 *int32
  subl = make([]int32,maxnumtrip)
  if len(subl) > 0 { _tmp270 = (*int32)(&subl[0]) }
  var _tmp271 *float64
  valkl = make([]float64,maxnumtrip)
  if len(valkl) > 0 { _tmp271 = (*float64)(&valkl[0]) }
  if _tmp272 := C.MSK_getafebarfblocktriplet(self.ptr(),C.int64_t(maxnumtrip),(*C.int64_t)(&numtrip),(*C.int64_t)(_tmp267),(*C.int32_t)(_tmp268),(*C.int32_t)(_tmp269),(*C.int32_t)(_tmp270),(*C.double)(_tmp271)); _tmp272 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp272)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains an upper bound on the number of elements in the block triplet form of barf.
//
// - numtrip int64
//   An upper bound on the number of elements in the block triplet form of barf.
func (self *Task) GetAfeBarfNumBlockTriplets() (numtrip int64,err error) {
  if _tmp273 := C.MSK_getafebarfnumblocktriplets(self.ptr(),(*C.int64_t)(&numtrip)); _tmp273 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp273)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of nonzero entries in a row of barF.
//
// - afeidx int64
//   Row index of barF.
// - numentr int32
//   Number of nonzero entries in a row of barF.
func (self *Task) GetAfeBarfNumRowEntries(afeidx int64) (numentr int32,err error) {
  if _tmp274 := C.MSK_getafebarfnumrowentries(self.ptr(),C.int64_t(afeidx),(*C.int32_t)(&numentr)); _tmp274 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp274)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains nonzero entries in one row of barF.
//
// - afeidx int64
//   Row index of barF.
// - barvaridx []int32
//   Semidefinite variable indices.
// - ptrterm []int64
//   Pointers to the description of entries.
// - numterm []int64
//   Number of terms in each entry.
// - termidx []int64
//   Indices of semidefinite matrices from E.
// - termweight []float64
//   Weights appearing in the weighted sum representation.
func (self *Task) GetAfeBarfRow(afeidx int64) (barvaridx []int32,ptrterm []int64,numterm []int64,termidx []int64,termweight []float64,err error) {
  var _tmp275 C.int32_t
  var _tmp276 C.int64_t
  if _tmp277 := C.MSK_getafebarfrowinfo(self.ptr(),(C.int64_t)(afeidx),(&_tmp275),(&_tmp276)); _tmp277 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp277)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp278 *int32
  barvaridx = make([]int32,_tmp275)
  if len(barvaridx) > 0 { _tmp278 = (*int32)(&barvaridx[0]) }
  var _tmp279 C.int32_t
  var _tmp280 C.int64_t
  if _tmp281 := C.MSK_getafebarfrowinfo(self.ptr(),(C.int64_t)(afeidx),(&_tmp279),(&_tmp280)); _tmp281 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp281)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp282 *int64
  ptrterm = make([]int64,_tmp279)
  if len(ptrterm) > 0 { _tmp282 = (*int64)(&ptrterm[0]) }
  var _tmp283 C.int32_t
  var _tmp284 C.int64_t
  if _tmp285 := C.MSK_getafebarfrowinfo(self.ptr(),(C.int64_t)(afeidx),(&_tmp283),(&_tmp284)); _tmp285 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp285)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp286 *int64
  numterm = make([]int64,_tmp283)
  if len(numterm) > 0 { _tmp286 = (*int64)(&numterm[0]) }
  var _tmp287 C.int32_t
  var _tmp288 C.int64_t
  if _tmp289 := C.MSK_getafebarfrowinfo(self.ptr(),(C.int64_t)(afeidx),(&_tmp287),(&_tmp288)); _tmp289 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp289)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp290 *int64
  termidx = make([]int64,_tmp288)
  if len(termidx) > 0 { _tmp290 = (*int64)(&termidx[0]) }
  var _tmp291 C.int32_t
  var _tmp292 C.int64_t
  if _tmp293 := C.MSK_getafebarfrowinfo(self.ptr(),(C.int64_t)(afeidx),(&_tmp291),(&_tmp292)); _tmp293 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp293)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp294 *float64
  termweight = make([]float64,_tmp292)
  if len(termweight) > 0 { _tmp294 = (*float64)(&termweight[0]) }
  if _tmp295 := C.MSK_getafebarfrow(self.ptr(),C.int64_t(afeidx),(*C.int32_t)(_tmp278),(*C.int64_t)(_tmp282),(*C.int64_t)(_tmp286),(*C.int64_t)(_tmp290),(*C.double)(_tmp294)); _tmp295 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp295)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains information about one row of barF.
//
// - afeidx int64
//   Row index of barF.
// - numentr int32
//   Number of nonzero entries in a row of barF.
// - numterm int64
//   Number of terms in the weighted sums representation of the row of barF.
func (self *Task) GetAfeBarfRowInfo(afeidx int64) (numentr int32,numterm int64,err error) {
  if _tmp296 := C.MSK_getafebarfrowinfo(self.ptr(),C.int64_t(afeidx),(*C.int32_t)(&numentr),(*C.int64_t)(&numterm)); _tmp296 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp296)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the total number of nonzeros in F.
//
// - numnz int64
//   Number of nonzeros in F.
func (self *Task) GetAfeFNumNz() (numnz int64,err error) {
  if _tmp297 := C.MSK_getafefnumnz(self.ptr(),(*C.int64_t)(&numnz)); _tmp297 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp297)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains one row of F in sparse format.
//
// - afeidx int64
//   Row index.
// - numnz int32
//   Number of non-zeros in the row obtained.
// - varidx []int32
//   Column indices of the non-zeros in the row obtained.
// - val []float64
//   Values of the non-zeros in the row obtained.
func (self *Task) GetAfeFRow(afeidx int64) (numnz int32,varidx []int32,val []float64,err error) {
  var _tmp298 C.int32_t
  if _tmp299 := C.MSK_getafefrownumnz(self.ptr(),(C.int64_t)(afeidx),(&_tmp298)); _tmp299 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp299)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp300 *int32
  varidx = make([]int32,_tmp298)
  if len(varidx) > 0 { _tmp300 = (*int32)(&varidx[0]) }
  var _tmp301 C.int32_t
  if _tmp302 := C.MSK_getafefrownumnz(self.ptr(),(C.int64_t)(afeidx),(&_tmp301)); _tmp302 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp302)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp303 *float64
  val = make([]float64,_tmp301)
  if len(val) > 0 { _tmp303 = (*float64)(&val[0]) }
  if _tmp304 := C.MSK_getafefrow(self.ptr(),C.int64_t(afeidx),(*C.int32_t)(&numnz),(*C.int32_t)(_tmp300),(*C.double)(_tmp303)); _tmp304 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp304)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of nonzeros in a row of F.
//
// - afeidx int64
//   Row index.
// - numnz int32
//   Number of non-zeros in the row.
func (self *Task) GetAfeFRowNumNz(afeidx int64) (numnz int32,err error) {
  if _tmp305 := C.MSK_getafefrownumnz(self.ptr(),C.int64_t(afeidx),(*C.int32_t)(&numnz)); _tmp305 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp305)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the F matrix in triplet format.
//
// - afeidx []int64
//   Row indices of nonzeros.
// - varidx []int32
//   Column indices of nonzeros.
// - val []float64
//   Values of nonzero entries.
func (self *Task) GetAfeFTrip() (afeidx []int64,varidx []int32,val []float64,err error) {
  var _tmp306 C.int64_t
  if _tmp307 := C.MSK_getafefnumnz(self.ptr(),(&_tmp306)); _tmp307 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp307)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp308 *int64
  afeidx = make([]int64,_tmp306)
  if len(afeidx) > 0 { _tmp308 = (*int64)(&afeidx[0]) }
  var _tmp309 C.int64_t
  if _tmp310 := C.MSK_getafefnumnz(self.ptr(),(&_tmp309)); _tmp310 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp310)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp311 *int32
  varidx = make([]int32,_tmp309)
  if len(varidx) > 0 { _tmp311 = (*int32)(&varidx[0]) }
  var _tmp312 C.int64_t
  if _tmp313 := C.MSK_getafefnumnz(self.ptr(),(&_tmp312)); _tmp313 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp313)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp314 *float64
  val = make([]float64,_tmp312)
  if len(val) > 0 { _tmp314 = (*float64)(&val[0]) }
  if _tmp315 := C.MSK_getafeftrip(self.ptr(),(*C.int64_t)(_tmp308),(*C.int32_t)(_tmp311),(*C.double)(_tmp314)); _tmp315 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp315)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a single coefficient in g.
//
// - afeidx int64
//   Element index.
// - g float64
//   The entry in g.
func (self *Task) GetAfeG(afeidx int64) (g float64,err error) {
  if _tmp316 := C.MSK_getafeg(self.ptr(),C.int64_t(afeidx),(*C.double)(&g)); _tmp316 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp316)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a sequence of coefficients from the vector g.
//
// - first int64
//   First index in the sequence.
// - last int64
//   Last index plus 1 in the sequence.
// - g []float64
//   The slice of g as a dense vector.
func (self *Task) GetAfeGSlice(first int64,last int64) (g []float64,err error) {
  var _tmp317 *float64
  g = make([]float64,(last - first))
  if len(g) > 0 { _tmp317 = (*float64)(&g[0]) }
  if _tmp318 := C.MSK_getafegslice(self.ptr(),C.int64_t(first),C.int64_t(last),(*C.double)(_tmp317)); _tmp318 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp318)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a single coefficient in linear constraint matrix.
//
// - i int32
//   Row index of the coefficient to be returned.
// - j int32
//   Column index of the coefficient to be returned.
// - aij float64
//   Returns the requested coefficient.
func (self *Task) GetAij(i int32,j int32) (aij float64,err error) {
  if _tmp319 := C.MSK_getaij(self.ptr(),C.int32_t(i),C.int32_t(j),(*C.double)(&aij)); _tmp319 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp319)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number non-zeros in a rectangular piece of the linear constraint matrix.
//
// - firsti int32
//   Index of the first row in the rectangular piece.
// - lasti int32
//   Index of the last row plus one in the rectangular piece.
// - firstj int32
//   Index of the first column in the rectangular piece.
// - lastj int32
//   Index of the last column plus one in the rectangular piece.
// - numnz int32
//   Number of non-zero elements in the rectangular piece of the linear constraint matrix.
func (self *Task) GetAPieceNumNz(firsti int32,lasti int32,firstj int32,lastj int32) (numnz int32,err error) {
  if _tmp320 := C.MSK_getapiecenumnz(self.ptr(),C.int32_t(firsti),C.int32_t(lasti),C.int32_t(firstj),C.int32_t(lastj),(*C.int32_t)(&numnz)); _tmp320 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp320)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains one row of the linear constraint matrix.
//
// - i int32
//   Index of the row.
// - nzi int32
//   Number of non-zeros in the row obtained.
// - subi []int32
//   Column indices of the non-zeros in the row obtained.
// - vali []float64
//   Numerical values of the row obtained.
func (self *Task) GetARow(i int32) (nzi int32,subi []int32,vali []float64,err error) {
  var _tmp321 C.int32_t
  if _tmp322 := C.MSK_getarownumnz(self.ptr(),(C.int32_t)(i),(&_tmp321)); _tmp322 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp322)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp323 *int32
  subi = make([]int32,_tmp321)
  if len(subi) > 0 { _tmp323 = (*int32)(&subi[0]) }
  var _tmp324 C.int32_t
  if _tmp325 := C.MSK_getarownumnz(self.ptr(),(C.int32_t)(i),(&_tmp324)); _tmp325 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp325)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp326 *float64
  vali = make([]float64,_tmp324)
  if len(vali) > 0 { _tmp326 = (*float64)(&vali[0]) }
  if _tmp327 := C.MSK_getarow(self.ptr(),C.int32_t(i),(*C.int32_t)(&nzi),(*C.int32_t)(_tmp323),(*C.double)(_tmp326)); _tmp327 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp327)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of non-zero elements in one row of the linear constraint matrix
//
// - i int32
//   Index of the row.
// - nzi int32
//   Number of non-zeros in the i'th row of `A`.
func (self *Task) GetARowNumNz(i int32) (nzi int32,err error) {
  if _tmp328 := C.MSK_getarownumnz(self.ptr(),C.int32_t(i),(*C.int32_t)(&nzi)); _tmp328 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp328)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a sequence of rows from the coefficient matrix.
//
// - first int32
//   Index of the first row in the sequence.
// - last int32
//   Index of the last row in the sequence plus one.
// - ptrb []int64
//   Row start pointers.
// - ptre []int64
//   Row end pointers.
// - sub []int32
//   Contains the column subscripts.
// - val []float64
//   Contains the coefficient values.
func (self *Task) GetARowSlice(first int32,last int32) (ptrb []int64,ptre []int64,sub []int32,val []float64,err error) {
  var _tmp329 C.int64_t
  if _tmp330 := C.MSK_getarowslicenumnz64(self.ptr(),(C.int32_t)(first),(C.int32_t)(last),(&_tmp329)); _tmp330 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp330)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp329)
  var _tmp331 *int64
  ptrb = make([]int64,(last - first))
  if len(ptrb) > 0 { _tmp331 = (*int64)(&ptrb[0]) }
  var _tmp332 *int64
  ptre = make([]int64,(last - first))
  if len(ptre) > 0 { _tmp332 = (*int64)(&ptre[0]) }
  var _tmp333 *int32
  sub = make([]int32,maxnumnz)
  if len(sub) > 0 { _tmp333 = (*int32)(&sub[0]) }
  var _tmp334 *float64
  val = make([]float64,maxnumnz)
  if len(val) > 0 { _tmp334 = (*float64)(&val[0]) }
  if _tmp335 := C.MSK_getarowslice64(self.ptr(),C.int32_t(first),C.int32_t(last),C.int64_t(maxnumnz),(*C.int64_t)(_tmp331),(*C.int64_t)(_tmp332),(*C.int32_t)(_tmp333),(*C.double)(_tmp334)); _tmp335 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp335)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of non-zeros in a slice of rows of the coefficient matrix.
//
// - first int32
//   Index of the first row in the sequence.
// - last int32
//   Index of the last row plus one in the sequence.
// - numnz int64
//   Number of non-zeros in the slice.
func (self *Task) GetARowSliceNumNz(first int32,last int32) (numnz int64,err error) {
  if _tmp336 := C.MSK_getarowslicenumnz64(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int64_t)(&numnz)); _tmp336 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp336)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a sequence of rows from the coefficient matrix in sparse triplet format.
//
// - first int32
//   Index of the first row in the sequence.
// - last int32
//   Index of the last row in the sequence plus one.
// - subi []int32
//   Constraint subscripts.
// - subj []int32
//   Column subscripts.
// - val []float64
//   Values.
func (self *Task) GetARowSliceTrip(first int32,last int32) (subi []int32,subj []int32,val []float64,err error) {
  var _tmp337 C.int64_t
  if _tmp338 := C.MSK_getarowslicenumnz64(self.ptr(),(C.int32_t)(first),(C.int32_t)(last),(&_tmp337)); _tmp338 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp338)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp337)
  var _tmp339 *int32
  subi = make([]int32,maxnumnz)
  if len(subi) > 0 { _tmp339 = (*int32)(&subi[0]) }
  var _tmp340 *int32
  subj = make([]int32,maxnumnz)
  if len(subj) > 0 { _tmp340 = (*int32)(&subj[0]) }
  var _tmp341 *float64
  val = make([]float64,maxnumnz)
  if len(val) > 0 { _tmp341 = (*float64)(&val[0]) }
  if _tmp342 := C.MSK_getarowslicetrip(self.ptr(),C.int32_t(first),C.int32_t(last),C.int64_t(maxnumnz),(*C.int32_t)(_tmp339),(*C.int32_t)(_tmp340),(*C.double)(_tmp341)); _tmp342 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp342)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the A matrix in sparse triplet format.
//
// - subi []int32
//   Constraint subscripts.
// - subj []int32
//   Column subscripts.
// - val []float64
//   Values.
func (self *Task) GetATrip() (subi []int32,subj []int32,val []float64,err error) {
  var _tmp343 C.int64_t
  if _tmp344 := C.MSK_getnumanz64(self.ptr(),(&_tmp343)); _tmp344 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp344)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp343)
  var _tmp345 *int32
  subi = make([]int32,maxnumnz)
  if len(subi) > 0 { _tmp345 = (*int32)(&subi[0]) }
  var _tmp346 *int32
  subj = make([]int32,maxnumnz)
  if len(subj) > 0 { _tmp346 = (*int32)(&subj[0]) }
  var _tmp347 *float64
  val = make([]float64,maxnumnz)
  if len(val) > 0 { _tmp347 = (*float64)(&val[0]) }
  if _tmp348 := C.MSK_getatrip(self.ptr(),C.int64_t(maxnumnz),(*C.int32_t)(_tmp345),(*C.int32_t)(_tmp346),(*C.double)(_tmp347)); _tmp348 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp348)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Gets the current A matrix truncation threshold.
//
// - tolzero []float64
//   Truncation tolerance.
func (self *Task) GetATruncateTol() (tolzero []float64,err error) {
  var _tmp349 *float64
  tolzero = make([]float64,1)
  if len(tolzero) > 0 { _tmp349 = (*float64)(&tolzero[0]) }
  if _tmp350 := C.MSK_getatruncatetol(self.ptr(),(*C.double)(_tmp349)); _tmp350 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp350)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains barA in block triplet form.
//
// - num int64
//   Number of elements in the block triplet form.
// - subi []int32
//   Constraint index.
// - subj []int32
//   Symmetric matrix variable index.
// - subk []int32
//   Block row index.
// - subl []int32
//   Block column index.
// - valijkl []float64
//   The numerical value associated with each block triplet.
func (self *Task) GetBaraBlockTriplet() (num int64,subi []int32,subj []int32,subk []int32,subl []int32,valijkl []float64,err error) {
  var _tmp351 C.int64_t
  if _tmp352 := C.MSK_getnumbarablocktriplets(self.ptr(),(&_tmp351)); _tmp352 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp352)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnum int64 = int64(_tmp351)
  var _tmp353 *int32
  subi = make([]int32,maxnum)
  if len(subi) > 0 { _tmp353 = (*int32)(&subi[0]) }
  var _tmp354 *int32
  subj = make([]int32,maxnum)
  if len(subj) > 0 { _tmp354 = (*int32)(&subj[0]) }
  var _tmp355 *int32
  subk = make([]int32,maxnum)
  if len(subk) > 0 { _tmp355 = (*int32)(&subk[0]) }
  var _tmp356 *int32
  subl = make([]int32,maxnum)
  if len(subl) > 0 { _tmp356 = (*int32)(&subl[0]) }
  var _tmp357 *float64
  valijkl = make([]float64,maxnum)
  if len(valijkl) > 0 { _tmp357 = (*float64)(&valijkl[0]) }
  if _tmp358 := C.MSK_getbarablocktriplet(self.ptr(),C.int64_t(maxnum),(*C.int64_t)(&num),(*C.int32_t)(_tmp353),(*C.int32_t)(_tmp354),(*C.int32_t)(_tmp355),(*C.int32_t)(_tmp356),(*C.double)(_tmp357)); _tmp358 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp358)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains information about an element in barA.
//
// - idx int64
//   Position of the element in the vectorized form.
// - i int32
//   Row index of the element at position idx.
// - j int32
//   Column index of the element at position idx.
// - num int64
//   Number of terms in weighted sum that forms the element.
// - sub []int64
//   A list indexes of the elements from symmetric matrix storage that appear in the weighted sum.
// - weights []float64
//   The weights associated with each term in the weighted sum.
func (self *Task) GetBaraIdx(idx int64) (i int32,j int32,num int64,sub []int64,weights []float64,err error) {
  var _tmp359 C.int64_t
  if _tmp360 := C.MSK_getbaraidxinfo(self.ptr(),(C.int64_t)(idx),(&_tmp359)); _tmp360 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp360)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnum int64 = int64(_tmp359)
  var _tmp361 *int64
  sub = make([]int64,maxnum)
  if len(sub) > 0 { _tmp361 = (*int64)(&sub[0]) }
  var _tmp362 *float64
  weights = make([]float64,maxnum)
  if len(weights) > 0 { _tmp362 = (*float64)(&weights[0]) }
  if _tmp363 := C.MSK_getbaraidx(self.ptr(),C.int64_t(idx),C.int64_t(maxnum),(*C.int32_t)(&i),(*C.int32_t)(&j),(*C.int64_t)(&num),(*C.int64_t)(_tmp361),(*C.double)(_tmp362)); _tmp363 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp363)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains information about an element in barA.
//
// - idx int64
//   Position of the element in the vectorized form.
// - i int32
//   Row index of the element at position idx.
// - j int32
//   Column index of the element at position idx.
func (self *Task) GetBaraIdxIJ(idx int64) (i int32,j int32,err error) {
  if _tmp364 := C.MSK_getbaraidxij(self.ptr(),C.int64_t(idx),(*C.int32_t)(&i),(*C.int32_t)(&j)); _tmp364 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp364)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of terms in the weighted sum that form a particular element in barA.
//
// - idx int64
//   The internal position of the element for which information should be obtained.
// - num int64
//   Number of terms in the weighted sum that form the specified element in barA.
func (self *Task) GetBaraIdxInfo(idx int64) (num int64,err error) {
  if _tmp365 := C.MSK_getbaraidxinfo(self.ptr(),C.int64_t(idx),(*C.int64_t)(&num)); _tmp365 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp365)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the sparsity pattern of the barA matrix.
//
// - numnz int64
//   Number of nonzero elements in barA.
// - idxij []int64
//   Position of each nonzero element in the vector representation of barA.
func (self *Task) GetBaraSparsity() (numnz int64,idxij []int64,err error) {
  var _tmp366 C.int64_t
  if _tmp367 := C.MSK_getnumbaranz(self.ptr(),(&_tmp366)); _tmp367 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp367)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp366)
  var _tmp368 *int64
  idxij = make([]int64,maxnumnz)
  if len(idxij) > 0 { _tmp368 = (*int64)(&idxij[0]) }
  if _tmp369 := C.MSK_getbarasparsity(self.ptr(),C.int64_t(maxnumnz),(*C.int64_t)(&numnz),(*C.int64_t)(_tmp368)); _tmp369 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp369)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains barC in block triplet form.
//
// - num int64
//   Number of elements in the block triplet form.
// - subj []int32
//   Symmetric matrix variable index.
// - subk []int32
//   Block row index.
// - subl []int32
//   Block column index.
// - valjkl []float64
//   The numerical value associated with each block triplet.
func (self *Task) GetBarcBlockTriplet() (num int64,subj []int32,subk []int32,subl []int32,valjkl []float64,err error) {
  var _tmp370 C.int64_t
  if _tmp371 := C.MSK_getnumbarcblocktriplets(self.ptr(),(&_tmp370)); _tmp371 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp371)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnum int64 = int64(_tmp370)
  var _tmp372 *int32
  subj = make([]int32,maxnum)
  if len(subj) > 0 { _tmp372 = (*int32)(&subj[0]) }
  var _tmp373 *int32
  subk = make([]int32,maxnum)
  if len(subk) > 0 { _tmp373 = (*int32)(&subk[0]) }
  var _tmp374 *int32
  subl = make([]int32,maxnum)
  if len(subl) > 0 { _tmp374 = (*int32)(&subl[0]) }
  var _tmp375 *float64
  valjkl = make([]float64,maxnum)
  if len(valjkl) > 0 { _tmp375 = (*float64)(&valjkl[0]) }
  if _tmp376 := C.MSK_getbarcblocktriplet(self.ptr(),C.int64_t(maxnum),(*C.int64_t)(&num),(*C.int32_t)(_tmp372),(*C.int32_t)(_tmp373),(*C.int32_t)(_tmp374),(*C.double)(_tmp375)); _tmp376 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp376)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains information about an element in barc.
//
// - idx int64
//   Index of the element for which information should be obtained.
// - j int32
//   Row index in barc.
// - num int64
//   Number of terms in the weighted sum.
// - sub []int64
//   Elements appearing the weighted sum.
// - weights []float64
//   Weights of terms in the weighted sum.
func (self *Task) GetBarcIdx(idx int64) (j int32,num int64,sub []int64,weights []float64,err error) {
  var _tmp377 C.int64_t
  if _tmp378 := C.MSK_getbarcidxinfo(self.ptr(),(C.int64_t)(idx),(&_tmp377)); _tmp378 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp378)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnum int64 = int64(_tmp377)
  var _tmp379 *int64
  sub = make([]int64,maxnum)
  if len(sub) > 0 { _tmp379 = (*int64)(&sub[0]) }
  var _tmp380 *float64
  weights = make([]float64,maxnum)
  if len(weights) > 0 { _tmp380 = (*float64)(&weights[0]) }
  if _tmp381 := C.MSK_getbarcidx(self.ptr(),C.int64_t(idx),C.int64_t(maxnum),(*C.int32_t)(&j),(*C.int64_t)(&num),(*C.int64_t)(_tmp379),(*C.double)(_tmp380)); _tmp381 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp381)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains information about an element in barc.
//
// - idx int64
//   Index of the element for which information should be obtained. The value is an index of a symmetric sparse variable.
// - num int64
//   Number of terms that appear in the weighted sum that forms the requested element.
func (self *Task) GetBarcIdxInfo(idx int64) (num int64,err error) {
  if _tmp382 := C.MSK_getbarcidxinfo(self.ptr(),C.int64_t(idx),(*C.int64_t)(&num)); _tmp382 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp382)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the row index of an element in barc.
//
// - idx int64
//   Index of the element for which information should be obtained.
// - j int32
//   Row index in barc.
func (self *Task) GetBarcIdxJ(idx int64) (j int32,err error) {
  if _tmp383 := C.MSK_getbarcidxj(self.ptr(),C.int64_t(idx),(*C.int32_t)(&j)); _tmp383 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp383)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Get the positions of the nonzero elements in barc.
//
// - numnz int64
//   Number of nonzero elements in barc.
// - idxj []int64
//   Internal positions of the nonzeros elements in barc.
func (self *Task) GetBarcSparsity() (numnz int64,idxj []int64,err error) {
  var _tmp384 C.int64_t
  if _tmp385 := C.MSK_getnumbarcnz(self.ptr(),(&_tmp384)); _tmp385 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp385)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp384)
  var _tmp386 *int64
  idxj = make([]int64,maxnumnz)
  if len(idxj) > 0 { _tmp386 = (*int64)(&idxj[0]) }
  if _tmp387 := C.MSK_getbarcsparsity(self.ptr(),C.int64_t(maxnumnz),(*C.int64_t)(&numnz),(*C.int64_t)(_tmp386)); _tmp387 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp387)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the dual solution for a semidefinite variable.
//
// - whichsol Soltype
//   Selects a solution.
// - j int32
//   Index of the semidefinite variable.
// - barsj []float64
//   Value of the j'th dual variable of barx.
func (self *Task) GetBarsJ(whichsol Soltype,j int32) (barsj []float64,err error) {
  var _tmp388 C.int64_t
  if _tmp389 := C.MSK_getlenbarvarj(self.ptr(),(C.int32_t)(j),(&_tmp388)); _tmp389 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp389)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp390 *float64
  barsj = make([]float64,_tmp388)
  if len(barsj) > 0 { _tmp390 = (*float64)(&barsj[0]) }
  if _tmp391 := C.MSK_getbarsj(self.ptr(),C.int32_t(whichsol),C.int32_t(j),(*C.double)(_tmp390)); _tmp391 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp391)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the dual solution for a sequence of semidefinite variables.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   Index of the first semidefinite variable in the slice.
// - last int32
//   Index of the last semidefinite variable in the slice plus one.
// - slicesize int64
//   Denotes the length of the array barsslice.
// - barsslice []float64
//   Dual solution values of symmetric matrix variables in the slice, stored sequentially.
func (self *Task) GetBarsSlice(whichsol Soltype,first int32,last int32,slicesize int64) (barsslice []float64,err error) {
  var _tmp392 *float64
  barsslice = make([]float64,slicesize)
  if len(barsslice) > 0 { _tmp392 = (*float64)(&barsslice[0]) }
  if _tmp393 := C.MSK_getbarsslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),C.int64_t(slicesize),(*C.double)(_tmp392)); _tmp393 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp393)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the name of a semidefinite variable.
//
// - i int32
//   Index of the variable.
// - name string
//   The requested name is copied to this buffer.
func (self *Task) GetBarvarName(i int32) (name string,err error) {
  var _tmp394 C.int32_t
  if _tmp395 := C.MSK_getbarvarnamelen(self.ptr(),(C.int32_t)(i),(&_tmp394)); _tmp395 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp395)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp394))
  _tmp396 := make([]byte,sizename)
  if _tmp397 := C.MSK_getbarvarname(self.ptr(),C.int32_t(i),C.int32_t(sizename),(*C.uchar)(&_tmp396[0])); _tmp397 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp397)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp396,byte(0)); p < 0 {
    name = string(_tmp396)
  } else {
    name = string(_tmp396[:p])
  }
  return
}

// Obtains the index of semidefinite variable from its name.
//
// - somename string
//   The name of the variable.
// - asgn int32
//   Non-zero if the name somename is assigned to some semidefinite variable.
// - index int32
//   The index of a semidefinite variable with the name somename (if one exists).
func (self *Task) GetBarvarNameIndex(somename string) (asgn int32,index int32,err error) {
  _tmp398 := C.CString(somename)
  if _tmp399 := C.MSK_getbarvarnameindex(self.ptr(),_tmp398,(*C.int32_t)(&asgn),(*C.int32_t)(&index)); _tmp399 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp399)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the length of the name of a semidefinite variable.
//
// - i int32
//   Index of the variable.
// - len int32
//   Returns the length of the indicated name.
func (self *Task) GetBarvarNameLen(i int32) (len int32,err error) {
  if _tmp400 := C.MSK_getbarvarnamelen(self.ptr(),C.int32_t(i),(*C.int32_t)(&len)); _tmp400 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp400)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the primal solution for a semidefinite variable.
//
// - whichsol Soltype
//   Selects a solution.
// - j int32
//   Index of the semidefinite variable.
// - barxj []float64
//   Value of the j'th variable of barx.
func (self *Task) GetBarxJ(whichsol Soltype,j int32) (barxj []float64,err error) {
  var _tmp401 C.int64_t
  if _tmp402 := C.MSK_getlenbarvarj(self.ptr(),(C.int32_t)(j),(&_tmp401)); _tmp402 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp402)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp403 *float64
  barxj = make([]float64,_tmp401)
  if len(barxj) > 0 { _tmp403 = (*float64)(&barxj[0]) }
  if _tmp404 := C.MSK_getbarxj(self.ptr(),C.int32_t(whichsol),C.int32_t(j),(*C.double)(_tmp403)); _tmp404 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp404)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the primal solution for a sequence of semidefinite variables.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   Index of the first semidefinite variable in the slice.
// - last int32
//   Index of the last semidefinite variable in the slice plus one.
// - slicesize int64
//   Denotes the length of the array barxslice.
// - barxslice []float64
//   Solution values of symmetric matrix variables in the slice, stored sequentially.
func (self *Task) GetBarxSlice(whichsol Soltype,first int32,last int32,slicesize int64) (barxslice []float64,err error) {
  var _tmp405 *float64
  barxslice = make([]float64,slicesize)
  if len(barxslice) > 0 { _tmp405 = (*float64)(&barxslice[0]) }
  if _tmp406 := C.MSK_getbarxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),C.int64_t(slicesize),(*C.double)(_tmp405)); _tmp406 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp406)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains all objective coefficients.
//
// - c []float64
//   Linear terms of the objective as a dense vector. The length is the number of variables.
func (self *Task) GetC() (c []float64,err error) {
  var _tmp407 C.int32_t
  if _tmp408 := C.MSK_getnumvar(self.ptr(),(&_tmp407)); _tmp408 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp408)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp409 *float64
  c = make([]float64,_tmp407)
  if len(c) > 0 { _tmp409 = (*float64)(&c[0]) }
  if _tmp410 := C.MSK_getc(self.ptr(),(*C.double)(_tmp409)); _tmp410 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp410)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the fixed term in the objective.
//
// - cfix float64
//   Fixed term in the objective.
func (self *Task) GetCfix() (cfix float64,err error) {
  if _tmp411 := C.MSK_getcfix(self.ptr(),(*C.double)(&cfix)); _tmp411 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp411)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains one objective coefficient.
//
// - j int32
//   Index of the variable for which the c coefficient should be obtained.
// - cj float64
//   The c coefficient value.
func (self *Task) GetCJ(j int32) (cj float64,err error) {
  if _tmp412 := C.MSK_getcj(self.ptr(),C.int32_t(j),(*C.double)(&cj)); _tmp412 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp412)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a sequence of coefficients from the objective.
//
// - subj []int32
//   A list of variable indexes.
// - c []float64
//   Linear terms of the requested list of the objective as a dense vector.
func (self *Task) GetCList(subj []int32) (c []float64,err error) {
  _tmp413 := len(subj)
  var num int32 = int32(_tmp413)
  var _tmp414 *int32
  if len(subj) > 0 { _tmp414 = (*int32)(&subj[0]) }
  var _tmp415 *float64
  c = make([]float64,num)
  if len(c) > 0 { _tmp415 = (*float64)(&c[0]) }
  if _tmp416 := C.MSK_getclist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp414),(*C.double)(_tmp415)); _tmp416 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp416)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains bound information for one constraint.
//
// - i int32
//   Index of the constraint for which the bound information should be obtained.
// - bk Boundkey
//   Bound keys.
// - bl float64
//   Values for lower bounds.
// - bu float64
//   Values for upper bounds.
func (self *Task) GetConBound(i int32) (bk Boundkey,bl float64,bu float64,err error) {
  var _tmp417 int32
  if _tmp418 := C.MSK_getconbound(self.ptr(),C.int32_t(i),(*C.int32_t)(&_tmp417),(*C.double)(&bl),(*C.double)(&bu)); _tmp418 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp418)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  bk = Boundkey(_tmp417)
  return
}

// Obtains bounds information for a slice of the constraints.
//
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - bk []Boundkey
//   Bound keys.
// - bl []float64
//   Values for lower bounds.
// - bu []float64
//   Values for upper bounds.
func (self *Task) GetConBoundSlice(first int32,last int32) (bk []Boundkey,bl []float64,bu []float64,err error) {
  var _tmp419 *Boundkey
  bk = make([]Boundkey,(last - first))
  if len(bk) > 0 { _tmp419 = (*Boundkey)(&bk[0]) }
  var _tmp420 *float64
  bl = make([]float64,(last - first))
  if len(bl) > 0 { _tmp420 = (*float64)(&bl[0]) }
  var _tmp421 *float64
  bu = make([]float64,(last - first))
  if len(bu) > 0 { _tmp421 = (*float64)(&bu[0]) }
  if _tmp422 := C.MSK_getconboundslice(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp419),(*C.double)(_tmp420),(*C.double)(_tmp421)); _tmp422 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp422)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a cone.
//
// - k int32
//   Index of the cone.
// - ct Conetype
//   Specifies the type of the cone.
// - conepar float64
//   For the power cone it denotes the exponent alpha. For other cone types it is unused and can be set to 0.
// - nummem int32
//   Number of member variables in the cone.
// - submem []int32
//   Variable subscripts of the members in the cone.
func (self *Task) GetCone(k int32) (ct Conetype,conepar float64,nummem int32,submem []int32,err error) {
  var _tmp423 int32
  var _tmp424 C.int32_t
  var _tmp425 C.double
  var _tmp426 C.int32_t
  if _tmp427 := C.MSK_getconeinfo(self.ptr(),(C.int32_t)(k),(&_tmp424),(&_tmp425),(&_tmp426)); _tmp427 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp427)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp428 *int32
  submem = make([]int32,_tmp426)
  if len(submem) > 0 { _tmp428 = (*int32)(&submem[0]) }
  if _tmp429 := C.MSK_getcone(self.ptr(),C.int32_t(k),(*C.int32_t)(&_tmp423),(*C.double)(&conepar),(*C.int32_t)(&nummem),(*C.int32_t)(_tmp428)); _tmp429 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp429)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  ct = Conetype(_tmp423)
  return
}

// Obtains information about a cone.
//
// - k int32
//   Index of the cone.
// - ct Conetype
//   Specifies the type of the cone.
// - conepar float64
//   For the power cone it denotes the exponent alpha. For other cone types it is unused and can be set to 0.
// - nummem int32
//   Number of member variables in the cone.
func (self *Task) GetConeInfo(k int32) (ct Conetype,conepar float64,nummem int32,err error) {
  var _tmp430 int32
  if _tmp431 := C.MSK_getconeinfo(self.ptr(),C.int32_t(k),(*C.int32_t)(&_tmp430),(*C.double)(&conepar),(*C.int32_t)(&nummem)); _tmp431 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp431)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  ct = Conetype(_tmp430)
  return
}

// Obtains the name of a cone.
//
// - i int32
//   Index of the cone.
// - name string
//   The required name.
func (self *Task) GetConeName(i int32) (name string,err error) {
  var _tmp432 C.int32_t
  if _tmp433 := C.MSK_getconenamelen(self.ptr(),(C.int32_t)(i),(&_tmp432)); _tmp433 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp433)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp432))
  _tmp434 := make([]byte,sizename)
  if _tmp435 := C.MSK_getconename(self.ptr(),C.int32_t(i),C.int32_t(sizename),(*C.uchar)(&_tmp434[0])); _tmp435 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp435)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp434,byte(0)); p < 0 {
    name = string(_tmp434)
  } else {
    name = string(_tmp434[:p])
  }
  return
}

// Checks whether the name has been assigned to any cone.
//
// - somename string
//   The name which should be checked.
// - asgn int32
//   Is non-zero if the name somename is assigned to some cone.
// - index int32
//   If the name somename is assigned to some cone, this is the index of the cone.
func (self *Task) GetConeNameIndex(somename string) (asgn int32,index int32,err error) {
  _tmp436 := C.CString(somename)
  if _tmp437 := C.MSK_getconenameindex(self.ptr(),_tmp436,(*C.int32_t)(&asgn),(*C.int32_t)(&index)); _tmp437 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp437)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the length of the name of a cone.
//
// - i int32
//   Index of the cone.
// - len int32
//   Returns the length of the indicated name.
func (self *Task) GetConeNameLen(i int32) (len int32,err error) {
  if _tmp438 := C.MSK_getconenamelen(self.ptr(),C.int32_t(i),(*C.int32_t)(&len)); _tmp438 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp438)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the name of a constraint.
//
// - i int32
//   Index of the constraint.
// - name string
//   The required name.
func (self *Task) GetConName(i int32) (name string,err error) {
  var _tmp439 C.int32_t
  if _tmp440 := C.MSK_getconnamelen(self.ptr(),(C.int32_t)(i),(&_tmp439)); _tmp440 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp440)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp439))
  _tmp441 := make([]byte,sizename)
  if _tmp442 := C.MSK_getconname(self.ptr(),C.int32_t(i),C.int32_t(sizename),(*C.uchar)(&_tmp441[0])); _tmp442 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp442)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp441,byte(0)); p < 0 {
    name = string(_tmp441)
  } else {
    name = string(_tmp441[:p])
  }
  return
}

// Checks whether the name has been assigned to any constraint.
//
// - somename string
//   The name which should be checked.
// - asgn int32
//   Is non-zero if the name somename is assigned to some constraint.
// - index int32
//   If the name somename is assigned to a constraint, then return the index of the constraint.
func (self *Task) GetConNameIndex(somename string) (asgn int32,index int32,err error) {
  _tmp443 := C.CString(somename)
  if _tmp444 := C.MSK_getconnameindex(self.ptr(),_tmp443,(*C.int32_t)(&asgn),(*C.int32_t)(&index)); _tmp444 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp444)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the length of the name of a constraint.
//
// - i int32
//   Index of the constraint.
// - len int32
//   Returns the length of the indicated name.
func (self *Task) GetConNameLen(i int32) (len int32,err error) {
  if _tmp445 := C.MSK_getconnamelen(self.ptr(),C.int32_t(i),(*C.int32_t)(&len)); _tmp445 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp445)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a sequence of coefficients from the objective.
//
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - c []float64
//   Linear terms of the requested slice of the objective as a dense vector.
func (self *Task) GetCSlice(first int32,last int32) (c []float64,err error) {
  var _tmp446 *float64
  c = make([]float64,(last - first))
  if len(c) > 0 { _tmp446 = (*float64)(&c[0]) }
  if _tmp447 := C.MSK_getcslice(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp446)); _tmp447 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp447)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the dimension of a symmetric matrix variable.
//
// - j int32
//   Index of the semidefinite variable whose dimension is requested.
// - dimbarvarj int32
//   The dimension of the j'th semidefinite variable.
func (self *Task) GetDimBarvarJ(j int32) (dimbarvarj int32,err error) {
  if _tmp448 := C.MSK_getdimbarvarj(self.ptr(),C.int32_t(j),(*C.int32_t)(&dimbarvarj)); _tmp448 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp448)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the list of affine expression indexes in a disjunctive constraint.
//
// - djcidx int64
//   Index of the disjunctive constraint.
// - afeidxlist []int64
//   List of affine expression indexes.
func (self *Task) GetDjcAfeIdxList(djcidx int64) (afeidxlist []int64,err error) {
  var _tmp449 C.int64_t
  if _tmp450 := C.MSK_getdjcnumafe(self.ptr(),(C.int64_t)(djcidx),(&_tmp449)); _tmp450 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp450)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp451 *int64
  afeidxlist = make([]int64,_tmp449)
  if len(afeidxlist) > 0 { _tmp451 = (*int64)(&afeidxlist[0]) }
  if _tmp452 := C.MSK_getdjcafeidxlist(self.ptr(),C.int64_t(djcidx),(*C.int64_t)(_tmp451)); _tmp452 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp452)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the optional constant term vector of a disjunctive constraint.
//
// - djcidx int64
//   Index of the disjunctive constraint.
// - b []float64
//   The vector b.
func (self *Task) GetDjcB(djcidx int64) (b []float64,err error) {
  var _tmp453 C.int64_t
  if _tmp454 := C.MSK_getdjcnumafe(self.ptr(),(C.int64_t)(djcidx),(&_tmp453)); _tmp454 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp454)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp455 *float64
  b = make([]float64,_tmp453)
  if len(b) > 0 { _tmp455 = (*float64)(&b[0]) }
  if _tmp456 := C.MSK_getdjcb(self.ptr(),C.int64_t(djcidx),(*C.double)(_tmp455)); _tmp456 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp456)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the list of domain indexes in a disjunctive constraint.
//
// - djcidx int64
//   Index of the disjunctive constraint.
// - domidxlist []int64
//   List of term sizes.
func (self *Task) GetDjcDomainIdxList(djcidx int64) (domidxlist []int64,err error) {
  var _tmp457 C.int64_t
  if _tmp458 := C.MSK_getdjcnumdomain(self.ptr(),(C.int64_t)(djcidx),(&_tmp457)); _tmp458 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp458)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp459 *int64
  domidxlist = make([]int64,_tmp457)
  if len(domidxlist) > 0 { _tmp459 = (*int64)(&domidxlist[0]) }
  if _tmp460 := C.MSK_getdjcdomainidxlist(self.ptr(),C.int64_t(djcidx),(*C.int64_t)(_tmp459)); _tmp460 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp460)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the name of a disjunctive constraint.
//
// - djcidx int64
//   Index of a disjunctive constraint.
// - name string
//   Returns the required name.
func (self *Task) GetDjcName(djcidx int64) (name string,err error) {
  var _tmp461 C.int32_t
  if _tmp462 := C.MSK_getdjcnamelen(self.ptr(),(C.int64_t)(djcidx),(&_tmp461)); _tmp462 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp462)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp461))
  _tmp463 := make([]byte,sizename)
  if _tmp464 := C.MSK_getdjcname(self.ptr(),C.int64_t(djcidx),C.int32_t(sizename),(*C.uchar)(&_tmp463[0])); _tmp464 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp464)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp463,byte(0)); p < 0 {
    name = string(_tmp463)
  } else {
    name = string(_tmp463[:p])
  }
  return
}

// Obtains the length of the name of a disjunctive constraint.
//
// - djcidx int64
//   Index of a disjunctive constraint.
// - len int32
//   Returns the length of the indicated name.
func (self *Task) GetDjcNameLen(djcidx int64) (len int32,err error) {
  if _tmp465 := C.MSK_getdjcnamelen(self.ptr(),C.int64_t(djcidx),(*C.int32_t)(&len)); _tmp465 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp465)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of affine expressions in the disjunctive constraint.
//
// - djcidx int64
//   Index of the disjunctive constraint.
// - numafe int64
//   Number of affine expressions in the disjunctive constraint.
func (self *Task) GetDjcNumAfe(djcidx int64) (numafe int64,err error) {
  if _tmp466 := C.MSK_getdjcnumafe(self.ptr(),C.int64_t(djcidx),(*C.int64_t)(&numafe)); _tmp466 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp466)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of affine expressions in all disjunctive constraints.
//
// - numafetot int64
//   Number of affine expressions in all disjunctive constraints.
func (self *Task) GetDjcNumAfeTot() (numafetot int64,err error) {
  if _tmp467 := C.MSK_getdjcnumafetot(self.ptr(),(*C.int64_t)(&numafetot)); _tmp467 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp467)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of domains in the disjunctive constraint.
//
// - djcidx int64
//   Index of the disjunctive constraint.
// - numdomain int64
//   Number of domains in the disjunctive constraint.
func (self *Task) GetDjcNumDomain(djcidx int64) (numdomain int64,err error) {
  if _tmp468 := C.MSK_getdjcnumdomain(self.ptr(),C.int64_t(djcidx),(*C.int64_t)(&numdomain)); _tmp468 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp468)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of domains in all disjunctive constraints.
//
// - numdomaintot int64
//   Number of domains in all disjunctive constraints.
func (self *Task) GetDjcNumDomainTot() (numdomaintot int64,err error) {
  if _tmp469 := C.MSK_getdjcnumdomaintot(self.ptr(),(*C.int64_t)(&numdomaintot)); _tmp469 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp469)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number terms in the disjunctive constraint.
//
// - djcidx int64
//   Index of the disjunctive constraint.
// - numterm int64
//   Number of terms in the disjunctive constraint.
func (self *Task) GetDjcNumTerm(djcidx int64) (numterm int64,err error) {
  if _tmp470 := C.MSK_getdjcnumterm(self.ptr(),C.int64_t(djcidx),(*C.int64_t)(&numterm)); _tmp470 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp470)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of terms in all disjunctive constraints.
//
// - numtermtot int64
//   Total number of terms in all disjunctive constraints.
func (self *Task) GetDjcNumTermTot() (numtermtot int64,err error) {
  if _tmp471 := C.MSK_getdjcnumtermtot(self.ptr(),(*C.int64_t)(&numtermtot)); _tmp471 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp471)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains full data of all disjunctive constraints.
//
// - domidxlist []int64
//   The concatenation of index lists of domains appearing in all disjunctive constraints.
// - afeidxlist []int64
//   The concatenation of index lists of affine expressions appearing in all disjunctive constraints.
// - b []float64
//   The concatenation of vectors b appearing in all disjunctive constraints.
// - termsizelist []int64
//   The concatenation of lists of term sizes appearing in all disjunctive constraints.
// - numterms []int64
//   The number of terms in each of the disjunctive constraints.
func (self *Task) GetDjcs() (domidxlist []int64,afeidxlist []int64,b []float64,termsizelist []int64,numterms []int64,err error) {
  var _tmp472 C.int64_t
  if _tmp473 := C.MSK_getdjcnumdomaintot(self.ptr(),(&_tmp472)); _tmp473 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp473)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp474 *int64
  domidxlist = make([]int64,_tmp472)
  if len(domidxlist) > 0 { _tmp474 = (*int64)(&domidxlist[0]) }
  var _tmp475 C.int64_t
  if _tmp476 := C.MSK_getdjcnumafetot(self.ptr(),(&_tmp475)); _tmp476 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp476)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp477 *int64
  afeidxlist = make([]int64,_tmp475)
  if len(afeidxlist) > 0 { _tmp477 = (*int64)(&afeidxlist[0]) }
  var _tmp478 C.int64_t
  if _tmp479 := C.MSK_getdjcnumafetot(self.ptr(),(&_tmp478)); _tmp479 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp479)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp480 *float64
  b = make([]float64,_tmp478)
  if len(b) > 0 { _tmp480 = (*float64)(&b[0]) }
  var _tmp481 C.int64_t
  if _tmp482 := C.MSK_getdjcnumtermtot(self.ptr(),(&_tmp481)); _tmp482 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp482)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp483 *int64
  termsizelist = make([]int64,_tmp481)
  if len(termsizelist) > 0 { _tmp483 = (*int64)(&termsizelist[0]) }
  var _tmp484 C.int64_t
  if _tmp485 := C.MSK_getnumdjc(self.ptr(),(&_tmp484)); _tmp485 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp485)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp486 *int64
  numterms = make([]int64,_tmp484)
  if len(numterms) > 0 { _tmp486 = (*int64)(&numterms[0]) }
  if _tmp487 := C.MSK_getdjcs(self.ptr(),(*C.int64_t)(_tmp474),(*C.int64_t)(_tmp477),(*C.double)(_tmp480),(*C.int64_t)(_tmp483),(*C.int64_t)(_tmp486)); _tmp487 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp487)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the list of term sizes in a disjunctive constraint.
//
// - djcidx int64
//   Index of the disjunctive constraint.
// - termsizelist []int64
//   List of term sizes.
func (self *Task) GetDjcTermSizeList(djcidx int64) (termsizelist []int64,err error) {
  var _tmp488 C.int64_t
  if _tmp489 := C.MSK_getdjcnumterm(self.ptr(),(C.int64_t)(djcidx),(&_tmp488)); _tmp489 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp489)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp490 *int64
  termsizelist = make([]int64,_tmp488)
  if len(termsizelist) > 0 { _tmp490 = (*int64)(&termsizelist[0]) }
  if _tmp491 := C.MSK_getdjctermsizelist(self.ptr(),C.int64_t(djcidx),(*C.int64_t)(_tmp490)); _tmp491 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp491)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the dimension of the domain.
//
// - domidx int64
//   Index of the domain.
// - n int64
//   Dimension of the domain.
func (self *Task) GetDomainN(domidx int64) (n int64,err error) {
  if _tmp492 := C.MSK_getdomainn(self.ptr(),C.int64_t(domidx),(*C.int64_t)(&n)); _tmp492 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp492)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the name of a domain.
//
// - domidx int64
//   Index of a domain.
// - name string
//   Returns the required name.
func (self *Task) GetDomainName(domidx int64) (name string,err error) {
  var _tmp493 C.int32_t
  if _tmp494 := C.MSK_getdomainnamelen(self.ptr(),(C.int64_t)(domidx),(&_tmp493)); _tmp494 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp494)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp493))
  _tmp495 := make([]byte,sizename)
  if _tmp496 := C.MSK_getdomainname(self.ptr(),C.int64_t(domidx),C.int32_t(sizename),(*C.uchar)(&_tmp495[0])); _tmp496 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp496)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp495,byte(0)); p < 0 {
    name = string(_tmp495)
  } else {
    name = string(_tmp495[:p])
  }
  return
}

// Obtains the length of the name of a domain.
//
// - domidx int64
//   Index of a domain.
// - len int32
//   Returns the length of the indicated name.
func (self *Task) GetDomainNameLen(domidx int64) (len int32,err error) {
  if _tmp497 := C.MSK_getdomainnamelen(self.ptr(),C.int64_t(domidx),(*C.int32_t)(&len)); _tmp497 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp497)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Returns the type of the domain.
//
// - domidx int64
//   Index of the domain.
// - domtype Domaintype
//   The type of the domain.
func (self *Task) GetDomainType(domidx int64) (domtype Domaintype,err error) {
  var _tmp498 int32
  if _tmp499 := C.MSK_getdomaintype(self.ptr(),C.int64_t(domidx),(*C.int32_t)(&_tmp498)); _tmp499 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp499)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  domtype = Domaintype(_tmp498)
  return
}

// Obtains a double information item.
//
// - whichdinf Dinfitem
//   Specifies a double information item.
// - dvalue float64
//   The value of the required double information item.
func (self *Task) GetDouInf(whichdinf Dinfitem) (dvalue float64,err error) {
  if _tmp500 := C.MSK_getdouinf(self.ptr(),C.int32_t(whichdinf),(*C.double)(&dvalue)); _tmp500 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp500)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a double parameter.
//
// - param Dparam
//   Which parameter.
// - parvalue float64
//   Parameter value.
func (self *Task) GetDouParam(param Dparam) (parvalue float64,err error) {
  if _tmp501 := C.MSK_getdouparam(self.ptr(),C.int32_t(param),(*C.double)(&parvalue)); _tmp501 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp501)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the dual objective value associated with the solution.
//
// - whichsol Soltype
//   Selects a solution.
// - dualobj float64
//   Objective value corresponding to the dual solution.
func (self *Task) GetDualObj(whichsol Soltype) (dualobj float64,err error) {
  if _tmp502 := C.MSK_getdualobj(self.ptr(),C.int32_t(whichsol),(*C.double)(&dualobj)); _tmp502 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp502)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Compute norms of the dual solution.
//
// - whichsol Soltype
//   Selects a solution.
// - nrmy float64
//   The norm of the y vector.
// - nrmslc float64
//   The norm of the slc vector.
// - nrmsuc float64
//   The norm of the suc vector.
// - nrmslx float64
//   The norm of the slx vector.
// - nrmsux float64
//   The norm of the sux vector.
// - nrmsnx float64
//   The norm of the snx vector.
// - nrmbars float64
//   The norm of the bars vector.
func (self *Task) GetDualSolutionNorms(whichsol Soltype) (nrmy float64,nrmslc float64,nrmsuc float64,nrmslx float64,nrmsux float64,nrmsnx float64,nrmbars float64,err error) {
  if _tmp503 := C.MSK_getdualsolutionnorms(self.ptr(),C.int32_t(whichsol),(*C.double)(&nrmy),(*C.double)(&nrmslc),(*C.double)(&nrmsuc),(*C.double)(&nrmslx),(*C.double)(&nrmsux),(*C.double)(&nrmsnx),(*C.double)(&nrmbars)); _tmp503 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp503)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the violation of the dual solution for set of affine conic constraints.
//
// - whichsol Soltype
//   Selects a solution.
// - accidxlist []int64
//   An array of indexes of conic constraints.
// - viol []float64
//   List of violations corresponding to sub.
func (self *Task) GetDviolAcc(whichsol Soltype,accidxlist []int64) (viol []float64,err error) {
  _tmp504 := len(accidxlist)
  var numaccidx int64 = int64(_tmp504)
  var _tmp505 *int64
  if len(accidxlist) > 0 { _tmp505 = (*int64)(&accidxlist[0]) }
  var _tmp506 *float64
  viol = make([]float64,numaccidx)
  if len(viol) > 0 { _tmp506 = (*float64)(&viol[0]) }
  if _tmp507 := C.MSK_getdviolacc(self.ptr(),C.int32_t(whichsol),C.int64_t(numaccidx),(*C.int64_t)(_tmp505),(*C.double)(_tmp506)); _tmp507 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp507)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the violation of dual solution for a set of semidefinite variables.
//
// - whichsol Soltype
//   Selects a solution.
// - sub []int32
//   An array of indexes of barx variables.
// - viol []float64
//   List of violations corresponding to sub.
func (self *Task) GetDviolBarvar(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp508 := len(sub)
  var num int32 = int32(_tmp508)
  var _tmp509 *int32
  if len(sub) > 0 { _tmp509 = (*int32)(&sub[0]) }
  var _tmp510 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp510 = (*float64)(&viol[0]) }
  if _tmp511 := C.MSK_getdviolbarvar(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp509),(*C.double)(_tmp510)); _tmp511 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp511)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the violation of a dual solution associated with a set of constraints.
//
// - whichsol Soltype
//   Selects a solution.
// - sub []int32
//   An array of indexes of constraints.
// - viol []float64
//   List of violations corresponding to sub.
func (self *Task) GetDviolCon(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp512 := len(sub)
  var num int32 = int32(_tmp512)
  var _tmp513 *int32
  if len(sub) > 0 { _tmp513 = (*int32)(&sub[0]) }
  var _tmp514 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp514 = (*float64)(&viol[0]) }
  if _tmp515 := C.MSK_getdviolcon(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp513),(*C.double)(_tmp514)); _tmp515 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp515)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the violation of a solution for set of dual conic constraints.
//
// - whichsol Soltype
//   Selects a solution.
// - sub []int32
//   An array of indexes of conic constraints.
// - viol []float64
//   List of violations corresponding to sub.
func (self *Task) GetDviolCones(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp516 := len(sub)
  var num int32 = int32(_tmp516)
  var _tmp517 *int32
  if len(sub) > 0 { _tmp517 = (*int32)(&sub[0]) }
  var _tmp518 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp518 = (*float64)(&viol[0]) }
  if _tmp519 := C.MSK_getdviolcones(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp517),(*C.double)(_tmp518)); _tmp519 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp519)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the violation of a dual solution associated with a set of scalar variables.
//
// - whichsol Soltype
//   Selects a solution.
// - sub []int32
//   An array of indexes of x variables.
// - viol []float64
//   List of violations corresponding to sub.
func (self *Task) GetDviolVar(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp520 := len(sub)
  var num int32 = int32(_tmp520)
  var _tmp521 *int32
  if len(sub) > 0 { _tmp521 = (*int32)(&sub[0]) }
  var _tmp522 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp522 = (*float64)(&viol[0]) }
  if _tmp523 := C.MSK_getdviolvar(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp521),(*C.double)(_tmp522)); _tmp523 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp523)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the index of a named information item.
//
// - inftype Inftype
//   Type of the information item.
// - infname string
//   Name of the information item.
// - infindex int32
//   The item index.
func (self *Task) GetInfIndex(inftype Inftype,infname string) (infindex int32,err error) {
  _tmp524 := C.CString(infname)
  if _tmp525 := C.MSK_getinfindex(self.ptr(),C.int32_t(inftype),_tmp524,(*C.int32_t)(&infindex)); _tmp525 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp525)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the maximum index of an information item of a given type.
//
// - inftype Inftype
//   Type of the information item.
// - infmax []int32
//   The maximum index (plus 1) requested.
func (self *Task) GetInfMax(inftype Inftype) (infmax []int32,err error) {
  var _tmp526 *int32
  infmax = make([]int32,MSK_MAX_STR_LEN)
  if len(infmax) > 0 { _tmp526 = (*int32)(&infmax[0]) }
  if _tmp527 := C.MSK_getinfmax(self.ptr(),C.int32_t(inftype),(*C.int32_t)(_tmp526)); _tmp527 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp527)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the name of an information item.
//
// - inftype Inftype
//   Type of the information item.
// - whichinf int32
//   An information item.
// - infname string
//   Name of the information item.
func (self *Task) GetInfName(inftype Inftype,whichinf int32) (infname string,err error) {
  _tmp528 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp529 := C.MSK_getinfname(self.ptr(),C.int32_t(inftype),C.int32_t(whichinf),(*C.uchar)(&_tmp528[0])); _tmp529 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp529)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp528,byte(0)); p < 0 {
    infname = string(_tmp528)
  } else {
    infname = string(_tmp528[:p])
  }
  return
}

// Obtains an integer information item.
//
// - whichiinf Iinfitem
//   Specifies an integer information item.
// - ivalue int32
//   The value of the required integer information item.
func (self *Task) GetIntInf(whichiinf Iinfitem) (ivalue int32,err error) {
  if _tmp530 := C.MSK_getintinf(self.ptr(),C.int32_t(whichiinf),(*C.int32_t)(&ivalue)); _tmp530 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp530)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains an integer parameter.
//
// - param Iparam
//   Which parameter.
// - parvalue int32
//   Parameter value.
func (self *Task) GetIntParam(param Iparam) (parvalue int32,err error) {
  if _tmp531 := C.MSK_getintparam(self.ptr(),C.int32_t(param),(*C.int32_t)(&parvalue)); _tmp531 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp531)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the length of one semidefinite variable.
//
// - j int32
//   Index of the semidefinite variable whose length if requested.
// - lenbarvarj int64
//   Number of scalar elements in the lower triangular part of the semidefinite variable.
func (self *Task) GetLenBarvarJ(j int32) (lenbarvarj int64,err error) {
  if _tmp532 := C.MSK_getlenbarvarj(self.ptr(),C.int32_t(j),(*C.int64_t)(&lenbarvarj)); _tmp532 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp532)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a long integer information item.
//
// - whichliinf Liinfitem
//   Specifies a long information item.
// - ivalue int64
//   The value of the required long integer information item.
func (self *Task) GetLintInf(whichliinf Liinfitem) (ivalue int64,err error) {
  if _tmp533 := C.MSK_getlintinf(self.ptr(),C.int32_t(whichliinf),(*C.int64_t)(&ivalue)); _tmp533 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp533)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains number of preallocated non-zeros in the linear constraint matrix.
//
// - maxnumanz int64
//   Number of preallocated non-zero linear matrix elements.
func (self *Task) GetMaxNumANz() (maxnumanz int64,err error) {
  if _tmp534 := C.MSK_getmaxnumanz64(self.ptr(),(*C.int64_t)(&maxnumanz)); _tmp534 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp534)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains maximum number of symmetric matrix variables for which space is currently preallocated.
//
// - maxnumbarvar int32
//   Maximum number of symmetric matrix variables for which space is currently preallocated.
func (self *Task) GetMaxNumBarvar() (maxnumbarvar int32,err error) {
  if _tmp535 := C.MSK_getmaxnumbarvar(self.ptr(),(*C.int32_t)(&maxnumbarvar)); _tmp535 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp535)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of preallocated constraints in the optimization task.
//
// - maxnumcon int32
//   Number of preallocated constraints in the optimization task.
func (self *Task) GetMaxNumCon() (maxnumcon int32,err error) {
  if _tmp536 := C.MSK_getmaxnumcon(self.ptr(),(*C.int32_t)(&maxnumcon)); _tmp536 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp536)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of preallocated cones in the optimization task.
//
// - maxnumcone int32
//   Number of preallocated conic constraints in the optimization task.
func (self *Task) GetMaxNumCone() (maxnumcone int32,err error) {
  if _tmp537 := C.MSK_getmaxnumcone(self.ptr(),(*C.int32_t)(&maxnumcone)); _tmp537 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp537)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of preallocated non-zeros for all quadratic terms in objective and constraints.
//
// - maxnumqnz int64
//   Number of non-zero elements preallocated in quadratic coefficient matrices.
func (self *Task) GetMaxNumQNz() (maxnumqnz int64,err error) {
  if _tmp538 := C.MSK_getmaxnumqnz64(self.ptr(),(*C.int64_t)(&maxnumqnz)); _tmp538 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp538)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the maximum number variables allowed.
//
// - maxnumvar int32
//   Number of preallocated variables in the optimization task.
func (self *Task) GetMaxNumVar() (maxnumvar int32,err error) {
  if _tmp539 := C.MSK_getmaxnumvar(self.ptr(),(*C.int32_t)(&maxnumvar)); _tmp539 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp539)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains information about the amount of memory used by a task.
//
// - meminuse int64
//   Amount of memory currently used by the task.
// - maxmemuse int64
//   Maximum amount of memory used by the task until now.
func (self *Task) GetMemUsage() (meminuse int64,maxmemuse int64,err error) {
  if _tmp540 := C.MSK_getmemusagetask(self.ptr(),(*C.int64_t)(&meminuse),(*C.int64_t)(&maxmemuse)); _tmp540 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp540)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of affine conic constraints.
//
// - num int64
//   The number of affine conic constraints.
func (self *Task) GetNumAcc() (num int64,err error) {
  if _tmp541 := C.MSK_getnumacc(self.ptr(),(*C.int64_t)(&num)); _tmp541 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp541)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of affine expressions.
//
// - numafe int64
//   Number of affine expressions.
func (self *Task) GetNumAfe() (numafe int64,err error) {
  if _tmp542 := C.MSK_getnumafe(self.ptr(),(*C.int64_t)(&numafe)); _tmp542 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp542)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of non-zeros in the coefficient matrix.
//
// - numanz int32
//   Number of non-zero elements in the linear constraint matrix.
func (self *Task) GetNumANz() (numanz int32,err error) {
  if _tmp543 := C.MSK_getnumanz(self.ptr(),(*C.int32_t)(&numanz)); _tmp543 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp543)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of non-zeros in the coefficient matrix.
//
// - numanz int64
//   Number of non-zero elements in the linear constraint matrix.
func (self *Task) GetNumANz64() (numanz int64,err error) {
  if _tmp544 := C.MSK_getnumanz64(self.ptr(),(*C.int64_t)(&numanz)); _tmp544 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp544)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains an upper bound on the number of scalar elements in the block triplet form of bara.
//
// - num int64
//   An upper bound on the number of elements in the block triplet form of bara.
func (self *Task) GetNumBaraBlockTriplets() (num int64,err error) {
  if _tmp545 := C.MSK_getnumbarablocktriplets(self.ptr(),(*C.int64_t)(&num)); _tmp545 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp545)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Get the number of nonzero elements in barA.
//
// - nz int64
//   The number of nonzero block elements in barA.
func (self *Task) GetNumBaraNz() (nz int64,err error) {
  if _tmp546 := C.MSK_getnumbaranz(self.ptr(),(*C.int64_t)(&nz)); _tmp546 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp546)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains an upper bound on the number of elements in the block triplet form of barc.
//
// - num int64
//   An upper bound on the number of elements in the block triplet form of barc.
func (self *Task) GetNumBarcBlockTriplets() (num int64,err error) {
  if _tmp547 := C.MSK_getnumbarcblocktriplets(self.ptr(),(*C.int64_t)(&num)); _tmp547 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp547)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of nonzero elements in barc.
//
// - nz int64
//   The number of nonzero elements in barc.
func (self *Task) GetNumBarcNz() (nz int64,err error) {
  if _tmp548 := C.MSK_getnumbarcnz(self.ptr(),(*C.int64_t)(&nz)); _tmp548 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp548)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of semidefinite variables.
//
// - numbarvar int32
//   Number of semidefinite variables in the problem.
func (self *Task) GetNumBarvar() (numbarvar int32,err error) {
  if _tmp549 := C.MSK_getnumbarvar(self.ptr(),(*C.int32_t)(&numbarvar)); _tmp549 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp549)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of constraints.
//
// - numcon int32
//   Number of constraints.
func (self *Task) GetNumCon() (numcon int32,err error) {
  if _tmp550 := C.MSK_getnumcon(self.ptr(),(*C.int32_t)(&numcon)); _tmp550 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp550)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of cones.
//
// - numcone int32
//   Number of conic constraints.
func (self *Task) GetNumCone() (numcone int32,err error) {
  if _tmp551 := C.MSK_getnumcone(self.ptr(),(*C.int32_t)(&numcone)); _tmp551 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp551)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of members in a cone.
//
// - k int32
//   Index of the cone.
// - nummem int32
//   Number of member variables in the cone.
func (self *Task) GetNumConeMem(k int32) (nummem int32,err error) {
  if _tmp552 := C.MSK_getnumconemem(self.ptr(),C.int32_t(k),(*C.int32_t)(&nummem)); _tmp552 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp552)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of disjunctive constraints.
//
// - num int64
//   The number of disjunctive constraints.
func (self *Task) GetNumDjc() (num int64,err error) {
  if _tmp553 := C.MSK_getnumdjc(self.ptr(),(*C.int64_t)(&num)); _tmp553 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp553)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtain the number of domains defined.
//
// - numdomain int64
//   Number of domains in the task.
func (self *Task) GetNumDomain() (numdomain int64,err error) {
  if _tmp554 := C.MSK_getnumdomain(self.ptr(),(*C.int64_t)(&numdomain)); _tmp554 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp554)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of integer-constrained variables.
//
// - numintvar int32
//   Number of integer variables.
func (self *Task) GetNumIntVar() (numintvar int32,err error) {
  if _tmp555 := C.MSK_getnumintvar(self.ptr(),(*C.int32_t)(&numintvar)); _tmp555 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp555)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of parameters of a given type.
//
// - partype Parametertype
//   Parameter type.
// - numparam int32
//   Returns the number of parameters of the requested type.
func (self *Task) GetNumParam(partype Parametertype) (numparam int32,err error) {
  if _tmp556 := C.MSK_getnumparam(self.ptr(),C.int32_t(partype),(*C.int32_t)(&numparam)); _tmp556 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp556)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of non-zero quadratic terms in a constraint.
//
// - k int32
//   Index of the constraint for which the number quadratic terms should be obtained.
// - numqcnz int64
//   Number of quadratic terms.
func (self *Task) GetNumQConKNz(k int32) (numqcnz int64,err error) {
  if _tmp557 := C.MSK_getnumqconknz64(self.ptr(),C.int32_t(k),(*C.int64_t)(&numqcnz)); _tmp557 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp557)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of non-zero quadratic terms in the objective.
//
// - numqonz int64
//   Number of non-zero elements in the quadratic objective terms.
func (self *Task) GetNumQObjNz() (numqonz int64,err error) {
  if _tmp558 := C.MSK_getnumqobjnz64(self.ptr(),(*C.int64_t)(&numqonz)); _tmp558 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp558)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of symmetric matrices stored.
//
// - num int64
//   The number of symmetric sparse matrices.
func (self *Task) GetNumSymMat() (num int64,err error) {
  if _tmp559 := C.MSK_getnumsymmat(self.ptr(),(*C.int64_t)(&num)); _tmp559 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp559)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the number of variables.
//
// - numvar int32
//   Number of variables.
func (self *Task) GetNumVar() (numvar int32,err error) {
  if _tmp560 := C.MSK_getnumvar(self.ptr(),(*C.int32_t)(&numvar)); _tmp560 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp560)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the name assigned to the objective function.
//
// - objname string
//   Assigned the objective name.
func (self *Task) GetObjName() (objname string,err error) {
  var _tmp561 C.int32_t
  if _tmp562 := C.MSK_getobjnamelen(self.ptr(),(&_tmp561)); _tmp562 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp562)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizeobjname int32 = int32((1 + _tmp561))
  _tmp563 := make([]byte,sizeobjname)
  if _tmp564 := C.MSK_getobjname(self.ptr(),C.int32_t(sizeobjname),(*C.uchar)(&_tmp563[0])); _tmp564 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp564)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp563,byte(0)); p < 0 {
    objname = string(_tmp563)
  } else {
    objname = string(_tmp563[:p])
  }
  return
}

// Obtains the length of the name assigned to the objective function.
//
// - len int32
//   Assigned the length of the objective name.
func (self *Task) GetObjNameLen() (len int32,err error) {
  if _tmp565 := C.MSK_getobjnamelen(self.ptr(),(*C.int32_t)(&len)); _tmp565 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp565)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Gets the objective sense.
//
// - sense Objsense
//   The returned objective sense.
func (self *Task) GetObjSense() (sense Objsense,err error) {
  var _tmp566 int32
  if _tmp567 := C.MSK_getobjsense(self.ptr(),(*C.int32_t)(&_tmp566)); _tmp567 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp567)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  sense = Objsense(_tmp566)
  return
}

// Obtains the maximum index of a parameter of a given type.
//
// - partype Parametertype
//   Parameter type.
// - parammax int32
//   The maximum index (plus 1) of the given parameter type.
func (self *Task) GetParamMax(partype Parametertype) (parammax int32,err error) {
  if _tmp568 := C.MSK_getparammax(self.ptr(),C.int32_t(partype),(*C.int32_t)(&parammax)); _tmp568 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp568)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the name of a parameter.
//
// - partype Parametertype
//   Parameter type.
// - param int32
//   Which parameter.
// - parname string
//   Parameter name.
func (self *Task) GetParamName(partype Parametertype,param int32) (parname string,err error) {
  _tmp569 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp570 := C.MSK_getparamname(self.ptr(),C.int32_t(partype),C.int32_t(param),(*C.uchar)(&_tmp569[0])); _tmp570 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp570)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp569,byte(0)); p < 0 {
    parname = string(_tmp569)
  } else {
    parname = string(_tmp569[:p])
  }
  return
}

// Obtains the exponent vector of a power domain.
//
// - domidx int64
//   Index of the domain.
// - alpha []float64
//   The exponent vector of the domain.
func (self *Task) GetPowerDomainAlpha(domidx int64) (alpha []float64,err error) {
  var _tmp571 C.int64_t
  var _tmp572 C.int64_t
  if _tmp573 := C.MSK_getpowerdomaininfo(self.ptr(),(C.int64_t)(domidx),(&_tmp571),(&_tmp572)); _tmp573 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp573)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp574 *float64
  alpha = make([]float64,_tmp572)
  if len(alpha) > 0 { _tmp574 = (*float64)(&alpha[0]) }
  if _tmp575 := C.MSK_getpowerdomainalpha(self.ptr(),C.int64_t(domidx),(*C.double)(_tmp574)); _tmp575 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp575)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains structural information about a power domain.
//
// - domidx int64
//   Index of the domain.
// - n int64
//   Dimension of the domain.
// - nleft int64
//   Number of variables on the left hand side.
func (self *Task) GetPowerDomainInfo(domidx int64) (n int64,nleft int64,err error) {
  if _tmp576 := C.MSK_getpowerdomaininfo(self.ptr(),C.int64_t(domidx),(*C.int64_t)(&n),(*C.int64_t)(&nleft)); _tmp576 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp576)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the primal objective value for the desired solution.
//
// - whichsol Soltype
//   Selects a solution.
// - primalobj float64
//   Objective value corresponding to the primal solution.
func (self *Task) GetPrimalObj(whichsol Soltype) (primalobj float64,err error) {
  if _tmp577 := C.MSK_getprimalobj(self.ptr(),C.int32_t(whichsol),(*C.double)(&primalobj)); _tmp577 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp577)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Compute norms of the primal solution.
//
// - whichsol Soltype
//   Selects a solution.
// - nrmxc float64
//   The norm of the xc vector.
// - nrmxx float64
//   The norm of the xx vector.
// - nrmbarx float64
//   The norm of the barX vector.
func (self *Task) GetPrimalSolutionNorms(whichsol Soltype) (nrmxc float64,nrmxx float64,nrmbarx float64,err error) {
  if _tmp578 := C.MSK_getprimalsolutionnorms(self.ptr(),C.int32_t(whichsol),(*C.double)(&nrmxc),(*C.double)(&nrmxx),(*C.double)(&nrmbarx)); _tmp578 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp578)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the problem type.
//
// - probtype Problemtype
//   The problem type.
func (self *Task) GetProbType() (probtype Problemtype,err error) {
  var _tmp579 int32
  if _tmp580 := C.MSK_getprobtype(self.ptr(),(*C.int32_t)(&_tmp579)); _tmp580 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp580)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  probtype = Problemtype(_tmp579)
  return
}

// Obtains the problem status.
//
// - whichsol Soltype
//   Selects a solution.
// - problemsta Prosta
//   Problem status.
func (self *Task) GetProSta(whichsol Soltype) (problemsta Prosta,err error) {
  var _tmp581 int32
  if _tmp582 := C.MSK_getprosta(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(&_tmp581)); _tmp582 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp582)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  problemsta = Prosta(_tmp581)
  return
}

// Computes the violation of a solution for set of affine conic constraints.
//
// - whichsol Soltype
//   Selects a solution.
// - accidxlist []int64
//   An array of indexes of conic constraints.
// - viol []float64
//   List of violations corresponding to sub.
func (self *Task) GetPviolAcc(whichsol Soltype,accidxlist []int64) (viol []float64,err error) {
  _tmp583 := len(accidxlist)
  var numaccidx int64 = int64(_tmp583)
  var _tmp584 *int64
  if len(accidxlist) > 0 { _tmp584 = (*int64)(&accidxlist[0]) }
  var _tmp585 *float64
  viol = make([]float64,numaccidx)
  if len(viol) > 0 { _tmp585 = (*float64)(&viol[0]) }
  if _tmp586 := C.MSK_getpviolacc(self.ptr(),C.int32_t(whichsol),C.int64_t(numaccidx),(*C.int64_t)(_tmp584),(*C.double)(_tmp585)); _tmp586 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp586)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the violation of a primal solution for a list of semidefinite variables.
//
// - whichsol Soltype
//   Selects a solution.
// - sub []int32
//   An array of indexes of barX variables.
// - viol []float64
//   List of violations corresponding to sub.
func (self *Task) GetPviolBarvar(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp587 := len(sub)
  var num int32 = int32(_tmp587)
  var _tmp588 *int32
  if len(sub) > 0 { _tmp588 = (*int32)(&sub[0]) }
  var _tmp589 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp589 = (*float64)(&viol[0]) }
  if _tmp590 := C.MSK_getpviolbarvar(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp588),(*C.double)(_tmp589)); _tmp590 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp590)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the violation of a primal solution associated to a constraint.
//
// - whichsol Soltype
//   Selects a solution.
// - sub []int32
//   An array of indexes of constraints.
// - viol []float64
//   List of violations corresponding to sub.
func (self *Task) GetPviolCon(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp591 := len(sub)
  var num int32 = int32(_tmp591)
  var _tmp592 *int32
  if len(sub) > 0 { _tmp592 = (*int32)(&sub[0]) }
  var _tmp593 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp593 = (*float64)(&viol[0]) }
  if _tmp594 := C.MSK_getpviolcon(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp592),(*C.double)(_tmp593)); _tmp594 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp594)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the violation of a solution for set of conic constraints.
//
// - whichsol Soltype
//   Selects a solution.
// - sub []int32
//   An array of indexes of conic constraints.
// - viol []float64
//   List of violations corresponding to sub.
func (self *Task) GetPviolCones(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp595 := len(sub)
  var num int32 = int32(_tmp595)
  var _tmp596 *int32
  if len(sub) > 0 { _tmp596 = (*int32)(&sub[0]) }
  var _tmp597 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp597 = (*float64)(&viol[0]) }
  if _tmp598 := C.MSK_getpviolcones(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp596),(*C.double)(_tmp597)); _tmp598 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp598)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the violation of a solution for set of disjunctive constraints.
//
// - whichsol Soltype
//   Selects a solution.
// - djcidxlist []int64
//   An array of indexes of disjunctive constraints.
// - viol []float64
//   List of violations corresponding to sub.
func (self *Task) GetPviolDjc(whichsol Soltype,djcidxlist []int64) (viol []float64,err error) {
  _tmp599 := len(djcidxlist)
  var numdjcidx int64 = int64(_tmp599)
  var _tmp600 *int64
  if len(djcidxlist) > 0 { _tmp600 = (*int64)(&djcidxlist[0]) }
  var _tmp601 *float64
  viol = make([]float64,numdjcidx)
  if len(viol) > 0 { _tmp601 = (*float64)(&viol[0]) }
  if _tmp602 := C.MSK_getpvioldjc(self.ptr(),C.int32_t(whichsol),C.int64_t(numdjcidx),(*C.int64_t)(_tmp600),(*C.double)(_tmp601)); _tmp602 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp602)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the violation of a primal solution for a list of scalar variables.
//
// - whichsol Soltype
//   Selects a solution.
// - sub []int32
//   An array of indexes of x variables.
// - viol []float64
//   List of violations corresponding to sub.
func (self *Task) GetPviolVar(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp603 := len(sub)
  var num int32 = int32(_tmp603)
  var _tmp604 *int32
  if len(sub) > 0 { _tmp604 = (*int32)(&sub[0]) }
  var _tmp605 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp605 = (*float64)(&viol[0]) }
  if _tmp606 := C.MSK_getpviolvar(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp604),(*C.double)(_tmp605)); _tmp606 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp606)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains all the quadratic terms in a constraint.
//
// - k int32
//   Which constraint.
// - numqcnz int64
//   Number of quadratic terms.
// - qcsubi []int32
//   Row subscripts for quadratic constraint matrix.
// - qcsubj []int32
//   Column subscripts for quadratic constraint matrix.
// - qcval []float64
//   Quadratic constraint coefficient values.
func (self *Task) GetQConK(k int32) (numqcnz int64,qcsubi []int32,qcsubj []int32,qcval []float64,err error) {
  var _tmp607 C.int64_t
  if _tmp608 := C.MSK_getnumqconknz64(self.ptr(),(C.int32_t)(k),(&_tmp607)); _tmp608 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp608)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumqcnz int64 = int64(_tmp607)
  var _tmp609 C.int64_t
  if _tmp610 := C.MSK_getnumqconknz64(self.ptr(),(C.int32_t)(k),(&_tmp609)); _tmp610 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp610)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp611 *int32
  qcsubi = make([]int32,_tmp609)
  if len(qcsubi) > 0 { _tmp611 = (*int32)(&qcsubi[0]) }
  var _tmp612 C.int64_t
  if _tmp613 := C.MSK_getnumqconknz64(self.ptr(),(C.int32_t)(k),(&_tmp612)); _tmp613 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp613)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp614 *int32
  qcsubj = make([]int32,_tmp612)
  if len(qcsubj) > 0 { _tmp614 = (*int32)(&qcsubj[0]) }
  var _tmp615 C.int64_t
  if _tmp616 := C.MSK_getnumqconknz64(self.ptr(),(C.int32_t)(k),(&_tmp615)); _tmp616 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp616)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp617 *float64
  qcval = make([]float64,_tmp615)
  if len(qcval) > 0 { _tmp617 = (*float64)(&qcval[0]) }
  if _tmp618 := C.MSK_getqconk64(self.ptr(),C.int32_t(k),C.int64_t(maxnumqcnz),(*C.int64_t)(&numqcnz),(*C.int32_t)(_tmp611),(*C.int32_t)(_tmp614),(*C.double)(_tmp617)); _tmp618 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp618)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains all the quadratic terms in the objective.
//
// - numqonz int64
//   Number of non-zero elements in the quadratic objective terms.
// - qosubi []int32
//   Row subscripts for quadratic objective coefficients.
// - qosubj []int32
//   Column subscripts for quadratic objective coefficients.
// - qoval []float64
//   Quadratic objective coefficient values.
func (self *Task) GetQObj() (numqonz int64,qosubi []int32,qosubj []int32,qoval []float64,err error) {
  var _tmp619 C.int64_t
  if _tmp620 := C.MSK_getnumqobjnz64(self.ptr(),(&_tmp619)); _tmp620 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp620)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumqonz int64 = int64(_tmp619)
  var _tmp621 *int32
  qosubi = make([]int32,maxnumqonz)
  if len(qosubi) > 0 { _tmp621 = (*int32)(&qosubi[0]) }
  var _tmp622 *int32
  qosubj = make([]int32,maxnumqonz)
  if len(qosubj) > 0 { _tmp622 = (*int32)(&qosubj[0]) }
  var _tmp623 *float64
  qoval = make([]float64,maxnumqonz)
  if len(qoval) > 0 { _tmp623 = (*float64)(&qoval[0]) }
  if _tmp624 := C.MSK_getqobj64(self.ptr(),C.int64_t(maxnumqonz),(*C.int64_t)(&numqonz),(*C.int32_t)(_tmp621),(*C.int32_t)(_tmp622),(*C.double)(_tmp623)); _tmp624 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp624)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains one coefficient from the quadratic term of the objective
//
// - i int32
//   Row index of the coefficient.
// - j int32
//   Column index of coefficient.
// - qoij float64
//   The required coefficient.
func (self *Task) GetQObjIJ(i int32,j int32) (qoij float64,err error) {
  if _tmp625 := C.MSK_getqobjij(self.ptr(),C.int32_t(i),C.int32_t(j),(*C.double)(&qoij)); _tmp625 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp625)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the reduced costs for a sequence of variables.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   The index of the first variable in the sequence.
// - last int32
//   The index of the last variable in the sequence plus 1.
// - redcosts []float64
//   Returns the requested reduced costs.
func (self *Task) GetReducedCosts(whichsol Soltype,first int32,last int32) (redcosts []float64,err error) {
  var _tmp626 *float64
  redcosts = make([]float64,(last - first))
  if len(redcosts) > 0 { _tmp626 = (*float64)(&redcosts[0]) }
  if _tmp627 := C.MSK_getreducedcosts(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp626)); _tmp627 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp627)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the status keys for the constraints.
//
// - whichsol Soltype
//   Selects a solution.
// - skc []Stakey
//   Status keys for the constraints.
func (self *Task) GetSkc(whichsol Soltype) (skc []Stakey,err error) {
  var _tmp628 C.int32_t
  if _tmp629 := C.MSK_getnumcon(self.ptr(),(&_tmp628)); _tmp629 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp629)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp630 *Stakey
  skc = make([]Stakey,_tmp628)
  if len(skc) > 0 { _tmp630 = (*Stakey)(&skc[0]) }
  if _tmp631 := C.MSK_getskc(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp630)); _tmp631 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp631)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the status keys for a slice of the constraints.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - skc []Stakey
//   Status keys for the constraints.
func (self *Task) GetSkcSlice(whichsol Soltype,first int32,last int32) (skc []Stakey,err error) {
  var _tmp632 *Stakey
  skc = make([]Stakey,(last - first))
  if len(skc) > 0 { _tmp632 = (*Stakey)(&skc[0]) }
  if _tmp633 := C.MSK_getskcslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp632)); _tmp633 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp633)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the status keys for the conic constraints.
//
// - whichsol Soltype
//   Selects a solution.
// - skn []Stakey
//   Status keys for the conic constraints.
func (self *Task) GetSkn(whichsol Soltype) (skn []Stakey,err error) {
  var _tmp634 C.int32_t
  if _tmp635 := C.MSK_getnumcone(self.ptr(),(&_tmp634)); _tmp635 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp635)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp636 *Stakey
  skn = make([]Stakey,_tmp634)
  if len(skn) > 0 { _tmp636 = (*Stakey)(&skn[0]) }
  if _tmp637 := C.MSK_getskn(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp636)); _tmp637 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp637)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the status keys for the scalar variables.
//
// - whichsol Soltype
//   Selects a solution.
// - skx []Stakey
//   Status keys for the variables.
func (self *Task) GetSkx(whichsol Soltype) (skx []Stakey,err error) {
  var _tmp638 C.int32_t
  if _tmp639 := C.MSK_getnumvar(self.ptr(),(&_tmp638)); _tmp639 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp639)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp640 *Stakey
  skx = make([]Stakey,_tmp638)
  if len(skx) > 0 { _tmp640 = (*Stakey)(&skx[0]) }
  if _tmp641 := C.MSK_getskx(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp640)); _tmp641 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp641)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the status keys for a slice of the scalar variables.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - skx []Stakey
//   Status keys for the variables.
func (self *Task) GetSkxSlice(whichsol Soltype,first int32,last int32) (skx []Stakey,err error) {
  var _tmp642 *Stakey
  skx = make([]Stakey,(last - first))
  if len(skx) > 0 { _tmp642 = (*Stakey)(&skx[0]) }
  if _tmp643 := C.MSK_getskxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp642)); _tmp643 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp643)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the slc vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - slc []float64
//   Dual variables corresponding to the lower bounds on the constraints.
func (self *Task) GetSlc(whichsol Soltype) (slc []float64,err error) {
  var _tmp644 C.int32_t
  if _tmp645 := C.MSK_getnumcon(self.ptr(),(&_tmp644)); _tmp645 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp645)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp646 *float64
  slc = make([]float64,_tmp644)
  if len(slc) > 0 { _tmp646 = (*float64)(&slc[0]) }
  if _tmp647 := C.MSK_getslc(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp646)); _tmp647 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp647)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a slice of the slc vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - slc []float64
//   Dual variables corresponding to the lower bounds on the constraints.
func (self *Task) GetSlcSlice(whichsol Soltype,first int32,last int32) (slc []float64,err error) {
  var _tmp648 *float64
  slc = make([]float64,(last - first))
  if len(slc) > 0 { _tmp648 = (*float64)(&slc[0]) }
  if _tmp649 := C.MSK_getslcslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp648)); _tmp649 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp649)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the slx vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - slx []float64
//   Dual variables corresponding to the lower bounds on the variables.
func (self *Task) GetSlx(whichsol Soltype) (slx []float64,err error) {
  var _tmp650 C.int32_t
  if _tmp651 := C.MSK_getnumvar(self.ptr(),(&_tmp650)); _tmp651 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp651)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp652 *float64
  slx = make([]float64,_tmp650)
  if len(slx) > 0 { _tmp652 = (*float64)(&slx[0]) }
  if _tmp653 := C.MSK_getslx(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp652)); _tmp653 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp653)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a slice of the slx vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - slx []float64
//   Dual variables corresponding to the lower bounds on the variables.
func (self *Task) GetSlxSlice(whichsol Soltype,first int32,last int32) (slx []float64,err error) {
  var _tmp654 *float64
  slx = make([]float64,(last - first))
  if len(slx) > 0 { _tmp654 = (*float64)(&slx[0]) }
  if _tmp655 := C.MSK_getslxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp654)); _tmp655 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp655)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the snx vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - snx []float64
//   Dual variables corresponding to the conic constraints on the variables.
func (self *Task) GetSnx(whichsol Soltype) (snx []float64,err error) {
  var _tmp656 C.int32_t
  if _tmp657 := C.MSK_getnumvar(self.ptr(),(&_tmp656)); _tmp657 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp657)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp658 *float64
  snx = make([]float64,_tmp656)
  if len(snx) > 0 { _tmp658 = (*float64)(&snx[0]) }
  if _tmp659 := C.MSK_getsnx(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp658)); _tmp659 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp659)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a slice of the snx vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - snx []float64
//   Dual variables corresponding to the conic constraints on the variables.
func (self *Task) GetSnxSlice(whichsol Soltype,first int32,last int32) (snx []float64,err error) {
  var _tmp660 *float64
  snx = make([]float64,(last - first))
  if len(snx) > 0 { _tmp660 = (*float64)(&snx[0]) }
  if _tmp661 := C.MSK_getsnxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp660)); _tmp661 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp661)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the solution status.
//
// - whichsol Soltype
//   Selects a solution.
// - solutionsta Solsta
//   Solution status.
func (self *Task) GetSolSta(whichsol Soltype) (solutionsta Solsta,err error) {
  var _tmp662 int32
  if _tmp663 := C.MSK_getsolsta(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(&_tmp662)); _tmp663 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp663)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  solutionsta = Solsta(_tmp662)
  return
}

// Obtains the complete solution.
//
// - whichsol Soltype
//   Selects a solution.
// - problemsta Prosta
//   Problem status.
// - solutionsta Solsta
//   Solution status.
// - skc []Stakey
//   Status keys for the constraints.
// - skx []Stakey
//   Status keys for the variables.
// - skn []Stakey
//   Status keys for the conic constraints.
// - xc []float64
//   Primal constraint solution.
// - xx []float64
//   Primal variable solution.
// - y []float64
//   Vector of dual variables corresponding to the constraints.
// - slc []float64
//   Dual variables corresponding to the lower bounds on the constraints.
// - suc []float64
//   Dual variables corresponding to the upper bounds on the constraints.
// - slx []float64
//   Dual variables corresponding to the lower bounds on the variables.
// - sux []float64
//   Dual variables corresponding to the upper bounds on the variables.
// - snx []float64
//   Dual variables corresponding to the conic constraints on the variables.
func (self *Task) GetSolution(whichsol Soltype) (problemsta Prosta,solutionsta Solsta,skc []Stakey,skx []Stakey,skn []Stakey,xc []float64,xx []float64,y []float64,slc []float64,suc []float64,slx []float64,sux []float64,snx []float64,err error) {
  var _tmp664 int32
  var _tmp665 int32
  var _tmp666 C.int32_t
  if _tmp667 := C.MSK_getnumcon(self.ptr(),(&_tmp666)); _tmp667 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp667)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp668 *Stakey
  skc = make([]Stakey,_tmp666)
  if len(skc) > 0 { _tmp668 = (*Stakey)(&skc[0]) }
  var _tmp669 C.int32_t
  if _tmp670 := C.MSK_getnumvar(self.ptr(),(&_tmp669)); _tmp670 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp670)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp671 *Stakey
  skx = make([]Stakey,_tmp669)
  if len(skx) > 0 { _tmp671 = (*Stakey)(&skx[0]) }
  var _tmp672 C.int32_t
  if _tmp673 := C.MSK_getnumcone(self.ptr(),(&_tmp672)); _tmp673 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp673)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp674 *Stakey
  skn = make([]Stakey,_tmp672)
  if len(skn) > 0 { _tmp674 = (*Stakey)(&skn[0]) }
  var _tmp675 C.int32_t
  if _tmp676 := C.MSK_getnumcon(self.ptr(),(&_tmp675)); _tmp676 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp676)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp677 *float64
  xc = make([]float64,_tmp675)
  if len(xc) > 0 { _tmp677 = (*float64)(&xc[0]) }
  var _tmp678 C.int32_t
  if _tmp679 := C.MSK_getnumvar(self.ptr(),(&_tmp678)); _tmp679 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp679)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp680 *float64
  xx = make([]float64,_tmp678)
  if len(xx) > 0 { _tmp680 = (*float64)(&xx[0]) }
  var _tmp681 C.int32_t
  if _tmp682 := C.MSK_getnumcon(self.ptr(),(&_tmp681)); _tmp682 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp682)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp683 *float64
  y = make([]float64,_tmp681)
  if len(y) > 0 { _tmp683 = (*float64)(&y[0]) }
  var _tmp684 C.int32_t
  if _tmp685 := C.MSK_getnumcon(self.ptr(),(&_tmp684)); _tmp685 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp685)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp686 *float64
  slc = make([]float64,_tmp684)
  if len(slc) > 0 { _tmp686 = (*float64)(&slc[0]) }
  var _tmp687 C.int32_t
  if _tmp688 := C.MSK_getnumcon(self.ptr(),(&_tmp687)); _tmp688 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp688)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp689 *float64
  suc = make([]float64,_tmp687)
  if len(suc) > 0 { _tmp689 = (*float64)(&suc[0]) }
  var _tmp690 C.int32_t
  if _tmp691 := C.MSK_getnumvar(self.ptr(),(&_tmp690)); _tmp691 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp691)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp692 *float64
  slx = make([]float64,_tmp690)
  if len(slx) > 0 { _tmp692 = (*float64)(&slx[0]) }
  var _tmp693 C.int32_t
  if _tmp694 := C.MSK_getnumvar(self.ptr(),(&_tmp693)); _tmp694 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp694)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp695 *float64
  sux = make([]float64,_tmp693)
  if len(sux) > 0 { _tmp695 = (*float64)(&sux[0]) }
  var _tmp696 C.int32_t
  if _tmp697 := C.MSK_getnumvar(self.ptr(),(&_tmp696)); _tmp697 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp697)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp698 *float64
  snx = make([]float64,_tmp696)
  if len(snx) > 0 { _tmp698 = (*float64)(&snx[0]) }
  if _tmp699 := C.MSK_getsolution(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(&_tmp664),(*C.int32_t)(&_tmp665),(*C.int32_t)(_tmp668),(*C.int32_t)(_tmp671),(*C.int32_t)(_tmp674),(*C.double)(_tmp677),(*C.double)(_tmp680),(*C.double)(_tmp683),(*C.double)(_tmp686),(*C.double)(_tmp689),(*C.double)(_tmp692),(*C.double)(_tmp695),(*C.double)(_tmp698)); _tmp699 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp699)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  problemsta = Prosta(_tmp664)
  solutionsta = Solsta(_tmp665)
  return
}

// Obtains information about of a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - pobj float64
//   The primal objective value.
// - pviolcon float64
//   Maximal primal bound violation for a xc variable.
// - pviolvar float64
//   Maximal primal bound violation for a xx variable.
// - pviolbarvar float64
//   Maximal primal bound violation for a barx variable.
// - pviolcone float64
//   Maximal primal violation of the solution with respect to the conic constraints.
// - pviolitg float64
//   Maximal violation in the integer constraints.
// - dobj float64
//   Dual objective value.
// - dviolcon float64
//   Maximal dual bound violation for a xc variable.
// - dviolvar float64
//   Maximal dual bound violation for a xx variable.
// - dviolbarvar float64
//   Maximal dual bound violation for a bars variable.
// - dviolcone float64
//   Maximum violation of the dual solution in the dual conic constraints.
func (self *Task) GetSolutionInfo(whichsol Soltype) (pobj float64,pviolcon float64,pviolvar float64,pviolbarvar float64,pviolcone float64,pviolitg float64,dobj float64,dviolcon float64,dviolvar float64,dviolbarvar float64,dviolcone float64,err error) {
  if _tmp700 := C.MSK_getsolutioninfo(self.ptr(),C.int32_t(whichsol),(*C.double)(&pobj),(*C.double)(&pviolcon),(*C.double)(&pviolvar),(*C.double)(&pviolbarvar),(*C.double)(&pviolcone),(*C.double)(&pviolitg),(*C.double)(&dobj),(*C.double)(&dviolcon),(*C.double)(&dviolvar),(*C.double)(&dviolbarvar),(*C.double)(&dviolcone)); _tmp700 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp700)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains information about of a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - pobj float64
//   The primal objective value.
// - pviolcon float64
//   Maximal primal bound violation for a xc variable.
// - pviolvar float64
//   Maximal primal bound violation for a xx variable.
// - pviolbarvar float64
//   Maximal primal bound violation for a barx variable.
// - pviolcone float64
//   Maximal primal violation of the solution with respect to the conic constraints.
// - pviolacc float64
//   Maximal primal violation of the solution with respect to the affine conic constraints.
// - pvioldjc float64
//   Maximal primal violation of the solution with respect to the disjunctive constraints.
// - pviolitg float64
//   Maximal violation in the integer constraints.
// - dobj float64
//   Dual objective value.
// - dviolcon float64
//   Maximal dual bound violation for a xc variable.
// - dviolvar float64
//   Maximal dual bound violation for a xx variable.
// - dviolbarvar float64
//   Maximal dual bound violation for a bars variable.
// - dviolcone float64
//   Maximum violation of the dual solution in the dual conic constraints.
// - dviolacc float64
//   Maximum violation of the dual solution in the dual affine conic constraints.
func (self *Task) GetSolutionInfoNew(whichsol Soltype) (pobj float64,pviolcon float64,pviolvar float64,pviolbarvar float64,pviolcone float64,pviolacc float64,pvioldjc float64,pviolitg float64,dobj float64,dviolcon float64,dviolvar float64,dviolbarvar float64,dviolcone float64,dviolacc float64,err error) {
  if _tmp701 := C.MSK_getsolutioninfonew(self.ptr(),C.int32_t(whichsol),(*C.double)(&pobj),(*C.double)(&pviolcon),(*C.double)(&pviolvar),(*C.double)(&pviolbarvar),(*C.double)(&pviolcone),(*C.double)(&pviolacc),(*C.double)(&pvioldjc),(*C.double)(&pviolitg),(*C.double)(&dobj),(*C.double)(&dviolcon),(*C.double)(&dviolvar),(*C.double)(&dviolbarvar),(*C.double)(&dviolcone),(*C.double)(&dviolacc)); _tmp701 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp701)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the complete solution.
//
// - whichsol Soltype
//   Selects a solution.
// - problemsta Prosta
//   Problem status.
// - solutionsta Solsta
//   Solution status.
// - skc []Stakey
//   Status keys for the constraints.
// - skx []Stakey
//   Status keys for the variables.
// - skn []Stakey
//   Status keys for the conic constraints.
// - xc []float64
//   Primal constraint solution.
// - xx []float64
//   Primal variable solution.
// - y []float64
//   Vector of dual variables corresponding to the constraints.
// - slc []float64
//   Dual variables corresponding to the lower bounds on the constraints.
// - suc []float64
//   Dual variables corresponding to the upper bounds on the constraints.
// - slx []float64
//   Dual variables corresponding to the lower bounds on the variables.
// - sux []float64
//   Dual variables corresponding to the upper bounds on the variables.
// - snx []float64
//   Dual variables corresponding to the conic constraints on the variables.
// - doty []float64
//   Dual variables corresponding to affine conic constraints.
func (self *Task) GetSolutionNew(whichsol Soltype) (problemsta Prosta,solutionsta Solsta,skc []Stakey,skx []Stakey,skn []Stakey,xc []float64,xx []float64,y []float64,slc []float64,suc []float64,slx []float64,sux []float64,snx []float64,doty []float64,err error) {
  var _tmp702 int32
  var _tmp703 int32
  var _tmp704 C.int32_t
  if _tmp705 := C.MSK_getnumcon(self.ptr(),(&_tmp704)); _tmp705 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp705)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp706 *Stakey
  skc = make([]Stakey,_tmp704)
  if len(skc) > 0 { _tmp706 = (*Stakey)(&skc[0]) }
  var _tmp707 C.int32_t
  if _tmp708 := C.MSK_getnumvar(self.ptr(),(&_tmp707)); _tmp708 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp708)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp709 *Stakey
  skx = make([]Stakey,_tmp707)
  if len(skx) > 0 { _tmp709 = (*Stakey)(&skx[0]) }
  var _tmp710 C.int32_t
  if _tmp711 := C.MSK_getnumcone(self.ptr(),(&_tmp710)); _tmp711 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp711)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp712 *Stakey
  skn = make([]Stakey,_tmp710)
  if len(skn) > 0 { _tmp712 = (*Stakey)(&skn[0]) }
  var _tmp713 C.int32_t
  if _tmp714 := C.MSK_getnumcon(self.ptr(),(&_tmp713)); _tmp714 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp714)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp715 *float64
  xc = make([]float64,_tmp713)
  if len(xc) > 0 { _tmp715 = (*float64)(&xc[0]) }
  var _tmp716 C.int32_t
  if _tmp717 := C.MSK_getnumvar(self.ptr(),(&_tmp716)); _tmp717 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp717)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp718 *float64
  xx = make([]float64,_tmp716)
  if len(xx) > 0 { _tmp718 = (*float64)(&xx[0]) }
  var _tmp719 C.int32_t
  if _tmp720 := C.MSK_getnumcon(self.ptr(),(&_tmp719)); _tmp720 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp720)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp721 *float64
  y = make([]float64,_tmp719)
  if len(y) > 0 { _tmp721 = (*float64)(&y[0]) }
  var _tmp722 C.int32_t
  if _tmp723 := C.MSK_getnumcon(self.ptr(),(&_tmp722)); _tmp723 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp723)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp724 *float64
  slc = make([]float64,_tmp722)
  if len(slc) > 0 { _tmp724 = (*float64)(&slc[0]) }
  var _tmp725 C.int32_t
  if _tmp726 := C.MSK_getnumcon(self.ptr(),(&_tmp725)); _tmp726 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp726)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp727 *float64
  suc = make([]float64,_tmp725)
  if len(suc) > 0 { _tmp727 = (*float64)(&suc[0]) }
  var _tmp728 C.int32_t
  if _tmp729 := C.MSK_getnumvar(self.ptr(),(&_tmp728)); _tmp729 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp729)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp730 *float64
  slx = make([]float64,_tmp728)
  if len(slx) > 0 { _tmp730 = (*float64)(&slx[0]) }
  var _tmp731 C.int32_t
  if _tmp732 := C.MSK_getnumvar(self.ptr(),(&_tmp731)); _tmp732 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp732)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp733 *float64
  sux = make([]float64,_tmp731)
  if len(sux) > 0 { _tmp733 = (*float64)(&sux[0]) }
  var _tmp734 C.int32_t
  if _tmp735 := C.MSK_getnumvar(self.ptr(),(&_tmp734)); _tmp735 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp735)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp736 *float64
  snx = make([]float64,_tmp734)
  if len(snx) > 0 { _tmp736 = (*float64)(&snx[0]) }
  var _tmp737 C.int64_t
  if _tmp738 := C.MSK_getaccntot(self.ptr(),(&_tmp737)); _tmp738 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp738)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp739 *float64
  doty = make([]float64,_tmp737)
  if len(doty) > 0 { _tmp739 = (*float64)(&doty[0]) }
  if _tmp740 := C.MSK_getsolutionnew(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(&_tmp702),(*C.int32_t)(&_tmp703),(*C.int32_t)(_tmp706),(*C.int32_t)(_tmp709),(*C.int32_t)(_tmp712),(*C.double)(_tmp715),(*C.double)(_tmp718),(*C.double)(_tmp721),(*C.double)(_tmp724),(*C.double)(_tmp727),(*C.double)(_tmp730),(*C.double)(_tmp733),(*C.double)(_tmp736),(*C.double)(_tmp739)); _tmp740 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp740)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  problemsta = Prosta(_tmp702)
  solutionsta = Solsta(_tmp703)
  return
}

// Obtains a slice of the solution.
//
// - whichsol Soltype
//   Selects a solution.
// - solitem Solitem
//   Which part of the solution is required.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - values []float64
//   The values of the requested solution elements.
func (self *Task) GetSolutionSlice(whichsol Soltype,solitem Solitem,first int32,last int32) (values []float64,err error) {
  var _tmp741 *float64
  values = make([]float64,(last - first))
  if len(values) > 0 { _tmp741 = (*float64)(&values[0]) }
  if _tmp742 := C.MSK_getsolutionslice(self.ptr(),C.int32_t(whichsol),C.int32_t(solitem),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp741)); _tmp742 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp742)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Gets a single symmetric matrix from the matrix store.
//
// - idx int64
//   Index of the matrix to retrieve.
// - subi []int32
//   Row subscripts of the matrix non-zero elements.
// - subj []int32
//   Column subscripts of the matrix non-zero elements.
// - valij []float64
//   Coefficients of the matrix non-zero elements.
func (self *Task) GetSparseSymMat(idx int64) (subi []int32,subj []int32,valij []float64,err error) {
  var _tmp743 C.int32_t
  var _tmp744 C.int64_t
  var _tmp745 C.int32_t
  if _tmp746 := C.MSK_getsymmatinfo(self.ptr(),(C.int64_t)(idx),(&_tmp743),(&_tmp744),(&_tmp745)); _tmp746 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp746)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxlen int64 = int64(_tmp744)
  var _tmp747 *int32
  subi = make([]int32,maxlen)
  if len(subi) > 0 { _tmp747 = (*int32)(&subi[0]) }
  var _tmp748 *int32
  subj = make([]int32,maxlen)
  if len(subj) > 0 { _tmp748 = (*int32)(&subj[0]) }
  var _tmp749 *float64
  valij = make([]float64,maxlen)
  if len(valij) > 0 { _tmp749 = (*float64)(&valij[0]) }
  if _tmp750 := C.MSK_getsparsesymmat(self.ptr(),C.int64_t(idx),C.int64_t(maxlen),(*C.int32_t)(_tmp747),(*C.int32_t)(_tmp748),(*C.double)(_tmp749)); _tmp750 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp750)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the value of a string parameter.
//
// - param Sparam
//   Which parameter.
// - len int32
//   The length of the parameter value.
// - parvalue string
//   If this is not a null pointer, the parameter value is stored here.
func (self *Task) GetStrParam(param Sparam) (len int32,parvalue string,err error) {
  var _tmp751 C.int32_t
  if _tmp752 := C.MSK_getstrparamlen(self.ptr(),(C.int32_t)(param),(&_tmp751)); _tmp752 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp752)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxlen int32 = int32((1 + _tmp751))
  _tmp753 := make([]byte,maxlen)
  if _tmp754 := C.MSK_getstrparam(self.ptr(),C.int32_t(param),C.int32_t(maxlen),(*C.int32_t)(&len),(*C.uchar)(&_tmp753[0])); _tmp754 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp754)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp753,byte(0)); p < 0 {
    parvalue = string(_tmp753)
  } else {
    parvalue = string(_tmp753[:p])
  }
  return
}

// Obtains the length of a string parameter.
//
// - param Sparam
//   Which parameter.
// - len int32
//   The length of the parameter value.
func (self *Task) GetStrParamLen(param Sparam) (len int32,err error) {
  if _tmp755 := C.MSK_getstrparamlen(self.ptr(),C.int32_t(param),(*C.int32_t)(&len)); _tmp755 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp755)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the suc vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - suc []float64
//   Dual variables corresponding to the upper bounds on the constraints.
func (self *Task) GetSuc(whichsol Soltype) (suc []float64,err error) {
  var _tmp756 C.int32_t
  if _tmp757 := C.MSK_getnumcon(self.ptr(),(&_tmp756)); _tmp757 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp757)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp758 *float64
  suc = make([]float64,_tmp756)
  if len(suc) > 0 { _tmp758 = (*float64)(&suc[0]) }
  if _tmp759 := C.MSK_getsuc(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp758)); _tmp759 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp759)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a slice of the suc vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - suc []float64
//   Dual variables corresponding to the upper bounds on the constraints.
func (self *Task) GetSucSlice(whichsol Soltype,first int32,last int32) (suc []float64,err error) {
  var _tmp760 *float64
  suc = make([]float64,(last - first))
  if len(suc) > 0 { _tmp760 = (*float64)(&suc[0]) }
  if _tmp761 := C.MSK_getsucslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp760)); _tmp761 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp761)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the sux vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - sux []float64
//   Dual variables corresponding to the upper bounds on the variables.
func (self *Task) GetSux(whichsol Soltype) (sux []float64,err error) {
  var _tmp762 C.int32_t
  if _tmp763 := C.MSK_getnumvar(self.ptr(),(&_tmp762)); _tmp763 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp763)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp764 *float64
  sux = make([]float64,_tmp762)
  if len(sux) > 0 { _tmp764 = (*float64)(&sux[0]) }
  if _tmp765 := C.MSK_getsux(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp764)); _tmp765 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp765)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a slice of the sux vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - sux []float64
//   Dual variables corresponding to the upper bounds on the variables.
func (self *Task) GetSuxSlice(whichsol Soltype,first int32,last int32) (sux []float64,err error) {
  var _tmp766 *float64
  sux = make([]float64,(last - first))
  if len(sux) > 0 { _tmp766 = (*float64)(&sux[0]) }
  if _tmp767 := C.MSK_getsuxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp766)); _tmp767 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp767)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains information about a matrix from the symmetric matrix storage.
//
// - idx int64
//   Index of the matrix for which information is requested.
// - dim int32
//   Returns the dimension of the requested matrix.
// - nz int64
//   Returns the number of non-zeros in the requested matrix.
// - mattype Symmattype
//   Returns the type of the requested matrix.
func (self *Task) GetSymMatInfo(idx int64) (dim int32,nz int64,mattype Symmattype,err error) {
  var _tmp768 int32
  if _tmp769 := C.MSK_getsymmatinfo(self.ptr(),C.int64_t(idx),(*C.int32_t)(&dim),(*C.int64_t)(&nz),(*C.int32_t)(&_tmp768)); _tmp769 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp769)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  mattype = Symmattype(_tmp768)
  return
}

// Obtains the task name.
//
// - taskname string
//   Returns the task name.
func (self *Task) GetTaskName() (taskname string,err error) {
  var _tmp770 C.int32_t
  if _tmp771 := C.MSK_gettasknamelen(self.ptr(),(&_tmp770)); _tmp771 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp771)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizetaskname int32 = int32((1 + _tmp770))
  _tmp772 := make([]byte,sizetaskname)
  if _tmp773 := C.MSK_gettaskname(self.ptr(),C.int32_t(sizetaskname),(*C.uchar)(&_tmp772[0])); _tmp773 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp773)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp772,byte(0)); p < 0 {
    taskname = string(_tmp772)
  } else {
    taskname = string(_tmp772[:p])
  }
  return
}

// Obtains the length the task name.
//
// - len int32
//   Returns the length of the task name.
func (self *Task) GetTaskNameLen() (len int32,err error) {
  if _tmp774 := C.MSK_gettasknamelen(self.ptr(),(*C.int32_t)(&len)); _tmp774 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp774)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains bound information for one variable.
//
// - i int32
//   Index of the variable for which the bound information should be obtained.
// - bk Boundkey
//   Bound keys.
// - bl float64
//   Values for lower bounds.
// - bu float64
//   Values for upper bounds.
func (self *Task) GetVarBound(i int32) (bk Boundkey,bl float64,bu float64,err error) {
  var _tmp775 int32
  if _tmp776 := C.MSK_getvarbound(self.ptr(),C.int32_t(i),(*C.int32_t)(&_tmp775),(*C.double)(&bl),(*C.double)(&bu)); _tmp776 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp776)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  bk = Boundkey(_tmp775)
  return
}

// Obtains bounds information for a slice of the variables.
//
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - bk []Boundkey
//   Bound keys.
// - bl []float64
//   Values for lower bounds.
// - bu []float64
//   Values for upper bounds.
func (self *Task) GetVarBoundSlice(first int32,last int32) (bk []Boundkey,bl []float64,bu []float64,err error) {
  var _tmp777 *Boundkey
  bk = make([]Boundkey,(last - first))
  if len(bk) > 0 { _tmp777 = (*Boundkey)(&bk[0]) }
  var _tmp778 *float64
  bl = make([]float64,(last - first))
  if len(bl) > 0 { _tmp778 = (*float64)(&bl[0]) }
  var _tmp779 *float64
  bu = make([]float64,(last - first))
  if len(bu) > 0 { _tmp779 = (*float64)(&bu[0]) }
  if _tmp780 := C.MSK_getvarboundslice(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp777),(*C.double)(_tmp778),(*C.double)(_tmp779)); _tmp780 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp780)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the name of a variable.
//
// - j int32
//   Index of a variable.
// - name string
//   Returns the required name.
func (self *Task) GetVarName(j int32) (name string,err error) {
  var _tmp781 C.int32_t
  if _tmp782 := C.MSK_getvarnamelen(self.ptr(),(C.int32_t)(j),(&_tmp781)); _tmp782 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp782)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp781))
  _tmp783 := make([]byte,sizename)
  if _tmp784 := C.MSK_getvarname(self.ptr(),C.int32_t(j),C.int32_t(sizename),(*C.uchar)(&_tmp783[0])); _tmp784 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp784)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp783,byte(0)); p < 0 {
    name = string(_tmp783)
  } else {
    name = string(_tmp783[:p])
  }
  return
}

// Checks whether the name has been assigned to any variable.
//
// - somename string
//   The name which should be checked.
// - asgn int32
//   Is non-zero if the name somename is assigned to a variable.
// - index int32
//   If the name somename is assigned to a variable, then return the index of the variable.
func (self *Task) GetVarNameIndex(somename string) (asgn int32,index int32,err error) {
  _tmp785 := C.CString(somename)
  if _tmp786 := C.MSK_getvarnameindex(self.ptr(),_tmp785,(*C.int32_t)(&asgn),(*C.int32_t)(&index)); _tmp786 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp786)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the length of the name of a variable.
//
// - i int32
//   Index of a variable.
// - len int32
//   Returns the length of the indicated name.
func (self *Task) GetVarNameLen(i int32) (len int32,err error) {
  if _tmp787 := C.MSK_getvarnamelen(self.ptr(),C.int32_t(i),(*C.int32_t)(&len)); _tmp787 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp787)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Gets the variable type of one variable.
//
// - j int32
//   Index of the variable.
// - vartype Variabletype
//   Variable type of variable index j.
func (self *Task) GetVarType(j int32) (vartype Variabletype,err error) {
  var _tmp788 int32
  if _tmp789 := C.MSK_getvartype(self.ptr(),C.int32_t(j),(*C.int32_t)(&_tmp788)); _tmp789 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp789)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  vartype = Variabletype(_tmp788)
  return
}

// Obtains the variable type for one or more variables.
//
// - subj []int32
//   A list of variable indexes.
// - vartype []Variabletype
//   Returns the variables types corresponding the variable indexes requested.
func (self *Task) GetVarTypeList(subj []int32) (vartype []Variabletype,err error) {
  _tmp790 := len(subj)
  var num int32 = int32(_tmp790)
  var _tmp791 *int32
  if len(subj) > 0 { _tmp791 = (*int32)(&subj[0]) }
  var _tmp792 *Variabletype
  vartype = make([]Variabletype,num)
  if len(vartype) > 0 { _tmp792 = (*Variabletype)(&vartype[0]) }
  if _tmp793 := C.MSK_getvartypelist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp791),(*C.int32_t)(_tmp792)); _tmp793 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp793)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the xc vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - xc []float64
//   Primal constraint solution.
func (self *Task) GetXc(whichsol Soltype) (xc []float64,err error) {
  var _tmp794 C.int32_t
  if _tmp795 := C.MSK_getnumcon(self.ptr(),(&_tmp794)); _tmp795 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp795)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp796 *float64
  xc = make([]float64,_tmp794)
  if len(xc) > 0 { _tmp796 = (*float64)(&xc[0]) }
  if _tmp797 := C.MSK_getxc(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp796)); _tmp797 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp797)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a slice of the xc vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - xc []float64
//   Primal constraint solution.
func (self *Task) GetXcSlice(whichsol Soltype,first int32,last int32) (xc []float64,err error) {
  var _tmp798 *float64
  xc = make([]float64,(last - first))
  if len(xc) > 0 { _tmp798 = (*float64)(&xc[0]) }
  if _tmp799 := C.MSK_getxcslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp798)); _tmp799 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp799)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the xx vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - xx []float64
//   Primal variable solution.
func (self *Task) GetXx(whichsol Soltype) (xx []float64,err error) {
  var _tmp800 C.int32_t
  if _tmp801 := C.MSK_getnumvar(self.ptr(),(&_tmp800)); _tmp801 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp801)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp802 *float64
  xx = make([]float64,_tmp800)
  if len(xx) > 0 { _tmp802 = (*float64)(&xx[0]) }
  if _tmp803 := C.MSK_getxx(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp802)); _tmp803 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp803)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a slice of the xx vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - xx []float64
//   Primal variable solution.
func (self *Task) GetXxSlice(whichsol Soltype,first int32,last int32) (xx []float64,err error) {
  var _tmp804 *float64
  xx = make([]float64,(last - first))
  if len(xx) > 0 { _tmp804 = (*float64)(&xx[0]) }
  if _tmp805 := C.MSK_getxxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp804)); _tmp805 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp805)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains the y vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - y []float64
//   Vector of dual variables corresponding to the constraints.
func (self *Task) GetY(whichsol Soltype) (y []float64,err error) {
  var _tmp806 C.int32_t
  if _tmp807 := C.MSK_getnumcon(self.ptr(),(&_tmp806)); _tmp807 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp807)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp808 *float64
  y = make([]float64,_tmp806)
  if len(y) > 0 { _tmp808 = (*float64)(&y[0]) }
  if _tmp809 := C.MSK_gety(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp808)); _tmp809 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp809)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a slice of the y vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - y []float64
//   Vector of dual variables corresponding to the constraints.
func (self *Task) GetYSlice(whichsol Soltype,first int32,last int32) (y []float64,err error) {
  var _tmp810 *float64
  y = make([]float64,(last - first))
  if len(y) > 0 { _tmp810 = (*float64)(&y[0]) }
  if _tmp811 := C.MSK_getyslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp810)); _tmp811 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp811)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Prints the infeasibility report to an output stream.
//
// - whichstream Streamtype
//   Index of the stream.
// - whichsol Soltype
//   Selects a solution.
func (self *Task) InfeasibilityReport(whichstream Streamtype,whichsol Soltype) (err error) {
  if _tmp812 := C.MSK_infeasibilityreport(self.ptr(),C.int32_t(whichstream),C.int32_t(whichsol)); _tmp812 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp812)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Prepare a task for basis solver.
//
// - basis []int32
//   The array of basis indexes to use.
func (self *Task) InitBasisSolve() (basis []int32,err error) {
  var _tmp813 C.int32_t
  if _tmp814 := C.MSK_getnumcon(self.ptr(),(&_tmp813)); _tmp814 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp814)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp815 *int32
  basis = make([]int32,_tmp813)
  if len(basis) > 0 { _tmp815 = (*int32)(&basis[0]) }
  if _tmp816 := C.MSK_initbasissolve(self.ptr(),(*C.int32_t)(_tmp815)); _tmp816 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp816)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Input the linear part of an optimization task in one function call.
//
// - maxnumcon int32
//   Number of preallocated constraints in the optimization task.
// - maxnumvar int32
//   Number of preallocated variables in the optimization task.
// - c []float64
//   Linear terms of the objective as a dense vector. The length is the number of variables.
// - cfix float64
//   Fixed term in the objective.
// - aptrb []int64
//   Row or column start pointers.
// - aptre []int64
//   Row or column end pointers.
// - asub []int32
//   Coefficient subscripts.
// - aval []float64
//   Coefficient values.
// - bkc []Boundkey
//   Bound keys for the constraints.
// - blc []float64
//   Lower bounds for the constraints.
// - buc []float64
//   Upper bounds for the constraints.
// - bkx []Boundkey
//   Bound keys for the variables.
// - blx []float64
//   Lower bounds for the variables.
// - bux []float64
//   Upper bounds for the variables.
func (self *Task) InputData(maxnumcon int32,maxnumvar int32,c []float64,cfix float64,aptrb []int64,aptre []int64,asub []int32,aval []float64,bkc []Boundkey,blc []float64,buc []float64,bkx []Boundkey,blx []float64,bux []float64) (err error) {
  _tmp817 := len(buc)
  if _tmp817 < len(bkc) { _tmp817 = len(bkc) }
  if _tmp817 < len(blc) { _tmp817 = len(blc) }
  var numcon int32 = int32(_tmp817)
  _tmp818 := len(aptrb)
  if _tmp818 < len(blx) { _tmp818 = len(blx) }
  if _tmp818 < len(bux) { _tmp818 = len(bux) }
  if _tmp818 < len(bkx) { _tmp818 = len(bkx) }
  if _tmp818 < len(aptre) { _tmp818 = len(aptre) }
  if _tmp818 < len(c) { _tmp818 = len(c) }
  var numvar int32 = int32(_tmp818)
  var _tmp819 *float64
  if len(c) > 0 { _tmp819 = (*float64)(&c[0]) }
  var _tmp820 *int64
  if len(aptrb) > 0 { _tmp820 = (*int64)(&aptrb[0]) }
  var _tmp821 *int64
  if len(aptre) > 0 { _tmp821 = (*int64)(&aptre[0]) }
  var _tmp822 *int32
  if len(asub) > 0 { _tmp822 = (*int32)(&asub[0]) }
  var _tmp823 *float64
  if len(aval) > 0 { _tmp823 = (*float64)(&aval[0]) }
  var _tmp824 *Boundkey
  if len(bkc) > 0 { _tmp824 = (*Boundkey)(&bkc[0]) }
  var _tmp825 *float64
  if len(blc) > 0 { _tmp825 = (*float64)(&blc[0]) }
  var _tmp826 *float64
  if len(buc) > 0 { _tmp826 = (*float64)(&buc[0]) }
  var _tmp827 *Boundkey
  if len(bkx) > 0 { _tmp827 = (*Boundkey)(&bkx[0]) }
  var _tmp828 *float64
  if len(blx) > 0 { _tmp828 = (*float64)(&blx[0]) }
  var _tmp829 *float64
  if len(bux) > 0 { _tmp829 = (*float64)(&bux[0]) }
  if _tmp830 := C.MSK_inputdata64(self.ptr(),C.int32_t(maxnumcon),C.int32_t(maxnumvar),C.int32_t(numcon),C.int32_t(numvar),(*C.double)(_tmp819),C.double(cfix),(*C.int64_t)(_tmp820),(*C.int64_t)(_tmp821),(*C.int32_t)(_tmp822),(*C.double)(_tmp823),(*C.int32_t)(_tmp824),(*C.double)(_tmp825),(*C.double)(_tmp826),(*C.int32_t)(_tmp827),(*C.double)(_tmp828),(*C.double)(_tmp829)); _tmp830 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp830)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Checks a double parameter name.
//
// - parname string
//   Parameter name.
// - param Dparam
//   Returns the parameter corresponding to the name, if one exists.
func (self *Task) IsDouParName(parname string) (param Dparam,err error) {
  _tmp831 := C.CString(parname)
  var _tmp832 int32
  if _tmp833 := C.MSK_isdouparname(self.ptr(),_tmp831,(*C.int32_t)(&_tmp832)); _tmp833 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp833)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  param = Dparam(_tmp832)
  return
}

// Checks an integer parameter name.
//
// - parname string
//   Parameter name.
// - param Iparam
//   Returns the parameter corresponding to the name, if one exists.
func (self *Task) IsIntParName(parname string) (param Iparam,err error) {
  _tmp834 := C.CString(parname)
  var _tmp835 int32
  if _tmp836 := C.MSK_isintparname(self.ptr(),_tmp834,(*C.int32_t)(&_tmp835)); _tmp836 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp836)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  param = Iparam(_tmp835)
  return
}

// Checks a string parameter name.
//
// - parname string
//   Parameter name.
// - param Sparam
//   Returns the parameter corresponding to the name, if one exists.
func (self *Task) IsStrParName(parname string) (param Sparam,err error) {
  _tmp837 := C.CString(parname)
  var _tmp838 int32
  if _tmp839 := C.MSK_isstrparname(self.ptr(),_tmp837,(*C.int32_t)(&_tmp838)); _tmp839 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp839)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  param = Sparam(_tmp838)
  return
}

// Directs all output from a task stream to a file.
//
// - whichstream Streamtype
//   Index of the stream.
// - filename string
//   A valid file name.
// - append int32
//   If this argument is 0 the output file will be overwritten, otherwise it will be appended to.
func (self *Task) LinkFileToStream(whichstream Streamtype,filename string,append int32) (err error) {
  _tmp840 := C.CString(filename)
  if _tmp841 := C.MSK_linkfiletotaskstream(self.ptr(),C.int32_t(whichstream),_tmp840,C.int32_t(append)); _tmp841 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp841)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Prints a short summary of a specified solution.
//
// - whichstream Streamtype
//   Index of the stream.
// - whichsol Soltype
//   Selects a solution.
func (self *Task) OneSolutionSummary(whichstream Streamtype,whichsol Soltype) (err error) {
  if _tmp842 := C.MSK_onesolutionsummary(self.ptr(),C.int32_t(whichstream),C.int32_t(whichsol)); _tmp842 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp842)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Prints a short summary with optimizer statistics from last optimization.
//
// - whichstream Streamtype
//   Index of the stream.
func (self *Task) OptimizerSummary(whichstream Streamtype) (err error) {
  if _tmp843 := C.MSK_optimizersummary(self.ptr(),C.int32_t(whichstream)); _tmp843 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp843)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Optimizes the problem.
//
// - trmcode Rescode
//   Is either OK or a termination response code.
func (self *Task) Optimize() (trmcode Rescode,err error) {
  var _tmp844 int32
  if _tmp845 := C.MSK_optimizetrm(self.ptr(),(*C.int32_t)(&_tmp844)); _tmp845 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp845)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  trmcode = Rescode(_tmp844)
  return
}

// Repairs a primal infeasible optimization problem by adjusting the bounds on the constraints and variables.
//
// - wlc []float64
//   Weights associated with relaxing lower bounds on the constraints.
// - wuc []float64
//   Weights associated with relaxing the upper bound on the constraints.
// - wlx []float64
//   Weights associated with relaxing the lower bounds of the variables.
// - wux []float64
//   Weights associated with relaxing the upper bounds of variables.
func (self *Task) PrimalRepair(wlc []float64,wuc []float64,wlx []float64,wux []float64) (err error) {
  var _tmp846 C.int32_t
  if _tmp847 := C.MSK_getnumcon(self.ptr(),(&_tmp846)); _tmp847 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp847)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if wlc != nil && int64(len(wlc)) < int64(_tmp846) {
    err = &ArrayLengthError{fun:"PrimalRepair",arg:"wlc"}
    return
  }
  var _tmp848 *float64
  if len(wlc) > 0 { _tmp848 = (*float64)(&wlc[0]) }
  var _tmp849 C.int32_t
  if _tmp850 := C.MSK_getnumcon(self.ptr(),(&_tmp849)); _tmp850 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp850)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if wuc != nil && int64(len(wuc)) < int64(_tmp849) {
    err = &ArrayLengthError{fun:"PrimalRepair",arg:"wuc"}
    return
  }
  var _tmp851 *float64
  if len(wuc) > 0 { _tmp851 = (*float64)(&wuc[0]) }
  var _tmp852 C.int32_t
  if _tmp853 := C.MSK_getnumvar(self.ptr(),(&_tmp852)); _tmp853 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp853)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if wlx != nil && int64(len(wlx)) < int64(_tmp852) {
    err = &ArrayLengthError{fun:"PrimalRepair",arg:"wlx"}
    return
  }
  var _tmp854 *float64
  if len(wlx) > 0 { _tmp854 = (*float64)(&wlx[0]) }
  var _tmp855 C.int32_t
  if _tmp856 := C.MSK_getnumvar(self.ptr(),(&_tmp855)); _tmp856 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp856)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if wux != nil && int64(len(wux)) < int64(_tmp855) {
    err = &ArrayLengthError{fun:"PrimalRepair",arg:"wux"}
    return
  }
  var _tmp857 *float64
  if len(wux) > 0 { _tmp857 = (*float64)(&wux[0]) }
  if _tmp858 := C.MSK_primalrepair(self.ptr(),(*C.double)(_tmp848),(*C.double)(_tmp851),(*C.double)(_tmp854),(*C.double)(_tmp857)); _tmp858 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp858)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Perform sensitivity analysis on bounds.
//
// - subi []int32
//   Indexes of constraints to analyze.
// - marki []Mark
//   Mark which constraint bounds to analyze.
// - subj []int32
//   Indexes of variables to analyze.
// - markj []Mark
//   Mark which variable bounds to analyze.
// - leftpricei []float64
//   Left shadow price for constraints.
// - rightpricei []float64
//   Right shadow price for constraints.
// - leftrangei []float64
//   Left range for constraints.
// - rightrangei []float64
//   Right range for constraints.
// - leftpricej []float64
//   Left shadow price for variables.
// - rightpricej []float64
//   Right shadow price for variables.
// - leftrangej []float64
//   Left range for variables.
// - rightrangej []float64
//   Right range for variables.
func (self *Task) PrimalSensitivity(subi []int32,marki []Mark,subj []int32,markj []Mark) (leftpricei []float64,rightpricei []float64,leftrangei []float64,rightrangei []float64,leftpricej []float64,rightpricej []float64,leftrangej []float64,rightrangej []float64,err error) {
  _tmp859 := len(subi)
  if _tmp859 < len(marki) { _tmp859 = len(marki) }
  var numi int32 = int32(_tmp859)
  var _tmp860 *int32
  if len(subi) > 0 { _tmp860 = (*int32)(&subi[0]) }
  var _tmp861 *Mark
  if len(marki) > 0 { _tmp861 = (*Mark)(&marki[0]) }
  _tmp862 := len(markj)
  if _tmp862 < len(subj) { _tmp862 = len(subj) }
  var numj int32 = int32(_tmp862)
  var _tmp863 *int32
  if len(subj) > 0 { _tmp863 = (*int32)(&subj[0]) }
  var _tmp864 *Mark
  if len(markj) > 0 { _tmp864 = (*Mark)(&markj[0]) }
  var _tmp865 *float64
  leftpricei = make([]float64,numi)
  if len(leftpricei) > 0 { _tmp865 = (*float64)(&leftpricei[0]) }
  var _tmp866 *float64
  rightpricei = make([]float64,numi)
  if len(rightpricei) > 0 { _tmp866 = (*float64)(&rightpricei[0]) }
  var _tmp867 *float64
  leftrangei = make([]float64,numi)
  if len(leftrangei) > 0 { _tmp867 = (*float64)(&leftrangei[0]) }
  var _tmp868 *float64
  rightrangei = make([]float64,numi)
  if len(rightrangei) > 0 { _tmp868 = (*float64)(&rightrangei[0]) }
  var _tmp869 *float64
  leftpricej = make([]float64,numj)
  if len(leftpricej) > 0 { _tmp869 = (*float64)(&leftpricej[0]) }
  var _tmp870 *float64
  rightpricej = make([]float64,numj)
  if len(rightpricej) > 0 { _tmp870 = (*float64)(&rightpricej[0]) }
  var _tmp871 *float64
  leftrangej = make([]float64,numj)
  if len(leftrangej) > 0 { _tmp871 = (*float64)(&leftrangej[0]) }
  var _tmp872 *float64
  rightrangej = make([]float64,numj)
  if len(rightrangej) > 0 { _tmp872 = (*float64)(&rightrangej[0]) }
  if _tmp873 := C.MSK_primalsensitivity(self.ptr(),C.int32_t(numi),(*C.int32_t)(_tmp860),(*C.int32_t)(_tmp861),C.int32_t(numj),(*C.int32_t)(_tmp863),(*C.int32_t)(_tmp864),(*C.double)(_tmp865),(*C.double)(_tmp866),(*C.double)(_tmp867),(*C.double)(_tmp868),(*C.double)(_tmp869),(*C.double)(_tmp870),(*C.double)(_tmp871),(*C.double)(_tmp872)); _tmp873 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp873)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a string containing the name of a given problem type.
//
// - probtype Problemtype
//   Problem type.
// - str string
//   String corresponding to the problem type key.
func (self *Task) ProbTypeToStr(probtype Problemtype) (str string,err error) {
  _tmp874 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp875 := C.MSK_probtypetostr(self.ptr(),C.int32_t(probtype),(*C.uchar)(&_tmp874[0])); _tmp875 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp875)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp874,byte(0)); p < 0 {
    str = string(_tmp874)
  } else {
    str = string(_tmp874[:p])
  }
  return
}

// Obtains a string containing the name of a given problem status.
//
// - problemsta Prosta
//   Problem status.
// - str string
//   String corresponding to the status key.
func (self *Task) ProStaToStr(problemsta Prosta) (str string,err error) {
  _tmp876 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp877 := C.MSK_prostatostr(self.ptr(),C.int32_t(problemsta),(*C.uchar)(&_tmp876[0])); _tmp877 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp877)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp876,byte(0)); p < 0 {
    str = string(_tmp876)
  } else {
    str = string(_tmp876[:p])
  }
  return
}

// Puts an affine conic constraint.
//
// - accidx int64
//   Affine conic constraint index.
// - domidx int64
//   Domain index.
// - afeidxlist []int64
//   List of affine expression indexes.
// - b []float64
//   The vector of constant terms added to affine expressions. Optional, can be NULL.
func (self *Task) PutAcc(accidx int64,domidx int64,afeidxlist []int64,b []float64) (err error) {
  _tmp878 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp878)
  var _tmp879 *int64
  if len(afeidxlist) > 0 { _tmp879 = (*int64)(&afeidxlist[0]) }
  if b != nil && int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"PutAcc",arg:"b"}
    return
  }
  var _tmp880 *float64
  if len(b) > 0 { _tmp880 = (*float64)(&b[0]) }
  if _tmp881 := C.MSK_putacc(self.ptr(),C.int64_t(accidx),C.int64_t(domidx),C.int64_t(numafeidx),(*C.int64_t)(_tmp879),(*C.double)(_tmp880)); _tmp881 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp881)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Puts the constant vector b in an affine conic constraint.
//
// - accidx int64
//   Affine conic constraint index.
// - b []float64
//   The vector of constant terms added to affine expressions. Optional, can be NULL.
func (self *Task) PutAccB(accidx int64,b []float64) (err error) {
  _tmp882 := len(b)
  var lengthb int64 = int64(_tmp882)
  var _tmp883 *float64
  if len(b) > 0 { _tmp883 = (*float64)(&b[0]) }
  if _tmp884 := C.MSK_putaccb(self.ptr(),C.int64_t(accidx),C.int64_t(lengthb),(*C.double)(_tmp883)); _tmp884 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp884)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets one element in the b vector of an affine conic constraint.
//
// - accidx int64
//   Affine conic constraint index.
// - j int64
//   The index of an element in b to change.
// - bj float64
//   The new value of b[j].
func (self *Task) PutAccBJ(accidx int64,j int64,bj float64) (err error) {
  if _tmp885 := C.MSK_putaccbj(self.ptr(),C.int64_t(accidx),C.int64_t(j),C.double(bj)); _tmp885 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp885)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Puts the doty vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - accidx int64
//   The index of the affine conic constraint.
// - doty []float64
//   The dual values for this affine conic constraint. The array should have length equal to the dimension of the constraint.
func (self *Task) PutAccDotY(whichsol Soltype,accidx int64) (doty []float64,err error) {
  var _tmp886 C.int64_t
  if _tmp887 := C.MSK_getaccn(self.ptr(),(C.int64_t)(accidx),(&_tmp886)); _tmp887 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp887)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp888 *float64
  doty = make([]float64,_tmp886)
  if len(doty) > 0 { _tmp888 = (*float64)(&doty[0]) }
  if _tmp889 := C.MSK_putaccdoty(self.ptr(),C.int32_t(whichsol),C.int64_t(accidx),(*C.double)(_tmp888)); _tmp889 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp889)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Puts a number of affine conic constraints.
//
// - accidxs []int64
//   Affine conic constraint indices.
// - domidxs []int64
//   Domain indices.
// - afeidxlist []int64
//   List of affine expression indexes.
// - b []float64
//   The vector of constant terms added to affine expressions. Optional, can be NULL.
func (self *Task) PutAccList(accidxs []int64,domidxs []int64,afeidxlist []int64,b []float64) (err error) {
  _tmp890 := len(accidxs)
  if _tmp890 < len(domidxs) { _tmp890 = len(domidxs) }
  var numaccs int64 = int64(_tmp890)
  var _tmp891 *int64
  if len(accidxs) > 0 { _tmp891 = (*int64)(&accidxs[0]) }
  var _tmp892 *int64
  if len(domidxs) > 0 { _tmp892 = (*int64)(&domidxs[0]) }
  _tmp893 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp893)
  var _tmp894 *int64
  if len(afeidxlist) > 0 { _tmp894 = (*int64)(&afeidxlist[0]) }
  if b != nil && int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"PutAccList",arg:"b"}
    return
  }
  var _tmp895 *float64
  if len(b) > 0 { _tmp895 = (*float64)(&b[0]) }
  if _tmp896 := C.MSK_putacclist(self.ptr(),C.int64_t(numaccs),(*C.int64_t)(_tmp891),(*C.int64_t)(_tmp892),C.int64_t(numafeidx),(*C.int64_t)(_tmp894),(*C.double)(_tmp895)); _tmp896 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp896)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the name of an affine conic constraint.
//
// - accidx int64
//   Index of the affine conic constraint.
// - name string
//   The name of the affine conic constraint.
func (self *Task) PutAccName(accidx int64,name string) (err error) {
  _tmp897 := C.CString(name)
  if _tmp898 := C.MSK_putaccname(self.ptr(),C.int64_t(accidx),_tmp897); _tmp898 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp898)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces all elements in one column of the linear constraint matrix.
//
// - j int32
//   Column index.
// - subj []int32
//   Row indexes of non-zero values in column.
// - valj []float64
//   New non-zero values of column.
func (self *Task) PutACol(j int32,subj []int32,valj []float64) (err error) {
  _tmp899 := len(valj)
  if _tmp899 < len(subj) { _tmp899 = len(subj) }
  var nzj int32 = int32(_tmp899)
  var _tmp900 *int32
  if len(subj) > 0 { _tmp900 = (*int32)(&subj[0]) }
  var _tmp901 *float64
  if len(valj) > 0 { _tmp901 = (*float64)(&valj[0]) }
  if _tmp902 := C.MSK_putacol(self.ptr(),C.int32_t(j),C.int32_t(nzj),(*C.int32_t)(_tmp900),(*C.double)(_tmp901)); _tmp902 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp902)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces all elements in several columns the linear constraint matrix.
//
// - sub []int32
//   Indexes of columns that should be replaced.
// - ptrb []int32
//   Array of pointers to the first element in the columns.
// - ptre []int32
//   Array of pointers to the last element plus one in the columns.
// - asub []int32
//   Row indexes
// - aval []float64
//   Coefficient values.
func (self *Task) PutAColList(sub []int32,ptrb []int32,ptre []int32,asub []int32,aval []float64) (err error) {
  _tmp903 := len(ptrb)
  if _tmp903 < len(ptre) { _tmp903 = len(ptre) }
  if _tmp903 < len(sub) { _tmp903 = len(sub) }
  var num int32 = int32(_tmp903)
  var _tmp904 *int32
  if len(sub) > 0 { _tmp904 = (*int32)(&sub[0]) }
  var _tmp905 *int32
  if len(ptrb) > 0 { _tmp905 = (*int32)(&ptrb[0]) }
  var _tmp906 *int32
  if len(ptre) > 0 { _tmp906 = (*int32)(&ptre[0]) }
  var _tmp907 *int32
  if len(asub) > 0 { _tmp907 = (*int32)(&asub[0]) }
  var _tmp908 *float64
  if len(aval) > 0 { _tmp908 = (*float64)(&aval[0]) }
  if _tmp909 := C.MSK_putacollist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp904),(*C.int32_t)(_tmp905),(*C.int32_t)(_tmp906),(*C.int32_t)(_tmp907),(*C.double)(_tmp908)); _tmp909 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp909)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces all elements in a sequence of columns the linear constraint matrix.
//
// - first int32
//   First column in the slice.
// - last int32
//   Last column plus one in the slice.
// - ptrb []int64
//   Array of pointers to the first element in the columns.
// - ptre []int64
//   Array of pointers to the last element plus one in the columns.
// - asub []int32
//   Row indexes
// - aval []float64
//   Coefficient values.
func (self *Task) PutAColSlice(first int32,last int32,ptrb []int64,ptre []int64,asub []int32,aval []float64) (err error) {
  var _tmp910 *int64
  if len(ptrb) > 0 { _tmp910 = (*int64)(&ptrb[0]) }
  var _tmp911 *int64
  if len(ptre) > 0 { _tmp911 = (*int64)(&ptre[0]) }
  var _tmp912 *int32
  if len(asub) > 0 { _tmp912 = (*int32)(&asub[0]) }
  var _tmp913 *float64
  if len(aval) > 0 { _tmp913 = (*float64)(&aval[0]) }
  if _tmp914 := C.MSK_putacolslice64(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int64_t)(_tmp910),(*C.int64_t)(_tmp911),(*C.int32_t)(_tmp912),(*C.double)(_tmp913)); _tmp914 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp914)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inputs barF in block triplet form.
//
// - afeidx []int64
//   Constraint index.
// - barvaridx []int32
//   Symmetric matrix variable index.
// - subk []int32
//   Block row index.
// - subl []int32
//   Block column index.
// - valkl []float64
//   The numerical value associated with each block triplet.
func (self *Task) PutAfeBarfBlockTriplet(afeidx []int64,barvaridx []int32,subk []int32,subl []int32,valkl []float64) (err error) {
  _tmp915 := len(subl)
  if _tmp915 < len(subk) { _tmp915 = len(subk) }
  if _tmp915 < len(barvaridx) { _tmp915 = len(barvaridx) }
  if _tmp915 < len(valkl) { _tmp915 = len(valkl) }
  if _tmp915 < len(afeidx) { _tmp915 = len(afeidx) }
  var numtrip int64 = int64(_tmp915)
  if int64(len(afeidx)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"afeidx"}
    return
  }
  var _tmp916 *int64
  if len(afeidx) > 0 { _tmp916 = (*int64)(&afeidx[0]) }
  if int64(len(barvaridx)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"barvaridx"}
    return
  }
  var _tmp917 *int32
  if len(barvaridx) > 0 { _tmp917 = (*int32)(&barvaridx[0]) }
  if int64(len(subk)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"subk"}
    return
  }
  var _tmp918 *int32
  if len(subk) > 0 { _tmp918 = (*int32)(&subk[0]) }
  if int64(len(subl)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"subl"}
    return
  }
  var _tmp919 *int32
  if len(subl) > 0 { _tmp919 = (*int32)(&subl[0]) }
  if int64(len(valkl)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"valkl"}
    return
  }
  var _tmp920 *float64
  if len(valkl) > 0 { _tmp920 = (*float64)(&valkl[0]) }
  if _tmp921 := C.MSK_putafebarfblocktriplet(self.ptr(),C.int64_t(numtrip),(*C.int64_t)(_tmp916),(*C.int32_t)(_tmp917),(*C.int32_t)(_tmp918),(*C.int32_t)(_tmp919),(*C.double)(_tmp920)); _tmp921 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp921)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inputs one entry in barF.
//
// - afeidx int64
//   Row index of barF.
// - barvaridx int32
//   Semidefinite variable index.
// - termidx []int64
//   Element indices in matrix storage.
// - termweight []float64
//   Weights in the weighted sum.
func (self *Task) PutAfeBarfEntry(afeidx int64,barvaridx int32,termidx []int64,termweight []float64) (err error) {
  _tmp922 := len(termidx)
  if _tmp922 < len(termweight) { _tmp922 = len(termweight) }
  var numterm int64 = int64(_tmp922)
  var _tmp923 *int64
  if len(termidx) > 0 { _tmp923 = (*int64)(&termidx[0]) }
  var _tmp924 *float64
  if len(termweight) > 0 { _tmp924 = (*float64)(&termweight[0]) }
  if _tmp925 := C.MSK_putafebarfentry(self.ptr(),C.int64_t(afeidx),C.int32_t(barvaridx),C.int64_t(numterm),(*C.int64_t)(_tmp923),(*C.double)(_tmp924)); _tmp925 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp925)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inputs a list of entries in barF.
//
// - afeidx []int64
//   Row indexes of barF.
// - barvaridx []int32
//   Semidefinite variable indexes.
// - numterm []int64
//   Number of terms in the weighted sums.
// - ptrterm []int64
//   Pointer to the terms forming each entry.
// - termidx []int64
//   Concatenated element indexes in matrix storage.
// - termweight []float64
//   Concatenated weights in the weighted sum.
func (self *Task) PutAfeBarfEntryList(afeidx []int64,barvaridx []int32,numterm []int64,ptrterm []int64,termidx []int64,termweight []float64) (err error) {
  _tmp926 := len(barvaridx)
  if _tmp926 < len(ptrterm) { _tmp926 = len(ptrterm) }
  if _tmp926 < len(numterm) { _tmp926 = len(numterm) }
  if _tmp926 < len(afeidx) { _tmp926 = len(afeidx) }
  var numafeidx int64 = int64(_tmp926)
  var _tmp927 *int64
  if len(afeidx) > 0 { _tmp927 = (*int64)(&afeidx[0]) }
  var _tmp928 *int32
  if len(barvaridx) > 0 { _tmp928 = (*int32)(&barvaridx[0]) }
  var _tmp929 *int64
  if len(numterm) > 0 { _tmp929 = (*int64)(&numterm[0]) }
  var _tmp930 *int64
  if len(ptrterm) > 0 { _tmp930 = (*int64)(&ptrterm[0]) }
  _tmp931 := len(termidx)
  if _tmp931 < len(termweight) { _tmp931 = len(termweight) }
  var lenterm int64 = int64(_tmp931)
  var _tmp932 *int64
  if len(termidx) > 0 { _tmp932 = (*int64)(&termidx[0]) }
  var _tmp933 *float64
  if len(termweight) > 0 { _tmp933 = (*float64)(&termweight[0]) }
  if _tmp934 := C.MSK_putafebarfentrylist(self.ptr(),C.int64_t(numafeidx),(*C.int64_t)(_tmp927),(*C.int32_t)(_tmp928),(*C.int64_t)(_tmp929),(*C.int64_t)(_tmp930),C.int64_t(lenterm),(*C.int64_t)(_tmp932),(*C.double)(_tmp933)); _tmp934 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp934)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inputs a row of barF.
//
// - afeidx int64
//   Row index of barF.
// - barvaridx []int32
//   Semidefinite variable indexes.
// - numterm []int64
//   Number of terms in the weighted sums.
// - ptrterm []int64
//   Pointer to the terms forming each entry.
// - termidx []int64
//   Concatenated element indexes in matrix storage.
// - termweight []float64
//   Concatenated weights in the weighted sum.
func (self *Task) PutAfeBarfRow(afeidx int64,barvaridx []int32,numterm []int64,ptrterm []int64,termidx []int64,termweight []float64) (err error) {
  _tmp935 := len(barvaridx)
  if _tmp935 < len(ptrterm) { _tmp935 = len(ptrterm) }
  if _tmp935 < len(numterm) { _tmp935 = len(numterm) }
  var numentr int32 = int32(_tmp935)
  var _tmp936 *int32
  if len(barvaridx) > 0 { _tmp936 = (*int32)(&barvaridx[0]) }
  var _tmp937 *int64
  if len(numterm) > 0 { _tmp937 = (*int64)(&numterm[0]) }
  var _tmp938 *int64
  if len(ptrterm) > 0 { _tmp938 = (*int64)(&ptrterm[0]) }
  _tmp939 := len(termidx)
  if _tmp939 < len(termweight) { _tmp939 = len(termweight) }
  var lenterm int64 = int64(_tmp939)
  var _tmp940 *int64
  if len(termidx) > 0 { _tmp940 = (*int64)(&termidx[0]) }
  var _tmp941 *float64
  if len(termweight) > 0 { _tmp941 = (*float64)(&termweight[0]) }
  if _tmp942 := C.MSK_putafebarfrow(self.ptr(),C.int64_t(afeidx),C.int32_t(numentr),(*C.int32_t)(_tmp936),(*C.int64_t)(_tmp937),(*C.int64_t)(_tmp938),C.int64_t(lenterm),(*C.int64_t)(_tmp940),(*C.double)(_tmp941)); _tmp942 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp942)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces all elements in one column of the F matrix in the affine expressions.
//
// - varidx int32
//   Column index.
// - afeidx []int64
//   Row indexes of non-zero values in the column.
// - val []float64
//   New non-zero values in the column.
func (self *Task) PutAfeFCol(varidx int32,afeidx []int64,val []float64) (err error) {
  _tmp943 := len(val)
  if _tmp943 < len(afeidx) { _tmp943 = len(afeidx) }
  var numnz int64 = int64(_tmp943)
  var _tmp944 *int64
  if len(afeidx) > 0 { _tmp944 = (*int64)(&afeidx[0]) }
  var _tmp945 *float64
  if len(val) > 0 { _tmp945 = (*float64)(&val[0]) }
  if _tmp946 := C.MSK_putafefcol(self.ptr(),C.int32_t(varidx),C.int64_t(numnz),(*C.int64_t)(_tmp944),(*C.double)(_tmp945)); _tmp946 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp946)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces one entry in F.
//
// - afeidx int64
//   Row index in F.
// - varidx int32
//   Column index in F.
// - value float64
//   Value of the entry.
func (self *Task) PutAfeFEntry(afeidx int64,varidx int32,value float64) (err error) {
  if _tmp947 := C.MSK_putafefentry(self.ptr(),C.int64_t(afeidx),C.int32_t(varidx),C.double(value)); _tmp947 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp947)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces a list of entries in F.
//
// - afeidx []int64
//   Row indices in F.
// - varidx []int32
//   Column indices in F.
// - val []float64
//   Values of the entries in F.
func (self *Task) PutAfeFEntryList(afeidx []int64,varidx []int32,val []float64) (err error) {
  _tmp948 := len(val)
  if _tmp948 < len(varidx) { _tmp948 = len(varidx) }
  if _tmp948 < len(afeidx) { _tmp948 = len(afeidx) }
  var numentr int64 = int64(_tmp948)
  var _tmp949 *int64
  if len(afeidx) > 0 { _tmp949 = (*int64)(&afeidx[0]) }
  var _tmp950 *int32
  if len(varidx) > 0 { _tmp950 = (*int32)(&varidx[0]) }
  var _tmp951 *float64
  if len(val) > 0 { _tmp951 = (*float64)(&val[0]) }
  if _tmp952 := C.MSK_putafefentrylist(self.ptr(),C.int64_t(numentr),(*C.int64_t)(_tmp949),(*C.int32_t)(_tmp950),(*C.double)(_tmp951)); _tmp952 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp952)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces all elements in one row of the F matrix in the affine expressions.
//
// - afeidx int64
//   Row index.
// - varidx []int32
//   Column indexes of non-zero values in the row.
// - val []float64
//   New non-zero values in the row.
func (self *Task) PutAfeFRow(afeidx int64,varidx []int32,val []float64) (err error) {
  _tmp953 := len(varidx)
  if _tmp953 < len(val) { _tmp953 = len(val) }
  var numnz int32 = int32(_tmp953)
  var _tmp954 *int32
  if len(varidx) > 0 { _tmp954 = (*int32)(&varidx[0]) }
  var _tmp955 *float64
  if len(val) > 0 { _tmp955 = (*float64)(&val[0]) }
  if _tmp956 := C.MSK_putafefrow(self.ptr(),C.int64_t(afeidx),C.int32_t(numnz),(*C.int32_t)(_tmp954),(*C.double)(_tmp955)); _tmp956 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp956)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces all elements in a number of rows of the F matrix in the affine expressions.
//
// - afeidx []int64
//   Row indices.
// - numnzrow []int32
//   Number of non-zeros in each row.
// - ptrrow []int64
//   Pointer to the first nonzero in each row.
// - varidx []int32
//   Column indexes of non-zero values.
// - val []float64
//   New non-zero values in the rows.
func (self *Task) PutAfeFRowList(afeidx []int64,numnzrow []int32,ptrrow []int64,varidx []int32,val []float64) (err error) {
  _tmp957 := len(ptrrow)
  if _tmp957 < len(numnzrow) { _tmp957 = len(numnzrow) }
  if _tmp957 < len(afeidx) { _tmp957 = len(afeidx) }
  var numafeidx int64 = int64(_tmp957)
  var _tmp958 *int64
  if len(afeidx) > 0 { _tmp958 = (*int64)(&afeidx[0]) }
  var _tmp959 *int32
  if len(numnzrow) > 0 { _tmp959 = (*int32)(&numnzrow[0]) }
  var _tmp960 *int64
  if len(ptrrow) > 0 { _tmp960 = (*int64)(&ptrrow[0]) }
  _tmp961 := len(varidx)
  if _tmp961 < len(val) { _tmp961 = len(val) }
  var lenidxval int64 = int64(_tmp961)
  var _tmp962 *int32
  if len(varidx) > 0 { _tmp962 = (*int32)(&varidx[0]) }
  var _tmp963 *float64
  if len(val) > 0 { _tmp963 = (*float64)(&val[0]) }
  if _tmp964 := C.MSK_putafefrowlist(self.ptr(),C.int64_t(numafeidx),(*C.int64_t)(_tmp958),(*C.int32_t)(_tmp959),(*C.int64_t)(_tmp960),C.int64_t(lenidxval),(*C.int32_t)(_tmp962),(*C.double)(_tmp963)); _tmp964 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp964)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces one element in the g vector in the affine expressions.
//
// - afeidx int64
//   Row index.
// - g float64
//   New value for the element of g.
func (self *Task) PutAfeG(afeidx int64,g float64) (err error) {
  if _tmp965 := C.MSK_putafeg(self.ptr(),C.int64_t(afeidx),C.double(g)); _tmp965 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp965)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces a list of elements in the g vector in the affine expressions.
//
// - afeidx []int64
//   Indices of entries in g.
// - g []float64
//   New values for the elements of g.
func (self *Task) PutAfeGList(afeidx []int64,g []float64) (err error) {
  _tmp966 := len(g)
  if _tmp966 < len(afeidx) { _tmp966 = len(afeidx) }
  var numafeidx int64 = int64(_tmp966)
  var _tmp967 *int64
  if len(afeidx) > 0 { _tmp967 = (*int64)(&afeidx[0]) }
  var _tmp968 *float64
  if len(g) > 0 { _tmp968 = (*float64)(&g[0]) }
  if _tmp969 := C.MSK_putafeglist(self.ptr(),C.int64_t(numafeidx),(*C.int64_t)(_tmp967),(*C.double)(_tmp968)); _tmp969 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp969)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Modifies a slice of the vector g.
//
// - first int64
//   First index in the sequence.
// - last int64
//   Last index plus 1 in the sequence.
// - slice []float64
//   The slice of g as a dense vector.
func (self *Task) PutAfeGSlice(first int64,last int64,slice []float64) (err error) {
  if int64(len(slice)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutAfeGSlice",arg:"slice"}
    return
  }
  var _tmp970 *float64
  if len(slice) > 0 { _tmp970 = (*float64)(&slice[0]) }
  if _tmp971 := C.MSK_putafegslice(self.ptr(),C.int64_t(first),C.int64_t(last),(*C.double)(_tmp970)); _tmp971 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp971)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes a single value in the linear coefficient matrix.
//
// - i int32
//   Constraint (row) index.
// - j int32
//   Variable (column) index.
// - aij float64
//   New coefficient.
func (self *Task) PutAij(i int32,j int32,aij float64) (err error) {
  if _tmp972 := C.MSK_putaij(self.ptr(),C.int32_t(i),C.int32_t(j),C.double(aij)); _tmp972 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp972)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes one or more coefficients in the linear constraint matrix.
//
// - subi []int32
//   Constraint (row) indices.
// - subj []int32
//   Variable (column) indices.
// - valij []float64
//   New coefficient values.
func (self *Task) PutAijList(subi []int32,subj []int32,valij []float64) (err error) {
  _tmp973 := len(subi)
  if _tmp973 < len(valij) { _tmp973 = len(valij) }
  if _tmp973 < len(subj) { _tmp973 = len(subj) }
  var num int64 = int64(_tmp973)
  var _tmp974 *int32
  if len(subi) > 0 { _tmp974 = (*int32)(&subi[0]) }
  var _tmp975 *int32
  if len(subj) > 0 { _tmp975 = (*int32)(&subj[0]) }
  var _tmp976 *float64
  if len(valij) > 0 { _tmp976 = (*float64)(&valij[0]) }
  if _tmp977 := C.MSK_putaijlist64(self.ptr(),C.int64_t(num),(*C.int32_t)(_tmp974),(*C.int32_t)(_tmp975),(*C.double)(_tmp976)); _tmp977 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp977)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces all elements in one row of the linear constraint matrix.
//
// - i int32
//   Row index.
// - subi []int32
//   Column indexes of non-zero values in row.
// - vali []float64
//   New non-zero values of row.
func (self *Task) PutARow(i int32,subi []int32,vali []float64) (err error) {
  _tmp978 := len(vali)
  if _tmp978 < len(subi) { _tmp978 = len(subi) }
  var nzi int32 = int32(_tmp978)
  var _tmp979 *int32
  if len(subi) > 0 { _tmp979 = (*int32)(&subi[0]) }
  var _tmp980 *float64
  if len(vali) > 0 { _tmp980 = (*float64)(&vali[0]) }
  if _tmp981 := C.MSK_putarow(self.ptr(),C.int32_t(i),C.int32_t(nzi),(*C.int32_t)(_tmp979),(*C.double)(_tmp980)); _tmp981 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp981)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces all elements in several rows of the linear constraint matrix.
//
// - sub []int32
//   Indexes of rows or columns that should be replaced.
// - ptrb []int64
//   Array of pointers to the first element in the rows.
// - ptre []int64
//   Array of pointers to the last element plus one in the rows.
// - asub []int32
//   Variable indexes.
// - aval []float64
//   Coefficient values.
func (self *Task) PutARowList(sub []int32,ptrb []int64,ptre []int64,asub []int32,aval []float64) (err error) {
  _tmp982 := len(ptrb)
  if _tmp982 < len(ptre) { _tmp982 = len(ptre) }
  if _tmp982 < len(sub) { _tmp982 = len(sub) }
  var num int32 = int32(_tmp982)
  var _tmp983 *int32
  if len(sub) > 0 { _tmp983 = (*int32)(&sub[0]) }
  var _tmp984 *int64
  if len(ptrb) > 0 { _tmp984 = (*int64)(&ptrb[0]) }
  var _tmp985 *int64
  if len(ptre) > 0 { _tmp985 = (*int64)(&ptre[0]) }
  var _tmp986 *int32
  if len(asub) > 0 { _tmp986 = (*int32)(&asub[0]) }
  var _tmp987 *float64
  if len(aval) > 0 { _tmp987 = (*float64)(&aval[0]) }
  if _tmp988 := C.MSK_putarowlist64(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp983),(*C.int64_t)(_tmp984),(*C.int64_t)(_tmp985),(*C.int32_t)(_tmp986),(*C.double)(_tmp987)); _tmp988 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp988)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces all elements in several rows the linear constraint matrix.
//
// - first int32
//   First row in the slice.
// - last int32
//   Last row plus one in the slice.
// - ptrb []int64
//   Array of pointers to the first element in the rows.
// - ptre []int64
//   Array of pointers to the last element plus one in the rows.
// - asub []int32
//   Column indexes of new elements.
// - aval []float64
//   Coefficient values.
func (self *Task) PutARowSlice(first int32,last int32,ptrb []int64,ptre []int64,asub []int32,aval []float64) (err error) {
  if int64(len(ptrb)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutARowSlice",arg:"ptrb"}
    return
  }
  var _tmp989 *int64
  if len(ptrb) > 0 { _tmp989 = (*int64)(&ptrb[0]) }
  if int64(len(ptre)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutARowSlice",arg:"ptre"}
    return
  }
  var _tmp990 *int64
  if len(ptre) > 0 { _tmp990 = (*int64)(&ptre[0]) }
  var _tmp991 *int32
  if len(asub) > 0 { _tmp991 = (*int32)(&asub[0]) }
  var _tmp992 *float64
  if len(aval) > 0 { _tmp992 = (*float64)(&aval[0]) }
  if _tmp993 := C.MSK_putarowslice64(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int64_t)(_tmp989),(*C.int64_t)(_tmp990),(*C.int32_t)(_tmp991),(*C.double)(_tmp992)); _tmp993 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp993)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Truncates all elements in A below a certain tolerance to zero.
//
// - tolzero float64
//   Truncation tolerance.
func (self *Task) PutATruncateTol(tolzero float64) (err error) {
  if _tmp994 := C.MSK_putatruncatetol(self.ptr(),C.double(tolzero)); _tmp994 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp994)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inputs barA in block triplet form.
//
// - subi []int32
//   Constraint index.
// - subj []int32
//   Symmetric matrix variable index.
// - subk []int32
//   Block row index.
// - subl []int32
//   Block column index.
// - valijkl []float64
//   The numerical value associated with each block triplet.
func (self *Task) PutBaraBlockTriplet(subi []int32,subj []int32,subk []int32,subl []int32,valijkl []float64) (err error) {
  _tmp995 := len(subl)
  if _tmp995 < len(valijkl) { _tmp995 = len(valijkl) }
  if _tmp995 < len(subk) { _tmp995 = len(subk) }
  if _tmp995 < len(subj) { _tmp995 = len(subj) }
  var num int64 = int64(_tmp995)
  if int64(len(subi)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"subi"}
    return
  }
  var _tmp996 *int32
  if len(subi) > 0 { _tmp996 = (*int32)(&subi[0]) }
  if int64(len(subj)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"subj"}
    return
  }
  var _tmp997 *int32
  if len(subj) > 0 { _tmp997 = (*int32)(&subj[0]) }
  if int64(len(subk)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"subk"}
    return
  }
  var _tmp998 *int32
  if len(subk) > 0 { _tmp998 = (*int32)(&subk[0]) }
  if int64(len(subl)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"subl"}
    return
  }
  var _tmp999 *int32
  if len(subl) > 0 { _tmp999 = (*int32)(&subl[0]) }
  if int64(len(valijkl)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"valijkl"}
    return
  }
  var _tmp1000 *float64
  if len(valijkl) > 0 { _tmp1000 = (*float64)(&valijkl[0]) }
  if _tmp1001 := C.MSK_putbarablocktriplet(self.ptr(),C.int64_t(num),(*C.int32_t)(_tmp996),(*C.int32_t)(_tmp997),(*C.int32_t)(_tmp998),(*C.int32_t)(_tmp999),(*C.double)(_tmp1000)); _tmp1001 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1001)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inputs an element of barA.
//
// - i int32
//   Row index of barA.
// - j int32
//   Column index of barA.
// - sub []int64
//   Element indexes in matrix storage.
// - weights []float64
//   Weights in the weighted sum.
func (self *Task) PutBaraIj(i int32,j int32,sub []int64,weights []float64) (err error) {
  _tmp1002 := len(weights)
  if _tmp1002 < len(sub) { _tmp1002 = len(sub) }
  var num int64 = int64(_tmp1002)
  var _tmp1003 *int64
  if len(sub) > 0 { _tmp1003 = (*int64)(&sub[0]) }
  var _tmp1004 *float64
  if len(weights) > 0 { _tmp1004 = (*float64)(&weights[0]) }
  if _tmp1005 := C.MSK_putbaraij(self.ptr(),C.int32_t(i),C.int32_t(j),C.int64_t(num),(*C.int64_t)(_tmp1003),(*C.double)(_tmp1004)); _tmp1005 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1005)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inputs list of elements of barA.
//
// - subi []int32
//   Row index of barA.
// - subj []int32
//   Column index of barA.
// - alphaptrb []int64
//   Start entries for terms in the weighted sum.
// - alphaptre []int64
//   End entries for terms in the weighted sum.
// - matidx []int64
//   Element indexes in matrix storage.
// - weights []float64
//   Weights in the weighted sum.
func (self *Task) PutBaraIjList(subi []int32,subj []int32,alphaptrb []int64,alphaptre []int64,matidx []int64,weights []float64) (err error) {
  _tmp1006 := len(alphaptrb)
  if _tmp1006 < len(subi) { _tmp1006 = len(subi) }
  if _tmp1006 < len(alphaptre) { _tmp1006 = len(alphaptre) }
  if _tmp1006 < len(subj) { _tmp1006 = len(subj) }
  var num int32 = int32(_tmp1006)
  var _tmp1007 *int32
  if len(subi) > 0 { _tmp1007 = (*int32)(&subi[0]) }
  var _tmp1008 *int32
  if len(subj) > 0 { _tmp1008 = (*int32)(&subj[0]) }
  var _tmp1009 *int64
  if len(alphaptrb) > 0 { _tmp1009 = (*int64)(&alphaptrb[0]) }
  var _tmp1010 *int64
  if len(alphaptre) > 0 { _tmp1010 = (*int64)(&alphaptre[0]) }
  var _tmp1011 *int64
  if len(matidx) > 0 { _tmp1011 = (*int64)(&matidx[0]) }
  var _tmp1012 *float64
  if len(weights) > 0 { _tmp1012 = (*float64)(&weights[0]) }
  if _tmp1013 := C.MSK_putbaraijlist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1007),(*C.int32_t)(_tmp1008),(*C.int64_t)(_tmp1009),(*C.int64_t)(_tmp1010),(*C.int64_t)(_tmp1011),(*C.double)(_tmp1012)); _tmp1013 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1013)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replace a set of rows of barA
//
// - subi []int32
//   Row indexes of barA.
// - ptrb []int64
//   Start of rows in barA.
// - ptre []int64
//   End of rows in barA.
// - subj []int32
//   Column index of barA.
// - nummat []int64
//   Number of entries in weighted sum of matrixes.
// - matidx []int64
//   Matrix indexes for weighted sum of matrixes.
// - weights []float64
//   Weights for weighted sum of matrixes.
func (self *Task) PutBaraRowList(subi []int32,ptrb []int64,ptre []int64,subj []int32,nummat []int64,matidx []int64,weights []float64) (err error) {
  _tmp1014 := len(ptrb)
  if _tmp1014 < len(ptre) { _tmp1014 = len(ptre) }
  if _tmp1014 < len(subi) { _tmp1014 = len(subi) }
  var num int32 = int32(_tmp1014)
  var _tmp1015 *int32
  if len(subi) > 0 { _tmp1015 = (*int32)(&subi[0]) }
  var _tmp1016 *int64
  if len(ptrb) > 0 { _tmp1016 = (*int64)(&ptrb[0]) }
  var _tmp1017 *int64
  if len(ptre) > 0 { _tmp1017 = (*int64)(&ptre[0]) }
  var _tmp1018 *int32
  if len(subj) > 0 { _tmp1018 = (*int32)(&subj[0]) }
  if int64(len(nummat)) < int64(len(subj)) {
    err = &ArrayLengthError{fun:"PutBaraRowList",arg:"nummat"}
    return
  }
  var _tmp1019 *int64
  if len(nummat) > 0 { _tmp1019 = (*int64)(&nummat[0]) }
  if int64(len(matidx)) < int64(sum(nummat)) {
    err = &ArrayLengthError{fun:"PutBaraRowList",arg:"matidx"}
    return
  }
  var _tmp1020 *int64
  if len(matidx) > 0 { _tmp1020 = (*int64)(&matidx[0]) }
  if int64(len(weights)) < int64(sum(nummat)) {
    err = &ArrayLengthError{fun:"PutBaraRowList",arg:"weights"}
    return
  }
  var _tmp1021 *float64
  if len(weights) > 0 { _tmp1021 = (*float64)(&weights[0]) }
  if _tmp1022 := C.MSK_putbararowlist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1015),(*C.int64_t)(_tmp1016),(*C.int64_t)(_tmp1017),(*C.int32_t)(_tmp1018),(*C.int64_t)(_tmp1019),(*C.int64_t)(_tmp1020),(*C.double)(_tmp1021)); _tmp1022 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1022)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inputs barC in block triplet form.
//
// - subj []int32
//   Symmetric matrix variable index.
// - subk []int32
//   Block row index.
// - subl []int32
//   Block column index.
// - valjkl []float64
//   The numerical value associated with each block triplet.
func (self *Task) PutBarcBlockTriplet(subj []int32,subk []int32,subl []int32,valjkl []float64) (err error) {
  _tmp1023 := len(valjkl)
  if _tmp1023 < len(subl) { _tmp1023 = len(subl) }
  if _tmp1023 < len(subk) { _tmp1023 = len(subk) }
  if _tmp1023 < len(subj) { _tmp1023 = len(subj) }
  var num int64 = int64(_tmp1023)
  if int64(len(subj)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBarcBlockTriplet",arg:"subj"}
    return
  }
  var _tmp1024 *int32
  if len(subj) > 0 { _tmp1024 = (*int32)(&subj[0]) }
  if int64(len(subk)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBarcBlockTriplet",arg:"subk"}
    return
  }
  var _tmp1025 *int32
  if len(subk) > 0 { _tmp1025 = (*int32)(&subk[0]) }
  if int64(len(subl)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBarcBlockTriplet",arg:"subl"}
    return
  }
  var _tmp1026 *int32
  if len(subl) > 0 { _tmp1026 = (*int32)(&subl[0]) }
  if int64(len(valjkl)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBarcBlockTriplet",arg:"valjkl"}
    return
  }
  var _tmp1027 *float64
  if len(valjkl) > 0 { _tmp1027 = (*float64)(&valjkl[0]) }
  if _tmp1028 := C.MSK_putbarcblocktriplet(self.ptr(),C.int64_t(num),(*C.int32_t)(_tmp1024),(*C.int32_t)(_tmp1025),(*C.int32_t)(_tmp1026),(*C.double)(_tmp1027)); _tmp1028 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1028)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes one element in barc.
//
// - j int32
//   Index of the element in barc` that should be changed.
// - sub []int64
//   sub is list of indexes of those symmetric matrices appearing in sum.
// - weights []float64
//   The weights of the terms in the weighted sum.
func (self *Task) PutBarcJ(j int32,sub []int64,weights []float64) (err error) {
  _tmp1029 := len(weights)
  if _tmp1029 < len(sub) { _tmp1029 = len(sub) }
  var num int64 = int64(_tmp1029)
  var _tmp1030 *int64
  if len(sub) > 0 { _tmp1030 = (*int64)(&sub[0]) }
  var _tmp1031 *float64
  if len(weights) > 0 { _tmp1031 = (*float64)(&weights[0]) }
  if _tmp1032 := C.MSK_putbarcj(self.ptr(),C.int32_t(j),C.int64_t(num),(*C.int64_t)(_tmp1030),(*C.double)(_tmp1031)); _tmp1032 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1032)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the dual solution for a semidefinite variable.
//
// - whichsol Soltype
//   Selects a solution.
// - j int32
//   Index of the semidefinite variable.
// - barsj []float64
//   Value of the j'th variable of barx.
func (self *Task) PutBarsJ(whichsol Soltype,j int32,barsj []float64) (err error) {
  var _tmp1033 C.int64_t
  if _tmp1034 := C.MSK_getlenbarvarj(self.ptr(),(C.int32_t)(j),(&_tmp1033)); _tmp1034 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1034)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(barsj)) < int64(_tmp1033) {
    err = &ArrayLengthError{fun:"PutBarsJ",arg:"barsj"}
    return
  }
  var _tmp1035 *float64
  if len(barsj) > 0 { _tmp1035 = (*float64)(&barsj[0]) }
  if _tmp1036 := C.MSK_putbarsj(self.ptr(),C.int32_t(whichsol),C.int32_t(j),(*C.double)(_tmp1035)); _tmp1036 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1036)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the name of a semidefinite variable.
//
// - j int32
//   Index of the variable.
// - name string
//   The variable name.
func (self *Task) PutBarvarName(j int32,name string) (err error) {
  _tmp1037 := C.CString(name)
  if _tmp1038 := C.MSK_putbarvarname(self.ptr(),C.int32_t(j),_tmp1037); _tmp1038 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1038)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the primal solution for a semidefinite variable.
//
// - whichsol Soltype
//   Selects a solution.
// - j int32
//   Index of the semidefinite variable.
// - barxj []float64
//   Value of the j'th variable of barx.
func (self *Task) PutBarxJ(whichsol Soltype,j int32,barxj []float64) (err error) {
  var _tmp1039 C.int64_t
  if _tmp1040 := C.MSK_getlenbarvarj(self.ptr(),(C.int32_t)(j),(&_tmp1039)); _tmp1040 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1040)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(barxj)) < int64(_tmp1039) {
    err = &ArrayLengthError{fun:"PutBarxJ",arg:"barxj"}
    return
  }
  var _tmp1041 *float64
  if len(barxj) > 0 { _tmp1041 = (*float64)(&barxj[0]) }
  if _tmp1042 := C.MSK_putbarxj(self.ptr(),C.int32_t(whichsol),C.int32_t(j),(*C.double)(_tmp1041)); _tmp1042 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1042)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces the fixed term in the objective.
//
// - cfix float64
//   Fixed term in the objective.
func (self *Task) PutCfix(cfix float64) (err error) {
  if _tmp1043 := C.MSK_putcfix(self.ptr(),C.double(cfix)); _tmp1043 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1043)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Modifies one linear coefficient in the objective.
//
// - j int32
//   Index of the variable whose objective coefficient should be changed.
// - cj float64
//   New coefficient value.
func (self *Task) PutCJ(j int32,cj float64) (err error) {
  if _tmp1044 := C.MSK_putcj(self.ptr(),C.int32_t(j),C.double(cj)); _tmp1044 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1044)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Modifies a part of the linear objective coefficients.
//
// - subj []int32
//   Indices of variables for which objective coefficients should be changed.
// - val []float64
//   New numerical values for the objective coefficients that should be modified.
func (self *Task) PutCList(subj []int32,val []float64) (err error) {
  _tmp1045 := len(val)
  if _tmp1045 < len(subj) { _tmp1045 = len(subj) }
  var num int32 = int32(_tmp1045)
  var _tmp1046 *int32
  if len(subj) > 0 { _tmp1046 = (*int32)(&subj[0]) }
  var _tmp1047 *float64
  if len(val) > 0 { _tmp1047 = (*float64)(&val[0]) }
  if _tmp1048 := C.MSK_putclist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1046),(*C.double)(_tmp1047)); _tmp1048 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1048)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes the bound for one constraint.
//
// - i int32
//   Index of the constraint.
// - bkc Boundkey
//   New bound key.
// - blc float64
//   New lower bound.
// - buc float64
//   New upper bound.
func (self *Task) PutConBound(i int32,bkc Boundkey,blc float64,buc float64) (err error) {
  if _tmp1049 := C.MSK_putconbound(self.ptr(),C.int32_t(i),C.int32_t(bkc),C.double(blc),C.double(buc)); _tmp1049 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1049)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes the bounds of a list of constraints.
//
// - sub []int32
//   List of constraint indexes.
// - bkc []Boundkey
//   Bound keys for the constraints.
// - blc []float64
//   Lower bounds for the constraints.
// - buc []float64
//   Upper bounds for the constraints.
func (self *Task) PutConBoundList(sub []int32,bkc []Boundkey,blc []float64,buc []float64) (err error) {
  _tmp1050 := len(bkc)
  if _tmp1050 < len(blc) { _tmp1050 = len(blc) }
  if _tmp1050 < len(buc) { _tmp1050 = len(buc) }
  if _tmp1050 < len(sub) { _tmp1050 = len(sub) }
  var num int32 = int32(_tmp1050)
  var _tmp1051 *int32
  if len(sub) > 0 { _tmp1051 = (*int32)(&sub[0]) }
  var _tmp1052 *Boundkey
  if len(bkc) > 0 { _tmp1052 = (*Boundkey)(&bkc[0]) }
  var _tmp1053 *float64
  if len(blc) > 0 { _tmp1053 = (*float64)(&blc[0]) }
  var _tmp1054 *float64
  if len(buc) > 0 { _tmp1054 = (*float64)(&buc[0]) }
  if _tmp1055 := C.MSK_putconboundlist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1051),(*C.int32_t)(_tmp1052),(*C.double)(_tmp1053),(*C.double)(_tmp1054)); _tmp1055 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1055)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes the bounds of a list of constraints.
//
// - sub []int32
//   List of constraint indexes.
// - bkc Boundkey
//   New bound key for all constraints in the list.
// - blc float64
//   New lower bound for all constraints in the list.
// - buc float64
//   New upper bound for all constraints in the list.
func (self *Task) PutConBoundListConst(sub []int32,bkc Boundkey,blc float64,buc float64) (err error) {
  _tmp1056 := len(sub)
  var num int32 = int32(_tmp1056)
  var _tmp1057 *int32
  if len(sub) > 0 { _tmp1057 = (*int32)(&sub[0]) }
  if _tmp1058 := C.MSK_putconboundlistconst(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1057),C.int32_t(bkc),C.double(blc),C.double(buc)); _tmp1058 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1058)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes the bounds for a slice of the constraints.
//
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - bkc []Boundkey
//   Bound keys for the constraints.
// - blc []float64
//   Lower bounds for the constraints.
// - buc []float64
//   Upper bounds for the constraints.
func (self *Task) PutConBoundSlice(first int32,last int32,bkc []Boundkey,blc []float64,buc []float64) (err error) {
  if int64(len(bkc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutConBoundSlice",arg:"bkc"}
    return
  }
  var _tmp1059 *Boundkey
  if len(bkc) > 0 { _tmp1059 = (*Boundkey)(&bkc[0]) }
  if int64(len(blc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutConBoundSlice",arg:"blc"}
    return
  }
  var _tmp1060 *float64
  if len(blc) > 0 { _tmp1060 = (*float64)(&blc[0]) }
  if int64(len(buc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutConBoundSlice",arg:"buc"}
    return
  }
  var _tmp1061 *float64
  if len(buc) > 0 { _tmp1061 = (*float64)(&buc[0]) }
  if _tmp1062 := C.MSK_putconboundslice(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp1059),(*C.double)(_tmp1060),(*C.double)(_tmp1061)); _tmp1062 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1062)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes the bounds for a slice of the constraints.
//
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - bkc Boundkey
//   New bound key for all constraints in the slice.
// - blc float64
//   New lower bound for all constraints in the slice.
// - buc float64
//   New upper bound for all constraints in the slice.
func (self *Task) PutConBoundSliceConst(first int32,last int32,bkc Boundkey,blc float64,buc float64) (err error) {
  if _tmp1063 := C.MSK_putconboundsliceconst(self.ptr(),C.int32_t(first),C.int32_t(last),C.int32_t(bkc),C.double(blc),C.double(buc)); _tmp1063 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1063)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces a conic constraint.
//
// - k int32
//   Index of the cone.
// - ct Conetype
//   Specifies the type of the cone.
// - conepar float64
//   For the power cone it denotes the exponent alpha. For other cone types it is unused and can be set to 0.
// - submem []int32
//   Variable subscripts of the members in the cone.
func (self *Task) PutCone(k int32,ct Conetype,conepar float64,submem []int32) (err error) {
  _tmp1064 := len(submem)
  var nummem int32 = int32(_tmp1064)
  var _tmp1065 *int32
  if len(submem) > 0 { _tmp1065 = (*int32)(&submem[0]) }
  if _tmp1066 := C.MSK_putcone(self.ptr(),C.int32_t(k),C.int32_t(ct),C.double(conepar),C.int32_t(nummem),(*C.int32_t)(_tmp1065)); _tmp1066 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1066)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the name of a cone.
//
// - j int32
//   Index of the cone.
// - name string
//   The name of the cone.
func (self *Task) PutConeName(j int32,name string) (err error) {
  _tmp1067 := C.CString(name)
  if _tmp1068 := C.MSK_putconename(self.ptr(),C.int32_t(j),_tmp1067); _tmp1068 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1068)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the name of a constraint.
//
// - i int32
//   Index of the constraint.
// - name string
//   The name of the constraint.
func (self *Task) PutConName(i int32,name string) (err error) {
  _tmp1069 := C.CString(name)
  if _tmp1070 := C.MSK_putconname(self.ptr(),C.int32_t(i),_tmp1069); _tmp1070 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1070)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the primal and dual solution information for a single constraint.
//
// - i int32
//   Index of the constraint.
// - whichsol Soltype
//   Selects a solution.
// - sk Stakey
//   Status key of the constraint.
// - x float64
//   Primal solution value of the constraint.
// - sl float64
//   Solution value of the dual variable associated with the lower bound.
// - su float64
//   Solution value of the dual variable associated with the upper bound.
func (self *Task) PutConSolutionI(i int32,whichsol Soltype,sk Stakey,x float64,sl float64,su float64) (err error) {
  if _tmp1071 := C.MSK_putconsolutioni(self.ptr(),C.int32_t(i),C.int32_t(whichsol),C.int32_t(sk),C.double(x),C.double(sl),C.double(su)); _tmp1071 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1071)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Modifies a slice of the linear objective coefficients.
//
// - first int32
//   First element in the slice of c.
// - last int32
//   Last element plus 1 of the slice in c to be changed.
// - slice []float64
//   New numerical values for the objective coefficients that should be modified.
func (self *Task) PutCSlice(first int32,last int32,slice []float64) (err error) {
  if int64(len(slice)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutCSlice",arg:"slice"}
    return
  }
  var _tmp1072 *float64
  if len(slice) > 0 { _tmp1072 = (*float64)(&slice[0]) }
  if _tmp1073 := C.MSK_putcslice(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1072)); _tmp1073 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1073)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inputs a disjunctive constraint.
//
// - djcidx int64
//   Index of the disjunctive constraint.
// - domidxlist []int64
//   List of domain indexes.
// - afeidxlist []int64
//   List of affine expression indexes.
// - b []float64
//   The vector of constant terms added to affine expressions.
// - termsizelist []int64
//   List of term sizes.
func (self *Task) PutDjc(djcidx int64,domidxlist []int64,afeidxlist []int64,b []float64,termsizelist []int64) (err error) {
  _tmp1074 := len(domidxlist)
  var numdomidx int64 = int64(_tmp1074)
  var _tmp1075 *int64
  if len(domidxlist) > 0 { _tmp1075 = (*int64)(&domidxlist[0]) }
  _tmp1076 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp1076)
  var _tmp1077 *int64
  if len(afeidxlist) > 0 { _tmp1077 = (*int64)(&afeidxlist[0]) }
  if b != nil && int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"PutDjc",arg:"b"}
    return
  }
  var _tmp1078 *float64
  if len(b) > 0 { _tmp1078 = (*float64)(&b[0]) }
  _tmp1079 := len(termsizelist)
  var numterms int64 = int64(_tmp1079)
  var _tmp1080 *int64
  if len(termsizelist) > 0 { _tmp1080 = (*int64)(&termsizelist[0]) }
  if _tmp1081 := C.MSK_putdjc(self.ptr(),C.int64_t(djcidx),C.int64_t(numdomidx),(*C.int64_t)(_tmp1075),C.int64_t(numafeidx),(*C.int64_t)(_tmp1077),(*C.double)(_tmp1078),C.int64_t(numterms),(*C.int64_t)(_tmp1080)); _tmp1081 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1081)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the name of a disjunctive constraint.
//
// - djcidx int64
//   Index of the disjunctive constraint.
// - name string
//   The name of the disjunctive constraint.
func (self *Task) PutDjcName(djcidx int64,name string) (err error) {
  _tmp1082 := C.CString(name)
  if _tmp1083 := C.MSK_putdjcname(self.ptr(),C.int64_t(djcidx),_tmp1082); _tmp1083 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1083)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inputs a slice of disjunctive constraints.
//
// - idxfirst int64
//   Index of the first disjunctive constraint in the slice.
// - idxlast int64
//   Index of the last disjunctive constraint in the slice plus 1.
// - domidxlist []int64
//   List of domain indexes.
// - afeidxlist []int64
//   List of affine expression indexes.
// - b []float64
//   The vector of constant terms added to affine expressions. Optional, may be NULL.
// - termsizelist []int64
//   List of term sizes.
// - termsindjc []int64
//   Number of terms in each of the disjunctive constraints in the slice.
func (self *Task) PutDjcSlice(idxfirst int64,idxlast int64,domidxlist []int64,afeidxlist []int64,b []float64,termsizelist []int64,termsindjc []int64) (err error) {
  _tmp1084 := len(domidxlist)
  var numdomidx int64 = int64(_tmp1084)
  var _tmp1085 *int64
  if len(domidxlist) > 0 { _tmp1085 = (*int64)(&domidxlist[0]) }
  _tmp1086 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp1086)
  var _tmp1087 *int64
  if len(afeidxlist) > 0 { _tmp1087 = (*int64)(&afeidxlist[0]) }
  if b != nil && int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"PutDjcSlice",arg:"b"}
    return
  }
  var _tmp1088 *float64
  if len(b) > 0 { _tmp1088 = (*float64)(&b[0]) }
  _tmp1089 := len(termsizelist)
  var numterms int64 = int64(_tmp1089)
  var _tmp1090 *int64
  if len(termsizelist) > 0 { _tmp1090 = (*int64)(&termsizelist[0]) }
  if int64(len(termsindjc)) < int64((idxlast - idxfirst)) {
    err = &ArrayLengthError{fun:"PutDjcSlice",arg:"termsindjc"}
    return
  }
  var _tmp1091 *int64
  if len(termsindjc) > 0 { _tmp1091 = (*int64)(&termsindjc[0]) }
  if _tmp1092 := C.MSK_putdjcslice(self.ptr(),C.int64_t(idxfirst),C.int64_t(idxlast),C.int64_t(numdomidx),(*C.int64_t)(_tmp1085),C.int64_t(numafeidx),(*C.int64_t)(_tmp1087),(*C.double)(_tmp1088),C.int64_t(numterms),(*C.int64_t)(_tmp1090),(*C.int64_t)(_tmp1091)); _tmp1092 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1092)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the name of a domain.
//
// - domidx int64
//   Index of the domain.
// - name string
//   The name of the domain.
func (self *Task) PutDomainName(domidx int64,name string) (err error) {
  _tmp1093 := C.CString(name)
  if _tmp1094 := C.MSK_putdomainname(self.ptr(),C.int64_t(domidx),_tmp1093); _tmp1094 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1094)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets a double parameter.
//
// - param Dparam
//   Which parameter.
// - parvalue float64
//   Parameter value.
func (self *Task) PutDouParam(param Dparam,parvalue float64) (err error) {
  if _tmp1095 := C.MSK_putdouparam(self.ptr(),C.int32_t(param),C.double(parvalue)); _tmp1095 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1095)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets an integer parameter.
//
// - param Iparam
//   Which parameter.
// - parvalue int32
//   Parameter value.
func (self *Task) PutIntParam(param Iparam,parvalue int32) (err error) {
  if _tmp1096 := C.MSK_putintparam(self.ptr(),C.int32_t(param),C.int32_t(parvalue)); _tmp1096 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1096)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the number of preallocated affine conic constraints.
//
// - maxnumacc int64
//   Number of preallocated affine conic constraints.
func (self *Task) PutMaxNumAcc(maxnumacc int64) (err error) {
  if _tmp1097 := C.MSK_putmaxnumacc(self.ptr(),C.int64_t(maxnumacc)); _tmp1097 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1097)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the number of preallocated affine expressions in the optimization task.
//
// - maxnumafe int64
//   Number of preallocated affine expressions.
func (self *Task) PutMaxNumAfe(maxnumafe int64) (err error) {
  if _tmp1098 := C.MSK_putmaxnumafe(self.ptr(),C.int64_t(maxnumafe)); _tmp1098 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1098)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the number of preallocated non-zero entries in the linear coefficient matrix.
//
// - maxnumanz int64
//   New size of the storage reserved for storing the linear coefficient matrix.
func (self *Task) PutMaxNumANz(maxnumanz int64) (err error) {
  if _tmp1099 := C.MSK_putmaxnumanz(self.ptr(),C.int64_t(maxnumanz)); _tmp1099 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1099)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the number of preallocated symmetric matrix variables.
//
// - maxnumbarvar int32
//   Number of preallocated symmetric matrix variables.
func (self *Task) PutMaxNumBarvar(maxnumbarvar int32) (err error) {
  if _tmp1100 := C.MSK_putmaxnumbarvar(self.ptr(),C.int32_t(maxnumbarvar)); _tmp1100 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1100)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the number of preallocated constraints in the optimization task.
//
// - maxnumcon int32
//   Number of preallocated constraints in the optimization task.
func (self *Task) PutMaxNumCon(maxnumcon int32) (err error) {
  if _tmp1101 := C.MSK_putmaxnumcon(self.ptr(),C.int32_t(maxnumcon)); _tmp1101 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1101)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the number of preallocated conic constraints in the optimization task.
//
// - maxnumcone int32
//   Number of preallocated conic constraints in the optimization task.
func (self *Task) PutMaxNumCone(maxnumcone int32) (err error) {
  if _tmp1102 := C.MSK_putmaxnumcone(self.ptr(),C.int32_t(maxnumcone)); _tmp1102 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1102)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the number of preallocated disjunctive constraints.
//
// - maxnumdjc int64
//   Number of preallocated disjunctive constraints in the task.
func (self *Task) PutMaxNumDjc(maxnumdjc int64) (err error) {
  if _tmp1103 := C.MSK_putmaxnumdjc(self.ptr(),C.int64_t(maxnumdjc)); _tmp1103 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1103)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the number of preallocated domains in the optimization task.
//
// - maxnumdomain int64
//   Number of preallocated domains.
func (self *Task) PutMaxNumDomain(maxnumdomain int64) (err error) {
  if _tmp1104 := C.MSK_putmaxnumdomain(self.ptr(),C.int64_t(maxnumdomain)); _tmp1104 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1104)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the number of preallocated non-zero entries in quadratic terms.
//
// - maxnumqnz int64
//   Number of non-zero elements preallocated in quadratic coefficient matrices.
func (self *Task) PutMaxNumQNz(maxnumqnz int64) (err error) {
  if _tmp1105 := C.MSK_putmaxnumqnz(self.ptr(),C.int64_t(maxnumqnz)); _tmp1105 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1105)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the number of preallocated variables in the optimization task.
//
// - maxnumvar int32
//   Number of preallocated variables in the optimization task.
func (self *Task) PutMaxNumVar(maxnumvar int32) (err error) {
  if _tmp1106 := C.MSK_putmaxnumvar(self.ptr(),C.int32_t(maxnumvar)); _tmp1106 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1106)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets a double parameter.
//
// - paramname string
//   Name of a parameter.
// - parvalue float64
//   Parameter value.
func (self *Task) PutNaDouParam(paramname string,parvalue float64) (err error) {
  _tmp1107 := C.CString(paramname)
  if _tmp1108 := C.MSK_putnadouparam(self.ptr(),_tmp1107,C.double(parvalue)); _tmp1108 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1108)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets an integer parameter.
//
// - paramname string
//   Name of a parameter.
// - parvalue int32
//   Parameter value.
func (self *Task) PutNaIntParam(paramname string,parvalue int32) (err error) {
  _tmp1109 := C.CString(paramname)
  if _tmp1110 := C.MSK_putnaintparam(self.ptr(),_tmp1109,C.int32_t(parvalue)); _tmp1110 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1110)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets a string parameter.
//
// - paramname string
//   Name of a parameter.
// - parvalue string
//   Parameter value.
func (self *Task) PutNaStrParam(paramname string,parvalue string) (err error) {
  _tmp1111 := C.CString(paramname)
  _tmp1112 := C.CString(parvalue)
  if _tmp1113 := C.MSK_putnastrparam(self.ptr(),_tmp1111,_tmp1112); _tmp1113 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1113)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Assigns a new name to the objective.
//
// - objname string
//   Name of the objective.
func (self *Task) PutObjName(objname string) (err error) {
  _tmp1114 := C.CString(objname)
  if _tmp1115 := C.MSK_putobjname(self.ptr(),_tmp1114); _tmp1115 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1115)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the objective sense.
//
// - sense Objsense
//   The objective sense of the task
func (self *Task) PutObjSense(sense Objsense) (err error) {
  if _tmp1116 := C.MSK_putobjsense(self.ptr(),C.int32_t(sense)); _tmp1116 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1116)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Specify an OptServer for remote calls.
//
// - host string
//   A URL specifying the optimization server to be used.
func (self *Task) PutOptserverHost(host string) (err error) {
  _tmp1117 := C.CString(host)
  if _tmp1118 := C.MSK_putoptserverhost(self.ptr(),_tmp1117); _tmp1118 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1118)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Modifies the value of parameter.
//
// - parname string
//   Parameter name.
// - parvalue string
//   Parameter value.
func (self *Task) PutParam(parname string,parvalue string) (err error) {
  _tmp1119 := C.CString(parname)
  _tmp1120 := C.CString(parvalue)
  if _tmp1121 := C.MSK_putparam(self.ptr(),_tmp1119,_tmp1120); _tmp1121 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1121)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces all quadratic terms in constraints.
//
// - qcsubk []int32
//   Constraint subscripts for quadratic coefficients.
// - qcsubi []int32
//   Row subscripts for quadratic constraint matrix.
// - qcsubj []int32
//   Column subscripts for quadratic constraint matrix.
// - qcval []float64
//   Quadratic constraint coefficient values.
func (self *Task) PutQCon(qcsubk []int32,qcsubi []int32,qcsubj []int32,qcval []float64) (err error) {
  _tmp1122 := len(qcval)
  if _tmp1122 < len(qcsubi) { _tmp1122 = len(qcsubi) }
  if _tmp1122 < len(qcsubj) { _tmp1122 = len(qcsubj) }
  var numqcnz int32 = int32(_tmp1122)
  var _tmp1123 *int32
  if len(qcsubk) > 0 { _tmp1123 = (*int32)(&qcsubk[0]) }
  var _tmp1124 *int32
  if len(qcsubi) > 0 { _tmp1124 = (*int32)(&qcsubi[0]) }
  var _tmp1125 *int32
  if len(qcsubj) > 0 { _tmp1125 = (*int32)(&qcsubj[0]) }
  var _tmp1126 *float64
  if len(qcval) > 0 { _tmp1126 = (*float64)(&qcval[0]) }
  if _tmp1127 := C.MSK_putqcon(self.ptr(),C.int32_t(numqcnz),(*C.int32_t)(_tmp1123),(*C.int32_t)(_tmp1124),(*C.int32_t)(_tmp1125),(*C.double)(_tmp1126)); _tmp1127 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1127)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces all quadratic terms in a single constraint.
//
// - k int32
//   The constraint in which the new quadratic elements are inserted.
// - qcsubi []int32
//   Row subscripts for quadratic constraint matrix.
// - qcsubj []int32
//   Column subscripts for quadratic constraint matrix.
// - qcval []float64
//   Quadratic constraint coefficient values.
func (self *Task) PutQConK(k int32,qcsubi []int32,qcsubj []int32,qcval []float64) (err error) {
  _tmp1128 := len(qcval)
  if _tmp1128 < len(qcsubi) { _tmp1128 = len(qcsubi) }
  if _tmp1128 < len(qcsubj) { _tmp1128 = len(qcsubj) }
  var numqcnz int32 = int32(_tmp1128)
  var _tmp1129 *int32
  if len(qcsubi) > 0 { _tmp1129 = (*int32)(&qcsubi[0]) }
  var _tmp1130 *int32
  if len(qcsubj) > 0 { _tmp1130 = (*int32)(&qcsubj[0]) }
  var _tmp1131 *float64
  if len(qcval) > 0 { _tmp1131 = (*float64)(&qcval[0]) }
  if _tmp1132 := C.MSK_putqconk(self.ptr(),C.int32_t(k),C.int32_t(numqcnz),(*C.int32_t)(_tmp1129),(*C.int32_t)(_tmp1130),(*C.double)(_tmp1131)); _tmp1132 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1132)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces all quadratic terms in the objective.
//
// - qosubi []int32
//   Row subscripts for quadratic objective coefficients.
// - qosubj []int32
//   Column subscripts for quadratic objective coefficients.
// - qoval []float64
//   Quadratic objective coefficient values.
func (self *Task) PutQObj(qosubi []int32,qosubj []int32,qoval []float64) (err error) {
  _tmp1133 := len(qosubi)
  if _tmp1133 < len(qoval) { _tmp1133 = len(qoval) }
  if _tmp1133 < len(qosubj) { _tmp1133 = len(qosubj) }
  var numqonz int32 = int32(_tmp1133)
  var _tmp1134 *int32
  if len(qosubi) > 0 { _tmp1134 = (*int32)(&qosubi[0]) }
  var _tmp1135 *int32
  if len(qosubj) > 0 { _tmp1135 = (*int32)(&qosubj[0]) }
  var _tmp1136 *float64
  if len(qoval) > 0 { _tmp1136 = (*float64)(&qoval[0]) }
  if _tmp1137 := C.MSK_putqobj(self.ptr(),C.int32_t(numqonz),(*C.int32_t)(_tmp1134),(*C.int32_t)(_tmp1135),(*C.double)(_tmp1136)); _tmp1137 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1137)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Replaces one coefficient in the quadratic term in the objective.
//
// - i int32
//   Row index for the coefficient to be replaced.
// - j int32
//   Column index for the coefficient to be replaced.
// - qoij float64
//   The new coefficient value.
func (self *Task) PutQObjIJ(i int32,j int32,qoij float64) (err error) {
  if _tmp1138 := C.MSK_putqobjij(self.ptr(),C.int32_t(i),C.int32_t(j),C.double(qoij)); _tmp1138 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1138)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the status keys for the constraints.
//
// - whichsol Soltype
//   Selects a solution.
// - skc []Stakey
//   Status keys for the constraints.
func (self *Task) PutSkc(whichsol Soltype,skc []Stakey) (err error) {
  var _tmp1139 C.int32_t
  if _tmp1140 := C.MSK_getnumcon(self.ptr(),(&_tmp1139)); _tmp1140 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1140)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(skc)) < int64(_tmp1139) {
    err = &ArrayLengthError{fun:"PutSkc",arg:"skc"}
    return
  }
  var _tmp1141 *Stakey
  if len(skc) > 0 { _tmp1141 = (*Stakey)(&skc[0]) }
  if _tmp1142 := C.MSK_putskc(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp1141)); _tmp1142 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1142)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the status keys for a slice of the constraints.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - skc []Stakey
//   Status keys for the constraints.
func (self *Task) PutSkcSlice(whichsol Soltype,first int32,last int32,skc []Stakey) (err error) {
  if skc != nil && int64(len(skc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSkcSlice",arg:"skc"}
    return
  }
  var _tmp1143 *Stakey
  if len(skc) > 0 { _tmp1143 = (*Stakey)(&skc[0]) }
  if _tmp1144 := C.MSK_putskcslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp1143)); _tmp1144 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1144)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the status keys for the scalar variables.
//
// - whichsol Soltype
//   Selects a solution.
// - skx []Stakey
//   Status keys for the variables.
func (self *Task) PutSkx(whichsol Soltype,skx []Stakey) (err error) {
  var _tmp1145 C.int32_t
  if _tmp1146 := C.MSK_getnumvar(self.ptr(),(&_tmp1145)); _tmp1146 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1146)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(skx)) < int64(_tmp1145) {
    err = &ArrayLengthError{fun:"PutSkx",arg:"skx"}
    return
  }
  var _tmp1147 *Stakey
  if len(skx) > 0 { _tmp1147 = (*Stakey)(&skx[0]) }
  if _tmp1148 := C.MSK_putskx(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp1147)); _tmp1148 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1148)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the status keys for a slice of the variables.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - skx []Stakey
//   Status keys for the variables.
func (self *Task) PutSkxSlice(whichsol Soltype,first int32,last int32,skx []Stakey) (err error) {
  if int64(len(skx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSkxSlice",arg:"skx"}
    return
  }
  var _tmp1149 *Stakey
  if len(skx) > 0 { _tmp1149 = (*Stakey)(&skx[0]) }
  if _tmp1150 := C.MSK_putskxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp1149)); _tmp1150 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1150)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the slc vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - slc []float64
//   Dual variables corresponding to the lower bounds on the constraints.
func (self *Task) PutSlc(whichsol Soltype,slc []float64) (err error) {
  var _tmp1151 C.int32_t
  if _tmp1152 := C.MSK_getnumcon(self.ptr(),(&_tmp1151)); _tmp1152 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1152)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(slc)) < int64(_tmp1151) {
    err = &ArrayLengthError{fun:"PutSlc",arg:"slc"}
    return
  }
  var _tmp1153 *float64
  if len(slc) > 0 { _tmp1153 = (*float64)(&slc[0]) }
  if _tmp1154 := C.MSK_putslc(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1153)); _tmp1154 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1154)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets a slice of the slc vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - slc []float64
//   Dual variables corresponding to the lower bounds on the constraints.
func (self *Task) PutSlcSlice(whichsol Soltype,first int32,last int32,slc []float64) (err error) {
  if int64(len(slc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSlcSlice",arg:"slc"}
    return
  }
  var _tmp1155 *float64
  if len(slc) > 0 { _tmp1155 = (*float64)(&slc[0]) }
  if _tmp1156 := C.MSK_putslcslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1155)); _tmp1156 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1156)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the slx vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - slx []float64
//   Dual variables corresponding to the lower bounds on the variables.
func (self *Task) PutSlx(whichsol Soltype,slx []float64) (err error) {
  var _tmp1157 C.int32_t
  if _tmp1158 := C.MSK_getnumvar(self.ptr(),(&_tmp1157)); _tmp1158 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1158)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(slx)) < int64(_tmp1157) {
    err = &ArrayLengthError{fun:"PutSlx",arg:"slx"}
    return
  }
  var _tmp1159 *float64
  if len(slx) > 0 { _tmp1159 = (*float64)(&slx[0]) }
  if _tmp1160 := C.MSK_putslx(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1159)); _tmp1160 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1160)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets a slice of the slx vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - slx []float64
//   Dual variables corresponding to the lower bounds on the variables.
func (self *Task) PutSlxSlice(whichsol Soltype,first int32,last int32,slx []float64) (err error) {
  if int64(len(slx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSlxSlice",arg:"slx"}
    return
  }
  var _tmp1161 *float64
  if len(slx) > 0 { _tmp1161 = (*float64)(&slx[0]) }
  if _tmp1162 := C.MSK_putslxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1161)); _tmp1162 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1162)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the snx vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - sux []float64
//   Dual variables corresponding to the upper bounds on the variables.
func (self *Task) PutSnx(whichsol Soltype,sux []float64) (err error) {
  var _tmp1163 C.int32_t
  if _tmp1164 := C.MSK_getnumvar(self.ptr(),(&_tmp1163)); _tmp1164 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1164)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(sux)) < int64(_tmp1163) {
    err = &ArrayLengthError{fun:"PutSnx",arg:"sux"}
    return
  }
  var _tmp1165 *float64
  if len(sux) > 0 { _tmp1165 = (*float64)(&sux[0]) }
  if _tmp1166 := C.MSK_putsnx(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1165)); _tmp1166 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1166)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets a slice of the snx vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - snx []float64
//   Dual variables corresponding to the conic constraints on the variables.
func (self *Task) PutSnxSlice(whichsol Soltype,first int32,last int32,snx []float64) (err error) {
  if int64(len(snx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSnxSlice",arg:"snx"}
    return
  }
  var _tmp1167 *float64
  if len(snx) > 0 { _tmp1167 = (*float64)(&snx[0]) }
  if _tmp1168 := C.MSK_putsnxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1167)); _tmp1168 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1168)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inserts a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - skc []Stakey
//   Status keys for the constraints.
// - skx []Stakey
//   Status keys for the variables.
// - skn []Stakey
//   Status keys for the conic constraints.
// - xc []float64
//   Primal constraint solution.
// - xx []float64
//   Primal variable solution.
// - y []float64
//   Vector of dual variables corresponding to the constraints.
// - slc []float64
//   Dual variables corresponding to the lower bounds on the constraints.
// - suc []float64
//   Dual variables corresponding to the upper bounds on the constraints.
// - slx []float64
//   Dual variables corresponding to the lower bounds on the variables.
// - sux []float64
//   Dual variables corresponding to the upper bounds on the variables.
// - snx []float64
//   Dual variables corresponding to the conic constraints on the variables.
func (self *Task) PutSolution(whichsol Soltype,skc []Stakey,skx []Stakey,skn []Stakey,xc []float64,xx []float64,y []float64,slc []float64,suc []float64,slx []float64,sux []float64,snx []float64) (err error) {
  var _tmp1169 *Stakey
  if len(skc) > 0 { _tmp1169 = (*Stakey)(&skc[0]) }
  var _tmp1170 *Stakey
  if len(skx) > 0 { _tmp1170 = (*Stakey)(&skx[0]) }
  var _tmp1171 *Stakey
  if len(skn) > 0 { _tmp1171 = (*Stakey)(&skn[0]) }
  var _tmp1172 *float64
  if len(xc) > 0 { _tmp1172 = (*float64)(&xc[0]) }
  var _tmp1173 *float64
  if len(xx) > 0 { _tmp1173 = (*float64)(&xx[0]) }
  var _tmp1174 *float64
  if len(y) > 0 { _tmp1174 = (*float64)(&y[0]) }
  var _tmp1175 *float64
  if len(slc) > 0 { _tmp1175 = (*float64)(&slc[0]) }
  var _tmp1176 *float64
  if len(suc) > 0 { _tmp1176 = (*float64)(&suc[0]) }
  var _tmp1177 *float64
  if len(slx) > 0 { _tmp1177 = (*float64)(&slx[0]) }
  var _tmp1178 *float64
  if len(sux) > 0 { _tmp1178 = (*float64)(&sux[0]) }
  var _tmp1179 *float64
  if len(snx) > 0 { _tmp1179 = (*float64)(&snx[0]) }
  if _tmp1180 := C.MSK_putsolution(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp1169),(*C.int32_t)(_tmp1170),(*C.int32_t)(_tmp1171),(*C.double)(_tmp1172),(*C.double)(_tmp1173),(*C.double)(_tmp1174),(*C.double)(_tmp1175),(*C.double)(_tmp1176),(*C.double)(_tmp1177),(*C.double)(_tmp1178),(*C.double)(_tmp1179)); _tmp1180 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1180)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inserts a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - skc []Stakey
//   Status keys for the constraints.
// - skx []Stakey
//   Status keys for the variables.
// - skn []Stakey
//   Status keys for the conic constraints.
// - xc []float64
//   Primal constraint solution.
// - xx []float64
//   Primal variable solution.
// - y []float64
//   Vector of dual variables corresponding to the constraints.
// - slc []float64
//   Dual variables corresponding to the lower bounds on the constraints.
// - suc []float64
//   Dual variables corresponding to the upper bounds on the constraints.
// - slx []float64
//   Dual variables corresponding to the lower bounds on the variables.
// - sux []float64
//   Dual variables corresponding to the upper bounds on the variables.
// - snx []float64
//   Dual variables corresponding to the conic constraints on the variables.
// - doty []float64
//   Dual variables corresponding to affine conic constraints.
func (self *Task) PutSolutionNew(whichsol Soltype,skc []Stakey,skx []Stakey,skn []Stakey,xc []float64,xx []float64,y []float64,slc []float64,suc []float64,slx []float64,sux []float64,snx []float64,doty []float64) (err error) {
  var _tmp1181 *Stakey
  if len(skc) > 0 { _tmp1181 = (*Stakey)(&skc[0]) }
  var _tmp1182 *Stakey
  if len(skx) > 0 { _tmp1182 = (*Stakey)(&skx[0]) }
  var _tmp1183 *Stakey
  if len(skn) > 0 { _tmp1183 = (*Stakey)(&skn[0]) }
  var _tmp1184 *float64
  if len(xc) > 0 { _tmp1184 = (*float64)(&xc[0]) }
  var _tmp1185 *float64
  if len(xx) > 0 { _tmp1185 = (*float64)(&xx[0]) }
  var _tmp1186 *float64
  if len(y) > 0 { _tmp1186 = (*float64)(&y[0]) }
  var _tmp1187 *float64
  if len(slc) > 0 { _tmp1187 = (*float64)(&slc[0]) }
  var _tmp1188 *float64
  if len(suc) > 0 { _tmp1188 = (*float64)(&suc[0]) }
  var _tmp1189 *float64
  if len(slx) > 0 { _tmp1189 = (*float64)(&slx[0]) }
  var _tmp1190 *float64
  if len(sux) > 0 { _tmp1190 = (*float64)(&sux[0]) }
  var _tmp1191 *float64
  if len(snx) > 0 { _tmp1191 = (*float64)(&snx[0]) }
  var _tmp1192 *float64
  if len(doty) > 0 { _tmp1192 = (*float64)(&doty[0]) }
  if _tmp1193 := C.MSK_putsolutionnew(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp1181),(*C.int32_t)(_tmp1182),(*C.int32_t)(_tmp1183),(*C.double)(_tmp1184),(*C.double)(_tmp1185),(*C.double)(_tmp1186),(*C.double)(_tmp1187),(*C.double)(_tmp1188),(*C.double)(_tmp1189),(*C.double)(_tmp1190),(*C.double)(_tmp1191),(*C.double)(_tmp1192)); _tmp1193 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1193)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Inputs the dual variable of a solution.
//
// - i int32
//   Index of the dual variable.
// - whichsol Soltype
//   Selects a solution.
// - y float64
//   Solution value of the dual variable.
func (self *Task) PutSolutionYI(i int32,whichsol Soltype,y float64) (err error) {
  if _tmp1194 := C.MSK_putsolutionyi(self.ptr(),C.int32_t(i),C.int32_t(whichsol),C.double(y)); _tmp1194 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1194)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets a string parameter.
//
// - param Sparam
//   Which parameter.
// - parvalue string
//   Parameter value.
func (self *Task) PutStrParam(param Sparam,parvalue string) (err error) {
  _tmp1195 := C.CString(parvalue)
  if _tmp1196 := C.MSK_putstrparam(self.ptr(),C.int32_t(param),_tmp1195); _tmp1196 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1196)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the suc vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - suc []float64
//   Dual variables corresponding to the upper bounds on the constraints.
func (self *Task) PutSuc(whichsol Soltype,suc []float64) (err error) {
  var _tmp1197 C.int32_t
  if _tmp1198 := C.MSK_getnumcon(self.ptr(),(&_tmp1197)); _tmp1198 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1198)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(suc)) < int64(_tmp1197) {
    err = &ArrayLengthError{fun:"PutSuc",arg:"suc"}
    return
  }
  var _tmp1199 *float64
  if len(suc) > 0 { _tmp1199 = (*float64)(&suc[0]) }
  if _tmp1200 := C.MSK_putsuc(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1199)); _tmp1200 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1200)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets a slice of the suc vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - suc []float64
//   Dual variables corresponding to the upper bounds on the constraints.
func (self *Task) PutSucSlice(whichsol Soltype,first int32,last int32,suc []float64) (err error) {
  if int64(len(suc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSucSlice",arg:"suc"}
    return
  }
  var _tmp1201 *float64
  if len(suc) > 0 { _tmp1201 = (*float64)(&suc[0]) }
  if _tmp1202 := C.MSK_putsucslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1201)); _tmp1202 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1202)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the sux vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - sux []float64
//   Dual variables corresponding to the upper bounds on the variables.
func (self *Task) PutSux(whichsol Soltype,sux []float64) (err error) {
  var _tmp1203 C.int32_t
  if _tmp1204 := C.MSK_getnumvar(self.ptr(),(&_tmp1203)); _tmp1204 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1204)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(sux)) < int64(_tmp1203) {
    err = &ArrayLengthError{fun:"PutSux",arg:"sux"}
    return
  }
  var _tmp1205 *float64
  if len(sux) > 0 { _tmp1205 = (*float64)(&sux[0]) }
  if _tmp1206 := C.MSK_putsux(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1205)); _tmp1206 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1206)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets a slice of the sux vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - sux []float64
//   Dual variables corresponding to the upper bounds on the variables.
func (self *Task) PutSuxSlice(whichsol Soltype,first int32,last int32,sux []float64) (err error) {
  if int64(len(sux)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSuxSlice",arg:"sux"}
    return
  }
  var _tmp1207 *float64
  if len(sux) > 0 { _tmp1207 = (*float64)(&sux[0]) }
  if _tmp1208 := C.MSK_putsuxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1207)); _tmp1208 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1208)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Assigns a new name to the task.
//
// - taskname string
//   Name assigned to the task.
func (self *Task) PutTaskName(taskname string) (err error) {
  _tmp1209 := C.CString(taskname)
  if _tmp1210 := C.MSK_puttaskname(self.ptr(),_tmp1209); _tmp1210 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1210)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes the bounds for one variable.
//
// - j int32
//   Index of the variable.
// - bkx Boundkey
//   New bound key.
// - blx float64
//   New lower bound.
// - bux float64
//   New upper bound.
func (self *Task) PutVarBound(j int32,bkx Boundkey,blx float64,bux float64) (err error) {
  if _tmp1211 := C.MSK_putvarbound(self.ptr(),C.int32_t(j),C.int32_t(bkx),C.double(blx),C.double(bux)); _tmp1211 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1211)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes the bounds of a list of variables.
//
// - sub []int32
//   List of variable indexes.
// - bkx []Boundkey
//   Bound keys for the variables.
// - blx []float64
//   Lower bounds for the variables.
// - bux []float64
//   Upper bounds for the variables.
func (self *Task) PutVarBoundList(sub []int32,bkx []Boundkey,blx []float64,bux []float64) (err error) {
  _tmp1212 := len(bkx)
  if _tmp1212 < len(blx) { _tmp1212 = len(blx) }
  if _tmp1212 < len(bux) { _tmp1212 = len(bux) }
  if _tmp1212 < len(sub) { _tmp1212 = len(sub) }
  var num int32 = int32(_tmp1212)
  var _tmp1213 *int32
  if len(sub) > 0 { _tmp1213 = (*int32)(&sub[0]) }
  var _tmp1214 *Boundkey
  if len(bkx) > 0 { _tmp1214 = (*Boundkey)(&bkx[0]) }
  var _tmp1215 *float64
  if len(blx) > 0 { _tmp1215 = (*float64)(&blx[0]) }
  var _tmp1216 *float64
  if len(bux) > 0 { _tmp1216 = (*float64)(&bux[0]) }
  if _tmp1217 := C.MSK_putvarboundlist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1213),(*C.int32_t)(_tmp1214),(*C.double)(_tmp1215),(*C.double)(_tmp1216)); _tmp1217 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1217)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes the bounds of a list of variables.
//
// - sub []int32
//   List of variable indexes.
// - bkx Boundkey
//   New bound key for all variables in the list.
// - blx float64
//   New lower bound for all variables in the list.
// - bux float64
//   New upper bound for all variables in the list.
func (self *Task) PutVarBoundListConst(sub []int32,bkx Boundkey,blx float64,bux float64) (err error) {
  _tmp1218 := len(sub)
  var num int32 = int32(_tmp1218)
  var _tmp1219 *int32
  if len(sub) > 0 { _tmp1219 = (*int32)(&sub[0]) }
  if _tmp1220 := C.MSK_putvarboundlistconst(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1219),C.int32_t(bkx),C.double(blx),C.double(bux)); _tmp1220 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1220)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes the bounds for a slice of the variables.
//
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - bkx []Boundkey
//   Bound keys for the variables.
// - blx []float64
//   Lower bounds for the variables.
// - bux []float64
//   Upper bounds for the variables.
func (self *Task) PutVarBoundSlice(first int32,last int32,bkx []Boundkey,blx []float64,bux []float64) (err error) {
  if int64(len(bkx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutVarBoundSlice",arg:"bkx"}
    return
  }
  var _tmp1221 *Boundkey
  if len(bkx) > 0 { _tmp1221 = (*Boundkey)(&bkx[0]) }
  if int64(len(blx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutVarBoundSlice",arg:"blx"}
    return
  }
  var _tmp1222 *float64
  if len(blx) > 0 { _tmp1222 = (*float64)(&blx[0]) }
  if int64(len(bux)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutVarBoundSlice",arg:"bux"}
    return
  }
  var _tmp1223 *float64
  if len(bux) > 0 { _tmp1223 = (*float64)(&bux[0]) }
  if _tmp1224 := C.MSK_putvarboundslice(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp1221),(*C.double)(_tmp1222),(*C.double)(_tmp1223)); _tmp1224 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1224)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Changes the bounds for a slice of the variables.
//
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - bkx Boundkey
//   New bound key for all variables in the slice.
// - blx float64
//   New lower bound for all variables in the slice.
// - bux float64
//   New upper bound for all variables in the slice.
func (self *Task) PutVarBoundSliceConst(first int32,last int32,bkx Boundkey,blx float64,bux float64) (err error) {
  if _tmp1225 := C.MSK_putvarboundsliceconst(self.ptr(),C.int32_t(first),C.int32_t(last),C.int32_t(bkx),C.double(blx),C.double(bux)); _tmp1225 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1225)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the name of a variable.
//
// - j int32
//   Index of the variable.
// - name string
//   The variable name.
func (self *Task) PutVarName(j int32,name string) (err error) {
  _tmp1226 := C.CString(name)
  if _tmp1227 := C.MSK_putvarname(self.ptr(),C.int32_t(j),_tmp1226); _tmp1227 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1227)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the primal and dual solution information for a single variable.
//
// - j int32
//   Index of the variable.
// - whichsol Soltype
//   Selects a solution.
// - sk Stakey
//   Status key of the variable.
// - x float64
//   Primal solution value of the variable.
// - sl float64
//   Solution value of the dual variable associated with the lower bound.
// - su float64
//   Solution value of the dual variable associated with the upper bound.
// - sn float64
//   Solution value of the dual variable associated with the conic constraint.
func (self *Task) PutVarSolutionJ(j int32,whichsol Soltype,sk Stakey,x float64,sl float64,su float64,sn float64) (err error) {
  if _tmp1228 := C.MSK_putvarsolutionj(self.ptr(),C.int32_t(j),C.int32_t(whichsol),C.int32_t(sk),C.double(x),C.double(sl),C.double(su),C.double(sn)); _tmp1228 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1228)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the variable type of one variable.
//
// - j int32
//   Index of the variable.
// - vartype Variabletype
//   The new variable type.
func (self *Task) PutVarType(j int32,vartype Variabletype) (err error) {
  if _tmp1229 := C.MSK_putvartype(self.ptr(),C.int32_t(j),C.int32_t(vartype)); _tmp1229 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1229)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the variable type for one or more variables.
//
// - subj []int32
//   A list of variable indexes for which the variable type should be changed.
// - vartype []Variabletype
//   A list of variable types.
func (self *Task) PutVarTypeList(subj []int32,vartype []Variabletype) (err error) {
  _tmp1230 := len(vartype)
  if _tmp1230 < len(subj) { _tmp1230 = len(subj) }
  var num int32 = int32(_tmp1230)
  var _tmp1231 *int32
  if len(subj) > 0 { _tmp1231 = (*int32)(&subj[0]) }
  var _tmp1232 *Variabletype
  if len(vartype) > 0 { _tmp1232 = (*Variabletype)(&vartype[0]) }
  if _tmp1233 := C.MSK_putvartypelist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1231),(*C.int32_t)(_tmp1232)); _tmp1233 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1233)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the xc vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - xc []float64
//   Primal constraint solution.
func (self *Task) PutXc(whichsol Soltype) (xc []float64,err error) {
  var _tmp1234 C.int32_t
  if _tmp1235 := C.MSK_getnumcon(self.ptr(),(&_tmp1234)); _tmp1235 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1235)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp1236 *float64
  xc = make([]float64,_tmp1234)
  if len(xc) > 0 { _tmp1236 = (*float64)(&xc[0]) }
  if _tmp1237 := C.MSK_putxc(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1236)); _tmp1237 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1237)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets a slice of the xc vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - xc []float64
//   Primal constraint solution.
func (self *Task) PutXcSlice(whichsol Soltype,first int32,last int32,xc []float64) (err error) {
  if int64(len(xc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutXcSlice",arg:"xc"}
    return
  }
  var _tmp1238 *float64
  if len(xc) > 0 { _tmp1238 = (*float64)(&xc[0]) }
  if _tmp1239 := C.MSK_putxcslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1238)); _tmp1239 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1239)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the xx vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - xx []float64
//   Primal variable solution.
func (self *Task) PutXx(whichsol Soltype,xx []float64) (err error) {
  var _tmp1240 C.int32_t
  if _tmp1241 := C.MSK_getnumvar(self.ptr(),(&_tmp1240)); _tmp1241 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1241)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(xx)) < int64(_tmp1240) {
    err = &ArrayLengthError{fun:"PutXx",arg:"xx"}
    return
  }
  var _tmp1242 *float64
  if len(xx) > 0 { _tmp1242 = (*float64)(&xx[0]) }
  if _tmp1243 := C.MSK_putxx(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1242)); _tmp1243 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1243)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets a slice of the xx vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - xx []float64
//   Primal variable solution.
func (self *Task) PutXxSlice(whichsol Soltype,first int32,last int32,xx []float64) (err error) {
  if int64(len(xx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutXxSlice",arg:"xx"}
    return
  }
  var _tmp1244 *float64
  if len(xx) > 0 { _tmp1244 = (*float64)(&xx[0]) }
  if _tmp1245 := C.MSK_putxxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1244)); _tmp1245 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1245)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets the y vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - y []float64
//   Vector of dual variables corresponding to the constraints.
func (self *Task) PutY(whichsol Soltype,y []float64) (err error) {
  var _tmp1246 C.int32_t
  if _tmp1247 := C.MSK_getnumcon(self.ptr(),(&_tmp1246)); _tmp1247 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1247)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(y)) < int64(_tmp1246) {
    err = &ArrayLengthError{fun:"PutY",arg:"y"}
    return
  }
  var _tmp1248 *float64
  if len(y) > 0 { _tmp1248 = (*float64)(&y[0]) }
  if _tmp1249 := C.MSK_puty(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1248)); _tmp1249 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1249)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Sets a slice of the y vector for a solution.
//
// - whichsol Soltype
//   Selects a solution.
// - first int32
//   First index in the sequence.
// - last int32
//   Last index plus 1 in the sequence.
// - y []float64
//   Vector of dual variables corresponding to the constraints.
func (self *Task) PutYSlice(whichsol Soltype,first int32,last int32,y []float64) (err error) {
  if int64(len(y)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutYSlice",arg:"y"}
    return
  }
  var _tmp1250 *float64
  if len(y) > 0 { _tmp1250 = (*float64)(&y[0]) }
  if _tmp1251 := C.MSK_putyslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1250)); _tmp1251 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1251)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Read a binary dump of the task solution and information items.
//
// - filename string
//   A valid file name.
// - compress Compresstype
//   Data compression type.
func (self *Task) ReadBSolution(filename string,compress Compresstype) (err error) {
  _tmp1252 := C.CString(filename)
  if _tmp1253 := C.MSK_readbsolution(self.ptr(),_tmp1252,C.int32_t(compress)); _tmp1253 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1253)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Reads problem data from a file.
//
// - filename string
//   A valid file name.
func (self *Task) ReadData(filename string) (err error) {
  _tmp1254 := C.CString(filename)
  if _tmp1255 := C.MSK_readdataautoformat(self.ptr(),_tmp1254); _tmp1255 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1255)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Reads problem data from a file.
//
// - filename string
//   A valid file name.
// - format Dataformat
//   File data format.
// - compress Compresstype
//   File compression type.
func (self *Task) ReadDataFormat(filename string,format Dataformat,compress Compresstype) (err error) {
  _tmp1256 := C.CString(filename)
  if _tmp1257 := C.MSK_readdataformat(self.ptr(),_tmp1256,C.int32_t(format),C.int32_t(compress)); _tmp1257 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1257)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Reads a solution from a JSOL file.
//
// - filename string
//   A valid file name.
func (self *Task) ReadJsonSol(filename string) (err error) {
  _tmp1258 := C.CString(filename)
  if _tmp1259 := C.MSK_readjsonsol(self.ptr(),_tmp1258); _tmp1259 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1259)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Load task data from a string in JSON format.
//
// - data string
//   Problem data in text format.
func (self *Task) ReadJsonString(data string) (err error) {
  _tmp1260 := C.CString(data)
  if _tmp1261 := C.MSK_readjsonstring(self.ptr(),_tmp1260); _tmp1261 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1261)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Load task data from a string in LP format.
//
// - data string
//   Problem data in text format.
func (self *Task) ReadLpString(data string) (err error) {
  _tmp1262 := C.CString(data)
  if _tmp1263 := C.MSK_readlpstring(self.ptr(),_tmp1262); _tmp1263 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1263)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Load task data from a string in OPF format.
//
// - data string
//   Problem data in text format.
func (self *Task) ReadOpfString(data string) (err error) {
  _tmp1264 := C.CString(data)
  if _tmp1265 := C.MSK_readopfstring(self.ptr(),_tmp1264); _tmp1265 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1265)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Reads a parameter file.
//
// - filename string
//   A valid file name.
func (self *Task) ReadParamFile(filename string) (err error) {
  _tmp1266 := C.CString(filename)
  if _tmp1267 := C.MSK_readparamfile(self.ptr(),_tmp1266); _tmp1267 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1267)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Load task data from a string in PTF format.
//
// - data string
//   Problem data in text format.
func (self *Task) ReadPtfString(data string) (err error) {
  _tmp1268 := C.CString(data)
  if _tmp1269 := C.MSK_readptfstring(self.ptr(),_tmp1268); _tmp1269 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1269)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Reads a solution from a file.
//
// - whichsol Soltype
//   Selects a solution.
// - filename string
//   A valid file name.
func (self *Task) ReadSolution(whichsol Soltype,filename string) (err error) {
  _tmp1270 := C.CString(filename)
  if _tmp1271 := C.MSK_readsolution(self.ptr(),C.int32_t(whichsol),_tmp1270); _tmp1271 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1271)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Read solution file in format determined by the filename
//
// - filename string
//   A valid file name.
func (self *Task) ReadSolutionFile(filename string) (err error) {
  _tmp1272 := C.CString(filename)
  if _tmp1273 := C.MSK_readsolutionfile(self.ptr(),_tmp1272); _tmp1273 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1273)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Prints information about last file read.
//
// - whichstream Streamtype
//   Index of the stream.
func (self *Task) ReadSummary(whichstream Streamtype) (err error) {
  if _tmp1274 := C.MSK_readsummary(self.ptr(),C.int32_t(whichstream)); _tmp1274 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1274)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Load task data from a file.
//
// - filename string
//   A valid file name.
func (self *Task) ReadTask(filename string) (err error) {
  _tmp1275 := C.CString(filename)
  if _tmp1276 := C.MSK_readtask(self.ptr(),_tmp1275); _tmp1276 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1276)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Removes a number of symmetric matrices.
//
// - subset []int32
//   Indexes of symmetric matrices which should be removed.
func (self *Task) RemoveBarvars(subset []int32) (err error) {
  _tmp1277 := len(subset)
  var num int32 = int32(_tmp1277)
  var _tmp1278 *int32
  if len(subset) > 0 { _tmp1278 = (*int32)(&subset[0]) }
  if _tmp1279 := C.MSK_removebarvars(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1278)); _tmp1279 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1279)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Removes a number of conic constraints from the problem.
//
// - subset []int32
//   Indexes of cones which should be removed.
func (self *Task) RemoveCones(subset []int32) (err error) {
  _tmp1280 := len(subset)
  var num int32 = int32(_tmp1280)
  var _tmp1281 *int32
  if len(subset) > 0 { _tmp1281 = (*int32)(&subset[0]) }
  if _tmp1282 := C.MSK_removecones(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1281)); _tmp1282 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1282)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Removes a number of constraints.
//
// - subset []int32
//   Indexes of constraints which should be removed.
func (self *Task) RemoveCons(subset []int32) (err error) {
  _tmp1283 := len(subset)
  var num int32 = int32(_tmp1283)
  var _tmp1284 *int32
  if len(subset) > 0 { _tmp1284 = (*int32)(&subset[0]) }
  if _tmp1285 := C.MSK_removecons(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1284)); _tmp1285 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1285)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Removes a number of variables.
//
// - subset []int32
//   Indexes of variables which should be removed.
func (self *Task) RemoveVars(subset []int32) (err error) {
  _tmp1286 := len(subset)
  var num int32 = int32(_tmp1286)
  var _tmp1287 *int32
  if len(subset) > 0 { _tmp1287 = (*int32)(&subset[0]) }
  if _tmp1288 := C.MSK_removevars(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1287)); _tmp1288 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1288)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Resizes an optimization task.
//
// - maxnumcon int32
//   New maximum number of constraints.
// - maxnumvar int32
//   New maximum number of variables.
// - maxnumcone int32
//   New maximum number of cones.
// - maxnumanz int64
//   New maximum number of linear non-zero elements.
// - maxnumqnz int64
//   New maximum number of quadratic non-zeros elements.
func (self *Task) ResizeTask(maxnumcon int32,maxnumvar int32,maxnumcone int32,maxnumanz int64,maxnumqnz int64) (err error) {
  if _tmp1289 := C.MSK_resizetask(self.ptr(),C.int32_t(maxnumcon),C.int32_t(maxnumvar),C.int32_t(maxnumcone),C.int64_t(maxnumanz),C.int64_t(maxnumqnz)); _tmp1289 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1289)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Creates a sensitivity report.
//
// - whichstream Streamtype
//   Index of the stream.
func (self *Task) SensitivityReport(whichstream Streamtype) (err error) {
  if _tmp1290 := C.MSK_sensitivityreport(self.ptr(),C.int32_t(whichstream)); _tmp1290 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1290)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Resets all parameter values.
func (self *Task) SetDefaults() (err error) {
  if _tmp1291 := C.MSK_setdefaults(self.ptr()); _tmp1291 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1291)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a status key abbreviation.
//
// - sk Stakey
//   A valid status key.
// - str string
//   Abbreviation string corresponding to the status key.
func (self *Task) SkToStr(sk Stakey) (str string,err error) {
  _tmp1292 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp1293 := C.MSK_sktostr(self.ptr(),C.int32_t(sk),(*C.uchar)(&_tmp1292[0])); _tmp1293 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1293)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp1292,byte(0)); p < 0 {
    str = string(_tmp1292)
  } else {
    str = string(_tmp1292[:p])
  }
  return
}

// Obtains a solution status string.
//
// - solutionsta Solsta
//   Solution status.
// - str string
//   String corresponding to the solution status.
func (self *Task) SolStaToStr(solutionsta Solsta) (str string,err error) {
  _tmp1294 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp1295 := C.MSK_solstatostr(self.ptr(),C.int32_t(solutionsta),(*C.uchar)(&_tmp1294[0])); _tmp1295 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1295)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp1294,byte(0)); p < 0 {
    str = string(_tmp1294)
  } else {
    str = string(_tmp1294[:p])
  }
  return
}

// Checks whether a solution is defined.
//
// - whichsol Soltype
//   Selects a solution.
// - isdef bool
//   Is non-zero if the requested solution is defined.
func (self *Task) SolutionDef(whichsol Soltype) (isdef bool,err error) {
  var _tmp1296 C.int
  if _tmp1297 := C.MSK_solutiondef(self.ptr(),C.int32_t(whichsol),(&_tmp1296)); _tmp1297 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1297)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  isdef = _tmp1296 != 0
  return
}

// Prints a short summary of the current solutions.
//
// - whichstream Streamtype
//   Index of the stream.
func (self *Task) SolutionSummary(whichstream Streamtype) (err error) {
  if _tmp1298 := C.MSK_solutionsummary(self.ptr(),C.int32_t(whichstream)); _tmp1298 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1298)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Solve a linear equation system involving a basis matrix.
//
// - transp bool
//   Controls which problem formulation is solved.
// - numnz int32
//   Input (number of non-zeros in right-hand side).
// - sub []int32
//   Input (indexes of non-zeros in right-hand side) and output (indexes of non-zeros in solution vector).
// - val []float64
//   Input (right-hand side values) and output (solution vector values).
// - numnzout int32
//   Output (number of non-zeros in solution vector).
func (self *Task) SolveWithBasis(transp bool,numnz int32,sub []int32,val []float64) (numnzout int32,err error) {
  var _tmp1299 C.int; if transp { _tmp1299 = 1; }
  var _tmp1300 C.int32_t
  if _tmp1301 := C.MSK_getnumcon(self.ptr(),(&_tmp1300)); _tmp1301 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1301)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(sub)) < int64(_tmp1300) {
    err = &ArrayLengthError{fun:"SolveWithBasis",arg:"sub"}
    return
  }
  var _tmp1302 *int32
  if len(sub) > 0 { _tmp1302 = (*int32)(&sub[0]) }
  var _tmp1303 C.int32_t
  if _tmp1304 := C.MSK_getnumcon(self.ptr(),(&_tmp1303)); _tmp1304 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1304)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(val)) < int64(_tmp1303) {
    err = &ArrayLengthError{fun:"SolveWithBasis",arg:"val"}
    return
  }
  var _tmp1305 *float64
  if len(val) > 0 { _tmp1305 = (*float64)(&val[0]) }
  if _tmp1306 := C.MSK_solvewithbasis(self.ptr(),_tmp1299,C.int32_t(numnz),(*C.int32_t)(_tmp1302),(*C.double)(_tmp1305),(*C.int32_t)(&numnzout)); _tmp1306 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1306)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a cone type code.
//
// - str string
//   String corresponding to the cone type code.
// - conetype Conetype
//   The cone type corresponding to str.
func (self *Task) StrToConeType(str string) (conetype Conetype,err error) {
  _tmp1307 := C.CString(str)
  var _tmp1308 int32
  if _tmp1309 := C.MSK_strtoconetype(self.ptr(),_tmp1307,(*C.int32_t)(&_tmp1308)); _tmp1309 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1309)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  conetype = Conetype(_tmp1308)
  return
}

// Obtains a status key.
//
// - str string
//   A status key abbreviation string.
// - sk Stakey
//   Status key corresponding to the string.
func (self *Task) StrToSk(str string) (sk Stakey,err error) {
  _tmp1310 := C.CString(str)
  var _tmp1311 int32
  if _tmp1312 := C.MSK_strtosk(self.ptr(),_tmp1310,(*C.int32_t)(&_tmp1311)); _tmp1312 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1312)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  sk = Stakey(_tmp1311)
  return
}

// In-place reformulation of a QCQO to a conic quadratic problem.
func (self *Task) Toconic() (err error) {
  if _tmp1313 := C.MSK_toconic(self.ptr()); _tmp1313 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1313)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Update the information items related to the solution.
//
// - whichsol Soltype
//   Selects a solution.
func (self *Task) UpdateSolutionInfo(whichsol Soltype) (err error) {
  if _tmp1314 := C.MSK_updatesolutioninfo(self.ptr(),C.int32_t(whichsol)); _tmp1314 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1314)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Write a binary dump of the task solution and information items.
//
// - filename string
//   A valid file name.
// - compress Compresstype
//   Data compression type.
func (self *Task) WriteBSolution(filename string,compress Compresstype) (err error) {
  _tmp1315 := C.CString(filename)
  if _tmp1316 := C.MSK_writebsolution(self.ptr(),_tmp1315,C.int32_t(compress)); _tmp1316 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1316)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Writes problem data to a file.
//
// - filename string
//   A valid file name.
func (self *Task) WriteData(filename string) (err error) {
  _tmp1317 := C.CString(filename)
  if _tmp1318 := C.MSK_writedata(self.ptr(),_tmp1317); _tmp1318 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1318)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Writes a solution to a JSON file.
//
// - filename string
//   A valid file name.
func (self *Task) WriteJsonSol(filename string) (err error) {
  _tmp1319 := C.CString(filename)
  if _tmp1320 := C.MSK_writejsonsol(self.ptr(),_tmp1319); _tmp1320 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1320)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Writes all the parameters to a parameter file.
//
// - filename string
//   A valid file name.
func (self *Task) WriteParamFile(filename string) (err error) {
  _tmp1321 := C.CString(filename)
  if _tmp1322 := C.MSK_writeparamfile(self.ptr(),_tmp1321); _tmp1322 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1322)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Write a solution to a file.
//
// - whichsol Soltype
//   Selects a solution.
// - filename string
//   A valid file name.
func (self *Task) WriteSolution(whichsol Soltype,filename string) (err error) {
  _tmp1323 := C.CString(filename)
  if _tmp1324 := C.MSK_writesolution(self.ptr(),C.int32_t(whichsol),_tmp1323); _tmp1324 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1324)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Write solution file in format determined by the filename
//
// - filename string
//   A valid file name.
func (self *Task) WriteSolutionFile(filename string) (err error) {
  _tmp1325 := C.CString(filename)
  if _tmp1326 := C.MSK_writesolutionfile(self.ptr(),_tmp1325); _tmp1326 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1326)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Write a complete binary dump of the task data.
//
// - filename string
//   A valid file name.
func (self *Task) WriteTask(filename string) (err error) {
  _tmp1328 := C.CString(filename)
  if _tmp1329 := C.MSK_writetask(self.ptr(),_tmp1328); _tmp1329 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1329)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes vector addition and multiplication by a scalar.
//
// - n int32
//   Length of the vectors.
// - alpha float64
//   The scalar that multiplies x.
// - x []float64
//   The x vector.
// - y []float64
//   The y vector.
func (self *Env) Axpy(n int32,alpha float64,x []float64,y []float64) (err error) {
  if int64(len(x)) < int64(n) {
    err = &ArrayLengthError{fun:"Axpy",arg:"x"}
    return
  }
  var _tmp1330 *float64
  if len(x) > 0 { _tmp1330 = (*float64)(&x[0]) }
  if int64(len(y)) < int64(n) {
    err = &ArrayLengthError{fun:"Axpy",arg:"y"}
    return
  }
  var _tmp1331 *float64
  if len(y) > 0 { _tmp1331 = (*float64)(&y[0]) }
  if _tmp1332 := C.MSK_axpy(self.ptr(),C.int32_t(n),C.double(alpha),(*C.double)(_tmp1330),(*C.double)(_tmp1331)); _tmp1332 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1332)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes vector addition and multiplication by a scalar.
//
// - n int32
//   Length of the vectors.
// - alpha float64
//   The scalar that multiplies x.
// - x []float64
//   The x vector.
// - y []float64
//   The y vector.
func Axpy(n int32,alpha float64,x []float64,y []float64) (err error) {
  self := &globalenv
  if int64(len(x)) < int64(n) {
    err = &ArrayLengthError{fun:"Axpy",arg:"x"}
    return
  }
  var _tmp1330 *float64
  if len(x) > 0 { _tmp1330 = (*float64)(&x[0]) }
  if int64(len(y)) < int64(n) {
    err = &ArrayLengthError{fun:"Axpy",arg:"y"}
    return
  }
  var _tmp1331 *float64
  if len(y) > 0 { _tmp1331 = (*float64)(&y[0]) }
  if _tmp1333 := C.MSK_axpy(self.ptr(),C.int32_t(n),C.double(alpha),(*C.double)(_tmp1330),(*C.double)(_tmp1331)); _tmp1333 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1333)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Check in all unused license features to the license token server.
func (self *Env) CheckInAll() (err error) {
  if _tmp1334 := C.MSK_checkinall(self.ptr()); _tmp1334 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1334)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Check in all unused license features to the license token server.
func CheckInAll() (err error) {
  self := &globalenv
  if _tmp1335 := C.MSK_checkinall(self.ptr()); _tmp1335 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1335)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Check in a license feature back to the license server ahead of time.
//
// - feature Feature
//   Feature to check in to the license system.
func (self *Env) CheckInLicense(feature Feature) (err error) {
  if _tmp1336 := C.MSK_checkinlicense(self.ptr(),C.int32_t(feature)); _tmp1336 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1336)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Check in a license feature back to the license server ahead of time.
//
// - feature Feature
//   Feature to check in to the license system.
func CheckInLicense(feature Feature) (err error) {
  self := &globalenv
  if _tmp1337 := C.MSK_checkinlicense(self.ptr(),C.int32_t(feature)); _tmp1337 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1337)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Check out a license feature from the license server ahead of time.
//
// - feature Feature
//   Feature to check out from the license system.
func (self *Env) CheckOutLicense(feature Feature) (err error) {
  if _tmp1338 := C.MSK_checkoutlicense(self.ptr(),C.int32_t(feature)); _tmp1338 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1338)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Check out a license feature from the license server ahead of time.
//
// - feature Feature
//   Feature to check out from the license system.
func CheckOutLicense(feature Feature) (err error) {
  self := &globalenv
  if _tmp1339 := C.MSK_checkoutlicense(self.ptr(),C.int32_t(feature)); _tmp1339 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1339)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the inner product of two vectors.
//
// - n int32
//   Length of the vectors.
// - x []float64
//   The x vector.
// - y []float64
//   The y vector.
// - xty float64
//   The result of the inner product.
func (self *Env) Dot(n int32,x []float64,y []float64) (xty float64,err error) {
  if int64(len(x)) < int64(n) {
    err = &ArrayLengthError{fun:"Dot",arg:"x"}
    return
  }
  var _tmp1340 *float64
  if len(x) > 0 { _tmp1340 = (*float64)(&x[0]) }
  if int64(len(y)) < int64(n) {
    err = &ArrayLengthError{fun:"Dot",arg:"y"}
    return
  }
  var _tmp1341 *float64
  if len(y) > 0 { _tmp1341 = (*float64)(&y[0]) }
  if _tmp1342 := C.MSK_dot(self.ptr(),C.int32_t(n),(*C.double)(_tmp1340),(*C.double)(_tmp1341),(*C.double)(&xty)); _tmp1342 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1342)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes the inner product of two vectors.
//
// - n int32
//   Length of the vectors.
// - x []float64
//   The x vector.
// - y []float64
//   The y vector.
// - xty float64
//   The result of the inner product.
func Dot(n int32,x []float64,y []float64) (xty float64,err error) {
  self := &globalenv
  if int64(len(x)) < int64(n) {
    err = &ArrayLengthError{fun:"Dot",arg:"x"}
    return
  }
  var _tmp1340 *float64
  if len(x) > 0 { _tmp1340 = (*float64)(&x[0]) }
  if int64(len(y)) < int64(n) {
    err = &ArrayLengthError{fun:"Dot",arg:"y"}
    return
  }
  var _tmp1341 *float64
  if len(y) > 0 { _tmp1341 = (*float64)(&y[0]) }
  if _tmp1343 := C.MSK_dot(self.ptr(),C.int32_t(n),(*C.double)(_tmp1340),(*C.double)(_tmp1341),(*C.double)(&xty)); _tmp1343 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1343)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Prints an intro to message stream.
//
// - longver int32
//   If non-zero, then the intro is slightly longer.
func (self *Env) EchoIntro(longver int32) (err error) {
  if _tmp1344 := C.MSK_echointro(self.ptr(),C.int32_t(longver)); _tmp1344 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1344)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Prints an intro to message stream.
//
// - longver int32
//   If non-zero, then the intro is slightly longer.
func EchoIntro(longver int32) (err error) {
  self := &globalenv
  if _tmp1345 := C.MSK_echointro(self.ptr(),C.int32_t(longver)); _tmp1345 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1345)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Reports when the first license feature expires.
//
// - expiry int64
//   If nonnegative, then it is the minimum number days to expiry of any feature that has been checked out.
func (self *Env) Expirylicenses() (expiry int64,err error) {
  if _tmp1346 := C.MSK_expirylicenses(self.ptr(),(*C.int64_t)(&expiry)); _tmp1346 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1346)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Reports when the first license feature expires.
//
// - expiry int64
//   If nonnegative, then it is the minimum number days to expiry of any feature that has been checked out.
func Expirylicenses() (expiry int64,err error) {
  self := &globalenv
  if _tmp1347 := C.MSK_expirylicenses(self.ptr(),(*C.int64_t)(&expiry)); _tmp1347 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1347)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Performs a dense matrix multiplication.
//
// - transa Transpose
//   Indicates whether the matrix A must be transposed.
// - transb Transpose
//   Indicates whether the matrix B must be transposed.
// - m int32
//   Indicates the number of rows of matrix C.
// - n int32
//   Indicates the number of columns of matrix C.
// - k int32
//   Specifies the common dimension along which op(A) and op(B) are multiplied.
// - alpha float64
//   A scalar value multiplying the result of the matrix multiplication.
// - a []float64
//   The pointer to the array storing matrix A in a column-major format.
// - b []float64
//   The pointer to the array storing matrix B in a column-major format.
// - beta float64
//   A scalar value that multiplies C.
// - c []float64
//   The pointer to the array storing matrix C in a column-major format.
func (self *Env) Gemm(transa Transpose,transb Transpose,m int32,n int32,k int32,alpha float64,a []float64,b []float64,beta float64,c []float64) (err error) {
  if int64(len(a)) < int64((m * k)) {
    err = &ArrayLengthError{fun:"Gemm",arg:"a"}
    return
  }
  var _tmp1348 *float64
  if len(a) > 0 { _tmp1348 = (*float64)(&a[0]) }
  if int64(len(b)) < int64((k * n)) {
    err = &ArrayLengthError{fun:"Gemm",arg:"b"}
    return
  }
  var _tmp1349 *float64
  if len(b) > 0 { _tmp1349 = (*float64)(&b[0]) }
  if int64(len(c)) < int64((m * n)) {
    err = &ArrayLengthError{fun:"Gemm",arg:"c"}
    return
  }
  var _tmp1350 *float64
  if len(c) > 0 { _tmp1350 = (*float64)(&c[0]) }
  if _tmp1351 := C.MSK_gemm(self.ptr(),C.int32_t(transa),C.int32_t(transb),C.int32_t(m),C.int32_t(n),C.int32_t(k),C.double(alpha),(*C.double)(_tmp1348),(*C.double)(_tmp1349),C.double(beta),(*C.double)(_tmp1350)); _tmp1351 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1351)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Performs a dense matrix multiplication.
//
// - transa Transpose
//   Indicates whether the matrix A must be transposed.
// - transb Transpose
//   Indicates whether the matrix B must be transposed.
// - m int32
//   Indicates the number of rows of matrix C.
// - n int32
//   Indicates the number of columns of matrix C.
// - k int32
//   Specifies the common dimension along which op(A) and op(B) are multiplied.
// - alpha float64
//   A scalar value multiplying the result of the matrix multiplication.
// - a []float64
//   The pointer to the array storing matrix A in a column-major format.
// - b []float64
//   The pointer to the array storing matrix B in a column-major format.
// - beta float64
//   A scalar value that multiplies C.
// - c []float64
//   The pointer to the array storing matrix C in a column-major format.
func Gemm(transa Transpose,transb Transpose,m int32,n int32,k int32,alpha float64,a []float64,b []float64,beta float64,c []float64) (err error) {
  self := &globalenv
  if int64(len(a)) < int64((m * k)) {
    err = &ArrayLengthError{fun:"Gemm",arg:"a"}
    return
  }
  var _tmp1348 *float64
  if len(a) > 0 { _tmp1348 = (*float64)(&a[0]) }
  if int64(len(b)) < int64((k * n)) {
    err = &ArrayLengthError{fun:"Gemm",arg:"b"}
    return
  }
  var _tmp1349 *float64
  if len(b) > 0 { _tmp1349 = (*float64)(&b[0]) }
  if int64(len(c)) < int64((m * n)) {
    err = &ArrayLengthError{fun:"Gemm",arg:"c"}
    return
  }
  var _tmp1350 *float64
  if len(c) > 0 { _tmp1350 = (*float64)(&c[0]) }
  if _tmp1352 := C.MSK_gemm(self.ptr(),C.int32_t(transa),C.int32_t(transb),C.int32_t(m),C.int32_t(n),C.int32_t(k),C.double(alpha),(*C.double)(_tmp1348),(*C.double)(_tmp1349),C.double(beta),(*C.double)(_tmp1350)); _tmp1352 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1352)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes dense matrix times a dense vector product.
//
// - transa Transpose
//   Indicates whether the matrix A must be transposed.
// - m int32
//   Specifies the number of rows of the matrix A.
// - n int32
//   Specifies the number of columns of the matrix A.
// - alpha float64
//   A scalar value multiplying the matrix A.
// - a []float64
//   A pointer to the array storing matrix A in a column-major format.
// - x []float64
//   A pointer to the array storing the vector x.
// - beta float64
//   A scalar value multiplying the vector y.
// - y []float64
//   A pointer to the array storing the vector y.
func (self *Env) Gemv(transa Transpose,m int32,n int32,alpha float64,a []float64,x []float64,beta float64,y []float64) (err error) {
  if int64(len(a)) < int64((n * m)) {
    err = &ArrayLengthError{fun:"Gemv",arg:"a"}
    return
  }
  var _tmp1353 *float64
  if len(a) > 0 { _tmp1353 = (*float64)(&a[0]) }
  var _tmp1354 C.int32_t
  if (transa == MSK_TRANSPOSE_NO) {
    _tmp1354 = (C.int32_t)(n)
  } else {
    _tmp1354 = (C.int32_t)(m)
  }
  if int64(len(x)) < int64(_tmp1354) {
    err = &ArrayLengthError{fun:"Gemv",arg:"x"}
    return
  }
  var _tmp1355 *float64
  if len(x) > 0 { _tmp1355 = (*float64)(&x[0]) }
  var _tmp1356 C.int32_t
  if (transa == MSK_TRANSPOSE_NO) {
    _tmp1356 = (C.int32_t)(m)
  } else {
    _tmp1356 = (C.int32_t)(n)
  }
  if int64(len(y)) < int64(_tmp1356) {
    err = &ArrayLengthError{fun:"Gemv",arg:"y"}
    return
  }
  var _tmp1357 *float64
  if len(y) > 0 { _tmp1357 = (*float64)(&y[0]) }
  if _tmp1358 := C.MSK_gemv(self.ptr(),C.int32_t(transa),C.int32_t(m),C.int32_t(n),C.double(alpha),(*C.double)(_tmp1353),(*C.double)(_tmp1355),C.double(beta),(*C.double)(_tmp1357)); _tmp1358 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1358)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes dense matrix times a dense vector product.
//
// - transa Transpose
//   Indicates whether the matrix A must be transposed.
// - m int32
//   Specifies the number of rows of the matrix A.
// - n int32
//   Specifies the number of columns of the matrix A.
// - alpha float64
//   A scalar value multiplying the matrix A.
// - a []float64
//   A pointer to the array storing matrix A in a column-major format.
// - x []float64
//   A pointer to the array storing the vector x.
// - beta float64
//   A scalar value multiplying the vector y.
// - y []float64
//   A pointer to the array storing the vector y.
func Gemv(transa Transpose,m int32,n int32,alpha float64,a []float64,x []float64,beta float64,y []float64) (err error) {
  self := &globalenv
  if int64(len(a)) < int64((n * m)) {
    err = &ArrayLengthError{fun:"Gemv",arg:"a"}
    return
  }
  var _tmp1353 *float64
  if len(a) > 0 { _tmp1353 = (*float64)(&a[0]) }
  var _tmp1354 C.int32_t
  if (transa == MSK_TRANSPOSE_NO) {
    _tmp1354 = (C.int32_t)(n)
  } else {
    _tmp1354 = (C.int32_t)(m)
  }
  if int64(len(x)) < int64(_tmp1354) {
    err = &ArrayLengthError{fun:"Gemv",arg:"x"}
    return
  }
  var _tmp1355 *float64
  if len(x) > 0 { _tmp1355 = (*float64)(&x[0]) }
  var _tmp1356 C.int32_t
  if (transa == MSK_TRANSPOSE_NO) {
    _tmp1356 = (C.int32_t)(m)
  } else {
    _tmp1356 = (C.int32_t)(n)
  }
  if int64(len(y)) < int64(_tmp1356) {
    err = &ArrayLengthError{fun:"Gemv",arg:"y"}
    return
  }
  var _tmp1357 *float64
  if len(y) > 0 { _tmp1357 = (*float64)(&y[0]) }
  if _tmp1359 := C.MSK_gemv(self.ptr(),C.int32_t(transa),C.int32_t(m),C.int32_t(n),C.double(alpha),(*C.double)(_tmp1353),(*C.double)(_tmp1355),C.double(beta),(*C.double)(_tmp1357)); _tmp1359 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1359)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Obtains a short description of a response code.
//
// - code Rescode
//   A valid response code.
// - symname string
//   Symbolic name corresponding to the code.
// - str string
//   Obtains a short description of a response code.
func GetCodeDesc(code Rescode) (symname string,str string,err error) {
  _tmp1360 := make([]byte,MSK_MAX_STR_LEN)
  _tmp1361 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp1362 := C.MSK_getcodedesc(C.int32_t(code),(*C.uchar)(&_tmp1360[0]),(*C.uchar)(&_tmp1361[0])); _tmp1362 != 0 {
    err = &MosekError{ code:Rescode(_tmp1362) }
    return
  }
  if p := bytes.IndexByte(_tmp1360,byte(0)); p < 0 {
    symname = string(_tmp1360)
  } else {
    symname = string(_tmp1360[:p])
  }
  if p := bytes.IndexByte(_tmp1361,byte(0)); p < 0 {
    str = string(_tmp1361)
  } else {
    str = string(_tmp1361[:p])
  }
  return
}

// Obtains MOSEK version information.
//
// - major int32
//   Major version number.
// - minor int32
//   Minor version number.
// - revision int32
//   Revision number.
func GetVersion() (major int32,minor int32,revision int32,err error) {
  if _tmp1363 := C.MSK_getversion((*C.int32_t)(&major),(*C.int32_t)(&minor),(*C.int32_t)(&revision)); _tmp1363 != 0 {
    err = &MosekError{ code:Rescode(_tmp1363) }
    return
  }
  return
}

// Stops all threads and delete all handles used by the license system.
func LicenseCleanup() (err error) {
  if _tmp1364 := C.MSK_licensecleanup(); _tmp1364 != 0 {
    err = &MosekError{ code:Rescode(_tmp1364) }
    return
  }
  return
}

// Directs all output from a stream to a file.
//
// - whichstream Streamtype
//   Index of the stream.
// - filename string
//   A valid file name.
// - append int32
//   If this argument is 0 the file will be overwritten, otherwise it will be appended to.
func (self *Env) Linkfiletostream(whichstream Streamtype,filename string,append int32) (err error) {
  _tmp1365 := C.CString(filename)
  if _tmp1366 := C.MSK_linkfiletoenvstream(self.ptr(),C.int32_t(whichstream),_tmp1365,C.int32_t(append)); _tmp1366 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1366)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Directs all output from a stream to a file.
//
// - whichstream Streamtype
//   Index of the stream.
// - filename string
//   A valid file name.
// - append int32
//   If this argument is 0 the file will be overwritten, otherwise it will be appended to.
func Linkfiletostream(whichstream Streamtype,filename string,append int32) (err error) {
  self := &globalenv
  _tmp1365 := C.CString(filename)
  if _tmp1367 := C.MSK_linkfiletoenvstream(self.ptr(),C.int32_t(whichstream),_tmp1365,C.int32_t(append)); _tmp1367 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1367)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes a Cholesky factorization of a dense matrix.
//
// - uplo Uplo
//   Indicates whether the upper or lower triangular part of the matrix is stored.
// - n int32
//   Dimension of the symmetric matrix.
// - a []float64
//   A symmetric matrix stored in column-major order.
func (self *Env) Potrf(uplo Uplo,n int32,a []float64) (err error) {
  if int64(len(a)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Potrf",arg:"a"}
    return
  }
  var _tmp1368 *float64
  if len(a) > 0 { _tmp1368 = (*float64)(&a[0]) }
  if _tmp1369 := C.MSK_potrf(self.ptr(),C.int32_t(uplo),C.int32_t(n),(*C.double)(_tmp1368)); _tmp1369 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1369)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes a Cholesky factorization of a dense matrix.
//
// - uplo Uplo
//   Indicates whether the upper or lower triangular part of the matrix is stored.
// - n int32
//   Dimension of the symmetric matrix.
// - a []float64
//   A symmetric matrix stored in column-major order.
func Potrf(uplo Uplo,n int32,a []float64) (err error) {
  self := &globalenv
  if int64(len(a)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Potrf",arg:"a"}
    return
  }
  var _tmp1368 *float64
  if len(a) > 0 { _tmp1368 = (*float64)(&a[0]) }
  if _tmp1370 := C.MSK_potrf(self.ptr(),C.int32_t(uplo),C.int32_t(n),(*C.double)(_tmp1368)); _tmp1370 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1370)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Input a runtime license code.
//
// - code []int32
//   A license key string.
func (self *Env) PutLicenseCode(code []int32) (err error) {
  if code != nil && int64(len(code)) < int64(MSK_LICENSE_BUFFER_LENGTH) {
    err = &ArrayLengthError{fun:"PutLicenseCode",arg:"code"}
    return
  }
  var _tmp1371 *int32
  if len(code) > 0 { _tmp1371 = (*int32)(&code[0]) }
  if _tmp1372 := C.MSK_putlicensecode(self.ptr(),(*C.int32_t)(_tmp1371)); _tmp1372 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1372)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Input a runtime license code.
//
// - code []int32
//   A license key string.
func PutLicenseCode(code []int32) (err error) {
  self := &globalenv
  if code != nil && int64(len(code)) < int64(MSK_LICENSE_BUFFER_LENGTH) {
    err = &ArrayLengthError{fun:"PutLicenseCode",arg:"code"}
    return
  }
  var _tmp1371 *int32
  if len(code) > 0 { _tmp1371 = (*int32)(&code[0]) }
  if _tmp1373 := C.MSK_putlicensecode(self.ptr(),(*C.int32_t)(_tmp1371)); _tmp1373 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1373)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Enables debug information for the license system.
//
// - licdebug int32
//   Enable output of license check-out debug information.
func (self *Env) PutLicenseDebug(licdebug int32) (err error) {
  if _tmp1374 := C.MSK_putlicensedebug(self.ptr(),C.int32_t(licdebug)); _tmp1374 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1374)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Enables debug information for the license system.
//
// - licdebug int32
//   Enable output of license check-out debug information.
func PutLicenseDebug(licdebug int32) (err error) {
  self := &globalenv
  if _tmp1375 := C.MSK_putlicensedebug(self.ptr(),C.int32_t(licdebug)); _tmp1375 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1375)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Set the path to the license file.
//
// - licensepath string
//   A path specifying where to search for the license.
func (self *Env) PutLicensePath(licensepath string) (err error) {
  _tmp1376 := C.CString(licensepath)
  if _tmp1377 := C.MSK_putlicensepath(self.ptr(),_tmp1376); _tmp1377 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1377)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Set the path to the license file.
//
// - licensepath string
//   A path specifying where to search for the license.
func PutLicensePath(licensepath string) (err error) {
  self := &globalenv
  _tmp1376 := C.CString(licensepath)
  if _tmp1378 := C.MSK_putlicensepath(self.ptr(),_tmp1376); _tmp1378 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1378)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Control whether mosek should wait for an available license if no license is available.
//
// - licwait int32
//   Enable waiting for a license.
func (self *Env) PutLicenseWait(licwait int32) (err error) {
  if _tmp1379 := C.MSK_putlicensewait(self.ptr(),C.int32_t(licwait)); _tmp1379 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1379)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Control whether mosek should wait for an available license if no license is available.
//
// - licwait int32
//   Enable waiting for a license.
func PutLicenseWait(licwait int32) (err error) {
  self := &globalenv
  if _tmp1380 := C.MSK_putlicensewait(self.ptr(),C.int32_t(licwait)); _tmp1380 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1380)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Reset the license expiry reporting startpoint.
func (self *Env) ResetExpiryLicenses() (err error) {
  if _tmp1381 := C.MSK_resetexpirylicenses(self.ptr()); _tmp1381 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1381)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Reset the license expiry reporting startpoint.
func ResetExpiryLicenses() (err error) {
  self := &globalenv
  if _tmp1382 := C.MSK_resetexpirylicenses(self.ptr()); _tmp1382 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1382)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes all eigenvalues of a symmetric dense matrix.
//
// - uplo Uplo
//   Indicates whether the upper or lower triangular part is used.
// - n int32
//   Dimension of the symmetric input matrix.
// - a []float64
//   Input matrix A.
// - w []float64
//   Array of length at least n containing the eigenvalues of A.
func (self *Env) Syeig(uplo Uplo,n int32,a []float64) (w []float64,err error) {
  if int64(len(a)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Syeig",arg:"a"}
    return
  }
  var _tmp1383 *float64
  if len(a) > 0 { _tmp1383 = (*float64)(&a[0]) }
  var _tmp1384 *float64
  w = make([]float64,n)
  if len(w) > 0 { _tmp1384 = (*float64)(&w[0]) }
  if _tmp1385 := C.MSK_syeig(self.ptr(),C.int32_t(uplo),C.int32_t(n),(*C.double)(_tmp1383),(*C.double)(_tmp1384)); _tmp1385 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1385)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes all eigenvalues of a symmetric dense matrix.
//
// - uplo Uplo
//   Indicates whether the upper or lower triangular part is used.
// - n int32
//   Dimension of the symmetric input matrix.
// - a []float64
//   Input matrix A.
// - w []float64
//   Array of length at least n containing the eigenvalues of A.
func Syeig(uplo Uplo,n int32,a []float64) (w []float64,err error) {
  self := &globalenv
  if int64(len(a)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Syeig",arg:"a"}
    return
  }
  var _tmp1383 *float64
  if len(a) > 0 { _tmp1383 = (*float64)(&a[0]) }
  var _tmp1384 *float64
  w = make([]float64,n)
  if len(w) > 0 { _tmp1384 = (*float64)(&w[0]) }
  if _tmp1386 := C.MSK_syeig(self.ptr(),C.int32_t(uplo),C.int32_t(n),(*C.double)(_tmp1383),(*C.double)(_tmp1384)); _tmp1386 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1386)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes all the eigenvalues and eigenvectors of a symmetric dense matrix, and thus its eigenvalue decomposition.
//
// - uplo Uplo
//   Indicates whether the upper or lower triangular part is used.
// - n int32
//   Dimension of the symmetric input matrix.
// - a []float64
//   Input matrix A.
// - w []float64
//   Array of length at least n containing the eigenvalues of A.
func (self *Env) Syevd(uplo Uplo,n int32,a []float64) (w []float64,err error) {
  if int64(len(a)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Syevd",arg:"a"}
    return
  }
  var _tmp1387 *float64
  if len(a) > 0 { _tmp1387 = (*float64)(&a[0]) }
  var _tmp1388 *float64
  w = make([]float64,n)
  if len(w) > 0 { _tmp1388 = (*float64)(&w[0]) }
  if _tmp1389 := C.MSK_syevd(self.ptr(),C.int32_t(uplo),C.int32_t(n),(*C.double)(_tmp1387),(*C.double)(_tmp1388)); _tmp1389 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1389)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Computes all the eigenvalues and eigenvectors of a symmetric dense matrix, and thus its eigenvalue decomposition.
//
// - uplo Uplo
//   Indicates whether the upper or lower triangular part is used.
// - n int32
//   Dimension of the symmetric input matrix.
// - a []float64
//   Input matrix A.
// - w []float64
//   Array of length at least n containing the eigenvalues of A.
func Syevd(uplo Uplo,n int32,a []float64) (w []float64,err error) {
  self := &globalenv
  if int64(len(a)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Syevd",arg:"a"}
    return
  }
  var _tmp1387 *float64
  if len(a) > 0 { _tmp1387 = (*float64)(&a[0]) }
  var _tmp1388 *float64
  w = make([]float64,n)
  if len(w) > 0 { _tmp1388 = (*float64)(&w[0]) }
  if _tmp1390 := C.MSK_syevd(self.ptr(),C.int32_t(uplo),C.int32_t(n),(*C.double)(_tmp1387),(*C.double)(_tmp1388)); _tmp1390 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1390)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Performs a rank-k update of a symmetric matrix.
//
// - uplo Uplo
//   Indicates whether the upper or lower triangular part of C is used.
// - trans Transpose
//   Indicates whether the matrix A must be transposed.
// - n int32
//   Specifies the order of C.
// - k int32
//   Indicates the number of rows or columns of A, and its rank.
// - alpha float64
//   A scalar value multiplying the result of the matrix multiplication.
// - a []float64
//   The pointer to the array storing matrix A in a column-major format.
// - beta float64
//   A scalar value that multiplies C.
// - c []float64
//   The pointer to the array storing matrix C in a column-major format.
func (self *Env) Syrk(uplo Uplo,trans Transpose,n int32,k int32,alpha float64,a []float64,beta float64,c []float64) (err error) {
  if int64(len(a)) < int64((n * k)) {
    err = &ArrayLengthError{fun:"Syrk",arg:"a"}
    return
  }
  var _tmp1391 *float64
  if len(a) > 0 { _tmp1391 = (*float64)(&a[0]) }
  if int64(len(c)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Syrk",arg:"c"}
    return
  }
  var _tmp1392 *float64
  if len(c) > 0 { _tmp1392 = (*float64)(&c[0]) }
  if _tmp1393 := C.MSK_syrk(self.ptr(),C.int32_t(uplo),C.int32_t(trans),C.int32_t(n),C.int32_t(k),C.double(alpha),(*C.double)(_tmp1391),C.double(beta),(*C.double)(_tmp1392)); _tmp1393 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1393)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

// Performs a rank-k update of a symmetric matrix.
//
// - uplo Uplo
//   Indicates whether the upper or lower triangular part of C is used.
// - trans Transpose
//   Indicates whether the matrix A must be transposed.
// - n int32
//   Specifies the order of C.
// - k int32
//   Indicates the number of rows or columns of A, and its rank.
// - alpha float64
//   A scalar value multiplying the result of the matrix multiplication.
// - a []float64
//   The pointer to the array storing matrix A in a column-major format.
// - beta float64
//   A scalar value that multiplies C.
// - c []float64
//   The pointer to the array storing matrix C in a column-major format.
func Syrk(uplo Uplo,trans Transpose,n int32,k int32,alpha float64,a []float64,beta float64,c []float64) (err error) {
  self := &globalenv
  if int64(len(a)) < int64((n * k)) {
    err = &ArrayLengthError{fun:"Syrk",arg:"a"}
    return
  }
  var _tmp1391 *float64
  if len(a) > 0 { _tmp1391 = (*float64)(&a[0]) }
  if int64(len(c)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Syrk",arg:"c"}
    return
  }
  var _tmp1392 *float64
  if len(c) > 0 { _tmp1392 = (*float64)(&c[0]) }
  if _tmp1394 := C.MSK_syrk(self.ptr(),C.int32_t(uplo),C.int32_t(trans),C.int32_t(n),C.int32_t(k),C.double(alpha),(*C.double)(_tmp1391),C.double(beta),(*C.double)(_tmp1392)); _tmp1394 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1394)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

