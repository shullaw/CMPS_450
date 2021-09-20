#author: Andrew Shullaw
#cid: c00161818

# Your programs should accept one line of input on file front.in that resides in the same
# directory as the source code. It should give the same output that the C++ does on the given input.
import sys

token = {'int_lit' :    {'code':10,
                         'char':''},
          'ident' :     {'code':11,
                         'char':''}, 
          'assign_op' : {'code':20,
                         'char':'='},
          'add_op' :    {'code':21,
                         'char':'+'},
          'sub_op' :    {'code':22,
                         'char':'-'},
          'mult_op' :   {'code':23,
                         'char':'*'},
          'div_op' :    {'code':24,
                         'char':'/'},
          'left_paren' : {'code':25,
                          'char':'('},
          'right_paren': {'code':26,
                          'char':')'}}
char_class = {'letter':False,
                'digit':False,
                'unknown':False}

def getchar(x):
    if (x.isalpha()):
        char_class['letter']=True
    elif (x.isdigit()):
        char_class['digit']=True
    else:
        char_class['unknown']=True
def lex():
# Your programs should accept one line of input on file front.in that resides in the same
# directory as the source code. It should give the same output that the C++ does on the given input.
    try:
        cwd = sys.path[0]
        with open("front.in", 'r') as file:
            print("Current working dir: {}".format(cwd),"\n")
            lexemes=[]
            current_char=''
            i=0
            for line in file:
                for x in line:
                    if(x != ' '):
                        lexemes+=x
                total = len(lexemes) - 1
                while(i < total+1):
                        getchar(lexemes[i])
                        current_char=''
                        if (char_class['letter']):
                            while (char_class['letter'] or char_class['digit'] and i < total):
                                char_class['letter']=False
                                char_class['digit']=False
                                nextToken = token['ident']['code']
                                current_char+=lexemes[i]
                                i+=1
                                if (i < total):
                                    getchar(lexemes[i])
                        elif (char_class['digit'] and i < total):
                           while(char_class['digit']):
                                char_class['digit']=False
                                nextToken=token['int_lit']['code']
                                current_char+=lexemes[i]
                                i+=1
                                if (i < len(lexemes)):
                                    getchar(lexemes[i])
                        elif (char_class['unknown'] and i < total):
                            char_class['unknown']=False
                            current_char=lexemes[i]
                            nextToken = [j['code'] for j in token.values() if j['char']==current_char][0]
                            i+=1
                        elif (i == total):
                            char_class['letter']=False
                            char_class['digit']=False
                            char_class['unknown']=False
                            nextToken=-1
                            current_char='EOF'
                            print("Next token is: {},".format(nextToken), "Next lexeme is {}".format(current_char))
                            break
                        print("Next token is: {},".format(nextToken), "Next lexeme is {}".format(current_char))
    except Exception as e:
        print(e)

if __name__ == "__main__":
    lex()