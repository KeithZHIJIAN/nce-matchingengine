#!/usr/bin/env python
from src.matchingengine import MatchingEngine

from utils.MQHelper import MQHelper

import sys, os

"""
Message format:
-   Add, Symbol, Type, Side, Quantity, Price, Owner ID, Wallet ID, Stop Price (Optional)
    add ETHUSD limit ask 100 64000 user1 Alice1 (60000)
    add ethusd market ask 100 0 user1 Alice1 (60000)
    


-   Modify, Symbol, Side, Order ID, prev Quantity, prev Price, new Quantity, new Price
    modify ETHUSD buy 0000000002 100 63000 100 64000 //change price to 64000 only



-   Cancel, Symbol, Side, Price, Order ID
    cancel ETHUSD buy 100 0000000001
"""

# The front end will send valid order msg to the matching engine

if __name__ == "__main__":
    try:
        me = MatchingEngine.get_instance()
        MQHelper.listen_rabbitmq(me)
    except KeyboardInterrupt:
        print("Interrupted")
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
