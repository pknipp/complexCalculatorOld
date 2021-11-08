phase, polar, rect, exp, log, log, & trig [3 * 2 (direct vs inv) * 2 (circ vs hyperbolic)]

create class which has 4 properties: real, imag, mod, phase

class complex(x, y):
    this.real = x
    this.imag = y
    this.mod = sqrt(x * x + y * y)
    this.phase = atan2(y, x)

unary operations:
cexp(z) = (exp(x) * cos(y), exp(x) * sin(y))
ln(z) = (ln(|x|), phase)
cos(z) = (cos x * cosh y, -sin x * sinh y), and analogously for sin (& tan = sin / cos)
cosh(z)= ...
inverse trig functions and inverse hyperbolic trig functions can be defined in terms of other functions.
to do:
    - logs to other bases
    - roots [= (...) ** (1/integer)]
binary operations:
    z1 + z2 = (x1 + x2, y1 + y2) & analogous for subtraction
    z1 * z2 = (x1 * x2 - y1 * y2, x1 * y2 + x2 * y1)
    z1 / z2 = ((x1 * x2 + y1 * y2)/den, (x2 * y1 - x1 * y2)/den), where den = (x2, y2).mod
    z1 ** n = ...

How to parse string:
ops = "+-*/^"
def parse(stri_in):
    stri = stri_in
    pairs = []
    z = stuff trimmed off stri, which equals either # or unary op or ()
    //unary arg is evaluated recursively, and then unary is evaluated
    // () is evaluated recursively

    while len(ops) > 1:
        if ops.indexOf(pairs[0].op) < ops.indexOf(pairs[1].op):
            pair = pairs[1] removed from array
            pairs[0].# = pairs[0].# pair.op pair.#
        else:
            pair = pairs[0] removed from array
            z = z pair.op pair.#
    return z ops[0].op ops[0].#

    helper functions:

    def trim_number(stri):
        char = 1st item trimmed from stri
        if char == "(":
            start a stack of open/close parens, as follows:
                if next char is something other than a paren:
                    pass
                else
                    if next item is an open paren:
                        put it on stack
                    else
                        remove last item on stack
            return recursive call to parse, for contents of ()
        build up #, one char at a time

    def trim_op(stri):
        return ops.indexOf(1st item trimmed from stri)
